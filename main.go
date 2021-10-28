package main

import (
    "github.com/juanfgarcia/LiveAnalysis/Analysis"
	"fmt"
)

func main() {
	x := Analysis.Variable("x")
	y := Analysis.Variable("y")

	d := Analysis.Assignment{4, nil, x, I(2)}
	c := Analysis.Assignment{3, nil, x, Sub{x, I(1)}}
	b := Analysis.Conditional{2, &c, &d, LessThan{y, I(0)}}
	a := Analysis.Assignment{1, &b, x, I(1)}
	c.next = &b
	cfg := []Block{a, b, c, d}

	fmt.Println(Analysis.Analysis(cfg))
}
