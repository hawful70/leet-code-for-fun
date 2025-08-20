package main

func shortestToChar(s string, c byte) []int {
	answer := make([]int, len(s))
	n := len(s)
	prev := -n

	for i := 0; i < n; i++ {
		if s[i] == c {
			prev = i
		}
		answer[i] = i - prev
	}

	prev = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		answer[i] = min(answer[i], prev-i)
	}
	return answer
}
