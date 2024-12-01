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
	log.Printf("Opening %v", filePath)
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
		i++
		split := strings.Split(fileScanner.Text(), "   ")
		left, _ := strconv.Atoi(split[0])
		leftNums = append(leftNums, left)
		right, _ := strconv.Atoi(split[1])
		rightNums = append(rightNums, right)
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
