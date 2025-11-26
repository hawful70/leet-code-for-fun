func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    i, j := 0, 0
    pos1, pos2 := 0, 0
    for i < len(word1) && j < len(word2) {
        if word1[i][pos1] != word2[j][pos2] {
            return false
        }

        pos1++
        pos2++
        if pos1 == len(word1[i]) {
            i++
            pos1 = 0
        }
        if pos2 == len(word2[j]) {
            j++
            pos2 = 0
        }
    }

    return i == len(word1) && j == len(word2)
}