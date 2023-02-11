package main

import (
	"os"
	"fmt"
	"time"
	"github.com/go-co-op/gocron"
)

//tarea
func tarea(texto string) {
	fmt.Println(texto)
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

//cronometrar
func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	//s.Every(1).Hours().Do(func() {
	s.Every(5).Seconds().Do(func() {
		tarea("Ejecutar...")
 	})
 	s.StartBlocking()
}

//principal
func main() {
	runCronJobs()
}
