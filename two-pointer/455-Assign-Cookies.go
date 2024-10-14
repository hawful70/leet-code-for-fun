func findContentChildren(g []int, s []int) int {
    if len(s) == 0 {
\t\treturn 0
\t}

\tsort.Ints(g)
\tsort.Ints(s)

\tgPointer := len(g) - 1
\tsPointer := len(s) - 1
\tans := 0

\tfor gPointer >= 0 && sPointer >= 0 {
\t\tif g[gPointer] <= s[sPointer] {
\t\t\tans++
\t\t\tsPointer--
\t\t}
\t\tgPointer--
\t}

\treturn ans
}