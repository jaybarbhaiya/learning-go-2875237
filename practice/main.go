package main

import (
	"fmt"
)

func main() {
	anInt := 33
	var p = &anInt
	fmt.Println("Memory address - pointer p:", p)
	fmt.Println("Value - pointer p:", *p)

	aFloat := 33.33
	pointer1 := &aFloat
	fmt.Println("aFloat", *pointer1)

	*pointer1 = *pointer1 / 11
	fmt.Println("aFloat pointer", *pointer1)
	fmt.Println("aFloat value changed via changing pointer1", aFloat)
}
