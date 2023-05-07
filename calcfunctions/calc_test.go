package calcfunctions

import (
	"testing"
)

var equation = []struct {
	in  string
	out float64
}{
	{"1+1", 2},
	{"12/2", 6},
	{"10/4-5", -2.5},
	{"5+5*5", 30},
	{"12/3*4-5/4", 14.75},
	{"144/(128-116)*12", 144},
	{"5*(5+5)", 50},
}

func TestCalc(t *testing.T) {
	for _, tt := range equation {
		t.Run(tt.in, func(t *testing.T) {
			if got, _ := EvalRPN(ConvertToRPN(tt.in)); got != tt.out {
				t.Error("want: ", tt.out, " got: ", tt.out)
			}
		})
	}
}
