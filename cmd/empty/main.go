package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

const screenWidth = 500
const screenHeight = 500

type EmptyScreenExample struct{}

func (a *EmptyScreenExample) Draw(screen *ebiten.Image) {}
func (a *EmptyScreenExample) Update() error             { return nil }
func (a *EmptyScreenExample) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	example := EmptyScreenExample{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("empty screen")
	ebiten.RunGame(&example) // omitted error handling
}
