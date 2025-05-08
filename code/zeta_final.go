package main

import (
	"fmt"
	"image/color"
	
	"math/cmplx"
	"os"
	

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
	fmt.Println("Создание финального графика ζ(0.5 + i·t)...")

	nMax := 1000
	step := 0.1

	var imRange []float64
	var values []float64
	var zeros []string
	var minima plotter.XYs

	for t := -50.0; t <= 50.0; t += step {
		s := complex(0.5, t)
		z := zetaApprox(s, nMax)
		abs := cmplx.Abs(z)
		imRange = append(imRange, t)
		values = append(values, abs)

		if len(values) >= 3 {
			i := len(values) - 2
			if values[i] < values[i-1] && values[i] < values[i+1] && values[i] < 1.0 {
				minima = append(minima, plotter.XY{X: t, Y: abs})
				zeros = append(zeros, fmt.Sprintf("t ≈ %.3f |ζ(s)| ≈ %.4f", t, abs))
			}
		}
	}

	// Сохраняем нули в файл
	f, err := os.Create("zeta_zeros.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, z := range zeros {
		f.WriteString(z + "\n")
	}
	fmt.Println("Список нулей сохранён: zeta_zeros.txt")

	// Построение графика
	p := plot.New()
	p.Title.Text = "|ζ(0.5 + it)| и отмеченные нули"
	p.X.Label.Text = "Im(s)"
	p.Y.Label.Text = "|ζ(s)|"

	pts := make(plotter.XYs, len(imRange))
	for i := range imRange {
		pts[i].X = imRange[i]
		pts[i].Y = values[i]
	}

	err = plotutil.AddLinePoints(p, "ζ", pts)
	if err != nil {
		panic(err)
	}

	scatter, _ := plotter.NewScatter(minima)
	scatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	scatter.GlyphStyle.Radius = vg.Points(3)

	p.Add(scatter)
	p.Legend.Add("нули", scatter)

	err = p.Save(16*vg.Inch, 7*vg.Inch, "zeta_final.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("Готово: zeta_final.png и zeta_zeros.txt")
}
