package funcs

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s. I'm %d years old. \n", p.name, p.age)
}

// 解耦合, 變成由一群子類負責對具體類的實例化
func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			age:  age,
			name: name,
		}
	}
}

// babyFactory := NewPersonFactory(1)
// teenagerFactory := NewPersonFactory(16)
