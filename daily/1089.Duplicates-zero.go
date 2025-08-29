package daily

func duplicateZerosV2(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 && i < n-1 {
			for j := n - 2; j >= i; j-- {
				arr[j+1] = arr[j]
			}
			i++
		}
	}
}

func duplicateZeros(arr []int) {
	n := len(arr)
	zeros := 0

	for _, num := range arr {
		if num == 0 {
			zeros++
		}
	}

	i, j := n-1, n+zeros-1
	for i < j {
		if arr[i] != 0 {
			if j < n {
				arr[j] = arr[i]
			}
		} else {
			if j < n {
				arr[j] = 0
			}
			j--
			if j < n {
				arr[j] = 0
			}
		}
		i--
		j--
	}
}
