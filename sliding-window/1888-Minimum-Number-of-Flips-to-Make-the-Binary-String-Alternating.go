func minFlips(s string) int {
	n := len(s)
	minFlip := n
	window_size := n

	// Extend the string for odd-length cases
	if window_size % 2 != 0 {
		s = s + s
	}

	dif1, dif2 := 0, 0
	left := 0

    fmt.Println("window_size", window_size)
	for right := 0; right < len(s); right++ {
		// Expected values for alternating patterns
		expected1 := byte('0' + (right % 2)) // '0' if even index, '1' if odd index
		expected2 := byte('1' - (right % 2)) // '1' if even index, '0' if odd index

		// Count differences for both patterns
		if s[right] != expected1 {
			dif1++
		}
		if s[right] != expected2 {
			dif2++
		}
        
		// Check if the window size is valid
		if right-left+1 == window_size {
            minFlip = min(minFlip, dif1, dif2)
			if s[left] != byte('0'+(left%2)) {
				dif1--
			}
			if s[left] != byte('1'-(left%2)) {
				dif2--
			}

			left++
		}
	}

	return minFlip
}

// Helper function to find the minimum of three integers
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}