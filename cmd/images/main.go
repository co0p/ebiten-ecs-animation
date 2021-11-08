package main

import (
	"bytes"
	ecsexample "ebiten-ecs-animation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images/flappy"
	"image"
	_ "image/png"
	"math/rand"
)

const screenWidth = 500
const screenHeight = 500

type MultipleImagesExample struct {
	spriteRenderSystem *ecsexample.SpriteRenderSystem
}

func (a *MultipleImagesExample) Draw(screen *ebiten.Image) {
	a.spriteRenderSystem.Draw(screen)
}

func (a *MultipleImagesExample) Update() error {
	return nil
}

func (a *MultipleImagesExample) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	// loading assets
	img, _ := LoadImage(flappy.Gopher_png)
	sprite := ebiten.NewImageFromImage(img)

	// entities
	registry := ecsexample.Registry{}

	for i := 0; i < 1000; i++ {
		x := rand.Intn(screenWidth)
		y := rand.Intn(screenWidth)
		randomGopher := registry.NewEntity()
		randomGopher.AddComponent(&ecsexample.TransformComponent{PosX: float64(x), PosY: float64(y)})
		randomGopher.AddComponent(&ecsexample.SpriteComponent{Image: sprite})
	}

	spriteRenderSystem := ecsexample.SpriteRenderSystem{Registry: &registry}

	// the game
	example := MultipleImagesExample{spriteRenderSystem: &spriteRenderSystem}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("render multiple images")
	ebiten.RunGame(&example) // omitted error handling
}

func LoadImage(input []byte) (image.Image, error) {
	if img, _, err := image.Decode(bytes.NewReader(input)); err != nil {
		return img, err
	} else {
		return img, nil
	}
}
