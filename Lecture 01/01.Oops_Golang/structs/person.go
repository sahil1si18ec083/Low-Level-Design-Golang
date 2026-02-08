package structs

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p *Person) Walk() string {
	return (p.firstName + " Likes walking")
}
func NewPerson(firstName string, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}

}
func (p *Person) SetFirstName(firstName string) {
	if firstName == "" {
		panic("firstname cannot be blank")
	}
	p.firstName = firstName

}

func (p *Person) GetFirstName() string {
	return p.firstName

}
