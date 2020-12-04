package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkPassword1(char string, lower int, upper int, password string) bool {
	charCount := 0
	for _, s := range strings.Split(password, "") {
		if s == char {
			charCount++
		}

	}

	if charCount < lower || charCount > upper {
		return false
	}
	return true
}

func checkPassword2(char string, a int, b int, password string) bool {
	passwordSlice := strings.Split(password, "")

	ba := passwordSlice[a] == char
	bb := passwordSlice[b] == char

	if ba && bb {
		return false
	}
	if ba || bb {
		return true
	}

	return false

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	valid1 := 0
	valid2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()

		lineSplit := strings.Split(val, ": ")
		policy := lineSplit[0]
		password := lineSplit[1]

		policySplit := strings.Split(policy, " ")
		limits := policySplit[0]
		char := policySplit[1]

		limitsSplit := strings.Split(limits, "-")
		lower, _ := strconv.Atoi(limitsSplit[0])
		upper, _ := strconv.Atoi(limitsSplit[1])

		if checkPassword1(char, lower, upper, password) {
			valid1++
		}
		if checkPassword2(char, lower-1, upper-1, password) {
			valid2++
		}

	}
	fmt.Println("Valid Passwords Policy 1: ", valid1)
	fmt.Println("Valid Passwords Policy 2: ", valid2)
}
