package daily

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true

		next_sum := 0
		for n > 0 {
			digits := n % 10
			next_sum += digits * digits
			n = n / 10
		}
		n = next_sum
	}
	return n == 1
}
