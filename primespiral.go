package main

import (
	"github.com/jbarham/primegen"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math"
	"os"
)

func main() {
	primespiral(os.Stdout)
}

func primespiral(out io.Writer) {
	const (
		size = 1000
	)
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
	g := primegen.New()
	p := g.Next()
	for p < 100000 {
			r := float64(p)
		    theta := float64(p)
			x := int(size + r * math.Sin(theta) / 100.0 + 0.5)
			y := int(size + r  * math.Cos(theta) / 100.0 + 0.5)
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
			p = g.Next()
	}
	png.Encode(out, img)
}
