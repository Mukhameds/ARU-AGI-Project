package main

import (
	"fmt"
	"math/cmplx"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func zetaApprox(s complex128, nMax int) complex128 {
	var sum complex128
	for n := 1; n <= nMax; n++ {
		ns := cmplx.Pow(complex(float64(n), 0), s)
		sum += 1 / ns
	}
	return sum
}

func main() {
	fmt.Println("Генерируем ζ(0.5 + i·t)...")

	nMax := 1000
	imStart := -50.0
	imEnd := 50.0
	step := 0.1

	var points plotter.XYs
	for im := imStart; im <= imEnd; im += step {
		s := complex(0.5, im)
		z := zetaApprox(s, nMax)
		points = append(points, plotter.XY{
			X: im,
			Y: cmplx.Abs(z),
		})
	}

	// Построение графика
	p := plot.New()
	p.Title.Text = "|ζ(0.5 + i·t)| — критическая прямая"
	p.X.Label.Text = "Im(s)"
	p.Y.Label.Text = "|ζ(s)|"
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{B: 255, A: 255} // Синий цвет

	p.Add(line)
	p.Legend.Add("ζ(0.5 + it)", line)

	err = p.Save(12*vg.Inch, 5*vg.Inch, "zeta_zeros.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("График сохранён: zeta_zeros.png")
}
