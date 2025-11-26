import (
\t\fmt\
\t\math\
)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    m, n := len(nums1), len(nums2)
    low, high := 0, m

    for low <= high {
        partitionSmallArray := (low + high) / 2
        partitionLargeArray := (m + n + 1) / 2 - partitionSmallArray

        maxLeftSmall := math.Inf(-1)
\t\tif partitionSmallArray > 0 {
\t\t\tmaxLeftSmall = float64(nums1[partitionSmallArray-1])
\t\t}

\t\tminRightSmall := math.Inf(1)
\t\tif partitionSmallArray < m {
\t\t\tminRightSmall = float64(nums1[partitionSmallArray])
\t\t}

\t\tmaxLeftLarge := math.Inf(-1)
\t\tif partitionLargeArray > 0 {
\t\t\tmaxLeftLarge = float64(nums2[partitionLargeArray-1])
\t\t}

\t\tminRightLarge := math.Inf(1)
\t\tif partitionLargeArray < n {
\t\t\tminRightLarge = float64(nums2[partitionLargeArray])
\t\t}

        if maxLeftSmall <= minRightLarge && maxLeftLarge <= minRightSmall {
\t\t\t// If total number of elements is even
\t\t\tif (m+n)%2 == 0 {
\t\t\t\treturn (math.Max(maxLeftSmall, maxLeftLarge) + math.Min(minRightSmall, minRightLarge)) / 2
\t\t\t}
\t\t\t// If total number of elements is odd
\t\t\treturn math.Max(maxLeftSmall, maxLeftLarge)
\t\t} else if maxLeftSmall > minRightLarge {
\t\t\t// Move partition left in the smaller array
\t\t\thigh = partitionSmallArray - 1
\t\t} else {
\t\t\t// Move partition right in the smaller array
\t\t\tlow = partitionSmallArray + 1
\t\t}
    }

    return 0.0
}