func mergeAlternately(word1 string, word2 string) string {
    var ans string
    for len(word1) > 0 && len(word2) > 0 {
        ans += string(word1[0])
        ans += string(word2[0])
        word1 = word1[1:]
        word2 = word2[1:]
    }

    ans += word1
    ans += word2
    return ans
}