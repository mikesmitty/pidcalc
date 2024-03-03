package pessenintegral

import "github.com/mikesmitty/pidcalc/pkg/pid"

type PessenIntegral struct {
	Kp, Ki, Kd float64
}

func (z *PessenIntegral) Reset() {
	z.Kp, z.Ki, z.Kd = 0, 0, 0
}

func (z *PessenIntegral) Values() (float64, float64, float64) {
	return z.Kp, z.Ki, z.Kd
}

func (z *PessenIntegral) FromKuTu(ku, tu float64) (float64, float64, float64) {
	z.Kp = 0.7 * ku
	z.Ki = 1.75 * ku / tu
	z.Kd = 0.105 * ku * tu
	return z.Values()
}

func (z *PessenIntegral) FromKuKi(ku, ki float64) (float64, float64, float64) {
	tu := 1.75 * ku / ki
	return z.FromKuTu(ku, tu)
}

func (z *PessenIntegral) FromKuKd(ku, kd float64) (float64, float64, float64) {
	tu := kd / 0.105 / ku
	return z.FromKuTu(ku, tu)
}

func (z *PessenIntegral) FromTuKp(tu, kp float64) (float64, float64, float64) {
	ku := kp / 0.7
	return z.FromKuTu(ku, tu)
}

func (z *PessenIntegral) FromTuKi(tu, ki float64) (float64, float64, float64) {
	ku := tu * ki / 1.75
	return z.FromKuTu(ku, tu)
}

func (z *PessenIntegral) FromTuKd(tu, kd float64) (float64, float64, float64) {
	ku := kd / 0.105 / tu
	return z.FromKuTu(ku, tu)
}

var _ pid.PID = &PessenIntegral{}
