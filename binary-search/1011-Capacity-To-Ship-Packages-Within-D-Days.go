
import "math"
func shipWithinDays(weights []int, days int) int {
    if len(weights) == 1 {
        return weights[0]
    }

    left, right := max(weights), sum(weights)
    ans := sum(weights)
    for left <= right {
        mid := (left + right) / 2
        if validCapacity(weights, days, mid) {
            ans = int(math.Min(float64(ans), float64(mid)))
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

func sum(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func max(numbers []int) int {
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}

func validCapacity(weights []int, days int, capacity int) bool {
    day_shipped := 1
    current_capacity := capacity
    for _, weight := range weights {
        if weight > current_capacity {
            day_shipped += 1
            current_capacity = capacity
        }
        current_capacity -= weight
    }
    return day_shipped <= days
}