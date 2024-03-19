package main

import (
	"fmt"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func testH1(max_n int) []float64 {
	results := make([]float64, 0, max_n)
	for i := 3; i < max_n; i++ {
		result1, _ := getCircuit(i)
		//result2, _ := getCircuit(i + 1)
		results = append(results, result1)
	}
	return results
}

func getCircuit(n int) (float64, bool) {
	if n < 3 {
		return 0.0, true
	} else {
		angle := 2 * math.Pi / float64(n)
		w := mat.NewDense(1, 2, []float64{math.Cos(angle) - 1, math.Sin(angle)})
		circuit := w.Norm(2)
		m := mat.NewDense(2, 2, []float64{math.Cos(angle),
			-math.Sin(angle),
			math.Sin(angle),
			math.Cos(angle)})
		for i := int(1); i < n-1; i++ {
			w.Mul(w, m)
			circuit += w.Norm(2)
		}
		return circuit, false
	}
}

func getPoint() [2]float64 {
	xsig := rand.Int()%2*2 - 1
	ysig := rand.Int()%2*2 - 1
	x := rand.Float64() * float64(xsig)
	y := rand.Float64() * float64(ysig)
	return [2]float64{x, y}
}

func monteCarlo(n int) (float64, bool) {
	points := make([][2]float64, 0, n)
	for i := 0; i < n; i++ {
		points = append(points, getPoint())
	}

}

func main() {
	results := testH1(30)
	_ = results
	n := 3
	if circuit, err := getCircuit(n); err {
		fmt.Println("got err")
	} else {
		fmt.Println(circuit)
		fmt.Println(2 * math.Pi)
	}
}
