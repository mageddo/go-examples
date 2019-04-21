package structs

import "testing"

func TestMusChangeModelWhenStructWasPassedByValue(t *testing.T){

	// arrange
	car := Car{model: "Compass"}

	// act

	changeModelByValue(car)

	// assert
	if car.model != "Compass" {
		panic("car model must not be changed because the struct was passed as value not as pointer")
	}

}