////////////////////////DUCK///////////////////////////////
package main

import "fmt"

func main() {
	h := human{firstName: "Vasya", lastName: "Pupkin", occupation: "student"}
	d := dog{name: "Spotty", breed: "labrador"}
	c := cat("Whiskers")
	fmt.Println(greet(h))
	fmt.Println(greet(d))
	fmt.Println(greet(c))
}

func greet(c creature) string {
	return fmt.Sprintf("Hello, my dear %s!", c.Name())
}

type creature interface {
	Name() string
}

type human struct {
	firstName  string
	lastName   string
	occupation string
}

func (h human) Name() string {
	return fmt.Sprintf("%s %s", h.firstName, h.lastName)
}

type dog struct {
	name  string
	breed string
}

func (d dog) Name() string {
	return d.name
}

type cat string

func (c cat) Name() string {
	return string(c)
}