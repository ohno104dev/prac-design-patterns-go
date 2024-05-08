package proxy

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("car is being driven")
}
