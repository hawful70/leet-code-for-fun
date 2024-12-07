func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    s1Count := make([]int, 26)
    s2Count := make([]int, 26)
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i] - 'a']++
        s2Count[s2[i] - 'a']++
    }

    for i := 0; i < len(s2) - len(s1); i++ {
        if match(s1Count, s2Count) {
            return true
        }

        s2Count[s2[i] - 'a']--
        s2Count[s2[i + len(s1)] - 'a']++
    }

    return match(s1Count, s2Count)
}

func match (a, b []int) bool {
    for i := 0; i < 26; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}