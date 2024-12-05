func minFlips(s string) int {
    minFlip := len(s)

    window_size := len(s)
    if window_size % 2 != 0 {
        s = s + s
    }

    dif1, dif2 := 0, 0
    left := 0
    for right := 0; right < len(s); right++ {
        var alternating1, alternating2 byte
        if right % 2 == 0 {
            alternating1 = '0'
            alternating2 = '1'
        } else {
            alternating1 = '1'
            alternating2 = '0'
        }

        if s[right] == alternating1 {
            dif1++
        }
        if s[right] == alternating2 {
            dif2++
        }

        if right - left + 1 == window_size {
            if left % 2 == 0 {
                alternating1 = '0'
                alternating2 = '1'
            } else {
                alternating1 = '1'
                alternating2 = '0'
            }

            minFlip = min(minFlip, dif1, dif2)
            if s[left] == alternating1 {
                dif1--
            }
            if s[left] == alternating2 {
                dif2--
            }
            left++
        }
    }
    return minFlip
}

func min(a int, b int, c int) int {
    min := a
    if min > b {
        min = b
    }

    if min > c {
        min = c
    }

    return min
}