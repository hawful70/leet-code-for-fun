import "fmt"
func maximumRemovals(s string, p string, removable []int) int {
    left, right := 0, len(removable)
    ans := 0
    for left <= right {
        mid := (left + right) / 2
        if isValidRemoval(mid, s, p, removable) {
            ans = max(mid, ans)
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return ans
}

func isValidRemoval(subsequence int, s string, p string, removable []int) bool {
    removal := make(map[int]bool)
    
    for _, r := range(removable[:subsequence]) {
        removal[r] = true
    }

    pPointer := 0
    for i := range s {
        if removal[i] {
            continue
        }
        if s[i] == p[pPointer] {
            pPointer += 1
            if pPointer == len(p) {
                return true
            }
        }
    }
    return false
}