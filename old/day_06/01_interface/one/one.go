package one

import "fmt"

type Test_Car interface {
	Car_Name() string
	Car_Run()
	Car_Stop()
}

type X1 struct {
	Name string
}

func (this X1) Car_Name() string {
	return this.Name + " X1"
}

func (this X1) Car_Run() {
	fmt.Println(this.Name, "X1, 正在行驶")
}

func (this X1) Car_Stop() {
	fmt.Println(this.Name, "X1, 停车")
}

type X2 struct {
	Name string
}

func (this X2) Car_Name() string {
	return this.Name + " X2"
}

func (this X2) Car_Run() {
	fmt.Println(this.Name, "X2, 正在行驶")
}

func (this X2) Car_Stop() {
	fmt.Println(this.Name, "X2, 停车")
}

func One() {
	var Car Test_Car

	var BMW_X1 X1 = X1{
		Name: "BMW",
	}

	Car = BMW_X1

	fmt.Println(Car.Car_Name())
	Car.Car_Run()
	Car.Car_Stop()

	fmt.Println(Car)

	var BMW_X2 X2 = X2{
		Name: "BMW",
	}

	Car = BMW_X2

	fmt.Println(Car.Car_Name())
	Car.Car_Run()
	Car.Car_Stop()
}
