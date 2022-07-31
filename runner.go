package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed images/runner.png
	imgFile []byte
)

type runnerState uint

const (
	runnerStateIdle runnerState = iota
	runnerStateMoving
)

// Runner
type Runner struct {
	posX int
	posY int

	count int

	state runnerState
	image *ebiten.Image
}

func NewRunner() *Runner {
	image, _, err := image.Decode(bytes.NewReader(imgFile))
	if err != nil {
		log.Fatal(err)
	}

	return &Runner{
		state: runnerStateIdle,
		image: ebiten.NewImageFromImage(image),
	}
}

func (r *Runner) Update() {
	r.count++
}

func (r *Runner) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}

	// 將人物圖片的中心點移動到 screen 的 (0, 0) 位置
	opt.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)

	// 將人物圖片的中心點移動到 screen 中心位置
	opt.GeoM.Translate(screenWidth/2, screenHeight/2)

	// runnerImage 的大小為 256 * 96
	// 跑步動作的部分有 8 張，每張大小為 32 * 32

	// 用 i 來控制何時要換張
	i := (r.count / 5) % frameCount
	// 用 sx, sy 來控制目前要取 8 張中的第幾張圖片
	// (sx, sy 為跑步動作圖片在 runnerImage 中的實際位置)
	sx, sy := frameOX+i*frameWidth, frameOY

	// 跑步動作在 runnerImage 中的區塊位置及大小
	portionImage := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)

	// 從 runnerImage 中根據取出 portionImage 的資訊取出需要的圖片
	runner := r.image.SubImage(portionImage)
	img := ebiten.NewImageFromImage(runner)

	screen.DrawImage(img, opt)
}

func (r *Runner) idleIMage() *ebiten.Image {
	return nil
}
