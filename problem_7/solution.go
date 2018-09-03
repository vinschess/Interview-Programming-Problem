package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Where m represents total ways of steps at a time
func stepsCount(n int, m int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = 0
		j := 1
		for j <= m && j <= i {
			arr[i] += arr[i-j]
			j++
		}
	}
	return arr[n]
}

func main() {
	file, err := os.Open("test_cases.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := strings.Split(scanner.Text(), " ")
		steps, err := strconv.Atoi(content[0])
		if err != nil {
			fmt.Println(err)
		}
		ans, err := strconv.Atoi(content[1])
		if err != nil {
			fmt.Println(err)
		}
		if stepsCount(steps, 2) == ans {
			fmt.Println("Passed")
		} else {
			fmt.Println("Failed")
		}
	}
}
