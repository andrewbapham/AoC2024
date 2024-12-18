package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Direction struct {
	X int
	Y int
}

var (
	directions = []Direction{
		{X: 1, Y: 0},   // Right
		{X: -1, Y: 0},  // Left
		{X: 0, Y: 1},   // Up
		{X: 0, Y: -1},  // Down
		{X: 1, Y: 1},   // Down Right
		{X: 1, Y: -1},  // Up Right
		{X: -1, Y: 1},  // Down Left
		{X: -1, Y: -1}, // Up Left
	}
)

func searchWord(grid []string, word string, index int, dir Direction, pos Direction) bool {
	if index == len(word) {
		// search complete
		return true
	} else if pos.X < 0 || pos.X >= len(grid[0]) || pos.Y < 0 || pos.Y >= len(grid) {
		// out of bounds case
		return false
	} else if grid[pos.Y][pos.X] != word[index] {
		// characters no longer match
		return false
	}

	pos.X += dir.X
	pos.Y += dir.Y
	return searchWord(grid, word, index+1, dir, pos)
}

func parseInput(filename string) []string {
	input_file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer input_file.Close()

	scanner := bufio.NewScanner(input_file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1Concurrent() {
	lines := parseInput("input.txt")

	count := 0
	var wg sync.WaitGroup
  var mu sync.Mutex

	for i := range lines {
		for j := range lines[i] {
			for _, dir := range directions {
				wg.Add(1)
				go func() {
					defer wg.Done()
					pos := Direction{X: j, Y: i}
					if searchWord(lines, "XMAS", 0, dir, pos) {
            mu.Lock()
						count++
            mu.Unlock()
					}
				}()
			}
		}
	}
	wg.Wait()
	fmt.Println("Part 1 total is: ", count)
}

func part1() {
	lines := parseInput("input.txt")

	count := 0

	for i := range lines {
		for j := range lines[i] {
			for _, dir := range directions {
				pos := Direction{X: j, Y: i}
				if searchWord(lines, "XMAS", 0, dir, pos) {
					count++
				}
			}
		}
	}
	fmt.Println("Part 1 total is: ", count)
}

func main() {
  start := time.Now()
	part1()
  elapsed := time.Since(start)
  fmt.Println("Part 1 ran in", elapsed)

  start = time.Now()
  part2()
  elapsed = time.Since(start)
  fmt.Println("Part 2 ran in", elapsed)
}
