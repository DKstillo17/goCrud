package main

import(

	"fmt"
)
func main(){

//Borrar datos

	err := Borrar(12)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Borrado correctamente")


}