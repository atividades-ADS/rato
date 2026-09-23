package rato

import (
	"rato/labirinto"
)

type RatoType struct {
	Posicao   labirinto.Posicao
	labirinto *labirinto.Labirinto
}

func Init(lab *labirinto.Labirinto) RatoType {
	return RatoType{
		Posicao:   lab.Entrada,
		labirinto: lab,
	}
}

func (r *RatoType) Mover(novaPosicao labirinto.Posicao) bool {
	if r.labirinto.PosicaoValida(novaPosicao) == false {
		return false
	}

	r.Posicao = novaPosicao
	r.labirinto.MarcarVisitado(r.Posicao)

	return true
}
