# prime benchmark

## prime

`c++` `openssl/bn.h`

|Benchmark    |             Time     |        CPU |  Iterations|
|-------------|-----------------------|--------------|---------|
|BM_gen_prime256   |  1970684 ns  |    1968494 ns |         306|
|BM_gen_prime512  |   8739779 ns  |    8733908 ns  |         68|
|BM_gen_prime1024 |  32514318 ns  |   32477103 ns  |         28|
|BM_gen_prime2048 | 216849501 ns  |  216550829 ns  |          3|

`golang` `golang.org/crypto/rand`

|Benchmark    |             Time     |        CPU |  Iterations|
|-------------|-----------------------|--------------|---------|
|BenchmarkPrime256-2       |            1987115 ns/op    |      213676 B/op       |10000 |
|BenchmarkPrime512-2      |            13006211 ns/op     |     449956 B/op      | 1000 |
|BenchmarkPrime1024-2    |              105809165 ns/op    |      951834 B/op   |    200|
|BenchmarkPrime2048-2   |               1735267086 ns/op    |     5114944 B/op |     30|

`js` `nodejs/crypto`

depends on openssl/bn.h

same.

`python`

interface not found.