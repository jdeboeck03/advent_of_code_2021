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

func horizontal_position(input []string) int {
	multiplication := 0
	horizontal_position := 0
	depth := 0

	re := regexp.MustCompile("[ ]")
	for _, line := range input {
		input_split := re.Split(line, -1)
		direction := input_split[0]
		direction_number := input_split[1]
		direction_number_int, err := strconv.Atoi(direction_number)
		if err != nil {
			log.Fatalf("Program failed: %s", err)
		}
		switch {
		case direction == "forward":
			horizontal_position += direction_number_int
		case direction == "down":
			depth += direction_number_int
		case direction == "up":
			depth -= direction_number_int
		}
	}
	multiplication = horizontal_position * depth
	return multiplication
}

func horizontal_position_2(input []string) int {
	multiplication := 0
	horizontal_position := 0
	depth := 0
	aim := 0

	re := regexp.MustCompile("[ ]")
	for _, line := range input {
		input_split := re.Split(line, -1)
		direction := input_split[0]
		direction_number := input_split[1]
		direction_number_int, err := strconv.Atoi(direction_number)
		if err != nil {
			log.Fatalf("Program failed: %s", err)
		}
		switch {
		case direction == "forward":
			horizontal_position += direction_number_int
			aim_inc := direction_number_int * aim
			depth += aim_inc
		case direction == "down":
			aim += direction_number_int
		case direction == "up":
			aim -= direction_number_int
		}
	}
	multiplication = horizontal_position * depth
	return multiplication
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	depth_cnt := horizontal_position(input)
	fmt.Println("Solution 1:")
	fmt.Printf("Horizontal Position: %d\n", depth_cnt)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	depth_cnt_2 := horizontal_position_2(input)
	fmt.Println("Solution 2:")
	fmt.Printf("Horizontal Position: %d\n", depth_cnt_2)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
