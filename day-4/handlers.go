package day_4

import (
	"bufio"
	"log"
	"os"
)

func PerformTaskOne(filePath, targetWord string) int {
	return findMatches(GetInput(filePath), targetWord)
}

func PerformTaskTwo(filePath, targetWord string) int {
	return findCrossShapeMatches(GetInput(filePath), targetWord)
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

func findCrossShapeMatches(grid [][]rune, word string) int {
	if len(word) != 3 {
		log.Fatal("word must have 3 characters")
	}

	rows := len(grid)
	cols := len(grid[0])
	centerChar := rune(word[1])
	diagonalPairs := [][2][2]int{
		{{-1, -1}, {1, 1}}, // top-left to bottom-right
		{{-1, 1}, {1, -1}}, // top-right to bottom-left
	}

	isCrossShapedMatch := func(x, y int) bool {
		diagonalMatches := 0
		// two pairs, each represents one diagonal line of the 'X'
		for _, pair := range diagonalPairs {
			d1, d2 := pair[0], pair[1]
			// apply diagonal transforms to grid coordinates
			nx1, ny1 := x+d1[0], y+d1[1]
			nx2, ny2 := x+d2[0], y+d2[1]

			// check bounds for new coordinates
			if nx1 >= 0 && ny1 >= 0 && nx1 < rows && ny1 < cols &&
				nx2 >= 0 && ny2 >= 0 && nx2 < rows && ny2 < cols {
				// ensure diagonal is valid, either forward or backwards
				if (grid[nx1][ny1] == rune(word[0]) && grid[nx2][ny2] == rune(word[2])) ||
					(grid[nx1][ny1] == rune(word[2]) && grid[nx2][ny2] == rune(word[0])) {
					diagonalMatches++
				}
			}
		}

		// both valid matches means a valid 'X'
		return diagonalMatches == len(diagonalPairs)
	}

	matchCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == centerChar && isCrossShapedMatch(i, j) {
				matchCount++
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
