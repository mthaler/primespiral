package main

import (
	"github.com/fogleman/gg"
	"github.com/jbarham/primegen"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	primespiral(w)
}

func primespiral(out io.Writer) {
	const (
		size = 1000
	)
	dc := gg.NewContext(2 * size + 1, 2 * size + 1)
	dc.DrawRectangle(0, 0, 2 * size + 1, 2 * size + 1)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	g := primegen.New()
	p := g.Next()
	for p < 100000 {
			r := float64(p)
		    theta := float64(p)
			x := size + r * math.Sin(theta) / 100.0
			y := size + r  * math.Cos(theta) / 100.0
			dc.DrawCircle(x, y, 3)
			dc.SetRGB(0, 0, 0)
			dc.Fill()
			p = g.Next()
	}
	img := dc.Image()
	png.Encode(out, img)
}
