package daily

import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	ans := []int{math.MaxInt64, 0, 0}
	left := 0

	formed := 0
	frequency_t := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		frequency_t[t[i]]++
	}
	required := len(frequency_t)
	current_window := make(map[byte]int)

	for right := range s {
		char_r := s[right]
		current_window[char_r]++

		if count_t, ok := frequency_t[char_r]; ok && current_window[char_r] == count_t {
			formed++
		}

		for left <= right && formed == required {
			char_l := s[left]

			if right-left+1 < ans[0] {
				ans = []int{right - left + 1, left, right}
			}
			current_window[char_l]--
			if count_t, ok := frequency_t[char_l]; ok && current_window[char_l] < count_t {
				formed--
			}
			left++
		}
	}
	if ans[0] == math.MaxInt64 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}
