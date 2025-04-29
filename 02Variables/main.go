package main

import "fmt"

func main() {
	var username string = "nabid"
	var age int = 25

	fmt.Println("hey ", username)
	fmt.Println("you are ", age, " years old")
	fmt.Println("type of username is ", fmt.Sprintf("%T", username))
	fmt.Println("type of age is ", fmt.Sprintf("%T", age))

	// boolean

	var isTwentyFiveYearsOld bool = true

	if age == 25 {
		fmt.Println(username+" isTwenty Five Years Old is ", isTwentyFiveYearsOld)
	} else {
		fmt.Println("isTwentyFiveYearsOld is ", !isTwentyFiveYearsOld)
	}

	// integer

	var intNumber1 int = 10
	var intNumber2 int = 20

	var intNumber3 int = intNumber1 + intNumber2

	fmt.Println("Sum is : ", intNumber3)

	// implicit type

	var name = "nabid"

	fmt.Println("name is : ", name)
}
