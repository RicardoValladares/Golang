package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ejecutar...")
	file,err := os.OpenFile("/opt/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("No se logro acceder a /opt/logs.txt")
		return
	}
	defer file.Close()
	_, err2 := file.WriteString("EJECUTADO... \n")
	if err2 != nil {
		fmt.Println("No se logro escribir en /opt/logs.txt")
	}
	file.Close()
}
