package daily

func maxValue(n int, index int, maxSum int) int {
	lo, hi := 1, maxSum
	ans := 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if feasible(mid, n, index, maxSum) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func feasible(x int, n int, index int, maxSum int) bool {
	sum := x
	sum += calcSide(x-1, index)
	sum += calcSide(x-1, n-index-1)

	return sum <= maxSum
}

func calcSide(start int, length int) int {
	if length == 0 {
		return 0
	}

	if start > length {
		end := start - length + 1
		return (start + end) * length / 2
	}

	sum := start * (start + 1) / 2
	sum += (length - start)
	return sum
}
