The type should be same 

```go
package main

import "fmt"

func main() {
	var num1 float64 = 8
	var num2 int = 3
	answer := num1 / num2 // invalid operation: num1 / num2 (mismatched types float64 and int)
	fmt.Printf("%d", answer)
}
```

solution is to convert it
```go
answer := num1 / float64(num2)
fmt.Printf("%f", answer) // 2.666667
```

when you divide int by int, you will get int.
```go
	num3 := 9
	num4 := 4
	answer2 := num3 / num4
	fmt.Printf("%d\n", answer2) // 2

	answer2 := float64(num3 / num4) 
	fmt.Printf("%f\n", answer2) // 2.000000
	fmt.Printf("%g\n", answer2) // 2
```

modulus
```go
	var num5 int = 9
	var num6 int = 4
	answer3 := num5 % num6
	fmt.Printf("%d\n", answer3) // 1
```

panic
```go
	var num5 int = 4
	var num6 int = 0
	answer3 := num5 / num6
	fmt.Printf("%d\n", answer3)

panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
	golang/go_tim/06_arithmetic_operations/arithmetic_operations.go:18 +0xbd
exit status 2

```

math
```go
import "math"
...
fmt.Printf("%f\n", math.Ceil(3.4))

```



### Precedence 
B -> Brackets
E -> Exponents
D -> Division
M -> Multiplication
A -> Addition
S -> Subtraction
