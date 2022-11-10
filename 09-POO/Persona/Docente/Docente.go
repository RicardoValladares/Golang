package Docente

// importacion de packetes
import "fmt"
import Persona "github.com/RicardoValladares/Golang/09-POO/Persona"

// inicializacion de packete
func init() {
	fmt.Println("Packete Docente Inicializado")
}

// herencia del packete Persona a Docente
type Informacion struct {
	Persona.Informacion
	Carnet string
}

// metodos get
func (info Informacion) GetCarnet() string {
	return info.Carnet
}

// metodos set
func (info *Informacion) SetCarnet(carnet string) {
	info.Carnet = carnet
}
