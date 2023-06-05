package main

import "fmt"

func main() {
	testStruct := new(WarBike)
	testStruct.On = true
	testStruct.Ammo = 1
	testStruct.Power = 1
	fmt.Println(*testStruct)
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
	fmt.Println(*testStruct)
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
}

type WarBike struct {
	On    bool
	Ammo  int
	Power int
}

func (w *WarBike) Shoot() bool {
	if w.On == false {
		return false
	}
	if w.Ammo > 0 {
		w.Ammo--
		return true
	}
	return false
}

func (w *WarBike) RideBike() bool {
	if w.On == false {
		return false
	}
	if w.Power > 0 {
		w.Power--
		return true
	}
	return false
}
