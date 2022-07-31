package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenHeight = 320
	screenWidth  = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

type Game struct {
	runner *Runner
}

// Update is called 60 times per second
func (g *Game) Update() error {
	g.runner.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.runner.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS  %0.2f\nTPS %0.2f\n", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")

	if err := ebiten.RunGame(&Game{
		runner: NewRunner(),
	}); err != nil {
		log.Fatal(err)
	}
}
