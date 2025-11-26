func suggestedProducts(products []string, searchWord string) [][]string {
    res := [][]string{}
	sort.Strings(products)

    left, right := 0, len(products) - 1
    for i := 0; i < len(searchWord); i++ {
        s := string(searchWord[i])

        for left <= right && (len(products[left]) <= i || string(products[left][i]) != s) {
			left++
		}
		for left <= right && (len(products[right]) <= i || string(products[right][i]) != s) {
			right--
		}

        res = append(res, []string{})
		remain := right - left + 1
        for j := 0; j < min(3, remain); j++ {
			res[len(res)-1] = append(res[len(res)-1], products[left+j])
		}
    }
    return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}