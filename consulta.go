package main

import(
	
	"fmt"
)
func main(){

	//Consultar datos

	js, err := Consultar()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(js)

}