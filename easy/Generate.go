package easy

//func main() {
//	a := generate(5)
//	for i := 0; i < 5; i++ {
//		for j := 0; j <= i; j++ {
//			print(a[i][j])
//			print(" ")
//		}
//		println()
//	}
//}

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		temp := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 {
				temp = append(temp, 1)

			} else if i == j {
				temp = append(temp, 1)
			} else {
				temp = append(temp, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, temp)
	}
	return result
}
func getRow(rowIndex int) []int {
	result := make([][]int, 0)
	for i := 0; i < rowIndex+1; i++ {
		temp := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 {
				temp = append(temp, 1)

			} else if i == j {
				temp = append(temp, 1)
			} else {
				temp = append(temp, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, temp)
	}
	return result[rowIndex]
}
