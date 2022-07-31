package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	age      int16
	employed bool
	salary   float32
}

func main() {
	bob := person{}
	fmt.Println(unsafe.Sizeof(bob))
}
