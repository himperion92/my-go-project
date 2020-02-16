package main

import (
	"fmt"
)

type info struct {
	result string
}

func main() {
	var introMssg string = "Hello from Go toolchain"
	fmt.Printf("Go reports: %+v\n", introMssg)
	var s = info{}
	sp := &s
	sp.result = "set a struct pointer value"
	fmt.Printf("Go reports: %+v\n", sp.result)
}
