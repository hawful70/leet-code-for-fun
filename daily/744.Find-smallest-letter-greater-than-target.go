package daily

func nextGreatestLetter(letters []byte, target byte) byte {
	var ans byte
	low, high := 0, len(letters)-1
	for low <= high {
		mid := (high + low) / 2
		if letters[mid] > target {
			ans = letters[mid]
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if ans != 0 {
		return ans
	}
	return letters[0]
}
