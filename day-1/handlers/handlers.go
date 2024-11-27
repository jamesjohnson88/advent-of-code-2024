package handlers

import (
	"bufio"
	"log"
	"os"
)

func PerformTaskOne(filePath string) string {
	log.Printf("Opening %v", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	for fileScanner.Scan() {
		i++
		log.Printf("Line %v is: '%v'", i, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ":)"
}
