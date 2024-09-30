package main

import (
	"fmt"

	"github.com/Amari05fc/practica-6/cmd/config"
	
)

func main() {
	fmt.Println("Main")
	fmt.Println(config.PASSWORD)
	fmt.Println("Variables de entorno: ", config.USERNAME, config.PASSWORD)
}