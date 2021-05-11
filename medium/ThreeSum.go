package medium


func threeSum1(nums []int) [][]int { //方案一错误，不能去重
	num := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					num = append(num, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return num
}
func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	num := make([][]int, 0)
	QuickSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		l, r := i+1, len(nums)-1 //定义左指针和右指针
		if nums[i] > 0 || nums[i]+nums[l] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] { //左指针去重
				l++
				continue
			}
			if r < len(nums)-2 && nums[r] == nums[r+1] { //右指针去重
				r--
				continue
			}
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				num = append(num, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return num
}
func QuickSort(array []int) {
	if len(array) < 2 {
		return
	}

	pivot := array[0]
	l, r := 0, len(array)-1

	for i := 1; i <= r; {
		if array[i] < pivot { //
			array[l], array[i] = array[i], array[l]
			l++
			i++
		} else {
			array[r], array[i] = array[i], array[r]
			r--
		}
	}

	QuickSort(array[:l])
	QuickSort(array[l+1:])
}
