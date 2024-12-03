package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var leftArr []int
	var rightArr []int
	file, err := os.Open("day_1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		res := strings.Split(line, " ")
		leftNum, err := strconv.Atoi(res[0])
		if err != nil {
			panic(err)
		}
		leftArr = append(leftArr, leftNum)

		rightNum, err := strconv.Atoi(res[len(res)-1])
		if err != nil {
			panic(err)
		}
		rightArr = append(rightArr, rightNum)
	}

	// Sort the array
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	totalDist := 0
	fmt.Println("Total Length of the array is: ", len(leftArr))
	for i := 0; i < len(leftArr); i++ {
		dist := rightArr[i] - leftArr[i]
		dist = int(math.Abs(float64(dist)))
		totalDist += dist
	}
	fmt.Println("Total distance is: ", totalDist)
}
