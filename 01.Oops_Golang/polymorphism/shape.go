package polymorphism

type Shape interface {
	Render()
}

type Circle struct{}
type Rectangle struct{}

func (c Circle) Render() {

}

func (r Rectangle) Render() {

}
