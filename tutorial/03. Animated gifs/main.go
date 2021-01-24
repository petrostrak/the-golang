package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palete = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palete
	blackIndex = 1 // next color in palete
)

var green = color.RGBA{uint8(11), uint8(156), uint8(49), uint8(1)}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 60    // number of animation frames
		delay   = 8     // delay between fra,es in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase differance

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palete)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Cos(t)
			y := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	lissajous(os.Stdout)
}
