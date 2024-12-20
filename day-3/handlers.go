package day_3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PerformTaskOne(filePath string) int {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := regex.FindAllString(GetInput(filePath), -1)

	total := 0
	for _, match := range matches {
		trimmed := strings.Trim(match, "mul()")
		split := strings.Split(trimmed, ",")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		total += num1 * num2
	}

	return total
}

func PerformTaskTwo(filePath string) int {
	regex := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	matches := regex.FindAllString(GetInput(filePath), -1)

	total := 0
	enabled := true

	for _, match := range matches {
		switch match {
		case "do()":
			enabled = true
			break
		case "don't()":
			enabled = false
			break
		default:
			if enabled {
				trimmed := strings.Trim(match, "mul()")
				split := strings.Split(trimmed, ",")
				num1, _ := strconv.Atoi(split[0])
				num2, _ := strconv.Atoi(split[1])
				total += num1 * num2
			}
		}
	}

	return total
}

func GetInput(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
