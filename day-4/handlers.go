package day_4

import (
	"bufio"
	"log"
	"os"
)

func PerformTaskOne(filePath, targetWord string) int {
	return findMatches(GetInput(filePath), targetWord)
}

func PerformTaskTwo(filePath string) int {
	return 0
}

func findMatches(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][2]int{
		{0, 1}, {0, -1}, // Left, Right
		{1, 0}, {-1, 0}, // Up, Down
		{1, 1}, {1, -1}, {-1, -1}, {-1, 1}, // Diagonals
	}

	isMatch := func(x, y, dx, dy int) bool {
		for i := 0; i < len(word); i++ {
			nx, ny := x+i*dx, y+i*dy
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	matchCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if isMatch(i, j, dx, dy) {
					matchCount++
				}
			}
		}
	}

	return matchCount
}

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
