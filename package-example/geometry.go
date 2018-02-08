package main

import (
	"fmt"
	"log"

	"github.com/davidch93/go-tutorial/package-example/rectangle"
)

var rectLen, rectWidth float64 = 6, 7

/*
 * Init function to check if length and width are greater than zero
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Println("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Println("Diagonal of rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
