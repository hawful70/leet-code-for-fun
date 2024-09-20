func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
\tif len(nums1) > len(nums2) {
\t\tnums1, nums2 = nums2, nums1
\t}

\tm, n := len(nums1), len(nums2)
\tlow, high := 0, m

\tfor low <= high {
\t\tpartitionX := (low + high) / 2
\t\tpartitionY := (m + n + 1) / 2 - partitionX

\t\tmaxX := math.MinInt64
\t\tif partitionX > 0 {
\t\t\tmaxX = nums1[partitionX-1]
\t\t}

\t\tminX := math.MaxInt64
\t\tif partitionX < m {
\t\t\tminX = nums1[partitionX]
\t\t}

\t\tmaxY := math.MinInt64
\t\tif partitionY > 0 {
\t\t\tmaxY = nums2[partitionY-1]
\t\t}

\t\tminY := math.MaxInt64
\t\tif partitionY < n {
\t\t\tminY = nums2[partitionY]
\t\t}

\t\tif maxX <= minY && maxY <= minX {
\t\t\tif (m+n)%2 == 0 {
\t\t\t\treturn (float64(max(maxX, maxY)) + float64(min(minX, minY))) / 2.0
\t\t\t}
\t\t\treturn float64(max(maxX, maxY))
\t\t} else if maxX > minY {
\t\t\thigh = partitionX - 1
\t\t} else {
\t\t\tlow = partitionX + 1
\t\t}
\t}

\treturn 0.0
}

func max(a, b int) int {
\tif a > b {
\t\treturn a
\t}
\treturn b
}

func min(a, b int) int {
\tif a < b {
\t\treturn a
\t}
\treturn b
}