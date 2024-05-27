package main

import "fmt"

func main() {
	val := 10
	fmt.Printf("val = %T %v\n", val, val)

	// var str_val string = fmt.Sprintf()

	fmt.Printf("val = %t\n", val)
	

	fmt.Printf("%b %o %d %X\n", val, val, val, val)

	fmt.Printf("val = %f\n", 3.142)
	fmt.Printf("val = %e\n", 0.142)
	fmt.Printf("val = %g\n", 3.14212334567890123567)
	fmt.Printf("val = %9f\n", 3.142)
	fmt.Printf("val = %.2f\n", 3.14612334567890123567)
	fmt.Printf("val = %.f\n", 3.14212334567890123567)
	
	fmt.Printf("val = %07d\n", 45)

	fmt.Printf("val = %s\n", "hello")
	fmt.Printf("val = %9s\n", "hello")
	fmt.Printf("val = %q\n", "hello")
	fmt.Printf("val = %9s world\n", "hello")
	fmt.Printf("val = %-9s world\n", "hello")

	var out string = fmt.Sprintf("Number: %07d is cool\n", 45)
	fmt.Printf("out = %s\n", out)

	
}