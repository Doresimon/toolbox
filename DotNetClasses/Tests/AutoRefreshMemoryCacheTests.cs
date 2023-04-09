using AboutMemoryCache;

namespace Tests
{
    [TestClass]
    public class AutoRefreshMemoryCacheTests
    {
        [TestMethod]
        public async Task Basic()
        {
            var slidingExipration = TimeSpan.FromSeconds(5);
            var refreshRatio = 0.2;
            var cts = new CancellationTokenSource();
            var cache = new AutoRefreshMemoryCache<string, DateTime>(slidingExipration, refreshRatio, cts.Token, async (key, token) =>
            {
                await Task.Delay(0, token);
                return DateTime.UtcNow;
            });

            var cts2 = new CancellationTokenSource();

            DateTime time = DateTime.UtcNow;
            var i = 0;
            while (i++ < 10)
            {
                await Task.Delay(1111);
                var value = await cache.GetAsync("bike", cts2.Token);
                Assert.IsTrue(value > time);
                time = value;
            }
        }
    }
}