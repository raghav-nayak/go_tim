`go_tim/01_introduction/tutorial.go`
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!!")
}
```

To run,
```sh
~/Developer/Programming/golang/go_tim/01_introduction ⌚ 22:59:12
$ go run tutorial.go
Hello world!!

~/Developer/Programming/golang/go_tim/01_introduction ⌚ 22:59:16
$ ll
total 8
drwxr-xr-x  4 raghavnayak  staff  128 May 24 22:59 .
-rw-r--r--  1 raghavnayak  staff   73 May 24 22:56 tutorial.go
drwxr-xr-x  3 raghavnayak  staff   96 May 24 22:55 .vscode
drwxr-xr-x  3 raghavnayak  staff   96 May 24 22:55 ..
```

We can also build and then run the executable
```shell
~/Developer/Programming/golang/go_tim/01_introduction ⌚ 23:00:37
$ go build tutorial.go

~/Developer/Programming/golang/go_tim/01_introduction ⌚ 23:00:47
$ ./tutorial
Hello world!!
```
