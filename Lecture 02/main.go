package main

import (
	"errors"
	"fmt"
)

type Car interface {
	StartEngine()
	ShiftGear(gear int)
	Accelerate()
	Brake()
	StopEngine()
}
type SportsCar struct {
	brand        string
	model        string
	isEngineOn   bool
	currentSpeed int
	currentGear  int
}

func NewSportsCar(brand, model string) (*SportsCar, error) {
	if brand == "" || model == "" {
		return nil, errors.New("brand or model cannot be blank")
	}
	return &SportsCar{
		brand:        brand,
		model:        model,
		isEngineOn:   false,
		currentSpeed: 0,
		currentGear:  0,
	}, nil

}
func (s *SportsCar) StartEngine() {
	s.isEngineOn = true
	fmt.Println(s.model, s.brand, " started engine")

}
func (s *SportsCar) ShiftGear(gear int) {
	if !s.isEngineOn {
		return

	}
	s.currentGear = gear
	fmt.Println(s.model, s.brand, " gear shifted to ", s.currentGear)

}
func (s *SportsCar) Accelerate() {
	if !s.isEngineOn {
		return

	}
	s.currentSpeed += 10
	fmt.Println(s.model, s.brand, " speed increased to ", s.currentSpeed)

}
func (s *SportsCar) Brake() {
	if !s.isEngineOn {
		return

	}
	s.currentSpeed -= 10
	fmt.Println(s.model, s.brand, " speed decreased to ", s.currentSpeed)

}
func (s *SportsCar) StopEngine() {
	s.isEngineOn = false
	s.currentSpeed = 0
	s.currentGear = 0
	fmt.Println(s.model, s.brand, " engine stopped")

}
func main() {

	fmt.Println("hello")

	mysportscar, _ := NewSportsCar("Ford", "Mustang")
	mysportscar.StartEngine()
}
