package main

import (
	"fmt"
	"sort"
)

func main() {
	weight := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := 5
	fmt.Println(shipWithinDays(weight, d))
	// result := sort.Search(len(weight), func(i int) bool {
	// 	return weight[i] >= d
	// })
	// if result < len(weight) && weight[result] == d {
	// 	fmt.Printf("i find it:%d\n",result)
	// } else {
	// 	fmt.Printf("i can't find it:%d\n",result)
	// }

}
/*
	题解：照抄的官方题解 
	假设当船的运载能力为 xx 时，我们可以在 DD 天内运送完所有包裹，那么只要运载能力大于 xx，我们同样可以在 DD 天内运送完所有包裹：我们只需要使用运载能力为 xx 时的运送方法即可。

这样一来，我们就得到了一个非常重要的结论：

存在一个运载能力的「下限」minX，使得当 x >= minX时，我们可以在 D 天内运送完所有包裹；当 x < minX时，我们无法在 DD 天内运送完所有包裹。
同时minX 即为我们需要求出的答案。因此，我们就可以使用二分查找的方法找出 minX的值。

在二分查找的每一步中，我们实际上需要解决一个判定问题：给定船的运载能力 x，我们是否可以在 D 天内运送完所有包裹呢？这个判定问题可以通过贪心的方法来解决：

由于我们必须按照数组 weights 中包裹的顺序进行运送，因此我们从数组 weights 的首元素开始遍历，将连续的包裹都安排在同一天进行运送。当这批包裹的重量大于运载能力 x 时，我们就需要将最后一个包裹拿出来，安排在新的一天，并继续往下遍历。当我们遍历完整个数组后，就得到了最少需要运送的天数。

我们将「最少需要运送的天数」与 D 进行比较，就可以解决这个判定问题。当其小于等于 D 时，我们就忽略二分的右半部分区间；当其大于 D 时，我们就忽略二分的左半部分区间。

细节

二分查找的初始左右边界应当如何计算呢？

对于左边界而言，由于我们不能「拆分」一个包裹，因此船的运载能力不能小于所有包裹中最重的那个的重量，即左边界为数组 weights 中元素的最大值。

对于右边界而言，船的运载能力也不会大于所有包裹的重量之和，即右边界为数组 weights 中元素的和。

我们从上述左右边界开始进行二分查找，就可以保证找到最终的答案。
	解题方法：先确定一下二分查找的左右边界，然后利用二分查找查找到right-left中符合要求的最小边界
*/
func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, weight := range weights {
		if left < weight {
			left = weight
		}
		right += weight
	}

	return left + sort.Search(right-left, func(x int) bool {
		x += left
		sum := 0
		day := 1
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}
