package main

import (
	"fmt"
	"strconv"
)

func main() {
	number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string
	if number1 < number2 {
		resultMessage = strconv.Itoa(number2) + " is bigger than " + strconv.Itoa(number1) + "." + "\nThe difference is " + strconv.Itoa(number2-number1) + "."
		fmt.Println(resultMessage)
	}
	if number1 == number2 {
		resultMessage = "Both numbers are equal. (" + strconv.Itoa(number1) + ")"
		fmt.Println(resultMessage)
	}
	if number1 > number2 {
		resultMessage = strconv.Itoa(number1) + " is bigger than " + strconv.Itoa(number2) + "." + "\nThe difference is " + strconv.Itoa(number1-number2) + "."
		fmt.Println(resultMessage)
	}

}
