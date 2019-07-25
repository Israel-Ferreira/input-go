package main

import (
	"fmt"
	"github.com/Israel-Ferreira/cadastro-clientes/cliente"
	"io/ioutil"
	"encoding/json"
)

func main(){
	clientes := []cliente.Cliente{}

	for len(clientes) <= 8 {
		clientes =  append(clientes,cadastrarCliente())
	}


	byteValueJSON, err :=  json.Marshal(clientes)

	err =  ioutil.WriteFile("clientes.json",byteValueJSON,0644)

	if err != nil {
		panic("Erro")
	}

	fmt.Println("Escrita Feita com sucesso")
}



func cadastrarCliente() cliente.Cliente{
	newCliente := cliente.Cliente{}

	fmt.Println("Digite o nome do cliente: ")
	fmt.Scanln(&newCliente.Name)

	fmt.Println("Digite a idade do cliente: ")
	fmt.Scanln(&newCliente.Age)


	return newCliente

}
