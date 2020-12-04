package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	puzzle1()
	puzzle2()
}

func puzzle1() {
	fmt.Println("Day 4, puzzle 1.\n\n Count the number of valid passports - those that have all required fields." +
		"Treat cid as optional. In your batch file, how many passports are valid?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	passports := strings.Split(string(content), "\n\n")

	/*
		For each passport, increment the counter if it's valid.
	*/
	counter := 0
	for _, passport := range passports {
		passport = strings.ReplaceAll(passport, "\n", " ")

		passportFields := make(map[string]bool)

		for _, field := range strings.Fields(passport) {
			passportFields[field[0:3]] = true
		}

		if len(passportFields) >= 7 &&
			passportFields["byr"] &&
			passportFields["iyr"] &&
			passportFields["eyr"] &&
			passportFields["hgt"] &&
			passportFields["hcl"] &&
			passportFields["ecl"] &&
			passportFields["pid"] {
			counter++

		}
	}
	fmt.Printf("There are %v valid passports.\n", counter)

}

func puzzle2() {
	fmt.Println("Day 4, puzzle 2.\n\n Count the number of valid passports - those that have all required fields and valid values." +
		"Continue to treat cid as optional. In your batch file, how many passports are valid?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	passports := strings.Split(string(content), "\n\n")

	/*
		For each passport, increment the counter if it's valid.
	*/
	counter := 0
	for _, passport := range passports {
		passport = strings.ReplaceAll(passport, "\n", " ")

		passportFields := make(map[string]string)

		for _, field := range strings.Fields(passport) {
			passportFields[field[0:3]] = field[4:]
		}

		if len(passportFields) >= 7 {

			passportFieldsConstraints := make(map[string]bool)

			// I can ignore error because 0 is not valid anyway

			// Birth Year
			byr, _ := strconv.Atoi(passportFields["byr"])
			if byr >= 1920 && byr <= 2002 {
				passportFieldsConstraints["byr"] = true
			}

			//Issue Year
			iyr, _ := strconv.Atoi(passportFields["iyr"])
			if iyr >= 2010 && iyr <= 2020 {
				passportFieldsConstraints["iyr"] = true
			}

			//Expiration Year
			eyr, _ := strconv.Atoi(passportFields["eyr"])
			if eyr >= 2020 && eyr <= 2030 {
				passportFieldsConstraints["eyr"] = true
			}

			// Height
			if strings.Contains(passportFields["hgt"], "in") {
				hgt, _ := strconv.Atoi(passportFields["hgt"][:len(passportFields["hgt"])-2])
				if hgt >= 59 && hgt <= 76 {
					passportFieldsConstraints["hgt"] = true
				}
			} else if strings.Contains(passportFields["hgt"], "cm") {
				hgt, _ := strconv.Atoi(passportFields["hgt"][:len(passportFields["hgt"])-2])
				if hgt >= 150 && hgt <= 193 {
					passportFieldsConstraints["hgt"] = true
				}
			}

			//Hair color
			validHairColor := regexp.MustCompile(`#[a-z0-9]{6}`)
			if validHairColor.MatchString(passportFields["hcl"]) {
				passportFieldsConstraints["hcl"] = true
			}

			// Eye color
			if passportFields["ecl"] == "amb" ||
				passportFields["ecl"] == "blu" ||
				passportFields["ecl"] == "brn" ||
				passportFields["ecl"] == "gry" ||
				passportFields["ecl"] == "grn" ||
				passportFields["ecl"] == "hzl" ||
				passportFields["ecl"] == "oth" {
				passportFieldsConstraints["ecl"] = true
			}

			//Passport ID
			validPassportID := regexp.MustCompile(`[0-9]{9}`)
			if validPassportID.MatchString(passportFields["pid"]) {
				passportFieldsConstraints["pid"] = true
			}

			if passportFieldsConstraints["byr"] &&
				passportFieldsConstraints["iyr"] &&
				passportFieldsConstraints["eyr"] &&
				passportFieldsConstraints["hgt"] &&
				passportFieldsConstraints["hcl"] &&
				passportFieldsConstraints["ecl"] &&
				passportFieldsConstraints["pid"] {
				counter++

			}
		}
	}
	fmt.Printf("There are %v valid passports.\n", counter)
}
