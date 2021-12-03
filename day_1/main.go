package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

// loads input of file into a slice of integer variables
func loadInput() ([]string, error) {
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		input = append(input, line)
	}
	return input, nil
}

// process fields, so find the valid passports
func processFields1(input []string) int {
	field_checker := make(map[string]bool)
	req_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "ecl", "pid", "cid"}
	// init field checker
	for _, req_field := range req_fields {
		field_checker[req_field] = false
	}
	var valid_passports = 0
	var valid_passport = true
	for _, line := range input {
		// Reached a new passport so req_fields need to be resetted
		if line == "" {
			//fmt.Println("New Passport!")
			//fmt.Println(field_checker)
			for _, req_field := range req_fields {
				if !field_checker[req_field] && req_field != "cid" {
					// one of the fields on passport is not filled in
					//fmt.Println("Oei Oei")
					valid_passport = false
				}
			}
			// reset req fields
			for _, req_field := range req_fields {
				field_checker[req_field] = false
				//fmt.Println(field_checker)
			}
			// Check if passport is valid
			if valid_passport {
				valid_passports += 1
			}
			valid_passport = true
		} else {
			// Now we need to process the entire field (use a regex and substring)
			re := regexp.MustCompile("[: ]")
			line_split := re.Split(line, -1)
			//fmt.Println(line_split)
			for i, field := range line_split {
				//fmt.Println(field)
				// Only check even indexes
				if i%2 == 0 {
					//fmt.Println(field)
					field_checker[field] = true
				}
			}
		}

	}
	// Check final passport
	//fmt.Println("New Passport!")
	//fmt.Println(field_checker)
	for _, req_field := range req_fields {
		if !field_checker[req_field] && req_field != "cid" {
			// one of the fields on passport is not filled in
			//fmt.Println("Oei Oei")
			valid_passport = false
		}
	}
	// reset req fields
	for _, req_field := range req_fields {
		field_checker[req_field] = false
		//fmt.Println(field_checker)
	}
	// Check if passport is valid
	if valid_passport {
		valid_passports += 1
	}
	valid_passport = true
	return valid_passports
}

// process fields, so find the valid passports
func processFields2(input []string) int {
	field_checker := make(map[string]bool)
	req_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "ecl", "pid", "cid"}
	// init field checker
	for _, req_field := range req_fields {
		field_checker[req_field] = false
	}
	var valid_passports = 0
	var valid_passport = true
	for _, line := range input {
		// Reached a new passport so req_fields need to be resetted
		if line == "" {
			//fmt.Println("New Passport!")
			//fmt.Println(field_checker)
			for _, req_field := range req_fields {
				if !field_checker[req_field] && req_field != "cid" {
					// one of the fields on passport is not filled in
					//fmt.Println("Oei Oei")
					valid_passport = false
				}
			}
			// reset req fields
			for _, req_field := range req_fields {
				field_checker[req_field] = false
				//fmt.Println(field_checker)
			}
			// Check if passport is valid
			if valid_passport {
				valid_passports += 1
			}
			valid_passport = true
		} else {
			// Now we need to process the entire field (use a regex and substring)
			re := regexp.MustCompile("[: ]")
			line_split := re.Split(line, -1)
			//fmt.Println(line_split)
			for i, field := range line_split {
				//fmt.Println(field)
				// Only check even indexes
				if i%2 == 0 {
					//fmt.Println(field)
					// Need to add a switch case with checks for every passport field
					//field_checker[field] = true
					value := line_split[i+1]
					switch {
					case field == "byr":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^[0-9]{4}$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						if valid_field {
							// Parse value to an integer
							value_int, err := strconv.Atoi(value)
							if err != nil {
								log.Fatalf("Program failed: %s", err)
							}
							if 1920 <= value_int && value_int <= 2002 {
								//fmt.Println("Correct!")
								field_checker[field] = true
							}
						}
					case field == "iyr":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^[0-9]{4}$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						if valid_field {
							// Parse value to an integer
							value_int, err := strconv.Atoi(value)
							if err != nil {
								log.Fatalf("Program failed: %s", err)
							}
							if 2010 <= value_int && value_int <= 2020 {
								//fmt.Println("Correct!")
								field_checker[field] = true
							}
						}
					case field == "eyr":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^[0-9]{4}$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						if valid_field {
							// Parse value to an integer
							value_int, err := strconv.Atoi(value)
							if err != nil {
								log.Fatalf("Program failed: %s", err)
							}
							if 2020 <= value_int && value_int <= 2030 {
								//fmt.Println("Correct!")
								field_checker[field] = true
							}
						}
					case field == "hgt":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^([0-9]*)cm|in$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						if valid_field {
							// Parse value to an integer
							re := regexp.MustCompile("[0-9]*")
							extracted_value := re.FindString(value)
							re = regexp.MustCompile("cm|in")
							hgt_type := re.FindString(value)
							value_int, err := strconv.Atoi(extracted_value)
							if err != nil {
								log.Fatalf("Program failed: %s", err)
							}
							if hgt_type == "cm" && 150 <= value_int && value_int <= 193 {
								//fmt.Println("Correct!")
								//fmt.Println(extracted_value)
								field_checker[field] = true
							}
							if hgt_type == "in" && 59 <= value_int && value_int <= 76 {
								//fmt.Println("Correct!")
								//fmt.Println(extracted_value)
								field_checker[field] = true
							}
						}
					case field == "hcl":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^#([0-9]|[a-f]){6}$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						field_checker[field] = valid_field
					case field == "ecl":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						field_checker[field] = valid_field
					case field == "pid":
						// First match exactly on 4 digits
						re := regexp.MustCompile("^[0-9]{9}$")
						valid_field := re.MatchString(value)
						//fmt.Println(valid_field)
						field_checker[field] = valid_field
					}
				}
			}
		}

	}
	// Check final passport
	//fmt.Println("New Passport!")
	//fmt.Println(field_checker)
	for _, req_field := range req_fields {
		if !field_checker[req_field] && req_field != "cid" {
			// one of the fields on passport is not filled in
			valid_passport = false
		}
	}
	// reset req fields
	for _, req_field := range req_fields {
		field_checker[req_field] = false
		//fmt.Println(field_checker)
	}
	// Check if passport is valid
	if valid_passport {
		valid_passports += 1
	}
	valid_passport = true
	return valid_passports
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	valid_passports_1 := processFields1(input)
	fmt.Println("Solution 1:")
	fmt.Printf("Number of valid passports: %d\n", valid_passports_1)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	valid_passports_2 := processFields2(input)
	fmt.Println("Solution 2:")
	fmt.Printf("Number of valid passports: %d\n", valid_passports_2)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
