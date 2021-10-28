package main

import (
    "github.com/juanfgarcia/LiveAnalysis/CFG"
    "github.com/juanfgarcia/LiveAnalysis/Analysis"
	"fmt"
)

func main() {
	x := CFG.Variable("x")
	y := CFG.Variable("y")

	d := CFG.Assignment{4, nil, x, I(2)}
	c := CFG.Assignment{3, nil, x, Sub{x, I(1)}}
	b := CFG.Conditional{2, &c, &d, LessThan{y, I(0)}}
	a := CFG.Assignment{1, &b, x, I(1)}
	c.next = &b
	cfg := []Block{a, b, c, d}

	fmt.Println(Analysis.Analysis(cfg))
}
