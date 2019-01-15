# prime benchmark

## prime

`c++` `openssl/bn.h` `bn_generate_prime_ex()`

|Benchmark    |             Time     |Time     |        CPU |  Iterations|
|-|-|-|-|-|
|BM_gen_prime256|     2054894 ns|0.002054894 ms|     2052461 ns|        20858|
|BM_gen_prime512|     7957842 ns|0.007957842 ms|     7950281 ns|         5173|
|BM_gen_prime1024|   39886462 ns|0.039886462 ms|    39872998 ns|         1104|
|BM_gen_prime2048|  262401459 ns|0.262401459 ms|   262291253 ns|          130|

`c++` `openssl/gmp.h` `find_next_prime()`

|Benchmark    |             Time     |Time     |        CPU |  Iterations|
|-|-|-|-|-|
|BM_gen_prime256|      463608 ns|0.000463608 ms|       463427 ns|        91107|
|BM_gen_prime512|     2873275 ns|0.002873275 ms|      2872962 ns|        14658|
|BM_gen_prime1024|   26029270 ns|0.026029270 ms|     26022030 ns|         1628|
|BM_gen_prime2048|  280237060 ns|0.280237060 ms|    280146581 ns|          160|

`golang` `golang.org/crypto/rand`

|Benchmark    |             Time     |Time     |  Iterations|
|-|-|-|-|
|BenchmarkPrime256|2228119 ns|0.002228119  ms|20000 |
|BenchmarkPrime512|13188122 ns|0.013188122  ms|3000 |
|BenchmarkPrime1024|120893304 ns|0.120893304  ms|300 |
|BenchmarkPrime2048|1217249497 ns|1.217249497  ms|30 |


`js` `nodejs/crypto`

depends on openssl/bn.h

same.

`python`

interface not found.