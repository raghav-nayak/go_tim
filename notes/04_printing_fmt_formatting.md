
### General
%v -> value in default format
%T -> type
%% -> literal % (this to print %)

### Boolean
%t -> true/false

### Integer
%b -> base2
%o -> base8
%d -> base10
%x -> base16

### Floating
%e -> scientific notation
%f / %F -> decimal no exponent
%g -> for large exponents

### Strings
%s -> default string
%q -> double quoted string

### Width and precision
%f -> default width and default precision
%9f -> width 9 width and default precision
%.2f -> default width and precision 2
%9.2f -> width 9 and precision 2
%9.f -> width 9 and precision 0

### Padding
%09d -> pads digit to length 9 with proceeding 0's
%-4d -> pad with spaces; width 4, left justified

### Methods
Sprintf() -> format without printing
Printf() -> format with printing


```go
package main 
import "fmt"

func main() {
	val := 10
	fmt.Printf("val = %T %v\n", val, val) // val = int 10

	fmt.Printf("val = %t\n", val) // val = %!t(int=10) 
								  // fmt.Printf format %t has arg val of wrong type int

	fmt.Printf("%b %o %d %x\n", val, val, val, val) // 1010 12 10 a
	fmt.Printf("%b %o %d %X\n", val, val, val, val) // 1010 12 10 A


	fmt.Printf("val = %e\n", 3456.2) // val = 3.456200e+03
	fmt.Printf("val = %e\n", 0.142) // val = 1.420000e-01

	fmt.Printf("val = %f\n", 3.142123345) // val = 3.142123
	fmt.Printf("val = %9f\n", 3.142) // val =  3.142000
	fmt.Printf("val = %.2f\n", 3.14612334567890123567) // val = 3.15
	fmt.Printf("val = %.f\n", 3.14212334567890123567) // val = 3
	
	
	fmt.Printf("val = %g\n", 3.1412334567890123567) // val = 3.1421233456789013
	

	fmt.Printf("val = %s\n", "hello") // val = hello
	fmt.Printf("val = %q\n", "hello") // val = "hello"
	fmt.Printf("val = %9s\n", "hello") // val =     hello
	fmt.Printf("val = %9s world\n", "hello") // val =     hello world -> padding done on left
	fmt.Printf("val = %-9s world\n", "hello") // val = hello     world -> padding done on right

	var out string = fmt.Sprintf("Number: %07d is cool\n", 45)
	fmt.Printf("out = %s\n", out) // out = Number: 0000045 is cool
}
```