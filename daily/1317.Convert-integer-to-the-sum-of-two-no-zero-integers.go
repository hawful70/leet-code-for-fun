package daily

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		a := i
		b := n - i
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return nil
}

func hasZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}
