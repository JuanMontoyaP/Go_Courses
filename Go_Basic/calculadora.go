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

	values := strings.Split(operation, "+")

	operator1, err1 := strconv.Atoi(values[0])

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(operator1)
	}

	operator2, _ := strconv.Atoi(values[1])

	fmt.Println(operator1 + operator2)
}
