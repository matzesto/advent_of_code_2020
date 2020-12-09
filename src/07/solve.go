package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rules map[string]map[string]int

func parse(textRules []string) rules {
	var rls = rules{}

	re := regexp.MustCompile("[1-9] [a-z]* [a-z]*")
	for _, line := range textRules {
		color := strings.Split(line, "bags")[0]
		rls[color] = make(map[string]int)
		// No need to continue if current Bag does not contain other bags
		if strings.Contains(line, "no other bag") {
			continue
		}

		for _, tok := range re.FindAllString(line, -1) {
			num, _ := strconv.Atoi(strings.Split(tok, " ")[0])
			canContainColor := strings.Split(tok, " ")[1] + " " + strings.Split(tok, " ")[2]
			rls[color][canContainColor] = num
		}
	}

	return rls

}

// Given Bag of Color color, can it contain a bag of Color clr?
func canContain(color string, clr string, rls rules) bool {
	for key := range rls[clr] {
		if key == color {
			fmt.Printf("Can %s contain %s?", key, color)
			return true
		}
		return canContain(color, key, rls)
	}
	return false
}

// Count how many bags can contain a Bag of Color color
func countAtLeast(color string, rls rules) int {
	count := 0

	for clr := range rls {
		if canContain(color, clr, rls) {
			count++
		}
	}

	return count

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
	rls := parse(lines)

	countShinyGold := countAtLeast("shiny gold", rls)
	fmt.Println("Task 1: ", countShinyGold)

}
