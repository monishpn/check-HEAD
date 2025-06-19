package main

import (
	"math"
	"testing"
)

type rec struct {
	width, height float64
	area          float64
	cost          float64
}

type circ struct {
	radius float64
	area   float64
	cost   float64
}

var rect = []rec{
	{10, 20, 200, 400},
	{12, 22, 264, 528},
}

var cir = []circ{
	{10, 314.15, 471.23},
	{8, 201.06, 301.59},
}

func TestRectangle_Area(t *testing.T) {
	for _, r := range rect {
		r_test := Rectangle{r.width, r.height}
		op := r_test.Area()
		if op != r.area {
			t.Errorf("RECTANGLE AREA [Width: %v || Height %v]: Expected %f, got %f", r.width, r.height, r.area, op)
		}
	}
}

func TestCircle_Area(t *testing.T) {
	for _, r := range cir {
		c_test := Circle{r.radius}
		op := c_test.Area()
		if 1e-2 < (math.Max(r.area, op) - math.Min(r.area, op)) {
			t.Errorf("CIRCLE AREA [Radius: %v]: Expected %f, got %f", r.radius, r.area, op)

		}
	}
}

func Test_Cost(t *testing.T) {

	for _, r := range rect {
		r_test := Rectangle{r.width, r.height}
		op := cost(r_test)

		if op != r.cost {
			t.Errorf("RECTANGLE COST [Width: %v || Height %v]: Expected cost of %v, but got %v", r.width, r.height, r.cost, op)
		}
	}

	for _, r := range cir {
		c_test := Circle{r.radius}
		op := cost(c_test)
		if 1e-2 < (math.Max(r.cost, op) - math.Min(r.cost, op)) {
			t.Errorf("CIRCLE AREA [Radius: %v]: Expected %f, got %f", r.radius, r.cost, op)
		}
	}

}
