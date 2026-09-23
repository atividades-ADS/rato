package animacao

import (
	"image"
	_ "image/png"
	"os"
	"time"

	"rato/labirinto"
	"rato/pilha"
	"rato/rato"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Lab   *labirinto.Labirinto
	Rato  *rato.RatoType
	Pilha *pilha.Pilha[labirinto.Posicao]

	RatoImagem *ebiten.Image

	ultimoPasso time.Time
	intervalo   time.Duration

	finalizado bool
}

func carregarImagem(caminho string) (*ebiten.Image, error) {
	arquivo, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	img, _, err := image.Decode(arquivo)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func NewGame(
	lab *labirinto.Labirinto,
	r *rato.RatoType,
	p *pilha.Pilha[labirinto.Posicao],
) (*Game, error) {

	img, err := carregarImagem("animacao/rato.png")
	if err != nil {
		return nil, err
	}

	return &Game{
		Lab:         lab,
		Rato:        r,
		Pilha:       p,
		RatoImagem:  img,
		ultimoPasso: time.Now(),
		intervalo:   300 * time.Millisecond,
	}, nil
}

func (g *Game) Passo() {

	if g.Rato.Posicao == g.Lab.Saida {
		g.finalizado = true
		return
	}

	if g.Pilha.IsEmpty() {
		g.finalizado = true
		return
	}

	if g.Pilha.IsEmpty() {
		g.Pilha.Push(g.Rato.Posicao)
	}

	// direita
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X + 1,
		Y: g.Rato.Posicao.Y,
	}) {
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// esquerda
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X - 1,
		Y: g.Rato.Posicao.Y,
	}) {
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// baixo
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X,
		Y: g.Rato.Posicao.Y + 1,
	}) {
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// cima
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X,
		Y: g.Rato.Posicao.Y - 1,
	}) {
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// nenhum caminho disponível
	_, _ = g.Pilha.Pop()

	if !g.Pilha.IsEmpty() {
		g.Rato.Posicao, _ = g.Pilha.Top()
	}
}
