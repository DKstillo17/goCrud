package main

import(
	"log"
	"fmt"
)
func main(){

	//Crear datos

	e := Juegos{
		Name: "Super Mario Bros",
		Age: 1998,
		Active: true,
	}

	err := Crear(e)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("¡Creado exitosamente!")
	

}