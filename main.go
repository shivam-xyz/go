package main

import (
	"fmt"
)

//Very Important, short handed will not work in Global Scope
// var age = 55
// acnummber := 987654 // This one

func main() {
	fmt.Println("Hello Bunny")
	fmt.Println("This is line 2 of bunny")

	// fmt.Print("This is line 3")
	// fmt.Print("This is line 4")

	//Variables
	var name = "Bunny"       // this one can work in both [in-function, global scope]
	firstname := "Chaudhary" // Short Handed only work inside the function
	isPresent := false

	var designation string = "Software Eng"
	var isSunday bool = true
	var age = -24326487436 // 8 bytes of RAM (int64)
	var ageNew int8 = 125

	// 	Go has five keywords/types of signed integers:

	// int 	Depends on platform:
	// 32 bits in 32 bit systems and
	// 64 bit in 64 bit systems 	-2147483648 to 2147483647 in 32 bit systems and
	// -9223372036854775808 to 9223372036854775807 in 64 bit systems

	// Type 	Size 			Range
	// int8 	8 bits/1 byte 	-128 to 127
	// int16 	16 bits/2 byte 	-32768 to 32767
	// int32 	32 bits/4 byte 	-2147483648 to 2147483647
	// int64 	64 bits/8 byte 	-9223372036854775808 to 9223372036854775807

	// above all is Signed Integer (+ , -)

	// ***********************************************************************************

	// below is unsigned integer ( only +)

	// 	Go has five keywords/types of unsigned integers:

	// uint 	Depends on platform:
	// 32 bits in 32 bit systems and
	// 64 bit in 64 bit systems 	0 to 4294967295 in 32 bit systems and
	// 0 to 18446744073709551615 in 64 bit systems

	// Type 	Size 	Range
	// uint8 	8 bits/1 byte 	0 to 255
	// uint16 	16 bits/2 byte 	0 to 65535
	// uint32 	32 bits/4 byte 	0 to 4294967295
	// uint64 	64 bits/8 byte 	0 to 18446744073709551615

	fmt.Println(name, firstname, isPresent, designation, isSunday, age, ageNew)

	//////////////////////////////////////////////////////////////////////////////////////////////

	// FLOATING POINTS :

	// score := 25.11     float32 / float64
	// var score float32 = 25.11

	// STATEMENT FLOWS :

	// if {} else if {} else if {} else {}

	// marks := 5

	// if marks > 60 {
	// 	fmt.Println("You're passed")
	// } else if marks < 60 && marks > 30 {
	// 	fmt.Println("You're passed with 2nd Grade")
	// } else if marks < 30 && marks > 15 {
	// 	fmt.Println("You're very poor")
	// } else {
	// 	fmt.Println("Please don't come to school again in future")
	// }

	if marks := 5; marks > 60 {
		fmt.Println("You're passed")
	} else if marks < 60 && marks > 30 {
		fmt.Println("You're passed with 2nd Grade")
	} else if marks < 30 && marks > 15 {
		fmt.Println("You're very poor")
	} else {
		fmt.Println("Please don't come to school again in future")
	}

	// SWITCH CASE :

	day := 6

	// switch {
	// 	case CONDITION_WILL_BE_HERE:
	// 		fmt.Println("Case 1 output will be here")
	// 	case CONDITION_WILL_BE_HERE:
	// 		fmt.Println("Case 2 output")
	// 	case CONDITION_THREE_WILL_BE_HERE, CONDITION_FOUR, CONDITION_FIVE, CONDITION_SO_ON:
	// 		fmt.Println("Outputs")
	// 	default:
	// 		fmt.Println("This is the default case")

	// }

	switch day {
	case 0:
		fmt.Println("This is Sunday, today no-office")
	case 1:
		fmt.Println("Today is Monday, Office-day")
	case 2, 3, 4, 5:
		fmt.Println("This all will be the labourie days!")
	case 6:
		fmt.Println("Today is Saturday, Foodie-day!")
	default:
		fmt.Println("You have mentioned the wrong day, correct properly")
	}

	score := 85

	switch {
	case score >= 90:
		fmt.Println("This is A Grade")
	case score >= 75:
		fmt.Println("This is Grade B")
	case score >= 50:
		fmt.Println("This is Grade C")
	default:
		fmt.Println("You failed!")
	}

	x := 10

	switch x % 2 {
	case 0:
		fmt.Println("This is Even number")
	case 1:
		fmt.Println("This is Odd Number")

	}

	var data interface{} = "Hello"

	switch v := data.(type) {
	case int:
		fmt.Println("This is integer", v)
	case string:
		fmt.Println("This is A String", v)
	case bool:
		fmt.Println("This is Boolean", v)
	default:
		fmt.Println("This is unknown data type", v)
	}

	///////////////////////////////////////////////////////////////////////////////////////////

	// LOOP :

	// for YOUR_CONDITIONS_HERE {

	// }

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	p := 0

	// while p < 5 {
	// 	fmt.Println(p);
	// 	p++
	// }

	for p < 5 {
		fmt.Println(p)
		p++
	}

	// INFINITE_LOOP :

	// for {
	// 	fmt.Println("Running in string")
	// }

	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	scores := [4]int{90, 60, 40, 20}

	for i, v := range scores {
		fmt.Println(i, v)
	}

	users := []string{"Bunny", "Arch", "Debian", "Ubuntu"}

	for _, v := range users {
		fmt.Println(v)
	}

	ages := map[string]int{"Sakib G": 62, "Bunny": 4, "Robin": 26, "Smart": 25}

	for name, age := range ages {
		fmt.Println(name, age)
	}

	studentName := "Bunny"

	for i, ch := range studentName {
		fmt.Println(i, ch, string(ch))
	}

	/////////////////////////////////////////////////////////////////////////////////

	// FUNCTIONS :

	// func sumTwoNumber (param type) returnType {
	// 	return value
	// }

	userName := "Bunny"

	message := greet(userName)

	// fmt.Println(greet(userName))

	fmt.Println(message)

	// fmt.Println(addition(2, 3))

	_, err := divide(25, 5)

	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(res)

	logMessage("This is a log")

	totalSum := total(1, 25, 98, 22, 87, 67)

	fmt.Println(totalSum)

}

func greet(name string) string {
	return "Hello " + name
}

func addition(a int, b int) int {
	return a + b
}

func multiply(a, b, c, d int) {}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

func logMessage(msg string) {
	fmt.Println(msg)
}

func total(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
