package main 


import "fmt"


func splitNumber(n int) []int {
	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	return digits
}




func sum(numbers int)int{
	digits := splitNumber(numbers)
	totalSum := 0;

	for _, number := range digits {
		totalSum += number
	}

	return totalSum
}



func isPrimeNumber(int) bool{

}




func main(){
  splitAndsum := sum(44444)
  fmt.Println(splitAndsum)


}