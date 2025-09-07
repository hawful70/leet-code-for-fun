package daily

func canTransform(start string, result string) bool {
	left, right := 0, 0
	n := len(start)
	for left < len(start) && right < len(result) {
		for left < len(start) && start[left] == 'X' {
			left++
		}
		for right < len(result) && result[right] == 'X' {
			right++
		}

		if left == n || right == n {
			break
		}

		if start[left] != result[right] {
			return false
		}
		if start[left] == 'L' && left < right {
			return false
		}
		if start[left] == 'R' && left > right {
			return false
		}

		left++
		right++
	}

	for left < n {
		if start[left] != 'X' {
			return false
		}
		left++
	}
	for right < n {
		if result[right] != 'X' {
			return false
		}
		right++
	}

	return true
}
