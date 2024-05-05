package abstract

import "fmt"

type IPerson interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s. I'm %d years old. \n", p.name, p.age)
}

// 返回interface, 在不公開內部實現的情況下讓調用者使用提供的func
func NewPerson(name string, age int) IPerson {
	return person{
		name: name,
		age:  age,
	}
}
