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
	http.HandleFunc("/primes", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	primespiral(w, 5000)
}

func primespiral(out io.Writer, max int) {
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
			x := size + r * math.Sin(theta) * float64(size) / float64(max)
			y := size + r  * math.Cos(theta) * float64(size) / float64(max)
			dc.DrawCircle(x, y, 3)
			dc.SetRGB(0, 0, 0)
			dc.Fill()
			p = g.Next()
	}
	img := dc.Image()
	png.Encode(out, img)
}
