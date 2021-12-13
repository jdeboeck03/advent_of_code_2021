package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func remove_index(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func contains(s []int, value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func remove_white_spaces(input []string) []string {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == "" {
			input = remove_index(input, i)
		}
	}
	return input
}

func process_input(input []string) []int {
	var processed_lines = []int{}

	//split input
	input_split := strings.Split(input[0], ",")
	for _, number := range input_split {
		// width of bingo_board
		number_int, _ := strconv.Atoi(number)
		processed_lines = append(processed_lines, number_int)
	}
	return processed_lines
}

func calculate_median(coords []int) int {
	sort.Ints(coords)
	median_index := len(coords) / 2
	return coords[median_index]
}

func calculate_mean(coords []int) int {
	sort.Ints(coords)
	sum := 0.0
	for _, pos := range coords {
		sum += float64(pos)
	}
	mean := sum / float64(len(coords))
	mean_round := math.Round(mean)
	return int(mean_round)
}

func calculate_fuel(processed_input []int, median int) int {
	fuel := 0
	for _, pos := range processed_input {
		if pos > median {
			fuel += (pos - median)
		} else {
			fuel += (median - pos)
		}
	}
	return fuel
}

func calculate_fuel_2(processed_input []int, target int) int {
	fuel := 0
	increment := 0
	for _, pos := range processed_input {
		if pos > target {
			increment = pos - target
			for i := 1; i <= increment; i++ {
				fuel += i
			}
		} else {
			increment = target - pos
			for i := 1; i <= increment; i++ {
				fuel += i
			}
		}
	}
	return fuel
}

func calculate_lowest_fuel(processed_input []int, median int) int {
	fuel := calculate_fuel_2(processed_input, median)
	fmt.Println("FUEL")
	fmt.Println(fuel)
	passed_left_fuel := fuel
	passed_right_fuel := fuel
	left_fuel := 0
	right_fuel := 0
	searching_left := true
	searching_right := true
	left := median - 1
	right := median + 1
	for searching_left || searching_right {
		//fmt.Println("LEFT")
		//fmt.Println(left)
		left_fuel = calculate_fuel_2(processed_input, left)
		//fmt.Println(left_fuel)
		if left_fuel <= passed_left_fuel {
			passed_left_fuel = left_fuel
			left--
		} else {
			searching_left = false
		}
		right_fuel = calculate_fuel_2(processed_input, right)
		//fmt.Println("RIGHT")
		//fmt.Println(right)
		//fmt.Println(right_fuel)
		if right_fuel <= passed_right_fuel {
			passed_right_fuel = right_fuel
			right++
		} else {
			searching_right = false
		}
	}
	if passed_left_fuel < passed_right_fuel {
		fuel = passed_left_fuel
	} else {
		fuel = passed_right_fuel
	}
	return fuel
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	fmt.Println("Solution 1:")
	processed_input := process_input(input)
	median := calculate_median(processed_input)
	fuel := calculate_fuel(processed_input, median)
	fmt.Printf("Fuel: %d\n", fuel)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	fuel_2 := calculate_lowest_fuel(processed_input, median)
	fmt.Printf("Fuel: %d\n", fuel_2)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
