using Microsoft.Extensions.Caching.Memory;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AboutMemoryCache
{
    internal class CacheState<TKey, TValue>
    {
        private readonly CancellationTokenSource tokenSource;
        private readonly TimeSpan lifetime;

        public CacheState(TKey key, CancellationToken token = default, TimeSpan? lifetime = null)
        {
            this.Key = key;
            this.Status = Status = CacheStatus.NotReady | CacheStatus.NotRefreshing;
            this.lifetime = lifetime ?? TimeSpan.MaxValue;
            this.tokenSource = CancellationTokenSource.CreateLinkedTokenSource(token);
        }

        public TKey Key { get; set; }

        public TValue Value { get; set; }

        public CacheStatus Status { get; set; }

        public DateTime LastRefresh { get; set; }

        public Task RefreshTask { get; set; }

        public bool IsReady => (this.Status & CacheStatus.Ready) == CacheStatus.Ready;

        public bool IsRefreshing => (this.Status & CacheStatus.Refreshing) == CacheStatus.Refreshing;

        public bool IsExpired => (DateTime.UtcNow - this.LastRefresh) > this.lifetime;

        public void TriggerRefresh(Func<TKey, CancellationToken, Task<TValue>> func)
        {
            if (this.RefreshTask == null ||
                this.RefreshTask.IsCompleted)
            {
                this.RefreshTask = this.Refresh(func);
            }
        }

        public async Task Refresh(Func<TKey, CancellationToken, Task<TValue>> func)
        {
            this.Status |= CacheStatus.Refreshing;

            var value = await func(this.Key, this.tokenSource.Token);

            this.Value = value;
            this.Status = CacheStatus.Ready | CacheStatus.NotRefreshing;
            this.LastRefresh = DateTime.UtcNow;
        }

        public void Dispose()
        {
            this.tokenSource.Cancel();
            this.tokenSource.Dispose();
        }
    }
}
