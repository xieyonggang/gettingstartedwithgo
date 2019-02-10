package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Printf("Please input a floating number : ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseFloat(scanner.Text(), 64)

	if err != nil {
		fmt.Println("Wrong input format, Please input number only.")
	} else {
		fmt.Println("The truncated number in digit is : " + strconv.Itoa(int(input)))
	}

}
