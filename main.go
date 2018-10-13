package main

import (
	parkinglot "./parkinglot"
)

func main() {
	parking.Init()
	//Code for Command Shell Input
	if parkinglot.Type_Of_Input == 1 {
		parkinglot.StartCommandOperation()
	} else if parkinglot.Type_Of_Input == 2 { //Code for file input

	} else {

	}
}
