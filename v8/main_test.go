package main

import "testing"

func BenchmarkA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateRandomString(4)
	}
}

func BenchmarkB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateRandomString(4)
	}
}
