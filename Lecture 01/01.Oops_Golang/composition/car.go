package composition

type engine struct {
	hp int
}
type wheel struct {
	WheelDimension int
}
type Car struct {
	engine
	wheel
	CarName string
}

func (w wheel) GetWheelDimension() int {
	return w.WheelDimension

}

func NewCar(carname string, hp int, wheeldimension int) Car {
	return Car{
		CarName: carname,
		engine: engine{
			hp: hp,
		},
		wheel: wheel{
			WheelDimension: wheeldimension,
		},
	}

}
