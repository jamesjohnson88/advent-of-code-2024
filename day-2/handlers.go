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
			if increasing && current > last && current <= last+offset {
				if i == len(split)-1 {
					safe++
				}
			} else if !increasing && current < last && last-current <= offset {
				if i == len(split)-1 {
					safe++
				}
			} else {
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

	const offset = 3
	row := 0
	safe := 0
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), " ")
		s1, _ := strconv.Atoi(split[0])
		s2, _ := strconv.Atoi(split[1])
		s3, _ := strconv.Atoi(split[2])
		s4, _ := strconv.Atoi(split[3])
		s5, _ := strconv.Atoi(split[4])
		last, usedSkip, iOffset := s1, false, 0
		increasing := s1 < (s3+s4+s5)/3
		log.Printf("%v v %v", s1, (s3+s4+s5)/3)
		if s1 == s2 || ((!increasing && s2 < s1-offset) || (!increasing && s1 < s2) || (increasing && s2 > s1+offset)) {
			log.Printf("%v v %v", s1, s2)
			log.Print("hit skipper")
			if (increasing && s1 > s2 && s3 > s2) || (!increasing && s1 < s2 && s3 < s2) {
				log.Print("hit index2 skipper")
				usedSkip = true
				iOffset = 1
				last = s1
			} else {
				log.Print("hit generic skipper")
				usedSkip = true
				iOffset = 1
				last = s2
			}
		}

		for i := 1 + iOffset; i < len(split); i++ {
			current, _ := strconv.Atoi(split[i])
			log.Print("----------")
			log.Printf("increasing: %v, s1: %v, s2: %v, current: %v, last: %v", increasing, s1, s2, current, last)
			if increasing && current > last && current <= last+offset {
				if i == len(split)-1 {
					log.Print("---SAFE---")
					safe++
				}
				last = current
			} else if !increasing && current < last && last-current <= offset {
				if i == len(split)-1 {
					log.Print("---SAFE---")
					safe++
				}
				last = current
			} else {
				if !usedSkip {
					usedSkip = true
					if i < 2 {
						last = current
					}
					log.Printf("Used Skip!")
					if i == len(split)-1 {
						log.Print("---SAFE---")
						safe++
					}
				} else {
					log.Printf("---UNSAFE: %v---", fileScanner.Text())
					break
				}
			}
		}
		row++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return safe
}
