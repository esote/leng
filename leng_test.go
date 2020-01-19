package leng

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func fill(a []int) {
	for i := range a {
		a[i] = rand.Intn(len(a))
	}
}

func TestSort(t *testing.T) {
	a := make([]int, 1000000)
	fill(a)
	Sort(a)
	if !sort.IntsAreSorted(a) {
		t.Fatal("not sorted")
	}
}

func TestSortInPlace(t *testing.T) {
	a := make([]int, 1000000)
	fill(a)
	SortInPlace(a)
	if !sort.IntsAreSorted(a) {
		t.Fatal("not sorted")
	}
}

func BenchmarkCmp(b *testing.B) {
	sizes := []int{
		1e3,
		1e4,
		1e5,
		1e6,
		1e7,
	}

	if testing.Short() {
		sizes = sizes[len(sizes)-1:]
	}

	for _, s := range sizes {
		n := strconv.Itoa(s)
		a := make([]int, s)
		b.Run("sort-"+n, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				fill(a)
				b.StartTimer()
				Sort(a)
			}
		})
		b.Run("inplace-"+n, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				fill(a)
				b.StartTimer()
				SortInPlace(a)
			}
		})
		b.Run("stl-"+n, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				fill(a)
				b.StartTimer()
				sort.Ints(a)
			}
		})
		fmt.Println()
	}
}
