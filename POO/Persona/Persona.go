package Persona

import "fmt"

// inicilizacion de packete
func init() {
	fmt.Println("Packete Persona Inicializado")
}

// estructura de datos Persona
type Informacion struct {
	Edad int
	Nombre1 string
	Nombre2 string
	Apellido1 string
	Apellido2 string
}

// metodos get
func (info Informacion) GetEdad() int {
	return info.Edad
}

func (info Informacion) GetNombre1() string {
	return info.Nombre1
}

func (info Informacion) GetNombre2() string {
	return info.Nombre2
}

func (info Informacion) GetApellido1() string {
	return info.Apellido1
}

func (info Informacion) GetApellido2() string {
	return info.Apellido2
}

// metodos set
func (info *Informacion) SetEdad(edad int) {
	info.Edad = edad
}

func (info *Informacion) SetNombre1(nombre1 string) {
	info.Nombre1 = nombre1
}

func (info *Informacion) SetNombre2(nombre2 string) {
	info.Nombre2 = nombre2
}

func (info *Informacion) SetApellido1(apellido1 string) {
	info.Apellido1 = apellido1
}

func (info *Informacion) SetApellido2(apellido2 string) {
	info.Apellido2 = apellido2
}
