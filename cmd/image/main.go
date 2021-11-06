package main

import (
	"bytes"
	ecsexample "ebiten-ecs-animation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images/flappy"
	"image"
	_ "image/png"
)

const screenWidth = 500
const screenHeight = 500

type SimpleImageExample struct {
	spriteRenderSystem *ecsexample.SpriteRenderSystem
}

func (a *SimpleImageExample) Draw(screen *ebiten.Image) {
	a.spriteRenderSystem.Draw(screen)
}

func (a *SimpleImageExample) Update() error {
	return nil
}

func (a *SimpleImageExample) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	// loading assets
	img, _ := LoadImage(flappy.Gopher_png)
	sprite := ebiten.NewImageFromImage(img)

	// entities
	registry := ecsexample.Registry{}
	gopher := registry.NewEntity()

	// components
	gopher.AddComponent(&ecsexample.TransformComponent{PosY: 200, PosX: 200})
	gopher.AddComponent(&ecsexample.SpriteComponent{Image: sprite})

	// systems
	spriteRenderSystem := ecsexample.SpriteRenderSystem{Registry: &registry}

	// the game
	example := SimpleImageExample{spriteRenderSystem: &spriteRenderSystem}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("render single image")
	ebiten.RunGame(&example) // omitted error handling
}

func LoadImage(input []byte) (image.Image, error) {
	if img, _, err := image.Decode(bytes.NewReader(input)); err != nil {
		return img, err
	} else {
		return img, nil
	}
}
