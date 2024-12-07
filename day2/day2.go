package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrewbapham/AoC2024/utils"
)

func parseInput(filename string) ([][]int, error) {
	input_lines, err := utils.GetLines(filename)
  if err != nil {
    return nil, err
  }

  line_numbers := make([][]int, 0, 10)

  for _, line := range input_lines {
    line_tokens := strings.Split(line, " ")
    numbers := make([]int, 0, 5)
    for _, token := range line_tokens {
      num, _ := strconv.Atoi(token)
      numbers = append(numbers, num)
    }
    line_numbers = append(line_numbers, numbers)
  }
  return line_numbers, nil
}

func isSafe(nums []int, increasing bool) bool {
    valid := true
    for i := 0; i < len(nums) - 1; i++ {
      diff := nums[i+1] - nums[i]
      if !((increasing && 1 <= diff && diff <= 3) || (!increasing && -3 <= diff && diff <= -1)){
        valid = false
        break
      }     
    }
    return valid
}


func part1() {
  line_numbers, err := parseInput("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  
  safe_count := 0
  for _, nums := range line_numbers {
    if len(nums) < 2 {
      safe_count += 1
    }
    
    var increasing bool
    if nums[0] < nums[1] {
      increasing = true
    }else{
      increasing = false
    }

    safe := isSafe(nums, increasing)
    if safe {
      safe_count += 1
    }
  }

  fmt.Println("Safe count: ", safe_count)
}

func main() {
	part1()
}
