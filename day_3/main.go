package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

func RemoveIndex(s []string, index int) []string {
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

func power_consumption(input []string) float64 {
	common_bits := make([]int, len(input[0]))
	fmt.Println(common_bits)
	gamma := 0.0
	epsilon := 0.0
	for _, line := range input {
		line_split := strings.Split(line, "")
		for i, bit := range line_split {
			if bit == "0" {
				common_bits[i] -= 1
			} else {
				common_bits[i] += 1
			}
		}
	}
	fmt.Println(common_bits)
	for i, common_bit := range common_bits {
		exp := float64(len(common_bits) - i - 1)
		fmt.Println(exp)
		if common_bit > 0 {
			gamma += math.Pow(2, exp)
		} else {
			epsilon += math.Pow(2, exp)
		}
	}
	fmt.Println(gamma)
	fmt.Println(epsilon)
	power_consumption := gamma * epsilon
	return power_consumption
}

func check_common_bit(input []string, index int) (string, string) {
	//Find the most common bit for the corresponding index
	common_bit := 0
	common_bit_result := "0"
	least_common_bit_result := "0"
	for _, line := range input {
		line_split := strings.Split(line, "")
		if line_split[index] == "0" {
			common_bit -= 1
		} else {
			common_bit += 1
		}
	}
	if common_bit >= 0 {
		common_bit_result = "1"
		least_common_bit_result = "0"
	} else {
		common_bit_result = "0"
		least_common_bit_result = "1"
	}
	return common_bit_result, least_common_bit_result
}

func check_oxygen_rating(input []string) float64 {
	oxygen_rating := 0.0
	bit_range := len(input[0])
	//fmt.Println(input)
	for i := 0; i < bit_range; i++ {
		most_common_bit_result, _ := check_common_bit(input, i)
		//fmt.Println(most_common_bit_result)
		for j := len(input) - 1; j >= 0; j-- {
			line_split := strings.Split(input[j], "")
			if line_split[i] != most_common_bit_result {
				//Remove value and substracting i
				input = RemoveIndex(input, j)
			}
		}
	}
	//fmt.Println(input)
	input_split := strings.Split(input[0], "")
	//fmt.Println(input_split)
	for i, bit := range input_split {
		exp := float64(len(input_split) - i - 1)
		//fmt.Println(exp)
		if bit == "1" {
			oxygen_rating += math.Pow(2, exp)
		}
	}
	return oxygen_rating
}

func check_co2_rating(input []string) float64 {
	co2_rating := 0.0
	//fmt.Println(input)
	bit_range := len(input[0])
	for i := 0; i < bit_range; i++ {
		_, least_common_bit_result := check_common_bit(input, i)
		//fmt.Println(least_common_bit_result)
		for j := len(input) - 1; j >= 0; j-- {
			line_split := strings.Split(input[j], "")
			if line_split[i] != least_common_bit_result && len(input) != 1 {
				//Remove value and substracting i
				input = RemoveIndex(input, j)
			}
		}
	}
	//fmt.Println(input)
	input_split := strings.Split(input[0], "")
	//fmt.Println(input_split)
	for i, bit := range input_split {
		exp := float64(len(input_split) - i - 1)
		//fmt.Println(exp)
		if bit == "1" {
			co2_rating += math.Pow(2, exp)
		}
	}
	return co2_rating
}

func main() {
	input, err := loadInput()
	input2 := []string{}
	for i := 0; i < len(input); i++ {
		input2 = append(input2, input[i])
	}
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	power_consumption := power_consumption(input)
	fmt.Println("Solution 1:")
	fmt.Printf("Horizontal Position: %f\n", power_consumption)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	oxygen_rating := check_oxygen_rating(input)
	fmt.Printf("Oxygen Rating: %f\n", oxygen_rating)
	co2_rating := check_co2_rating(input2)
	fmt.Printf("co2 Rating: %f\n", co2_rating)
	multiplication := oxygen_rating * co2_rating
	fmt.Printf("Life support rating: %f", multiplication)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
