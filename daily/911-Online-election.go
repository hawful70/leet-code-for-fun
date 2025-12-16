package daily

type TopVotedCandidate struct {
	persons []int
	times   []int
	leaders []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	count := make(map[int]int)
	maxVotes := 0
	leader := -1
	leaders := make([]int, len(persons))
	for i, p := range persons {
		count[p]++
		if count[p] >= maxVotes {
			maxVotes = count[p]
			leader = p
		}
		leaders[i] = leader
	}
	return TopVotedCandidate{
		persons,
		times,
		leaders,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	lo, hi := 0, len(this.times)-1
	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if this.times[mid] <= t {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return this.leaders[ans]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
