package main 

import "fmt"

func main(){

	myArray := []string{"hello","Hi","sourov","arif"}

	for i, item := range myArray{
		fmt. Println(i,"--", item)
			for _, innerItem := range item{
				fmt.Printf(" %q\n" ,innerItem)
			}
	}

	type myVariable struct{
		Name string
		Email string
	},{
		
	}

}