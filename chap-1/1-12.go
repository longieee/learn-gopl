// Server4 prints an animation to the browser
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// Default values for lissajous animation
var cycles = 5
var res = 0.001
var size = 100
var nframes = 64
var delay = 8

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Printf("URL.Path = %q\n", r.URL.Path)
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("Header[%q] = %q\n", k, v)
	}
	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// Naive way to handle the query parameters for lissajous
	for k, v := range r.Form {
		if k == "cycles" {
			v_int, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println(err)
			} else {
				cycles = v_int
			}
		} else if k == "res" {
			v_float, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				fmt.Println(err)
			} else {
				res = v_float
			}
		} else if k == "size" {
			v_int, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println(err)
			} else {
				size = v_int
			}
		} else if k == "nframes" {
			v_int, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println(err)
			} else {
				nframes = v_int
			}
		} else if k == "delay" {
			v_int, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println(err)
			} else {
				delay = v_int
			}
		} else {
			fmt.Printf("Form[%q] = %q unrecognized for lissajous animation parameter. Skipping.\n", k, v)
		}
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	fmt.Printf("Lissajous parameters:\ncycles: %d\nres: %f\nsize: %d\nnframes: %d\ndelay: %d\n", cycles, res, size, nframes, delay)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
