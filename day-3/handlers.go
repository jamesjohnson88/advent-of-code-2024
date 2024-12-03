package day_3

import (
	"bufio"
	"log"
	"os"
)

func PerformTaskOne(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	for fileScanner.Scan() {
		// todo
		i++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0
}

func PerformTaskTwo(filePath string) int {
	// todo

	return 0
}
