package structs

type Car struct {
	model string
}


func changeModelByValue(car Car){
	car.model = "Civic"
}