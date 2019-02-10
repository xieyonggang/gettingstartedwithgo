package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Please input a floating number, type 'q' or 'quit' to exit the program")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//fmt.Print(scanner.Text() + "\n")
		if scanner.Text() == "quit" || scanner.Text() == "q" {
			break
		}

		input, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			fmt.Println("Wrong input format, Please input number only.")
			continue
		} else {
			fmt.Println("The truncated number in digit is " + strconv.Itoa(int(input)))
		}

	}

}
