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

func countOutOfPlace(nums []int, increasing bool) int {
  out_of_place_count := 0
  for i := 0; i < len(nums) - 1; i++ {
    diff := nums[i+1] - nums[i]
    if !((increasing && 1 <= diff && diff <= 3) || (!increasing && -3 <= diff && diff <= -1)){
      out_of_place_count++
    }
  }
  return out_of_place_count
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func isSafeWithReplacement(nums []int, increasing bool) bool {
  valid := true
  for i := 0; i < len(nums) - 1; i++ {
      diff := nums[i+1] - nums[i]
      if !((increasing && 1 <= diff && diff <= 3) || (!increasing && -3 <= diff && diff <= -1)){
        // create new array removing the bad number; if it is still bad then return false
        new_nums := append([]int{}, nums...)
        new_nums = remove(new_nums, i)
        if !isSafe(new_nums, increasing){
          valid = false
        } else {
          valid = true
        }
      }    
  }
  // catches edge case where bad number is at the very end
  new_nums := append([]int{}, nums...)
  new_nums = remove(new_nums, len(nums)-1)
  if isSafe(new_nums, increasing){
    valid = true
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
      safe_count++
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

func part2() {
  line_numbers, err := parseInput("input.txt")
  if err != nil {
    fmt.Println(err)
  }

  safe_count := 0
  for _, nums := range line_numbers {
    if len(nums) < 2 {
      safe_count++
    }

    if isSafeWithReplacement(nums, true) || isSafeWithReplacement(nums, false) {
      safe_count++
      // fmt.Println(nums)
    }
  }

  fmt.Println("Part 2 safe count: ", safe_count)
}

func main() {
	part1()
  part2()
}
