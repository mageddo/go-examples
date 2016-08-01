package main

import "fmt"

func main(){

	var v *int
	var tmp int = 1
	v = &tmp;
	for ;;{
		fmt.Printf("v=%d, p=%p, pv=%p\n", *v, &v, v)
		fmt.Scanln(&tmp);
		if(tmp != -1){
			*v = tmp
		}
	}
}
