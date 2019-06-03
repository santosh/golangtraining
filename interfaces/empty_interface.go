package main

import "fmt"

type vehicles interface {
}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func (v vehicle) Specs() {
	fmt.Printf("Seats %v, max speed %v, color %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boing747 := plane{}
	boing757 := plane{}
	boing767 := plane{}
	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	rides := []vehicles{prius, tacoma, bmw528, boing747, boing757, boing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
