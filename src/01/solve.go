package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// This is the naive implementation: Iterate over the list
func solution1(numbers []int, c chan int) {
	for indexI, i := range numbers {
		for indexJ, j := range numbers {
			if indexI == indexJ {
				continue
			}
			if (i + j) == 2020 {
				// fmt.Printf("Solution1: %d\n", i*j)
				c <- i * j
				return
			}
		}
	}
}

// Alternativ Solution with maps
func solution12(numbers map[int]bool, c chan int) {
	for i := range numbers {
		_, ok := numbers[2020-i]
		if !ok {
			continue
		}
		res := 2020 - i
		fmt.Println("Alternative: ", res*i)
		return
	}
}

// This is the same naive 0(n**3) approach
func solution2(numbers []int, c chan int) {
	for indexI, i := range numbers {
		for indexJ, j := range numbers {
			if indexI == indexJ {
				continue
			}
			for indexK, k := range numbers {
				if indexI == indexK {
					continue
				}
				if i+j+k == 2020 {
					c <- i * j * k
					return
				}
			}
		}
	}
}

func main() {
	var numbers = make([]int, 0)
	var numbers2 = make(map[int]bool)

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, val)
		numbers2[val] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(numbers)
	c1, c2 := make(chan int), make(chan int)

	// Not necessary, but why not use concurrency at this point
	go solution1(numbers, c1)
	go solution2(numbers, c2)

	solution1 := <-c1
	solution2 := <-c2

	fmt.Printf("Solution1: %d\n", solution1)
	fmt.Printf("Solution2: %d\n", solution2)
	solution12(numbers2, c1)
}
