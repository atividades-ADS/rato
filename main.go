package main

import (
	"log"

	"rato/animacao"
	"rato/labirinto"
	"rato/pilha"
	"rato/rato"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	lab := labirinto.InitLabirinto()

	r := rato.Init(lab)

	pilhaCaminhos := pilha.Init[labirinto.Posicao](nil)

	game, err := animacao.NewGame(
		lab,
		&r,
		pilhaCaminhos,
	)

	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetMonitor(ebiten.Monitor())
	ebiten.SetWindowTitle("Rato no Labirinto")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
