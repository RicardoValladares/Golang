package main

import "fmt"
import Docente "github.com/RicardoValladares/Golang/09-POO/Persona/Docente"

func main() {
	//creamos objeto profe1
	Profe1 := new(Docente.Informacion)
	//seteamos datos
	Profe1.SetEdad(28)
	Profe1.SetNombre1("Ricardo")
	Profe1.SetNombre2("Antonio")
	Profe1.SetApellido1("Valladares")
	Profe1.SetApellido2("Renderos")
	Profe1.SetCarnet("VR201104")
	//obtenemos datos
	fmt.Println("\tEdad:",Profe1.GetEdad())
	fmt.Println("\tNombre:",Profe1.GetNombre1(),Profe1.GetNombre2())
	fmt.Println("\tApellido:",Profe1.GetApellido1(),Profe1.GetApellido2())
	fmt.Println("\tCarnet:",Profe1.GetCarnet())
}
