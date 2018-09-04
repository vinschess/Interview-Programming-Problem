package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(count_arr [26]int, k int) bool {
	val := 0
	for _, count := range count_arr {
		if count > 0 {
			val += 1
		}
	}
	return k >= val
}

func validateString(str string, k int) bool {
	var count_arr [26]int
	unique_char := 0
	for _, val := range str {
		if count_arr[val-'a'] == 0 {
			unique_char += 1
		}
		count_arr[val-'a'] += 1
	}
	return unique_char >= k
}

func longestSubstrKUniq(str string, k int) int {
	if !validateString(str, k) {
		return -1
	}

	curr_start := 0
	curr_end := 0

	max_window_size := 1
	// max_window_start := 0

	var count_arr [26]int

	count_arr[str[0]-'a'] = 1

	for i := 1; i < len(str); i++ {
		count_arr[str[i]-'a'] += 1
		curr_end += 1

		for !isValid(count_arr, k) {
			count_arr[str[curr_start]-'a'] -= 1
			curr_start += 1
		}

		if curr_end-curr_start+1 > max_window_size {
			max_window_size = curr_end - curr_start + 1
			// max_window_start = curr_start
		}
	}
	// fmt.Println("Max size: ", str[max_window_start:max_window_start+max_window_size])

	return max_window_size
}

func main() {
	file, err := os.Open("test_cases.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents := strings.Split(scanner.Text(), " ")
		str := contents[0]
		k, err := strconv.Atoi(contents[1])
		if err != nil {
			fmt.Println(err)
		}

		ans, err := strconv.Atoi(contents[2])
		if err != nil {
			fmt.Println(err)
		}

		if longestSubstrKUniq(str, k) == ans {
			fmt.Println("Passed")
		} else {
			fmt.Println("Failed")
		}

	}
}
