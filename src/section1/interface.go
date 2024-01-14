package section1

import (
	"fmt"
	"unsafe"
)

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bicycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (b *bicycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}

func (b *bicycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}

func (v vehicle) String() string {
	return fmt.Sprintf("Vehicle current speed is %v (engine power is %v)", v.speed, v.enginePower)
}

func speedUpAndDown(c controller) {
	c.speedUp()
	c.speedDown()
}

func interfaceFunc() {
	// structを定義して&でポインタを渡す
	v := &vehicle{0, 5}
	speedUpAndDown(v)

	b := &bicycle{0, 5}
	speedUpAndDown(b)

	fmt.Println(v)

	// この２つは同じ
	var i1 interface{}
	var i2 any

	fmt.Printf("%[1]v %[1]T\n %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T\n %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)

	i2 = 1
	checkType(i2)

	i2 = "hello"
	checkType(i2)
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
