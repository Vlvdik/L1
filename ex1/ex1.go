package ex1

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	anotherField string
	Human
}

func NewHuman(name string, age int) *Human {
	return &Human{
		name: name,
		age:  age,
	}
}

func (h *Human) GetHumanInfo() string {
	return fmt.Sprintf("Human\nName: %s\tAge: %d", h.name, h.age)
}

func NewAction(anotherField string, h Human) *Action {
	return &Action{
		anotherField: anotherField,
		Human:        h,
	}
}

func (a *Action) GetActionInfo() string {
	return fmt.Sprintf("Action\nanotherField: %s\n%v", a.anotherField, a.GetHumanInfo())
}

func Test() {
	h := NewHuman("Vlad", 20)
	a := NewAction("some info", *h)
	fmt.Println(a.GetActionInfo())
}
