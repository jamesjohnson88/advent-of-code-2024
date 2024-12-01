package handlers

import (
	"bufio"
	"log"
	"os"
	"sort"
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

	var leftNums, rightNums = []int{}, []int{}
	i := 0
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), "   ")
		left, _ := strconv.Atoi(split[0])
		leftNums = append(leftNums, left)
		right, _ := strconv.Atoi(split[1])
		rightNums = append(rightNums, right)
		i++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)
	total := 0
	for i := 0; i < len(leftNums); i++ {
		if left, right := leftNums[i], rightNums[i]; left > right {
			total += left - right
		} else {
			total += right - left
		}
	}

	return total
}

func PerformTaskTwo(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var leftNums = []int{}
	counts := make(map[int]int)
	i := 0
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), "   ")
		left, _ := strconv.Atoi(split[0])
		leftNums = append(leftNums, left)
		right, _ := strconv.Atoi(split[1])
		counts[right]++
		i++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, num := range leftNums {
		total += num * counts[num]
	}

	return total
}
