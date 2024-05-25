based on the value assigned, the type of a variable is assumed.

```go
	var number = 2
	fmt.Printf("number: %T\n", number) // number: int

	var number = 200.98
	fmt.Printf("number: %T\n", number) // number: float64
```


```go
	num := 300
	fmt.Printf("num: %T\n", num) // num: int
	num = "str" // error: cannot use "str" (untyped string constant) as int value in assignment
	
```

`num := 5` is equivalent to `var num int = 5`

#### default values
```go
	var myNum uint64
	var myBool bool
	fmt.Println("myNum: ", myNum) // myNum:  0
	fmt.Println("myBool: ", myBool) // myBool:  false
```

A default value is assigned to each variable we define based on the type of the variable.