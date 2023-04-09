using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AboutMemoryCache
{
    public enum CacheStatus
    {
        NotReady = 0b00000000,
        Ready = 0b00000001,

        NotRefreshing = 0b00000000,
        Refreshing = 0b00000010,
    }
}
