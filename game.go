package main

import(
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {

}

func NewGame() *Game {
	game := &Game{}

	game.Init()

	return game
}

func (g *Game) Init() {

}

func (g *Game) Update() error {
	var err error
	
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(w, h int) (int, int) {
	return GAMEWIDTH, GAMEHEIGHT
}