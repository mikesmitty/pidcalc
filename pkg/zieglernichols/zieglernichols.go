package zieglernichols

import "github.com/mikesmitty/pidcalc/pkg/pid"

type ZieglerNichols struct {
	Kp, Ki, Kd float64
}

func (z *ZieglerNichols) Reset() {
	z.Kp, z.Ki, z.Kd = 0, 0, 0
}

func (z *ZieglerNichols) Values() (float64, float64, float64) {
	return z.Kp, z.Ki, z.Kd
}

func (z *ZieglerNichols) FromKuTu(ku, tu float64) (float64, float64, float64) {
	z.Kp = ku / 1.7
	z.Ki = 2 / 1.7 * ku / tu
	z.Kd = ku * tu / 1.7 / 8
	return z.Values()
}

func (z *ZieglerNichols) FromKuKi(ku, ki float64) (float64, float64, float64) {
	tu := 2 / 1.7 * ku / ki
	return z.FromKuTu(ku, tu)
}

func (z *ZieglerNichols) FromKuKd(ku, kd float64) (float64, float64, float64) {
	tu := kd / ku * 1.7 * 8
	return z.FromKuTu(ku, tu)
}

func (z *ZieglerNichols) FromTuKp(tu, kp float64) (float64, float64, float64) {
	ku := kp * 1.7
	return z.FromKuTu(ku, tu)
}

func (z *ZieglerNichols) FromTuKi(tu, ki float64) (float64, float64, float64) {
	ku := ki * tu * 1.7 / 2
	return z.FromKuTu(ku, tu)
}

func (z *ZieglerNichols) FromTuKd(tu, kd float64) (float64, float64, float64) {
	ku := kd * 1.7 * 8 / tu
	return z.FromKuTu(ku, tu)
}

var _ pid.PID = &ZieglerNichols{}
