package main

import (
	"commandLineArgs"
	"fmt"
)

func main() {
	test := commandLineArgs.New()
	fmt.Print(test.HasValue("test"))
	val, err := test.GetValue("more")
	if err != nil {
		fmt.Print("no more")
	} else {
		fmt.Print(val)
	}
}
