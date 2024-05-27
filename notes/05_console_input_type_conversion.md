bufio package

whatever we type, the Scanner takes it as string.
`scanner := bufio.NewScanner(os.Stdin)`
`scanner.Scan()`

To use int values we need to convert the input values. So, we need "strconv" package

`input, _ := strconv.ParseInt(scanner.Text(), 10, 64)`

`func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)`
ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.


```go
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
```