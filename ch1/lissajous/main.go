// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Modified by Filipp Zapolskikh

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
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

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//!+main

var palette = [1 + shades]color.Color{color.Black}

const (
	backgroundIndex = 0
	foregroundIndex = 1
	shades          = 16
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			lissajous(w, r.Form)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, nil)
}

func intParam(form url.Values, name string, defaultValue int) int {
	if value, ok := form[name]; ok {
		if num, err := strconv.Atoi(value[0]); err == nil {
			return num
		}
	}
	return defaultValue
}

func lissajous(out io.Writer, form url.Values) {
	var (
		cycles  = intParam(form, "cycles", 5)   // number of complete x oscillator revolutions
		res = 0.001                                             // angular resolution
		size    = intParam(form, "size", 100)   // image canvas covers [-size..+size]
		nframes = intParam(form, "nframes", 64) // number of animation frames
		delay   = intParam(form, "delay", 8)    // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < shades; i++ {
		palette[i+1] = color.RGBA{0x00, uint8(0x0f + i*0x20), 0x00, 0xff}
		palette[shades-i] = palette[i+1]
	}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette[:])
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				uint8(1+i%shades))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
