package daily

import "sort"

func threeSumMulti(arr []int, target int) int {
	const mod = 1000000007
	counted := make(map[int]int)

	for _, num := range arr {
		counted[num]++
	}

	keys := []int{}
	for k := range counted {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	ans := int64(0)
	for i := 0; i < len(keys); i++ {
		x := keys[i]
		for j := i; j < len(keys); j++ {
			y := keys[j]
			z := target - x - y

			if y > z {
				continue
			}

			if _, ok := counted[z]; !ok {
				continue
			}

			cx, cy, cz := counted[x], counted[y], counted[z]

			if x == y && y == z {
				ans += comb3(cx)
			} else if x == y && y != z {
				ans += comb2(cx) * int64(cz)
			} else if x != y && y == z {
				ans += int64(cx) * comb2(cy)
			} else {
				ans += int64(cx) * int64(cy) * int64(cz)
			}
		}
	}
	return int(ans % mod)
}

func comb2(n int) int64 {
	return int64(n) * int64(n-1) / 2
}

func comb3(n int) int64 {
	return int64(n) * int64(n-1) * int64(n-2) / 6
}
