package daily

func nthUglyNumber(n int, a int, b int, c int) int {
	ab := lcm(a, b)
	ac := lcm(a, c)
	bc := lcm(b, c)
	abc := lcm(ab, c)

	lo, hi := 1, min(a, min(b, c))*n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countUglyNumbers(mid, a, b, c, ab, ac, bc, abc) >= n {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func countUglyNumbers(x, a, b, c, ab, ac, bc, abc int) int {
	return x/a + x/b + x/c -
		x/ab - x/ac - x/bc +
		x/abc
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
