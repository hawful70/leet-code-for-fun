
import "sort"
func binarySearch(spell int, potions []int, target int64) int {
    left, right := 0, len(potions) - 1
    ans := -1
    for left <= right {
        mid := (left + right) / 2
        strength := spell * potions[mid]
        if int(strength) >= int(target) {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

func successfulPairs(spells []int, potions []int, success int64) []int {
    n := len(potions)
    sort.Ints(potions)

    var ans []int
    for _, spell := range spells {
        lower_index := binarySearch(spell, potions, success)
        if lower_index == -1 {
            lower_index = 0
        } else {
            lower_index = n - lower_index
        }
        ans = append(ans, lower_index)
    }
    return ans
}