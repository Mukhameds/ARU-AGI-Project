package main

import (
	"fmt"
	"image/color"
	
	"math/cmplx"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
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
	fmt.Println("Создаём график ζ(0.5 + i·t) с подписями нулей...")

	nMax := 1000
	step := 0.1

	var imRange []float64
	var values []float64

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

	var minima plotter.XYs
	var labelData plotter.XYLabels

	for i := 1; i < len(values)-1; i++ {
		if values[i] < values[i-1] && values[i] < values[i+1] && values[i] < 1.0 {
			x := imRange[i]
			y := values[i]
			minima = append(minima, plotter.XY{X: x, Y: y})
			labelData.XYs = append(labelData.XYs, plotter.XY{X: x, Y: y + 1})
			labelData.Labels = append(labelData.Labels, "t ≈ "+strconv.FormatFloat(x, 'f', 2, 64))
		}
	}

	p := plot.New()
	p.Title.Text = "|ζ(0.5 + it)| с минимумами и подписями"
	p.X.Label.Text = "Im(s)"
	p.Y.Label.Text = "|ζ(s)|"

	err := plotutil.AddLinePoints(p, "ζ", pts)
	if err != nil {
		panic(err)
	}

	scatter, _ := plotter.NewScatter(minima)
	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	scatter.GlyphStyle.Radius = vg.Points(3)

	labels, err := plotter.NewLabels(labelData)
	if err != nil {
		panic(err)
	}

	p.Add(scatter, labels)
	p.Legend.Add("нули", scatter)

	err = p.Save(16*vg.Inch, 7*vg.Inch, "zeta_annotated.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("Готово! Файл сохранён: zeta_annotated.png")
}
