package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()

	operator := "*"
	values := strings.Split(operation, operator)

	operator1, err1 := strconv.Atoi(values[0])

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(operator1)
	}

	operator2, err2 := strconv.Atoi(values[1])

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(operator2)
	}

	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
	case "-":
		fmt.Println(operator1 - operator2)
	case "*":
		fmt.Println(operator1 * operator2)
	case "/":
		fmt.Println(operator1 / operator2)
	default:
		fmt.Println(operator, "Not supported")
	}
}
