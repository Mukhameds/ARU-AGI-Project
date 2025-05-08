package main

import (
	"fmt"
	"math/cmplx"
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

)

func zetaApprox(s complex128, nMax int) complex128 {
	var sum complex128
	for n := 1; n <= nMax; n++ {
		ns := cmplx.Pow(complex(float64(n), 0), s)
		sum += 1 / ns
	}
	return sum
}

func computeZetaModulus(imRange []float64, re float64, nMax int) plotter.XYs {
	points := make(plotter.XYs, len(imRange))
	for i, im := range imRange {
		s := complex(re, im)
		z := zetaApprox(s, nMax)
		points[i].X = im
		points[i].Y = cmplx.Abs(z)
	}
	return points
}

func main() {
	nMax := 1000
	imRange := make([]float64, 0)
	for im := -50.0; im <= 50.0; im += 1.0 {
		imRange = append(imRange, im)
	}

	reValues := []float64{0.3, 0.5, 0.7}
	p := plot.New()
	p.Title.Text = "|ζ(s)| при разных Re(s)"
	p.X.Label.Text = "Im(s)"
	p.Y.Label.Text = "|ζ(s)|"

	colors := []string{"red", "blue", "green"}

	for i, re := range reValues {
		pts := computeZetaModulus(imRange, re, nMax)
		line, err := plotter.NewLine(pts)
		if err != nil {
			panic(err)
		}
		line.Color = plotutilColor(colors[i])
		p.Add(line)
		p.Legend.Add(fmt.Sprintf("Re(s)=%.1f", re), line)
	}

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "zeta_field.png"); err != nil {
		panic(err)
	}

	fmt.Println("График сохранён в файл zeta_field.png")
}

// Простая функция цвета
func plotutilColor(name string) color.Color {
	switch name {
	case "red":
		return color.RGBA{R: 255, A: 255}
	case "blue":
		return color.RGBA{B: 255, A: 255}
	case "green":
		return color.RGBA{G: 255, A: 255}
	default:
		return color.Black
	}
}
