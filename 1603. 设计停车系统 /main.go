package main

import "fmt"

func main() {
	obj := Constructor(1, 1, 0)
	fmt.Print(obj.AddCar(1))
}



type ParkingSystem struct {
	SmallCapacity int
	BigCapacity int
	MediumCapacity int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	system := ParkingSystem{small, big, medium}
	return system
}


func (this *ParkingSystem) AddCar(carType int) bool {
	res := true
	if carType == 1 {
		res =  this.BigCapacity > 0
		this.BigCapacity--
	}else if carType == 2 {
		res = this.MediumCapacity > 0
		this.MediumCapacity--
	}else {
		res = this.SmallCapacity > 0
		this.SmallCapacity--
	}
	return res
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */