package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operacion := scanner.Text()
	log.Println(operacion)

	valores := strings.Split(operacion, "+")
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])
	resultado := operador1 + operador2
	log.Println(resultado)
}
