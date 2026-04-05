package daily

import "math"

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	firstCriticalPoint := -1
	lastCriticalPoint := -1
	minDistance := math.MaxInt32

	prev := head
	curr := head.Next
	currentIndex := 1

	for curr.Next != nil {
		next := curr.Next
		maxima := curr.Val > prev.Val && curr.Val > next.Val
		minima := curr.Val < prev.Val && curr.Val < next.Val

		if maxima || minima {
			if firstCriticalPoint == -1 {
				firstCriticalPoint = currentIndex
			} else {
				dist := currentIndex - lastCriticalPoint
				if dist < minDistance {
					minDistance = dist
				}
			}
			lastCriticalPoint = currentIndex
		}

		prev = curr
		curr = next
		currentIndex++
	}

	if firstCriticalPoint == lastCriticalPoint {
		return []int{-1, -1}
	}

	return []int{minDistance, lastCriticalPoint - firstCriticalPoint}
}
