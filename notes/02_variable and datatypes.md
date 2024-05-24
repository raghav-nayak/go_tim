variable name cannot 
- have spaces
- starts with a number
- keywords

```go
var name string = "raghav"

var lastName string
lastName = "Nayak"
```

basic datatypes
- bool (true/false)
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr (`u` means unsigned -> only positive numbers)
- byte // alias for uint8
- rune // alias for int32 - represents a Unicode code point
- float32 float64
- complex64 complex128

machine dependent types
- uint (32/64 bits)
- int(same size as uint)
- uintptr (an unsigned integer to store the uninterpreted bits of a point value)

`var number uint8 = 260`
cannot use 260 (untyped int constant) as uint8 value in variable declaration (overflows)compiler[NumericOverflow](https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#NumericOverflow)

`var number uint8 = "260"`
cannot use "260" (untyped string constant) as uint8 value in variable declarationcompiler[IncompatibleAssign](https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#IncompatibleAssign)

