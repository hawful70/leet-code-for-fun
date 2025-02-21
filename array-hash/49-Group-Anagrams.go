func groupAnagrams(strs []string) [][]string {
	var results [][]string
	anagramHash := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		orderedStr := strings.Join(chars, "")
		anagramHash[orderedStr] = append(anagramHash[orderedStr], str)
	}

	for _, value := range anagramHash {
		results = append(results, value)
	}

	return results
}
