package main

import(
	
	"fmt"
)
func main(){


	//Actualizar Datos

	jj := Juegos {
		ID: 13,
		Name: "Flappy Bird",
		Age: 2014,
		Active: true,
	}

	err := Actualizar(jj)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Actualizado correctamente")




}