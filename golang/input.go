package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func greeting(input string, name string) {
	name = strings.TrimSpace(name)
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	if number > 0 && number < 10 && name != "" {
		for i := 0; i < number; i++ {
			fmt.Printf("Nice to meet you %s \n", name)
		}
	} else {
		fmt.Println("Number must be greater than 0 and less than 10 or name is empty")
	}
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(0)
	}
	number := arguments[1]
	fmt.Println("Your name please? Press the enter key when done.")
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	greeting(number, input)
}
