package day_5

import (
	"bufio"
	"log"
	"os"
)

func PerformTaskOne(filePath string) int {

	return 0
}

func PerformTaskTwo(filePath string) int { return 0 }

func GetInput(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var grid [][]rune
	for fileScanner.Scan() {
		grid = append(grid, []rune(fileScanner.Text()))
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
