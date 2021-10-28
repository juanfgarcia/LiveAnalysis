package Analysis

import (
    "github.com/juanfgarcia/LiveAnalysis/CFG"
)

func Analysis(cfg []CFG.Block) ([]Set, []Set) {
	LVOut := make([]CFG.Set, len(cfg))
	LVIn := make([]CFG.Set, len(cfg))
	OLVIn := make([]CFG.Set, len(cfg))
	OLVOut := make([]CFG.Set, len(cfg))

	for _, n := range cfg {
		LVOut[n.GetLabel()-1] = Set{}
		LVIn[n.GetLabel()-1] = Set{}
	}

	for {
		copy(OLVIn, LVIn)
		copy(OLVOut, LVOut)

		for _, n := range cfg {
			LVOut[n.GetLabel()-1] = n.LVOut(LVIn)
			LVIn[n.GetLabel()-1] = n.LVIn(OLVOut[n.GetLabel()-1])
		}
		fixed := true
		for _, n := range cfg {
			if !Equals(OLVIn[n.GetLabel()-1], LVIn[n.GetLabel()-1]) || !Equals(OLVOut[n.GetLabel()-1], LVOut[n.GetLabel()-1]) {
				fixed = false
				break
			}
		}
		if fixed {
			break
		}
	}
	return LVOut, LVIn
}
