package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passport Type
type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      int64
}

func parsePassport(pass string) (Passport, bool, bool) {

	var passport Passport
	pass = strings.Replace(pass, "\n", " ", -1)
	passFields := strings.Split(pass, " ")
	for _, field := range passFields {

		keyVal := strings.Split(field, ":")
		if len(keyVal) != 2 {
			continue
		}
		key := keyVal[0]
		val := keyVal[1]
		switch key {
		case "byr":
			passport.BirthYear, _ = strconv.Atoi(val)
		case "iyr":
			passport.IssueYear, _ = strconv.Atoi(val)
		case "eyr":
			passport.ExpirationYear, _ = strconv.Atoi(val)
		case "hgt":
			passport.Height = val
		case "hcl":
			passport.HairColor = val
		case "ecl":
			passport.EyeColor = val
		case "pid":
			passport.PassportID = val
		case "cid":
			passport.CountryID, _ = strconv.ParseInt(val, 10, 64)
		}
	}

	complete := true
	correct := true

	if passport.BirthYear != 0 {
		if passport.BirthYear < 1920 || passport.BirthYear > 2002 {
			correct = false
		}
	} else {
		return passport, false, false
	}

	if passport.IssueYear != 0 {
		if passport.IssueYear < 2010 || passport.IssueYear > 2020 {
			correct = false
		}
	} else {
		return passport, false, false
	}

	if passport.ExpirationYear != 0 {
		if passport.ExpirationYear < 2020 || passport.ExpirationYear > 2030 {
			correct = false
		}
	} else {
		return passport, false, false
	}

	if passport.Height != "" {
		if strings.Contains(passport.Height, "cm") {
			height, _ := strconv.Atoi(strings.Split(passport.Height, "cm")[0])
			if height < 150 || height > 193 {
				correct = false
			}
		} else {
			if strings.Contains(passport.Height, "in") {
				height, _ := strconv.Atoi(strings.Split(passport.Height, "in")[0])
				if height < 59 || height > 76 {
					correct = false
				}
			}
		}
	} else {
		return passport, false, false
	}

	if passport.HairColor != "" {
		matched, err := regexp.MatchString(`^#([0-9a-z]){6}$`, passport.HairColor)
		if !matched || err != nil {
			correct = false
		}
	} else {
		return passport, false, false
	}

	if passport.EyeColor != "" {
		matched, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, passport.EyeColor)
		if !matched || err != nil {
			correct = false
		}
	} else {
		return passport, false, false
	}

	if passport.PassportID != "" {
		matched, err := regexp.MatchString(`^([0-9]){9}$`, passport.PassportID)
		if !matched || err != nil {
			correct = false
		}
	} else {
		return passport, false, false
	}

	return passport, complete, correct
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	text, err := ioutil.ReadAll(r)

	completes := 0
	corrects := 0
	passports := strings.Split(string(text), "\n\n")
	for _, pass := range passports {
		_, complete, correct := parsePassport(pass)
		if complete {
			completes++
		}
		if correct {
			corrects++
		}
	}
	fmt.Println("Complete Passports: ", completes)
	fmt.Println("Correct Passports: ", corrects)
}
