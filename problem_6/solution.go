package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largestNonAdjSum(arr []int) int {
	first := 0
	next := 0
	for _, i := range arr {
		temp := findMax(first, next)
		first = next + i
		next = temp
	}
	return findMax(first, next)

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
		str_arr := strings.Split(content[0], ",")
		arr := []int{}
		for _, i := range str_arr {
			j, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
			}
			arr = append(arr, j)
		}

		ans, err := strconv.Atoi(content[1])
		if err != nil {
			fmt.Println(err)
		}

		if largestNonAdjSum(arr) == ans {
			fmt.Println("Passed")
		} else {
			fmt.Println("Failed")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
