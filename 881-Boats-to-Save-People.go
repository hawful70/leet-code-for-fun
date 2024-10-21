func numRescueBoats(people []int, limit int) int {
    if len(people) == 1 && people[0] <= limit {
        return 1
    }
    if len(people) == 1 && people[0] > limit {
        return 0
    }

    sort.Ints(people)
    left, right := 0, len(people) - 1
    ans := 0
    for left <= right {
        if people[left] + people[right] <= limit {
            left++
            right--
        } else {
            right--
        }
        ans += 1
    }
    return ans
}