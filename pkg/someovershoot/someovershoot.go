package someovershoot

import "github.com/mikesmitty/pidcalc/pkg/pid"

type SomeOvershoot struct {
	Kp, Ki, Kd float64
}

func (z *SomeOvershoot) Reset() {
	z.Kp, z.Ki, z.Kd = 0, 0, 0
}

func (z *SomeOvershoot) Values() (float64, float64, float64) {
	return z.Kp, z.Ki, z.Kd
}

func (z *SomeOvershoot) FromKuTu(ku, tu float64) (float64, float64, float64) {
	z.Kp = 1.0 / 3.0 * ku
	z.Ki = 2.0 / 3.0 * ku / tu
	z.Kd = 1.0 / 9.0 * ku * tu
	return z.Values()
}

func (z *SomeOvershoot) FromKuKi(ku, ki float64) (float64, float64, float64) {
	tu := 2.0 / 3.0 * ku / ki
	return z.FromKuTu(ku, tu)
}

func (z *SomeOvershoot) FromKuKd(ku, kd float64) (float64, float64, float64) {
	tu := kd * 9.0 / ku
	return z.FromKuTu(ku, tu)
}

func (z *SomeOvershoot) FromTuKp(tu, kp float64) (float64, float64, float64) {
	ku := kp * 3.0
	return z.FromKuTu(ku, tu)
}

func (z *SomeOvershoot) FromTuKi(tu, ki float64) (float64, float64, float64) {
	ku := ki * tu * 1.5
	return z.FromKuTu(ku, tu)
}

func (z *SomeOvershoot) FromTuKd(tu, kd float64) (float64, float64, float64) {
	ku := kd * 9.0 / tu
	return z.FromKuTu(ku, tu)
}

var _ pid.PID = &SomeOvershoot{}
