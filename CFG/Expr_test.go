package CFG

import (
    "testing"
)

// TestGetVariablesA calls GetVariablesA with a expression, cheking
// for a valid return value
func TestGetVariablesA(t *testing.T){
    t.Run("Basic case: Variable ", func(t *testing.T) {
        exp  := Variable("x")
        got  := GetVariablesA(exp)
        want := Set{exp:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
    t.Run("Composite case: Sum ", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := Add{x,y}
        got  := GetVariablesA(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
    t.Run("Composite case: Mul ", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := Mul{x,y}
        got  := GetVariablesA(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
    t.Run("Composite case: Sub ", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := Sub{x,y}
        got  := GetVariablesA(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
    t.Run("Composite case: Add(Mul(x,y),1) ", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := Add{Mul{x,y},I(1)}
        got  := GetVariablesA(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
}


// TestGetVariablesB calls GetVariablesB with a expression, cheking
// for a valid return value
func TestGetVariablesB(t *testing.T){
    t.Run("Case: Less than ", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := LessThan{x,y}
        got  := GetVariablesB(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
    t.Run("Case: Equals", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := Equal{x,y}
        got  := GetVariablesB(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
    t.Run("Case: And", func(t *testing.T) {
        x := Variable("x")
        y := Variable("y")
        exp := And{LessThan{x,y},Equal{x,y}}
        got  := GetVariablesB(exp)
        want := Set{x:true,y:true}

        if !Equals(got,want) {
            t.Errorf("Got %v but want %v", got, want)
        }
    })
}

