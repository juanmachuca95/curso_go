package main
/*
VARIABLES DE ENTORNO
Para esto utilizaremos una libreria https://github.com/joho/godotenv
A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file)

1. go mod init [Nombre del proyecto]
2. go get github.com/joho/godotenv 
*/

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	loadEnv() // loading environment's variables 

	secret := os.Getenv("SECRET")

	fmt.Printf("The variable SECRET is %s", secret)

}
	
func loadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}