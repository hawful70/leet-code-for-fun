package daily

func longestMountain(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	ans := 0
	left := 0
	for left+2 < n {
		right := left
		if arr[right] < arr[right+1] {
			for right+1 < n && arr[right] < arr[right+1] {
				right++
			}

			if right+1 < n && arr[right] > arr[right+1] {
				for right+1 < n && arr[right] > arr[right+1] {
					right++
				}
				ans = max(ans, right-left+1)
			}
		}

		if right > left {
			left = right
		} else {
			left++
		}
	}
	return ans
}
