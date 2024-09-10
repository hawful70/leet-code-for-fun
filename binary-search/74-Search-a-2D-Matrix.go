func searchMatrix(matrix [][]int, target int) bool {
    ROWS, COLS := len(matrix), len(matrix[0])

    top, bottom := 0, ROWS - 1
    for top <= bottom {
        mid := (top + bottom) / 2
        if matrix[mid][COLS - 1] < target {
            top = mid + 1
        } else if matrix[mid][0] > target {
            bottom = mid - 1
        } else {
            break
        }
    }

    if top > bottom {
        return false
    }

    row := (top + bottom) / 2
    left, right := 0, COLS - 1
    for left <= right {
        mid := (left + right) / 2
        if matrix[row][mid] == target {
            return true
        } else if matrix[row][mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}