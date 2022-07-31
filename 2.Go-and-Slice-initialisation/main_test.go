package main

import (
	"testing"
)

const (
	sliceLen = 10
)

func BenchmarkToStrSlice_Append(b *testing.B) {
	in := prepareInput(sliceLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapAppend(in)
	}
}

func BenchmarkToStrSlice_MakeAppend(b *testing.B) {
	in := prepareInput(sliceLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapMakeAppend(in)
	}
}

func BenchmarkToStrSlice_MakeLoop(b *testing.B) {
	in := prepareInput(sliceLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapMakeLoop(in)
	}
}

func prepareInput(n int) []int {
	slice := make([]int, n, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}
