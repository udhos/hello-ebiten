// Package main implements the game.
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!  ttttttttttttttt")
	screen.DrawImage(img, nil)
}

func (g *game) Layout(_ /*outsideWidth*/, _ /*outsideHeight*/ int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
