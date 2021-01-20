package builder

import "fmt"

// Phone Actual class
type Phone struct {
	screenSize int
	battery    int
	os         string
}

func (p Phone) GetSpec() string {
	return fmt.Sprintf("Sceen : %d\nBattery : %d\nOS : %s", p.screenSize, p.battery, p.os)
}

// PhoneBuilder Builder class
type PhoneBuilder struct {
	Phone
}

func (pb PhoneBuilder) SetSceenSize(screenSize int) PhoneBuilder {
	pb.Phone.screenSize = screenSize
	return pb
}

func (pb PhoneBuilder) SetBattery(battery int) PhoneBuilder {
	pb.Phone.battery = battery
	return pb
}

func (pb PhoneBuilder) SetOS(os string) PhoneBuilder {
	pb.Phone.os = os
	return pb
}

func (pb PhoneBuilder) GetPhone() Phone {
	return pb.Phone
}

func main() {

	androidPhoneBuilder := new(PhoneBuilder)
	myPhone := androidPhoneBuilder.setBattery(2000).setOS("Android 10").setSceenSize(5).getPhone()
	fmt.Println(myPhone.getSpec())

}
