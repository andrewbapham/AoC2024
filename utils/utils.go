package utils

import (
	"bufio"
	"os"
)

func GetLines(filepath string) ([]string, error) {
	lines := make([]string, 0, 100)
	file, err := os.Open(filepath)
	if err != nil {
		return lines, err
	}
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan(){
    lines = append(lines, scanner.Text())
  }
  return lines, nil
}

func Map[T, V any](ts []T, fn func(T) V) []V {
  res := make([]V, len(ts))
  for i, t:= range ts {
    res[i] = fn(t)
  }
  return res
}

func AbsInt(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
