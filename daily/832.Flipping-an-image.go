package daily

func flipAndInvertImageV2(image [][]int) [][]int {
	rows := len(image)
	columns := len(image[0])

	for i := 0; i < rows; i++ {
		left := 0
		right := columns - 1

		for left <= right {
			image[i][left], image[i][right] = image[i][right], image[i][left]
			left++
			right--
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			switch image[i][j] {
			case 0:
				image[i][j] = 1
			case 1:
				image[i][j] = 0
			}
		}
	}

	return image
}

func flipAndInvertImage(image [][]int) [][]int {
	rows := len(image)
	columns := len(image[0])

	for i := 0; i < rows; i++ {
		left, right := 0, columns-1
		for left <= right {
			if image[i][left] == image[i][right] {
				image[i][left] ^= 1
				image[i][right] = image[i][left]
			}
			left++
			right--
		}
	}

	return image
}
