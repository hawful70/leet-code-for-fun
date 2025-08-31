package daily

func diStringMatch(s string) []int {
    n := len(s)
    low, high := 0, n
    res := make([]int, n+1)
    for i := 0; i < n; i++ {
        if s[i] == 'I' {
            res[i] = low
            low++
        } else if s[i] == 'D' {
            res[i] = high
            high--
        }
    }

    if s[n-1] == 'I' {
        res[n] = high
    } else {
        res[n] = low
    }
    return res
}

