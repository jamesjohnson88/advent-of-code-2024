package day_2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func PerformTaskOne(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	const offset = 3
	row := 0
	safe := 0
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), " ")
		s1, _ := strconv.Atoi(split[0])
		s2, _ := strconv.Atoi(split[1])
		last := s2
		increasing := s1 < s2

		for i := 2; i < len(split); i++ {
			if s1 == s2 || (s2 < s1-offset || s2 > s1+offset) {
				break
			}
			current, _ := strconv.Atoi(split[i])
			log.Print("----------")
			log.Printf("increasing: %v, s1: %v, s2: %v, current: %v, last: %v", increasing, s1, s2, current, last)
			if increasing && current > last && current <= last+offset {
				log.Printf("%v was larger than %v but less than %v.", current, last, current+offset)
				if i == len(split)-1 {
					log.Print("---SAFE---")
					safe++
				}
			} else if !increasing && current < last && last-current <= offset {
				log.Printf("%v was smaller than %v but larger than %v", current, last, current-offset)
				if i == len(split)-1 {
					log.Print("---SAFE---")
					safe++
				}
			} else {
				log.Printf("---UNSAFE ln%v---", row)
				break
			}
			last = current
		}
		row++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return safe
}

func PerformTaskTwo(filePath string) int {
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
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0
}
