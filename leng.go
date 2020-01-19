// Package leng provides some sorting algorithms.
package leng

import (
	"sync"

	"github.com/esote/leng/sn"
)

const shellCutoff = 1 << 10

// Sort an array of integers quickly.
func Sort(a []int) {
	l := len(a)
	switch {
	case l <= 1:
		return
	case l <= sn.MaxSN:
		sn.Jump(l)(a)
	case l < shellCutoff:
		shellSort(a, l)
	default:
		mid := l / 2
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			Sort(a[:mid])
			wg.Done()
		}()
		Sort(a[mid:])
		wg.Wait()
		merge(a, mid)
	}
}

func merge(a []int, mid int) {
	s := make([]int, len(a))
	copy(s, a)

	left := 0
	right := mid
	cur := 0
	high := len(s) - 1

	for left < mid && right <= high {
		if s[left] <= s[right] {
			a[cur] = s[left]
			left++
		} else {
			a[cur] = s[right]
			right++
		}
		cur++
	}

	copy(a[cur:], s[left:mid])
}

// SortInPlace performs in-place merge sort with O(1) space. From "Practical
// In-Place Mergesort" by Katajainen, Pasanen, and Teuhola. This implementation
// does no memory allocation.
func SortInPlace(a []int) {
	l := len(a)
	switch {
	case l <= 1:
		return
	case l <= sn.MaxSN:
		sn.Jump(l)(a)
	case l < shellCutoff:
		shellSort(a, l)
	default:
		mid := l / 2
		w := l - mid
		wsort(a, 0, mid, w)
		for w > 2 {
			n := w
			w = (n + 1) / 2
			wsort(a, w, n, 0)
			wmerge(a, 0, n-w, n, l, w)
		}
		for ; w > 0; w-- {
			for j := w; j < l && a[j] < a[j-1]; j++ {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func wsort(a []int, start, end, w int) {
	mid := start + (end-start)/2
	SortInPlace(a[start:mid])
	SortInPlace(a[mid:end])
	wmerge(a, start, mid, mid, end, w)
}

func wmerge(a []int, i, m, j, n, w int) {
	for i < m && j < n {
		if a[i] < a[j] {
			a[w], a[i] = a[i], a[w]
			i++
		} else {
			a[w], a[j] = a[j], a[w]
			j++
		}
		w++
	}
	for i < m {
		a[w], a[i] = a[i], a[w]
		w++
		i++
	}
	for j < n {
		a[w], a[j] = a[j], a[w]
		w++
		j++
	}
}

func shellSort(a []int, l int) {
	for gap := l; gap > 0; gap >>= 1 {
		for i := gap; i < l; i++ {
			j, tmp := i, a[i]
			for ; j >= gap && a[j-gap] > tmp; j -= gap {
				a[j] = a[j-gap]
			}
			a[j] = tmp
		}
	}
}

// Parallel O((lg n)^2) merge algorithm. This does not perform well, even when
// using serial merge as a base case for small inputs.
/*
func mergeParallel(a, b, c []int, i, j, k, l, p, q int) {
	m := j - i
	n := l - k
	if m < n {
		a, b = b, a
		m, n = n, m

		i, j, k, l = k, l, i, j
	}
	if m <= 0 {
		return
	}
	r := (i + j) / 2
	s := sort.Search(n, func(x int) bool {
		return b[k+x] >= a[r]
	})
	s += k
	t := p + (r - i) + (s - k)
	c[t] = a[r]
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		mergeParallel(a, b, c, i, r, k, s, p, t)
		wg.Done()
	}()
	go func() {
		mergeParallel(a, b, c, r+1, j, s, l, t+1, q)
		wg.Done()
	}()
	wg.Wait()
}
*/
