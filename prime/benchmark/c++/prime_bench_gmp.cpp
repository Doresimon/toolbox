#include <benchmark/benchmark.h>
#include <gmp.h>

static void BM_gen_prime256(benchmark::State &state)
{
  mpz_t rop;
  mpz_init(rop);
  mpz_t p;
  mpz_init(p);
  gmp_randstate_t rands;
  gmp_randinit_default(rands);
  gmp_randseed_ui(rands, 256L);

  for (auto _ : state)
  {
    state.PauseTiming();
    mpz_urandomb(rop, rands, 256);
    state.ResumeTiming();

    mpz_nextprime(p, rop);
  }
}

static void BM_gen_prime512(benchmark::State &state)
{
  mpz_t rop;
  mpz_init(rop);
  mpz_t p;
  mpz_init(p);
  gmp_randstate_t rands;
  gmp_randinit_default(rands);
  gmp_randseed_ui(rands, 512L);

  for (auto _ : state)
  {
    state.PauseTiming();
    mpz_urandomb(rop, rands, 512);
    state.ResumeTiming();
    mpz_nextprime(p, rop);
  }
}

static void BM_gen_prime1024(benchmark::State &state)
{
  mpz_t rop;
  mpz_init(rop);
  mpz_t p;
  mpz_init(p);
  gmp_randstate_t rands;
  gmp_randinit_default(rands);
  gmp_randseed_ui(rands, 1024L);

  for (auto _ : state)
  {
    state.PauseTiming();
    mpz_urandomb(rop, rands, 1024);
    state.ResumeTiming();
    mpz_nextprime(p, rop);
  }
}

static void BM_gen_prime2048(benchmark::State &state)
{
  mpz_t rop;
  mpz_init(rop);
  mpz_t p;
  mpz_init(p);
  gmp_randstate_t rands;
  gmp_randinit_default(rands);
  gmp_randseed_ui(rands, 2048L);

  for (auto _ : state)
  {
    state.PauseTiming();
    mpz_urandomb(rop, rands, 2048);
    state.ResumeTiming();
    mpz_nextprime(p, rop);
  }
}

// Register the function as a benchmark
BENCHMARK(BM_gen_prime256);
BENCHMARK(BM_gen_prime512);
BENCHMARK(BM_gen_prime1024);
BENCHMARK(BM_gen_prime2048);

BENCHMARK_MAIN();

// g++ prime_bench_gmp.cpp -o prime_bench_gmp -lcrypto -lbenchmark -lbenchmark_main -pthread -lgmp

// > ./prime_bench_gmp --benchmark_min_time=30
