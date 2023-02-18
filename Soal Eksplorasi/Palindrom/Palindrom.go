package main

import(
	"fmt"
)

func main() {
	var kata, reverse string

	fmt.Scan(&kata)
	
	for i := len(kata)-1; i >=0 ; i-- {
		reverse += string(kata[i])
	}

	if kata == reverse {
		fmt.Printf("%s sama dengan %s", kata, reverse)
		fmt.Println("")
		fmt.Print("palindrom")
	}else{
		fmt.Printf("%s tidak sama dengan %s", kata,reverse)
		fmt.Println("")
		fmt.Print("bukan palindrom")
	}
}