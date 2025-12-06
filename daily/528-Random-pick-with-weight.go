package daily

import "math/rand"

type Solution struct {
	prefix []int
	total  int
}

func Constructor(w []int) Solution {
	prefix := make([]int, len(w))
	prefix[0] = w[0]
	for i := 1; i < len(w); i++ {
		prefix[i] = prefix[i-1] + w[i]
	}
	return Solution{prefix, prefix[len(prefix)-1]}
}

func (s *Solution) PickIndex() int {
	target := rand.Intn(s.total) + 1 // random number from 1..total

	lo, hi := 0, len(s.prefix)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if s.prefix[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
