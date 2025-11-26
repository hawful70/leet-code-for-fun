func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(strs[i]) < len(prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		for strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}