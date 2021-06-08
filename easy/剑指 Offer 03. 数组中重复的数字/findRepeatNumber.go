package main


func main() {

}

func findRepeatNumber(nums []int) int {
	mp := map[int]bool{}
	for _,num := range nums {
		if _,ok := mp[num];ok {
			return num
		}else{
			mp[num] = true
		}
	}
	return 0
}
