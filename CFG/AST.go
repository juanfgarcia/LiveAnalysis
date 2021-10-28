package CFG

import (
	"fmt"
)

type Block interface {
	GetLabel() int
	LVIn(LVOut Set) Set
	LVOut(LVIn []Set) Set
}

type Assignment struct {
	label int
	next  Block
	lhs   Variable
	rhs   AExp
}

func (a Assignment) String() string {
	return fmt.Sprintf("Block %d\n%v := %v", a.label, a.lhs, a.rhs)
}

func (a Assignment) GetLabel() int {
	return a.label
}

func (a Assignment) LVIn(LVOut Set) Set {
	Gen := GetVariablesA(a.rhs)
	Kill := Set{a.lhs: true}
	return Union(Gen, Subtract(LVOut, Kill))
}

func (a Assignment) LVOut(LVIn []Set) Set {
	if a.next != nil {
		return LVIn[a.next.GetLabel()-1]
	}
	return nil
}

type Conditional struct {
	label     int
	yes       Block
	no        Block
	condition BExp
}

func (c Conditional) String() string {
	return fmt.Sprintf("Block %d\n%v", c.label, c.condition)
}

func (c Conditional) LVIn(LVOut Set) Set {
	Gen := GetVariablesB(c.condition)
	return Union(LVOut, Gen)
}

func (c Conditional) GetLabel() int {
	return c.label
}

func (c Conditional) LVOut(LVIn []Set) Set {
	return Union(LVIn[c.yes.GetLabel()-1], LVIn[c.no.GetLabel()-1])
}
