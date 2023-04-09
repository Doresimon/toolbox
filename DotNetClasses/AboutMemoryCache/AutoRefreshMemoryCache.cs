using Microsoft.Extensions.Caching.Memory;

namespace AboutMemoryCache
{
    /// <summary>
    /// Auto refresh memory cache.
    /// </summary>
    /// <typeparam name="TKey"></typeparam>
    /// <typeparam name="TValue"></typeparam>
    public class AutoRefreshMemoryCache<TKey, TValue>
    {
        private readonly MemoryCache cache;
        private readonly Func<TKey, CancellationToken, Task<TValue>> valueFactory;
        private readonly SemaphoreSlim asyncLock;
        private readonly double refreshRatio;
        private readonly TimeSpan slidingExpiration;
        private readonly CancellationToken refreshTaskCancellationToken;

        public AutoRefreshMemoryCache(
            TimeSpan slidingExpiration,
            double refreshRatio,
            CancellationToken refreshTaskCancellationToken,
            Func<TKey, CancellationToken, Task<TValue>> valueFactory)
        {
            this.slidingExpiration = slidingExpiration;
            this.refreshTaskCancellationToken = refreshTaskCancellationToken;
            this.refreshRatio = refreshRatio;
            this.valueFactory = valueFactory;

            this.asyncLock = new SemaphoreSlim(1);
            this.cache = new MemoryCache(new MemoryCacheOptions());
        }

        /// <summary>
        /// Get value from cache, if not exist, create a new one.
        /// </summary>
        /// <param name="key"></param>
        /// <param name="cancellationToken"></param>
        /// <returns></returns>
        public async Task<TValue> GetAsync(TKey key, CancellationToken cancellationToken)
        {
            if (!this.cache.TryGetValue<CacheState<TKey, TValue>>(key, out var state))
            {
                await this.asyncLock.WaitAsync(cancellationToken);
                try
                {
                    if (!this.cache.TryGetValue(key, out state))
                    {
                        state = new CacheState<TKey, TValue>(key, refreshTaskCancellationToken, slidingExpiration * refreshRatio);

                        var options = new MemoryCacheEntryOptions();
                        options.SlidingExpiration = slidingExpiration;
                        options.RegisterPostEvictionCallback((key, value, reason, state) =>
                        {
                            if (value != null)
                            {
                                ((CacheState<TKey, TValue>)value)?.Dispose();
                            }
                        }, state);

                        this.cache.Set(key, state, new MemoryCacheEntryOptions { SlidingExpiration = slidingExpiration });
                    }
                }
                finally
                {
                    this.asyncLock.Release();
                }
            }

            if (!state.IsReady)
            {
                if (!state.IsRefreshing)
                {
                    await this.asyncLock.WaitAsync(cancellationToken);
                    try
                    {
                        state.TriggerRefresh(this.valueFactory);
                    }
                    finally
                    {
                        this.asyncLock.Release();
                    }
                }

                // wait for refreshing
                using var cts = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken, this.refreshTaskCancellationToken);
                try
                {
                    await Task.WhenAny(Task.Delay(TimeSpan.FromMinutes(5), cts.Token), state.RefreshTask);
                }
                catch (Exception)
                {
                }
                finally
                {
                    cts.Cancel();
                }
            }
            else
            {
                if (state.IsExpired)
                {
                    state.TriggerRefresh(this.valueFactory);
                }
            }

            return state.Value;
        }


        public void Remove(TKey key)
        {
            this.cache.Remove(key);
        }

        public void Clear()
        {
            this.cache.Clear();
        }
    }
}