package main

import "fmt"

func main() {

	//exercise
	for i := 0; i < 50; i++ {
		if i%3 == 0 && i %5 == 0 {
			fmt.Println("fizz buzz")
		}else if i%3 == 0 {
			fmt.Println("fizz")
		}else if i%5 == 0 {
			fmt.Println("buzz")
		}else{
			fmt.Println(i)

		}
	}
}