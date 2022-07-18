package main

func main() {

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	column := len(matrix[0])
	row := 0
	for row < rows && column >= 0 {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] < target {
			row++
		} else {
			column--
		}
	}
	return false
}
