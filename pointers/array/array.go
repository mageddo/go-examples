package main;

import "fmt"

func main(){
	numbers := make([]int, 5)
	changeValue(numbers, 3, 99)
	fmt.Printf("numbers=%+v\n", numbers)

}

func changeValue(arr []int, pos int, value int){
	arr[pos] = value
}