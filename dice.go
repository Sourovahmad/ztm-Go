package main 
import "math/rand"
import "time"
import "fmt"


func roll(sides int) int{
	return rand.Intn(sides) + 1;
}

func main(){	

 	rand.New(rand.NewSource(time.Now().UnixNano()))
	dice, sides := 2, 6
	rolls := 2


	for r := 1; r <= rolls; r++ {

		sum := 0
		for i := 1; i <= dice; i++ {

			randomSide := roll(sides)
			sum += randomSide

			fmt.Println("Dice number", i, "Dice result", randomSide)
		}	



		fmt.Println("total sum is", sum)
	} 








}