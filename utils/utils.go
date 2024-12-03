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
