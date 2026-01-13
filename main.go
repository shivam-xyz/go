package main

import "fmt"

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

}
