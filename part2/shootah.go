package main

import "fmt"

type Shootah struct {
	On    bool
	Ammo  int
	Power int
}

func (sh *Shootah) Shoot() bool {
	if sh.On == false || sh.Ammo <= 0 {
		return false
	} else {
		sh.Ammo--
		return true
	}
}

func (sh *Shootah) RideBike() bool {
	if sh.On == false || sh.Power <= 0 {
		return false
	} else {
		sh.Power--
		return true
	}
}

func main() {
	shootah := Shootah{true, 1, 1}
	testStruct := &shootah
	fmt.Println(testStruct.Shoot())
}
