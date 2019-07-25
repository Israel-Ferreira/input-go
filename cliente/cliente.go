package cliente

import (
	"fmt"
)

// Cliente represent a customer entity
type Cliente struct {
	Name string
	Age int
	City string
	State string
}


// ImprimirCliente returns a name and city of customer
func (cliente *Cliente) ImprimirCliente() string{
	return fmt.Sprintf("Nome: %s - Cidade: %s", cliente.Name,cliente.City)
}