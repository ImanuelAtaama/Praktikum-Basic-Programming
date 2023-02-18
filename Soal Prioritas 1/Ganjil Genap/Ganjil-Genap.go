package main

import (
	"fmt"
)	

func main(){
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan % 2 == 0 {
		fmt.Printf("genap")
	} else{
		fmt.Printf("ganjil")
	}
	

};