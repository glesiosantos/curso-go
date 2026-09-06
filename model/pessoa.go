package model

import "time"

type Pessoa struct {
	Nome           string
	Endereco       Endereco
	DataNascimento time.Time
	Idade          int
}

func (p *Pessoa) CalcularIdade() {
	anoNascimento := p.DataNascimento.Year()
	anaoAtual := time.Now().Year()

	p.Idade = anaoAtual - anoNascimento
}

// func CalcularIdade(p Pessoa) int {
// 	anoNascimento := p.DataNascimento.Year()
// 	anaoAtual := time.Now().Year()

// 	return anaoAtual - anoNascimento
// }
