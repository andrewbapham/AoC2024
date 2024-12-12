package main

import "fmt"

var (
	d_r_dir = Direction{X: 1, Y: 1}
	d_l_dir = Direction{X: -1, Y: 1}
)

func isTargetLetter(s string, target []string) bool {
	for _, t := range target {
		if s == t {
			return true
		}
	}
	return false
}

func checkCorners(grid []string, pos Direction) bool {
  if pos.X < 1 || pos.X >= len(grid[0]) - 1 || pos.Y < 1 || pos.Y >= len(grid) - 1{
    return false
  }
	// check top left to bottom right
	top_left := string(grid[pos.Y-1][pos.X-1])
	top_right := string(grid[pos.Y-1][pos.X+1])
	bottom_left := string(grid[pos.Y+1][pos.X-1])
	bottom_right := string(grid[pos.Y+1][pos.X+1])

	target := []string{"M", "S"}
	if !isTargetLetter(top_left, target) || !isTargetLetter(top_right, target) || !isTargetLetter(bottom_left, target) || !isTargetLetter(bottom_right, target) {
		return false
	}
  if top_left == bottom_right || top_right == bottom_left {
    return false
  }
  return true
}

func part2() {
	lines := parseInput("input.txt")
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
      if string(lines[i][j]) == "A"{
        if checkCorners(lines, Direction{X: j, Y: i}){
          count++
        }
      }
		}
	}

	fmt.Println("Part 2 total is: ", count)
}
