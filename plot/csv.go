package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"testing"

	"github.com/esote/leng"
)

var a []int

func fill() {
	for i := range a {
		a[i] = rand.Intn(len(a))
	}
}

func BenchmarkSort(b *testing.B, f func([]int)) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fill()
		b.StartTimer()
		f(a)
	}
}

func main() {
	var sizes []int
	for i := uint(9); i < uint(25); i++ {
		sizes = append(sizes, 1<<i)
		sizes = append(sizes, 1<<i+1<<(i-1)-1<<(i-2)-1<<(i-3))
		sizes = append(sizes, 1<<i+1<<(i-1)-1<<(i-2))
		sizes = append(sizes, 1<<i+1<<(i-1)-1<<(i-2)+1<<(i-3))
		sizes = append(sizes, 1<<i+1<<(i-1))
		sizes = append(sizes, 1<<i+1<<(i-1)+1<<(i-2)-1<<(i-3))
		sizes = append(sizes, 1<<i+1<<(i-1)+1<<(i-2))
		sizes = append(sizes, 1<<i+1<<(i-1)+1<<(i-2)+1<<(i-3))
	}

	time, _ := os.OpenFile("time.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	defer time.Close()

	fmt.Fprintln(time, "size,lengSort,lengSortInPlace,stl")

	for _, size := range sizes {
		a = make([]int, size)
		lengSort := testing.Benchmark(func(b *testing.B) {
			BenchmarkSort(b, leng.Sort)
		})
		lengSortInPlace := testing.Benchmark(func(b *testing.B) {
			BenchmarkSort(b, leng.SortInPlace)
		})
		stl := testing.Benchmark(func(b *testing.B) {
			BenchmarkSort(b, sort.Ints)
		})
		fmt.Fprintf(time, "%d,%d,%d,%d\n", size, lengSort.NsPerOp(),
			lengSortInPlace.NsPerOp(), stl.NsPerOp())
	}
}
