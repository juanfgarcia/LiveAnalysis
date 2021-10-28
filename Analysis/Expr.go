package Analysis

type AExp interface {
    isAExpr()
}

type I int
func (i I) isAExpr() {}

type Variable string
func (v Variable) isAExpr() {}

type Add struct { left AExp; right AExp }
func (a Add) isAExpr() {}

type Mul struct { left AExp; right AExp}
func (m Mul) isAExpr() {}

type Sub struct { left AExp; right AExp}
func (s Sub) isAExpr() {}

func GetVariablesA (e AExp) Set {
    switch exp := e.(type) {
        case I: return nil
        case Variable:
            return Set{exp: true}
        case Add:
            return Union(GetVariablesA(exp.left), GetVariablesA(exp.right))
        case Sub:
            return Union(GetVariablesA(exp.left), GetVariablesA(exp.right))
        case Mul:
            return Union(GetVariablesA(exp.left), GetVariablesA(exp.right))
    }
    return nil
}

type BExp interface {
   isBExp()
}

type LessThan struct {left AExp; right AExp}
func (lt LessThan) isBExp() {}

type Equal struct {left AExp; right AExp}
func (e Equal) isBExp() {}

type And struct {left BExp; right BExp}
func (a And) isBExp() {}

func GetVariablesB (b BExp) Set {
    switch exp := b.(type) {
        case LessThan:
            return Union(GetVariablesA(exp.left), GetVariablesA(exp.right))
        case Equal:
            return Union(GetVariablesA(exp.left), GetVariablesA(exp.right))
        case And:
            return Union(GetVariablesB(exp.left), GetVariablesB(exp.right))
    }
    return nil
}


