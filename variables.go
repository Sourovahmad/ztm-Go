package main

import "fmt"

func main() {


		// basic way of declaration 
		var myFavouriteColour =  "green"
		fmt.Println("my favourite colour is", myFavouriteColour)

		// create and declaration together
		myAgeinYears := 22
		fmt.Printf("I am %v years old\n", myAgeinYears)


		// block declartion 
		var (
			myHouseNumber = "habiganj sylhet bangladesh"
			myCityName = "sylhet habiganj"
		)
		fmt.Println(myHouseNumber,myCityName)


		//single declaration 
		var myFavouriteFood string
		myFavouriteFood = "mango"
		fmt.Println("my favourite food is",myFavouriteFood)


		// dual declaration 
		first, second := 10,20
		third, second := 30, 35
		fmt.Println(first,second,third)

}
