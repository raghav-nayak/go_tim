package main

import "fmt"

func main() {
	var number = 200.98
	fmt.Printf("number: %T\n", number)

	num := 300
	fmt.Printf("num: %T\n", num)
	//num = "str"

	var myNum uint64
	var myBool bool
	fmt.Println("myNum: ", myNum)
	fmt.Println("myBool: ", myBool)
}