package main

import (
	"crypto/rand"
	"testing"
)

func BenchmarkPrime256(b *testing.B) {
	var bit int = 256

	for i := 0; i < b.N; i++ {
		rand.Prime(rand.Reader, bit)
	}
}
func BenchmarkPrime512(b *testing.B) {
	var bit int = 512

	for i := 0; i < b.N; i++ {
		rand.Prime(rand.Reader, bit)
	}
}
func BenchmarkPrime1024(b *testing.B) {
	var bit int = 1024

	for i := 0; i < b.N; i++ {
		rand.Prime(rand.Reader, bit)
	}
}
func BenchmarkPrime2048(b *testing.B) {
	var bit int = 2048

	for i := 0; i < b.N; i++ {
		rand.Prime(rand.Reader, bit)
	}
}

// go test -test.bench=".*" -benchmem -benchtime 10s -count 1
