package day_4

import (
	"log"
	"os"
)

func PerformTaskOne(filePath string) int {
	return 0
}

func PerformTaskTwo(filePath string) int {
	return 0
}

// note: will most likely need to go back to a line scanner for day 4
func GetInput(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
