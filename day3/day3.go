package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	mul_regex   = "mul\\(([1-9][0-9]*),([1-9][0-9]*)\\)"
	do_regex    = "do\\(\\)"
	donot_regex = "don't\\(\\)"
)

type Range struct {
	Start int
	End   int
}

type RangeSlice []Range

func (arr RangeSlice) covers(i int) bool {
  for _, r := range arr {
    if r.Start <= i && i <= r.End{
      return true
    }
  }
  return false
}

type MulMatch struct {
	Value int
	Index int
}

func getAcceptableRanges(input_text string) RangeSlice {
	do_re := regexp.MustCompile(do_regex)
	donot_re := regexp.MustCompile(donot_regex)

	acceptable_ranges := make([]Range, 0)
	do_matches := do_re.FindAllStringIndex(string(input_text), -1)
	donot_matches := donot_re.FindAllStringIndex(string(input_text), -1)

	curr_range_start := 0
	var last_donot_idx int

	do_idx := 0
	donot_idx := 0

	for donot_idx < len(donot_matches) {
		if donot_matches[donot_idx][1] > curr_range_start {
			// current legal range has ended, append it to the allowed ranges
			last_donot_idx = donot_idx
			acceptable_ranges = append(acceptable_ranges, Range{Start: curr_range_start, End: donot_matches[last_donot_idx][0]})

			// move do pointer forward until it passes the last donot (there could be multiple do's in a row, so this updates so that
			// we start the next range after the last donot)
			for do_matches[do_idx][1] < donot_matches[last_donot_idx][1] {
				do_idx++
			}
			// make sure we haven't crossed the end of the do match list
			if do_idx < len(do_matches) {
				curr_range_start = do_matches[do_idx][1]
			} else {
				break
			}
		}
		donot_idx++
	}

	// covers case where string ends on an acceptable range (last token was a do)
	if curr_range_start > donot_matches[last_donot_idx][1] && curr_range_start < len(input_text) {
		acceptable_ranges = append(acceptable_ranges, Range{curr_range_start, len(input_text) - 1})
	}

	return acceptable_ranges
}

func getMulMatches(input_text string) []MulMatch {
	mul_re := regexp.MustCompile(mul_regex)

	matches := mul_re.FindAllSubmatchIndex([]byte(input_text), -1)
  mul_matches := make([]MulMatch, 0)
	for _, match := range matches {
		nums_as_strings := make([]string, 2)
		nums_as_strings[0] = input_text[match[2]:match[3]]
		nums_as_strings[1] = input_text[match[4]:match[5]]
		num1, _ := strconv.Atoi(nums_as_strings[0])
		num2, _ := strconv.Atoi(nums_as_strings[1])
    mul := MulMatch{
      Value: num1 * num2,
      Index: match[0],
    }
    mul_matches = append(mul_matches, mul)
	}
	return mul_matches

}

func part1() {
	input_text, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

  mul_matches := getMulMatches(string(input_text))
  total := 0
  for _, mul := range mul_matches {
    total += mul.Value
  }
  fmt.Println("Part 1 total: ", total)
}

func part2() {
	input_text, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mul_re := regexp.MustCompile(mul_regex)
	acceptable_ranges := getAcceptableRanges(string(input_text))
  mul_matches := getMulMatches(string(input_text))

  total := 0
  for _, mul := range mul_matches {
    if acceptable_ranges.covers(mul.Index){
      total += mul.Value
    }
  }

  fmt.Println("Part 2 total: ", total)
}

func main() {
	part1()
	part2()
}
