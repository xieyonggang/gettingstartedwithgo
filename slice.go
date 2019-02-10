package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	sli := make([]int, 3)
	//fmt.Println(len(sli))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "X" {
			break
		}

		inputInt, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("please input a integer only.")
			continue
		}

		if len(sli) < 3 {
			for i := range sli {
				if i == 0 {

				}
			}
		} else {
			sli = append(sli, inputInt)
		}

		sort.Ints(sli)

		fmt.Println(sli)
	}

}
