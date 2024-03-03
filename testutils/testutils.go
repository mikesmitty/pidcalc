package testutils

import (
	"math"
	"testing"
)

func Test_floatEqual(t *testing.T) {
	// Anything smaller than 1e-7 is considered 0
	for _, v := range []float64{0.0, 1e-8, 1e-9} {
		if !FloatEqual(v, 0.0) {
			t.Errorf("floatEqual(%v, 0.0) = false, want true", v)
		}
		if !FloatEqual(-v, 0.0) {
			t.Errorf("floatEqual(%v, 0.0) = false, want true", -v)
		}
	}
	for _, v := range []float64{1e-7, 1e-6, 1e-5, 0.0001, 0.001, 0.01, 0.1, 1.0} {
		if FloatEqual(v, 0.0) {
			t.Errorf("floatEqual(%v, 0.0) = true, want false", v)
		}
		if FloatEqual(-v, 0.0) {
			t.Errorf("floatEqual(%v, 0.0) = true, want false", -v)
		}
	}
}

func FloatEqual(x, y float64) bool {
	return FloatCompare(x, y, 1e-7)
}

func FloatCompare(x, y, diff float64) bool {
	return math.Abs(x-y) < diff
}

// RunTestValues(t, testValues, name, "ku", "ki", p.Reset, p.FromKuKi)
func RunTestValues(t *testing.T, testValues []map[string]float64, name, x, y string, reset func(), f func(float64, float64) (float64, float64, float64)) {
	for _, v := range testValues {
		reset()
		kp, ki, kd := f(v[x], v[y])
		if !FloatEqual(kp, v["kp"]) || !FloatEqual(ki, v["ki"]) || !FloatEqual(kd, v["kd"]) {
			t.Errorf("%s(%v, %v) = %v, %v, %v, want %v, %v, %v", name, v[x], v[y], kp, ki, kd, v["kp"], v["ki"], v["kd"])
		}
	}
}

func AlgoInputValidation(t *testing.T, name string, testValues []map[string]float64, f func(float64, float64, float64, float64, float64) (float64, float64, float64, error)) {
	// We're only testing input validation here so we only need to test one set of values
	v := testValues[0]
	kp, ki, kd, err := f(v["ku"], v["tu"], 0, 0, 0)
	CheckAlgoResult(t, name, v, kp, ki, kd, err)

	_, _, _, err = f(v["ku"], 0, v["kp"], 0, 0)
	if err == nil {
		t.Error("ki and kd cannot be calculated from ku and kp, but received no error")
	}

	kp, ki, kd, err = f(v["ku"], 0, 0, v["ki"], 0)
	CheckAlgoResult(t, name, v, kp, ki, kd, err)

	kp, ki, kd, err = f(v["ku"], 0, 0, 0, v["kd"])
	CheckAlgoResult(t, name, v, kp, ki, kd, err)

	kp, ki, kd, err = f(0, v["tu"], v["kp"], 0, 0)
	CheckAlgoResult(t, name, v, kp, ki, kd, err)

	kp, ki, kd, err = f(0, v["tu"], 0, v["ki"], 0)
	CheckAlgoResult(t, name, v, kp, ki, kd, err)

	kp, ki, kd, err = f(0, v["tu"], 0, 0, v["kd"])
	CheckAlgoResult(t, name, v, kp, ki, kd, err)
}

func CheckAlgoResult(t *testing.T, name string, v map[string]float64, kp, ki, kd float64, err error) {
	if !FloatEqual(kp, v["kp"]) || !FloatEqual(ki, v["ki"]) || !FloatEqual(kd, v["kd"]) || err != nil {
		t.Errorf("%s(%v, %v, 0, 0, 0) = %v, %v, %v, %v, want %v, %v, %v, nil", name, v["ku"], v["tu"], kp, ki, kd, err, v["kp"], v["ki"], v["kd"])
	}
}
