package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born : ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Errorf("Error while parsing %s err=%s", scanner.Text(), err.Error())
	}
	fmt.Printf("You will be %d years old at the end of 2024\n", 2024-input)
}