package model

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Automovel
	Cilindrada int
}

type Carro struct {
	Automovel        // seria a forma de herança
	QuantidadePortas int
	Potencia         int
	PossuiArCondic   bool
}
