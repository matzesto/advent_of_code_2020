package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func invalid(nums []int) int {
	for i, val := range nums {
		if i < 25 {
			continue
		}
		isFound := false
		for j, jval := range nums[i-25 : i] {
			for k, kval := range nums[i-25 : i] {
				if j == k {
					continue
				}

				if jval+kval == val {
					isFound = true
				}
			}
		}
		if isFound != true {
			return val
		}
	}
	return -1
}

func contiguous(nums []int, sum int) int {
	for i := range nums {
		partialSum := 0
		smallest := math.MaxInt64
		largest := 0
		for _, jval := range nums[i:] {
			partialSum += jval
			if partialSum > sum {
				continue
			}

			if jval > largest {
				largest = jval
			}

			if jval < smallest {
				smallest = jval
			}

			if partialSum == sum {
				return smallest + largest
			}
		}
	}

	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	nums := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, val)
	}
	res1 := invalid(nums)
	fmt.Println("Task 1: ", res1)

	res2 := contiguous(nums, res1)
	fmt.Println("Task 2: ", res2)
}
