package main

import (
	"b2d/internal/converter"
	"b2d/internal/validation"
	"fmt"
)

func main() {
	fmt.Println("B2C: Binary to Decimal")
	fmt.Println("CLI program, convert binary to decimal ('q' to exit)")
	fmt.Println("Only accept binary number (`0` and `1`)")
	fmt.Println("\nExample:")
	fmt.Println("input: 011 --> output: 3")
	fmt.Println("input: 100 --> output: 4")
	fmt.Println()
	fmt.Println()

	input := ""
	value := 0

	for input == "" {
		fmt.Print("Input: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("An error occurred, try again...")
			input = ""
			continue
		}

		if input == "q" {
			break
		}

		if err = validation.Binary(input); err != nil {
			fmt.Println(err.Error())
			input = ""
			continue
		}

		value = converter.Binary(input)
		fmt.Println("Output: ", value)
		fmt.Println()
		input = ""
	}
}
