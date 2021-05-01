package main

func main() {
	
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	importMap := map[int]*Employee{}
	for _, employe := range employees {
		importMap[employe.Id] = employe
	}
	sum := 0
	queue := []int{id}
	for len(queue) > 0 {
		employe := importMap[queue[0]]
		sum += employe.Importance
		queue = queue[1:]
		for _, sub := range employe.Subordinates {
			queue = append(queue, sub)
		}
	}
	return sum
}
