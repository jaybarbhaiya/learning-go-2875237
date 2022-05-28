package main

import "fmt"

const aConst int = 88

func main() {
	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable type of aString is %T\n", aString)

	var anInteger int = 45
	fmt.Println(anInteger);

	var defaultInteger int;
	fmt.Println(defaultInteger)

	var anotherString = "Jay"
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 56
	fmt.Printf("The variable's type is %T\n", anotherInt)

	myString := "My string"
	fmt.Println(myString);
	fmt.Printf("The variable's type is %T\n", myString)

	fmt.Println(aConst);
	fmt.Printf("The variable's type is %T\n", aConst)
}