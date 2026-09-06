package main

import (
	"cursogo/model"
	"fmt"
	"time"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua Des. Caio Oliveira",
		Bairro: "Loteamento Mocambinho",
		Numero: 8115,
		Cidade: "Teresina",
	}

	pessoa := model.Pessoa{
		Nome:           "Glêsio Santos",
		Endereco:       endereco,
		DataNascimento: time.Date(1982, 11, 16, 0, 0, 0, 0, time.Local),
	}

	pessoa.CalcularIdade()

	fmt.Printf("Paciente: %s com idade de %d\n", pessoa.Nome, pessoa.Idade)

	motoAutomovel := model.Automovel{
		Ano:    2022,
		Placa:  "XTPO2520",
		Modelo: "Honda",
	}

	moto := model.Moto{
		Automovel:  motoAutomovel,
		Cilindrada: 125,
	}

	fmt.Println(moto)

}
