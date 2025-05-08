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

func main() {
	fmt.Println("Ищем нули ζ(0.5 + i·t)...")

	nMax := 1000
	step := 0.1
	var (
		imRange []float64
		values  []float64
	)

	for t := -50.0; t <= 50.0; t += step {
		s := complex(0.5, t)
		z := zetaApprox(s, nMax)
		imRange = append(imRange, t)
		values = append(values, cmplx.Abs(z))
	}

	pts := make(plotter.XYs, len(imRange))
	for i := range imRange {
		pts[i].X = imRange[i]
		pts[i].Y = values[i]
	}

	// Поиск минимумов вручную (локальные впадины)
	var minima plotter.XYs
	for i := 1; i < len(values)-1; i++ {
		if values[i] < values[i-1] && values[i] < values[i+1] && values[i] < 1.0 {
			minima = append(minima, plotter.XY{X: imRange[i], Y: values[i]})
		}
	}

	// Построение графика
	p := plot.New()
	p.Title.Text = "|ζ(0.5 + it)| и нули"
	p.X.Label.Text = "Im(s)"
	p.Y.Label.Text = "|ζ(s)|"

	line, _ := plotter.NewLine(pts)
	line.Color = color.RGBA{B: 255, A: 255}

	minDots, _ := plotter.NewScatter(minima)
	minDots.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	minDots.GlyphStyle.Radius = vg.Points(3)

	p.Add(line, minDots)
	p.Legend.Add("ζ", line)
	p.Legend.Add("нули", minDots)

	if err := p.Save(14*vg.Inch, 6*vg.Inch, "zeta_zeros_marked.png"); err != nil {
		panic(err)
	}

	fmt.Println("График сохранён: zeta_zeros_marked.png")
}
