package main

import(
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle(GAMENAME)
	ebiten.SetWindowSize(GAMEWIDTH, GAMEHEIGHT)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil && err.Error() != "quit" {
		panic(err)
	}
}