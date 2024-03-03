package pid

type PID interface {
	FromKuTu(ku, tu float64) (float64, float64, float64)
	FromKuKi(ku, ki float64) (float64, float64, float64)
	FromKuKd(ku, kd float64) (float64, float64, float64)
	FromTuKp(tu, kp float64) (float64, float64, float64)
	FromTuKi(tu, ki float64) (float64, float64, float64)
	FromTuKd(tu, kd float64) (float64, float64, float64)
	Reset()
	Values() (float64, float64, float64)
}
