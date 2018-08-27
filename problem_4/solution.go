package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func segregate(arr []int) int {
	n := len(arr)
	j := 0
	for i := 0; i < n; i++ {
		if arr[i] <= 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return j

}

func findMissingPos(arr []int) int {
	size := len(arr)
	neg_size := segregate(arr)
	pos_size := size - neg_size
	for i := neg_size; i < size; i++ {
		index := int(math.Abs(float64(arr[i]))) - 1 + neg_size
		if index < size && arr[index] > 0 {
			arr[index] = -arr[index]
		}
	}

	for i := neg_size; i < size; i++ {
		if arr[i] > 0 {
			return i - neg_size + 1
		}
	}

	return pos_size + 1
}

func main() {
	file, err := os.Open("test_cases.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Reading text_cases.txt in variables
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
		ans, ans_err := strconv.Atoi(content[1])
		if ans_err != nil {
			fmt.Println(ans_err)
		}

		if findMissingPos(arr) == ans {
			fmt.Println("Passed")
		} else {
			fmt.Println("Failed")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
