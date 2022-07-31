package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
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

var (
	runnerImage *ebiten.Image

	//go:embed images/runner.png
	imgFile []byte
)

type Game struct {
	count int
}

// Update is called 60 times per second
func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}

	// 將人物圖片的中心點移動到 screen 的 (0, 0) 位置
	opt.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)

	// 將人物圖片的中心點移動到 screen 中心位置
	opt.GeoM.Translate(screenWidth/2, screenHeight/2)

	// runnerImage 的大小為 256 * 96
	// 跑步動作的部分有 8 張，每張大小為 32 * 32

	// 用 i 來控制何時要換張
	i := (g.count / 5) % frameCount
	// 用 sx, sy 來控制目前要取 8 張中的第幾張圖片
	// (sx, sy 為跑步動作圖片在 runnerImage 中的實際位置)
	sx, sy := frameOX+i*frameWidth, frameOY

	// 跑步動作在 runnerImage 中的區塊位置及大小
	portionImage := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)

	// 從 runnerImage 中根據取出 portionImage 的資訊取出需要的圖片
	runner := runnerImage.SubImage(portionImage)
	img := ebiten.NewImageFromImage(runner)

	screen.DrawImage(img, opt)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS  %0.2f\nTPS %0.2f\n", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(imgFile))
	if err != nil {
		log.Fatal(err)
	}

	// init runnerImage
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
