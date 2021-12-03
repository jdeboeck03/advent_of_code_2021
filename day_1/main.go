package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func depth_measurement(input []string) int {
	depth_cnt := 0
	previous_measurement := input[0]
	previous_measurement_int, err := strconv.Atoi(previous_measurement)
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	for _, current_measurement := range input {
		current_measurement_int, err := strconv.Atoi(current_measurement)
		if err != nil {
			log.Fatalf("Program failed: %s", err)
		}
		if current_measurement_int > previous_measurement_int {
			depth_cnt += 1
		}
		previous_measurement_int = current_measurement_int
	}
	return depth_cnt
}

func depth_measurement_2(input []string) int {
	depth_cnt := 0
	previous_sum := 0
	current_sum := 0
	fmt.Println(len(input))
	for i, _ := range input {
		if i < len(input)-2 {
			current_measurement_int_1, err_1 := strconv.Atoi(input[i])
			current_measurement_int_2, err_2 := strconv.Atoi(input[i+1])
			current_measurement_int_3, err_3 := strconv.Atoi(input[i+2])
			if err_1 != nil || err_2 != nil || err_3 != nil {
				log.Fatalf("Program failed!!")
			}
			current_sum = current_measurement_int_1 + current_measurement_int_2 + current_measurement_int_3
			if current_sum > previous_sum && i != 0 {
				depth_cnt += 1
			}
			previous_sum = current_sum
		}
	}

	return depth_cnt
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	depth_cnt := depth_measurement(input)
	fmt.Println("Solution 1:")
	fmt.Printf("Number of depth measurements: %d\n", depth_cnt)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	depth_cnt_2 := depth_measurement_2(input)
	fmt.Println("Solution 2:")
	fmt.Printf("Number of depth measurements: %d\n", depth_cnt_2)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
