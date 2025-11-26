func reverseWords(s string) string {
\t// Convert string to a slice of characters (runes)
\tsList := []rune(s)
\t
\t// Function to reverse characters between indices l and h
\tutilReverseWord := func(l, h int) {
\t\tfor l < h {
\t\t\tsList[l], sList[h] = sList[h], sList[l]
\t\t\tl++
\t\t\th--
\t\t}
\t}
\t
\tlow := 0
\tfor high := 1; high < len(sList); high++ {
\t\tif sList[high] == ' ' {
\t\t\tutilReverseWord(low, high-1)
\t\t\tlow = high + 1
\t\t} else if high == len(sList)-1 {
\t\t\tutilReverseWord(low, len(sList)-1)
\t\t}
\t}
\t
\t// Convert the rune slice back to a string and return it
\treturn string(sList)
}