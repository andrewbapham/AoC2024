package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	mul_regex = "mul\\(([1-9][0-9]*),([1-9][0-9]*)\\)"
)

func part1() {
	input_text, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	r := regexp.MustCompile(mul_regex)

	matches := r.FindAllSubmatch(input_text, -1)
	total := 0
	for _, match := range matches {
		nums_as_strings := make([]string, 2)
		nums_as_strings[0] = string(match[1])
		nums_as_strings[1] = string(match[2])
		num1, _ := strconv.Atoi(nums_as_strings[0])
		num2, _ := strconv.Atoi(nums_as_strings[1])
		total += num1 * num2
	}
	fmt.Println("Part 1 total is: ", total)
}

func main() {
	part1()
}
