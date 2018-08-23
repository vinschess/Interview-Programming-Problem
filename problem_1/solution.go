package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	a, b int
}

var exists = struct{}{}

type set struct {
	m map[int]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *set) Add(value int) {
	s.m[value] = exists
}

func (s *set) Contains(value int) bool {
	_, c := s.m[value]
	return c
}

// Method that returns all pair values which sum is equal to given number
func getPairs(numbers []int, sum_val int) []Pair {
	s := NewSet()
	pairs := []Pair{}
	for _, val := range numbers {
		diff := sum_val - val
		if s.Contains(diff) {
			pairs = append(pairs, Pair{diff, val})
		}
		s.Add(val)
	}
	return pairs
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
		k, k_err := strconv.Atoi(content[1])
		if k_err != nil {
			fmt.Println(err)
		}

		pair_vals := getPairs(arr, k)
		if len(pair_vals) > 0 {
			fmt.Printf("Following are the pairs which sum is equal to %v: %v.\n", k, pair_vals)
		} else {
			fmt.Printf("No pair is present which sum is euqal to %v.\n", k)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
