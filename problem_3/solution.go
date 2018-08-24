package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func productArray(arr []int) []int {
	n := len(arr)
	prod_arr := []int{}
	for i := 0; i < n; i++ {
		prod_arr = append(prod_arr, 1)
	}

	temp := 1
	for i := 0; i < n; i++ {
		prod_arr[i] = temp
		temp *= arr[i]
	}

	temp = 1
	for i := n - 1; i >= 0; i-- {
		prod_arr[i] *= temp
		temp *= arr[i]
	}

	return prod_arr

}

func main() {
	file, err := os.Open("test_cases.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str_arr := strings.Split(scanner.Text(), ",")
		arr := []int{}
		for _, i := range str_arr {
			j, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
			}
			arr = append(arr, j)
		}

		fmt.Printf("Product array: %v\n", productArray(arr))

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
