// Аналог наследования
package main

import "fmt"

type Human struct {
	Name string
}

type Action struct {
	Human
}

func (h Human) tellYourName() string {
	return "my name is " + h.Name
}
func (a Action) whatDoYouKnow() string {
	return "my name is still " + a.Name + " and I know tropinki volshebnih polyan"
}
func (a Action) drink() string {
	return "my name is still " + a.Name + " and I am non-drinker, sorryan"
}
func main() {
	kolyan := Human{Name: "Kolyan"}
	actionsWithKolyan := Action{Human: kolyan}
	fmt.Println(kolyan.tellYourName())
	fmt.Println(actionsWithKolyan.whatDoYouKnow())
	fmt.Println(actionsWithKolyan.drink())
}
