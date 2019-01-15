#include <benchmark/benchmark.h>
#include <openssl/bn.h>

static void BM_gen_prime256(benchmark::State& state) {
  auto ret = BN_new();
  int bit = 256;
  for (auto _ : state)
    BN_generate_prime_ex(ret, bit, false, NULL,NULL, NULL);
}
static void BM_gen_prime512(benchmark::State& state) {
  auto ret = BN_new();
  int bit = 512;
  for (auto _ : state)
    BN_generate_prime_ex(ret, bit, false, NULL,NULL, NULL);
}
static void BM_gen_prime1024(benchmark::State& state) {
  auto ret = BN_new();
  int bit = 1024;
  for (auto _ : state)
    BN_generate_prime_ex(ret, bit, false, NULL,NULL, NULL);
}
static void BM_gen_prime2048(benchmark::State& state) {
  auto ret = BN_new();
  int bit = 2048;
  for (auto _ : state)
    BN_generate_prime_ex(ret, bit, false, NULL,NULL, NULL);
}
// Register the function as a benchmark
BENCHMARK(BM_gen_prime256);
BENCHMARK(BM_gen_prime512);
BENCHMARK(BM_gen_prime1024);
BENCHMARK(BM_gen_prime2048);

BENCHMARK_MAIN();

 
// g++ prime_bench.cpp -lcrypto -lbenchmark -lbenchmark_main -pthread