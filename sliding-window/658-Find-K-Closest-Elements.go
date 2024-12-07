func findClosestElements(arr []int, k int, x int) []int {
    currentDistance := 0
    for i := 0; i < k; i++ {
        currentDistance += abs(arr[i] - x)
    }

    left := 0
    minDistance := currentDistance
    for i := 1; i <= len(arr) - k; i++ {
        currentDistance = currentDistance - abs(arr[i - 1] - x) + abs(arr[i+k-1] - x)
        if currentDistance < minDistance {
            minDistance = currentDistance
            left = i
        }
    }
    
    return arr[left : left+k]
}

func abs (a int) int {
    if a < 0 {
        return -a
    }
    return a
}