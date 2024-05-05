package simple

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Greet() {
	fmt.Printf("Hi! My name is %s. I'm %d years old. \n", p.name, p.age)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
