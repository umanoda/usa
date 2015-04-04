# About for Usa

Go lang liblary, like shell's pipe stream.

Usa is meaning in rabbit :rabbit: in japanese.

# Usage

```go
package main

import (
	"github.com/umanoda/usa"
	"fmt"
)

func main() {
	u := usa.Filter(seq15).Pipe(fizzbuzz).Pipe(stdout)
	for _ = range u.Run() {} // Wait until completed pipe line.
	//=> 1
	//   2
	//   Fizz
	//   4
	//   Buzz
	//   :
	//   14
	//   FizzBuzz

	// Same
	u.Run().Wait()

	// Result to array
	u2 := usa.Filter(seq15).Pipe(fizzbuzz).Run().ToArray()
	fmt.Println(u2)
	//=> [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz]
}

func seq15(_, out usa.Ch) {
	for n := 0; n < 15; n++ {
		out <- n + 1
	}
	close(out)
}

func fizzbuzz(val int) string {
	var str string
	if val%15 == 0 {
		str = "FizzBuzz"
	} else if val%3 == 0 {
		str = "Fizz"
	} else if val%5 == 0 {
		str = "Buzz"
	} else {
		str = fmt.Sprint(val)
	}
	return str
}

func stdout(val string) bool {
	fmt.Println(val)
	return true // Dummy return value
}
```
