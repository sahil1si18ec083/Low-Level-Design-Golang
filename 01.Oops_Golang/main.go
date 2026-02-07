// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"golang-oops/composition"
	"golang-oops/polymorphism"
	"golang-oops/structs"
)

func main() {
	fmt.Println("Try programiz.pro")
	p := structs.NewPerson("sahil", "kumar", 26)
	p.SetFirstName("jethiya")
	fmt.Println(p.Walk())
	fmt.Println(p.GetFirstName())

	// polymorphism

	var c polymorphism.Shape = polymorphism.Circle{}
	c.Render()
	var r polymorphism.Shape = polymorphism.Rectangle{}
	r.Render()

	car := composition.NewCar("honda", 89, 2000)
	fmt.Println(car)
	fmt.Println(car.GetWheelDimension())

}
