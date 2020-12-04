package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countTrees1(lines []string) int {
	currIndex := 0
	trees := 0
	for i, l := range lines {
		// First Line seems no to count
		if i == 0 {
			currIndex += 3
			continue
		}

		// Check for Tree
		ls := strings.Split(l, "")
		if ls[currIndex] == "#" {
			trees++
		}

		currIndex += 3
		currIndex %= len(ls)
	}
	return trees
}

func countTrees2(lines []string) int {
	indexPolicy1 := 1 // Right 1, down 1.
	indexPolicy2 := 3 // Right 3, down 1.
	indexPolicy3 := 5 // Right 5, down 1.
	indexPolicy4 := 7 // Right 7, down 1.
	indexPolicy5 := 1 // Right 1, down 2.

	treesPolicy1 := 0 // Right 1, down 1.
	treesPolicy2 := 0 // Right 3, down 1.
	treesPolicy3 := 0 // Right 5, down 1.
	treesPolicy4 := 0 // Right 7, down 1.
	treesPolicy5 := 0 // Right 1, down 2.

	for i, l := range lines {
		if i == 0 {
			continue
		}
		ls := strings.Split(l, "")
		if ls[indexPolicy1] == "#" {
			treesPolicy1++
		}
		if ls[indexPolicy2] == "#" {
			treesPolicy2++
		}
		if ls[indexPolicy3] == "#" {
			treesPolicy3++
		}
		if ls[indexPolicy4] == "#" {
			treesPolicy4++
		}
		if (i%2 == 0) && (ls[indexPolicy5] == "#") {
			treesPolicy5++
		}

		indexPolicy1++
		indexPolicy1 %= len(ls)

		indexPolicy2 += 3
		indexPolicy2 %= len(ls)

		indexPolicy3 += 5
		indexPolicy3 %= len(ls)

		indexPolicy4 += 7
		indexPolicy4 %= len(ls)

		if i%2 == 0 {
			indexPolicy5++
			indexPolicy5 %= len(ls)
		}
	}
	fmt.Println(treesPolicy1, treesPolicy2, treesPolicy3, treesPolicy4, treesPolicy5)

	return treesPolicy1 * treesPolicy2 * treesPolicy3 * treesPolicy4 * treesPolicy5

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	trees1 := countTrees1(lines)
	fmt.Println("Trees 1: ", trees1)

	trees2 := countTrees2(lines)
	fmt.Println("Trees 2: ", trees2)
}
