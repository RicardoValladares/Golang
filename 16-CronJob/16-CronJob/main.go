package main

import (
	"fmt"
	"time"
	"github.com/go-co-op/gocron"
)

//tarea
func tarea(texto string) {
	fmt.Println(texto)
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
