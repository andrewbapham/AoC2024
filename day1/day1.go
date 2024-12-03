package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func absInt(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left_list := make([]int, 0, 100)
	right_list := make([]int, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, "   ")
		left_num, _ := strconv.Atoi(nums[0])
		right_num, _ := strconv.Atoi(nums[1])
		left_list = append(left_list, left_num)
		right_list = append(right_list, right_num)
	}
  
  slices.Sort(left_list)
  slices.Sort(right_list)

  total_distance := 0

  for i := 0; i < len(left_list); i++ {
    total_distance += absInt(left_list[i] - right_list[i])
  }

  fmt.Println("Answer for part 1:", total_distance)

}

func part2() {
  file, _ := os.Open("input.txt")
  defer file.Close()

  scanner := bufio.NewScanner(file)
  left_list := make([]int, 0, 100)
  right_counts := make(map[int]int)
  for scanner.Scan(){
    line := scanner.Text()

		nums := strings.Split(line, "   ")
		left_num, _ := strconv.Atoi(nums[0])
		right_num, _ := strconv.Atoi(nums[1])
		left_list = append(left_list, left_num)
    right_counts[right_num] += 1
  }

  total := 0
  for _, x := range left_list {
    total += right_counts[x] * x
  }

  fmt.Println("Answer for part 2:", total)
}

func main() {
	part1()
  part2()
}
