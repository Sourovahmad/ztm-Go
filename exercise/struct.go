package main

import "fmt"

type Rectangle struct {
	widtht float64
	height float64
}

func mesaureArea(r Rectangle) float64 {
	return r.height * r.widtht

}

func main() {

	rectangleData := Rectangle{
		height: 20,
		widtht: 40,
	}

	rectalgeArea := mesaureArea(rectangleData)
	fmt.Println(rectalgeArea)
	


}

