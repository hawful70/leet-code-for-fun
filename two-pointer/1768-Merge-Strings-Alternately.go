func mergeAlternately(word1 string, word2 string) string {
    word1_pointer, word2_pointer := 0, 0
    var res []rune
    for word1_pointer < len(word1) && word2_pointer < len(word2) {
        res = append(res, rune(word1[word1_pointer]))
        res = append(res, rune(word2[word2_pointer]))
        word1_pointer++
        word2_pointer++
    }

    for word1_pointer < len(word1) {
        res = append(res, rune(word1[word1_pointer]))
        word1_pointer++
    }

    for word2_pointer < len(word2) {
        res = append(res, rune(word2[word2_pointer]))
        word2_pointer++
    }

    return string(res)
}