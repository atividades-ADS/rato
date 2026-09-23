package labirinto

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Posicao struct {
	X        int
	Y        int
	visitado bool
	Valor    rune
}

type Labirinto struct {
	Mapa    [][]Posicao
	Entrada Posicao
	Saida   Posicao
}

func InitLabirinto() *Labirinto {

	labirinto := make([][]Posicao, 0)
	linha := make([]Posicao, 0)

	var entrada, saida Posicao

	file, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		char, _, err := reader.ReadRune()

		if err == io.EOF {
			if len(linha) > 0 {
				labirinto = append(labirinto, linha)
			}
			break
		}

		if err != nil {
			panic(err)
		}

		switch {
		case char == 'm':
			entrada = Posicao{X: len(linha), Y: len(labirinto), visitado: true}
			linha = append(linha, Posicao{X: len(linha), Y: len(labirinto), visitado: true, Valor: 'M'})

		case char == 'e':
			saida = Posicao{X: len(linha), Y: len(labirinto), visitado: false}
			linha = append(linha, Posicao{X: len(linha), Y: len(labirinto), visitado: false, Valor: 'E'})

		case char == '1':
			linha = append(linha, Posicao{X: len(linha), Y: len(labirinto), visitado: false, Valor: '1'})

		case char == '0':
			linha = append(linha, Posicao{X: len(linha), Y: len(labirinto), visitado: false, Valor: '0'})

		case char == '\n':
			labirinto = append(labirinto, linha)
			linha = make([]Posicao, 0)

		}
	}

	return &Labirinto{
		Mapa:    labirinto,
		Entrada: entrada,
		Saida:   saida,
	}
}

func (labirinto *Labirinto) ImprimirLabirinto() {
	for _, linha := range labirinto.Mapa {
		for _, char := range linha {
			fmt.Printf("%c", char.Valor)
		}
		fmt.Println()
	}
}

func (labirinto *Labirinto) PosicaoValida(posicao Posicao) bool {
	if posicao.X < 0 || posicao.Y < 0 || posicao.Y >= len(labirinto.Mapa) || posicao.X >= len(labirinto.Mapa[posicao.Y]) {
		return false
	}

	if labirinto.Mapa[posicao.Y][posicao.X].Valor == '1' {
		return false
	}

	if labirinto.Mapa[posicao.Y][posicao.X].visitado {
		return false
	}

	return true
}

func (labirinto *Labirinto) MarcarVisitado(posicao Posicao) {
	labirinto.Mapa[posicao.Y][posicao.X].visitado = true
	labirinto.Mapa[posicao.Y][posicao.X].Valor = '.'

}
