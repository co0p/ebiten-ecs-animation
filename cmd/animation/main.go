package main

import (
	"bytes"
	ecsexample "ebiten-ecs-animation"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:generate file2byteslice -package=main -input=spritesheet.png -output=spritesheet.go -var=Spritesheet

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
	images, _ := LoadSpritesheet(Spritesheet, 4, 250, 260)

	// entities
	registry := ecsexample.Registry{}

	e := registry.NewEntity()
	e.AddComponent(&ecsexample.SpriteComponent{Image: images[0]})
	e.AddComponent(&ecsexample.TransformComponent{PosX: 100, PosY: 200})

	spriteRenderSystem := ecsexample.SpriteRenderSystem{Registry: &registry}

	// the game
	example := SimpleImageExample{spriteRenderSystem: &spriteRenderSystem}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("render single image")
	ebiten.RunGame(&example) // omitted error handling
}

// LoadSpritesheet returns n sub images from the given input
func LoadSpritesheet(input []byte, n int, width int, height int) ([]*ebiten.Image, error) {
	sprites := []*ebiten.Image{}

	spritesheet, _, _ := image.Decode(bytes.NewReader(input))
	ebitenImage := ebiten.NewImageFromImage(spritesheet)

	for i := 0; i <= n; i++ {
		r := image.Rect(i*width, 0, (i+1)*width, height)
		sprite := ebitenImage.SubImage(r).(*ebiten.Image)
		sprites = append(sprites, sprite)
	}

	return sprites, nil
}
