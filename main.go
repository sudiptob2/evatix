package main

import (
	"evatix/builder"
	"evatix/inheritance"
	"fmt"
)

func main() {

	// Driver for builder example
	androidPhoneBuilder := new(builder.PhoneBuilder)
	myPhone := androidPhoneBuilder.SetBattery(2000).SetOS("Android 10").SetSceenSize(5).GetPhone()
	fmt.Println(myPhone.GetSpec())

	// Driver for inheritance example
	aPerson := new(inheritance.Person)
	aPerson.Animal.SetLegs(2)
	fmt.Println(aPerson.Walk())

	aDog := new(inheritance.Dog)
	aDog.Animal.SetLegs(2)
	fmt.Println(aDog.Walk())

}
