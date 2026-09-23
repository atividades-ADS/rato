package main

import (
	"fmt"
	"rato/labirinto"
	"rato/pilha"
	"rato/rato"
)

func update(lab *labirinto.Labirinto, rato *rato.RatoType, pilhaCaminhos *pilha.Pilha[labirinto.Posicao]) {
	pilhaCaminhos.Push(rato.Posicao)

	for rato.Posicao != lab.Saida || !pilhaCaminhos.IsEmpty() {
		if rato.Mover(labirinto.Posicao{X: rato.Posicao.X + 1, Y: rato.Posicao.Y}) {
			fmt.Println("Rato se moveu para a direita:", rato.Posicao)
			pilhaCaminhos.Push(rato.Posicao)
		} else if rato.Mover(labirinto.Posicao{X: rato.Posicao.X - 1, Y: rato.Posicao.Y}) {
			fmt.Println("Rato se moveu para a esquerda:", rato.Posicao)
			pilhaCaminhos.Push(rato.Posicao)
		} else if rato.Mover(labirinto.Posicao{X: rato.Posicao.X, Y: rato.Posicao.Y + 1}) {
			fmt.Println("Rato se moveu para baixo:", rato.Posicao)
			pilhaCaminhos.Push(rato.Posicao)
		} else if rato.Mover(labirinto.Posicao{X: rato.Posicao.X, Y: rato.Posicao.Y - 1}) {
			fmt.Println("Rato se moveu para cima:", rato.Posicao)
			pilhaCaminhos.Push(rato.Posicao)
		} else {
			fmt.Println("Rato não pode se mover. Voltando para a posição anterior.")
			pilhaCaminhos.Pop()
			if pilhaCaminhos.IsEmpty() {
				fmt.Println("Não há mais posições para voltar. O rato está preso!")
				lab.ImprimirLabirinto()
				break
			}
			rato.Posicao, _ = pilhaCaminhos.Top()
		}
		lab.ImprimirLabirinto()
	}
	if rato.Posicao == lab.Saida {
		fmt.Println("Rato chegou na saída!")
		lab.ImprimirLabirinto()
	}

}

func main() {
	lab := labirinto.InitLabirinto()
	rato := rato.Init(lab)

	pilhaCaminhos := pilha.Init[labirinto.Posicao](nil)

	update(lab, &rato, pilhaCaminhos)

}
