package animacao

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"time"

	"rato/labirinto"
	"rato/pilha"
	"rato/rato"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Lab   *labirinto.Labirinto
	Rato  *rato.RatoType
	Pilha *pilha.Pilha[labirinto.Posicao]

	RatoImagem     *ebiten.Image
	ParedeImagem   *ebiten.Image
	CaminhoImagem  *ebiten.Image
	VisitadoImagem *ebiten.Image
	EntradaImagem  *ebiten.Image
	SaidaImagem    *ebiten.Image

	ultimoPasso time.Time
	intervalo   time.Duration

	finalizado bool
	pensando   bool
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

	img, err := carregarImagem("animacao/frames/rato_correndo.png")
	if err != nil {
		return nil, err
	}
	parede, err := carregarImagem("animacao/labirinto/parede.png")
	if err != nil {
		return nil, err
	}
	caminho, err := carregarImagem("animacao/labirinto/caminho.png")
	if err != nil {
		return nil, err
	}
	visitado, err := carregarImagem("animacao/labirinto/visitado.png")
	if err != nil {
		return nil, err
	}

	entrada, err := carregarImagem("animacao/labirinto/entrada.png")
	if err != nil {
		return nil, err
	}

	saida, err := carregarImagem("animacao/labirinto/saida.png")
	if err != nil {
		return nil, err
	}

	p.Push(r.Posicao)

	return &Game{
		Lab:            lab,
		Rato:           r,
		Pilha:          p,
		RatoImagem:     img,
		ParedeImagem:   parede,
		CaminhoImagem:  caminho,
		VisitadoImagem: visitado,
		EntradaImagem:  entrada,
		SaidaImagem:    saida,
		ultimoPasso:    time.Now(),
		intervalo:      300 * time.Millisecond,
	}, nil
}

func desenharImagemCelula(
	screen *ebiten.Image,
	imagem *ebiten.Image,
	x, y int,
) {
	op := &ebiten.DrawImageOptions{}

	op.Filter = ebiten.FilterLinear

	largura := imagem.Bounds().Dx()
	altura := imagem.Bounds().Dy()

	escalaX := float64(tamanhoCelula) / float64(largura)
	escalaY := float64(tamanhoCelula) / float64(altura)

	op.GeoM.Scale(
		escalaX,
		escalaY,
	)

	op.GeoM.Translate(
		float64(x*tamanhoCelula),
		float64(y*tamanhoCelula),
	)

	screen.DrawImage(imagem, op)
}

func mesmaPosicao(a, b labirinto.Posicao) bool {
	return a.X == b.X && a.Y == b.Y
}

func (g *Game) Passo() {

	if g.pensando {
		g.pensando = false

		_, _ = g.Pilha.Pop()

		if !g.Pilha.IsEmpty() {
			g.Rato.Posicao, _ = g.Pilha.Top()
		}

		img, _ := carregarImagem("animacao/frames/rato_correndo.png")
		g.RatoImagem = img

		return
	}

	if mesmaPosicao(g.Rato.Posicao, g.Lab.Saida) {
		g.finalizado = true
		img, _ := carregarImagem("animacao/frames/rato_comemorando.png")

		g.RatoImagem = img
		return
	}

	if g.Pilha.IsEmpty() {
		g.finalizado = true
		img, _ := carregarImagem("animacao/frames/rato_triste.png")
		g.RatoImagem = img
		return
	}

	// direita
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X + 1,
		Y: g.Rato.Posicao.Y,
	}) {
		img, _ := carregarImagem("animacao/frames/rato_correndo.png")
		g.RatoImagem = img
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// esquerda
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X - 1,
		Y: g.Rato.Posicao.Y,
	}) {
		img, _ := carregarImagem("animacao/frames/rato_correndo.png")
		g.RatoImagem = img
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// baixo
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X,
		Y: g.Rato.Posicao.Y + 1,
	}) {
		img, _ := carregarImagem("animacao/frames/rato_correndo.png")
		g.RatoImagem = img
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// cima
	if g.Rato.Mover(labirinto.Posicao{
		X: g.Rato.Posicao.X,
		Y: g.Rato.Posicao.Y - 1,
	}) {
		img, _ := carregarImagem("animacao/frames/rato_correndo.png")
		g.RatoImagem = img
		g.Pilha.Push(g.Rato.Posicao)
		return
	}

	// nenhum caminho disponível
	img, _ := carregarImagem("animacao/frames/rato_pensativo.png")
	g.RatoImagem = img
	g.pensando = true

	if !g.Pilha.IsEmpty() {
		g.Rato.Posicao, _ = g.Pilha.Top()

	}

}

