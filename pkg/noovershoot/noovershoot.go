package noovershoot

import "github.com/mikesmitty/pidcalc/pkg/pid"

type NoOvershoot struct {
	Kp, Ki, Kd float64
}

func (z *NoOvershoot) Reset() {
	z.Kp, z.Ki, z.Kd = 0, 0, 0
}

func (z *NoOvershoot) Values() (float64, float64, float64) {
	return z.Kp, z.Ki, z.Kd
}

func (z *NoOvershoot) FromKuTu(ku, tu float64) (float64, float64, float64) {
	z.Kp = 0.2 * ku
	z.Ki = 0.4 * ku / tu
	z.Kd = 0.2 / 3 * ku * tu
	return z.Values()
}

func (z *NoOvershoot) FromKuKi(ku, ki float64) (float64, float64, float64) {
	tu := 0.4 * ku / ki
	return z.FromKuTu(ku, tu)
}

func (z *NoOvershoot) FromKuKd(ku, kd float64) (float64, float64, float64) {
	tu := kd * 15 / ku
	return z.FromKuTu(ku, tu)
}

func (z *NoOvershoot) FromTuKp(tu, kp float64) (float64, float64, float64) {
	ku := kp / 0.2
	return z.FromKuTu(ku, tu)
}

func (z *NoOvershoot) FromTuKi(tu, ki float64) (float64, float64, float64) {
	ku := ki * tu / 0.4
	return z.FromKuTu(ku, tu)
}

func (z *NoOvershoot) FromTuKd(tu, kd float64) (float64, float64, float64) {
	ku := kd * 15 / tu
	return z.FromKuTu(ku, tu)
}

var _ pid.PID = &NoOvershoot{}
