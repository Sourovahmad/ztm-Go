package main

import (
	"fmt"
	"structs"
)


func main(){

	type Sample struct{
		name string 
		email string
		phone int 
		amount float64
	}

	data := Sample{
		name: "sourov",
		email: "sourov@gmail.com",
		phone: 1244587,
		amount: 125.215,
	}


	// create and assign together
	p1 := struct {
		Name   string
		Age    int
		Gender string
	}{
		Name:   "John",
		Age:    30,
		Gender: "Male",
	
	
	}


	
	fmt.Println(data)
}