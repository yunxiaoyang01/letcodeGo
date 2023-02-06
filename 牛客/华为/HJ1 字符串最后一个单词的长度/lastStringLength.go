package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 填fmt.scanf的坑
func Scanf(input *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*input = string(data)
}

func main() {
	var input string
	for {
		Scanf(&input)
		lastIndex := strings.LastIndex(input, " ")
		fmt.Printf("%d\n", len(input[lastIndex+1:]))
	}

}
