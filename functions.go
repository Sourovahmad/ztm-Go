package main

import "fmt"

func main() {

	// learn case
	number := 1
	switch number {
	case 1:
		fmt.Println("the number is 1")
		fallthrough
	case 2,3 :
		fmt.Println("the number is", number)
	default: 
		fmt.Println("the number is something else")
	}

}