package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"math"
	"os"
	"time"
)

type point struct {
	x int
	y int
}

func lower_color(colour color.RGBA, alpha float64) color.RGBA {
	return color.RGBA{
		uint8(float64(colour.R) * alpha),
		uint8(float64(colour.G) * alpha),
		uint8(float64(colour.B) * alpha),
		colour.A}
}

func dist(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2.0) +
		math.Pow(float64(p1.y-p2.y), 2.0))
}

func avg(p1, p2, p3 point) point{
	avgx := int((p1.x + p2.x + p3.x)/3)
	avgy := int((p1.y + p2.y + p3.y)/3)
	return point{avgx, avgy}
}

func mix(c1, c2 color.RGBA) color.RGBA {
	return color.RGBA{
		uint8(float64(c1.R + (c2.R - c1.R)) * alpha),
		uint8(float64(c1.G + (c2.G - c1.G)) * alpha),
		uint8(float64(c1.B + (c2.B - c1.B)) * alpha),
		255
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	im := image.NewRGBA(image.Rect(0, 0, 800, 600))
	var (
		points []point = []point{
			point{rand.Intn(800), rand.Intn(600)},
			point{rand.Intn(800), rand.Intn(600)},
			point{rand.Intn(800), rand.Intn(600)}}
		colors []color.RGBA = []color.RGBA{
			color.RGBA{255, 0, 0, 255},
			color.RGBA{0, 255, 0, 255},
			color.RGBA{0, 0, 255, 255}}
		x int
		y int
		centre point = avg(points[0], points[1], points[2])
	)
	origin := point{243, 235}
	for x = 0; x < 800; x++ {
		for y = 0; y < 600; y++ {
			im.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	iters := 1000000
	for iters > 0 {
		num := rand.Int() % 3
		origin.x = (origin.x + points[num].x) / 2
		origin.y = (origin.y + points[num].y) / 2
		alpha := make([]float64, 3, 3)
		for i:= 0; i<3; i++ {
			alpha[i] = dist(origin, centre)/dist(points[i], centre)
			if alpha[i] > 1 {
				alpha[i] = 0.0
			}
		}
		pix_color := mix(
			lower_color(colors[0], alpha[0]),
			lower_color(colors[1], alpha[1]),
			lower_color(colors[2], alpha[2]))
		im.Set(origin.x, origin.y, pix_color)
		iters--
	}
	output, err := os.Create("trig.png")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer output.Close()
	err = png.Encode(output, im)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}
