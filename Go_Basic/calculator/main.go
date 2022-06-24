package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string, operator string) int {
	cleanEntries := strings.Split(input, operator)

	operator1 := parser(cleanEntries[0])
	operator2 := parser(cleanEntries[1])

	switch operator {
	case "+":
		operation := operator1 + operator2
		fmt.Println(operation)
		return operation
	case "-":
		operation := operator1 - operator2
		fmt.Println(operation)
		return operation
	case "*":
		operation := operator1 * operator2
		fmt.Println(operation)
		return operation
	case "/":
		operation := operator1 / operator2
		fmt.Println(operation)
		return operation
	default:
		fmt.Println(operator, "Not supported")
		return 0
	}
}

func parser(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	input := readInput()
	operator := readInput()

	c := calc{}
	fmt.Println(c.operate(input, operator))
}