func (g *Game) Update() error {

	if g.finalizado {
		return nil
	}

	if time.Since(g.ultimoPasso) < g.intervalo {
		return nil
	}

	g.Passo()

	g.ultimoPasso = time.Now()

	return nil
}

const tamanhoCelula = 140
const painelLargura = 300

func desenharCelula(
	screen *ebiten.Image,
	x, y int,
	cor color.Color,
) {
	vector.DrawFilledRect(
		screen,
		float32(x*tamanhoCelula),
		float32(y*tamanhoCelula),
		tamanhoCelula,
		tamanhoCelula,
		cor,
		false,
	)
}

func (g *Game) desenharRato(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear

	op.GeoM.Scale(
		-0.7,
		0.7,
	)

	op.GeoM.Translate(
		float64((g.Rato.Posicao.X+1)*tamanhoCelula),
		float64(g.Rato.Posicao.Y*tamanhoCelula),
	)

	screen.DrawImage(g.RatoImagem, op)
}

func (g *Game) desenharPilha(screen *ebiten.Image) {

	larguraLabirinto := len(g.Lab.Mapa[0]) * tamanhoCelula

	xPainel := larguraLabirinto

	// Fundo do painel
	vector.DrawFilledRect(
		screen,
		float32(xPainel),
		0,
		painelLargura,
		float32(len(g.Lab.Mapa)*tamanhoCelula),
		color.RGBA{25, 25, 30, 255},
		false,
	)

	// Título
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("PILHA (%d)", g.Pilha.Size()),
		xPainel+20,
		20,
	)

	itens := g.Pilha.Items()

	// Espaçamento entre os elementos
	y := 60
	alturaItem := 50

	// Mostra primeiro o topo da pilha
	for i := len(itens) - 1; i >= 0; i-- {

		// Se o painel ficar cheio, paramos de desenhar
		if y+alturaItem > len(g.Lab.Mapa)*tamanhoCelula-20 {
			break
		}

		// O topo recebe uma cor diferente
		if i == len(itens)-1 {

			vector.DrawFilledRect(
				screen,
				float32(xPainel+20),
				float32(y),
				260,
				float32(alturaItem),
				color.RGBA{100, 150, 220, 255},
				false,
			)

		} else {

			vector.DrawFilledRect(
				screen,
				float32(xPainel+20),
				float32(y),
				260,
				float32(alturaItem),
				color.RGBA{60, 60, 70, 255},
				false,
			)
		}

		texto := fmt.Sprintf("(%d, %d)", itens[i].X, itens[i].Y)

		ebitenutil.DebugPrintAt(
			screen,
			texto,
			xPainel+35,
			y+17,
		)

		y += alturaItem + 10
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{30, 30, 30, 255})

	for y, linha := range g.Lab.Mapa {

		for x, pos := range linha {

			switch pos.Valor {

			case '1':
				desenharImagemCelula(
					screen,
					g.ParedeImagem,
					x,
					y,
				)

			case '0':
				desenharImagemCelula(
					screen,
					g.CaminhoImagem,
					x,
					y,
				)

			case 'M':
				desenharImagemCelula(
					screen,
					g.EntradaImagem,
					x,
					y,
				)

			case 'E':
				desenharImagemCelula(
					screen,
					g.SaidaImagem,
					x,
					y,
				)

			case '.':
				desenharImagemCelula(
					screen,
					g.VisitadoImagem,
					x,
					y,
				)
			}
		}
	}

	g.desenharRato(screen)
	g.desenharPilha(screen)
}

func (g *Game) Layout(
	outsideWidth,
	outsideHeight int,
) (int, int) {

	largura := len(g.Lab.Mapa[0]) * tamanhoCelula
	altura := len(g.Lab.Mapa) * tamanhoCelula

	largura += painelLargura

	return largura, altura
}
