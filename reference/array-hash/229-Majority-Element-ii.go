// func majorityElement(nums []int) []int {
//     frequency_item := make(map[int]int)
//     for _, num := range nums {
//         frequency_item[num]++
//     }

//     var results []int
//     for key, value := range frequency_item {
//         if value > (len(nums) / 3) {
//             results = append(results, key)
//         }
//     }
//     return results
// }

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	candidate1, candidate2, count1, count2 := 0, 0, 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1, count1 = num, 1
		} else if count2 == 0 {
			candidate2, count2 = num, 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	var results []int
	if count1 > len(nums)/3 {
		results = append(results, candidate1)
	}
	if count2 > len(nums)/3 {
		results = append(results, candidate2)
	}

	return results
}