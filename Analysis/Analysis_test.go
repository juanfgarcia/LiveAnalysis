package Analysis

import (
    "testing"
)

func TestAnalysis(t *testing.T) {
    x := Variable("x")
	y := Variable("y")

	d := Assignment{4, nil, x, I(2)}
	c := Assignment{3, nil, x, Sub{x, I(1)}}
	b := Conditional{2, &c, &d, LessThan{y, I(0)}}
	a := Assignment{1, &b, x, I(1)}
	c.next = &b
	cfg := []Block{a, b, c, d}

    Analysis(cfg)
}
