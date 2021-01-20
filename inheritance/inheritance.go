package inheritance

import "fmt"

type animal struct {
	mouth int
	legs  int
}

func (a animal) SetLegs(leg int) {
	a.legs = leg
}
func (a animal) SetMouth(mouth int) {
	a.mouth = mouth
}

// Dog : Instead of thinking Dog is an animal,
// we think like Dog has the functionalities of an animal.
// Similarly for Person
type Dog struct {
	Animal animal
}

func (d Dog) Walk() string {
	return fmt.Sprintf("Dog is walking with %d legs", d.Animal.legs)
}

type Person struct {
	Animal animal
}

func (p Person) Walk() string {
	return fmt.Sprintf("Person is walking with %d legs", p.Animal.legs)
}

// func main() {

// 	aPerson := new(Person)
// 	fmt.Println(aPerson.walk())

// 	aDog := new(Dog)
// 	fmt.Println(aDog.walk())

// }
