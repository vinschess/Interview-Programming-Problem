package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decodeCount(str_digit string) int {
	count_2nd_last := 1
	count_last := 1
	count_curr := 1
	n := len(str_digit)
	for i := 2; i <= n; i++ {
		count_curr = 0

		if str_digit[i-1] > '0' {
			count_curr = count_last
		}

		if str_digit[i-2] == '1' || (str_digit[i-2] == '2' && str_digit[i-1] < '7') {
			count_curr += count_2nd_last
		}

		count_2nd_last, count_last = count_last, count_curr
	}

	return count_curr

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
		str_digit := content[0]
		ans, ans_err := strconv.Atoi(content[1])
		if ans_err != nil {
			fmt.Println(ans_err)
		}

		if decodeCount(str_digit) == ans {
			fmt.Println("Passed")
		} else {
			fmt.Println("Failed")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
