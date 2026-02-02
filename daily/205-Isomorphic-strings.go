package daily

func isIsomorphic(s string, t string) bool {
	map_s := make(map[byte]byte)
	seen_t := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		char_s := s[i]
		char_t := t[i]

		if _, ok := map_s[char_s]; ok {
			if map_s[char_s] != char_t {
				return false
			}
		} else {
			if _, ok := seen_t[char_t]; ok {
				return false
			}

			map_s[char_s] = char_t
			seen_t[char_t] = true
		}
	}
	return true
}
