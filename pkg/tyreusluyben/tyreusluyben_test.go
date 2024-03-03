package tyreusluyben

import (
	"testing"

	"github.com/mikesmitty/pidcalc/testutils"
)

var testValues = []map[string]float64{
	{"ku": 10, "tu": 5, "kp": 4.545454545454545, "ki": 0.4132231404958677, "kd": 3.6075036075036073},
	{"ku": 5, "tu": 30, "kp": 2.2727272727272725, "ki": 0.034435261707988975, "kd": 10.822510822510822},
	{"ku": 50, "tu": 1, "kp": 22.727272727272727, "ki": 10.330578512396693, "kd": 3.6075036075036073},
}

/*
func Test_PrintValues(t *testing.T) {
	p := new(TyreusLuyben)
	testInputs := []struct{ ku, tu float64 }{
		{10, 5},
		{5, 30},
		{50, 1},
	}

	for _, ti := range testInputs {
		kp, ki, kd := p.FromKuTu(ti.ku, ti.tu)
		t.Logf("{\"ku\": %v, \"tu\": %v, \"kp\": %v, \"ki\": %v, \"kd\": %v},\n", ti.ku, ti.tu, kp, ki, kd)
	}
}
*/

func Test_FromKuTu(t *testing.T) {
	z := new(TyreusLuyben)
	testutils.RunTestValues(t, testValues, "FromKuTu", "ku", "tu", z.Reset, z.FromKuTu)
}

func Test_FromKuKi(t *testing.T) {
	z := new(TyreusLuyben)
	testutils.RunTestValues(t, testValues, "FromKuKi", "ku", "ki", z.Reset, z.FromKuKi)
}

func Test_FromKuKd(t *testing.T) {
	z := new(TyreusLuyben)
	testutils.RunTestValues(t, testValues, "FromKuKd", "ku", "kd", z.Reset, z.FromKuKd)
}

func Test_FromTuKp(t *testing.T) {
	z := new(TyreusLuyben)
	testutils.RunTestValues(t, testValues, "FromTuKp", "tu", "kp", z.Reset, z.FromTuKp)
}

func Test_FromTuKi(t *testing.T) {
	z := new(TyreusLuyben)
	testutils.RunTestValues(t, testValues, "FromTuKi", "tu", "ki", z.Reset, z.FromTuKi)
}

func Test_FromTuKd(t *testing.T) {
	z := new(TyreusLuyben)
	testutils.RunTestValues(t, testValues, "FromTuKd", "tu", "kd", z.Reset, z.FromTuKd)
}
