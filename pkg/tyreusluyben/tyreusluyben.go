package tyreusluyben

import "github.com/mikesmitty/pidcalc/pkg/pid"

type TyreusLuyben struct {
	Kp, Ki, Kd float64
}

func (z *TyreusLuyben) Reset() {
	z.Kp, z.Ki, z.Kd = 0, 0, 0
}

func (z *TyreusLuyben) Values() (float64, float64, float64) {
	return z.Kp, z.Ki, z.Kd
}

func (z *TyreusLuyben) FromKuTu(ku, tu float64) (float64, float64, float64) {
	z.Kp = ku / 2.2
	z.Ki = ku / 2.2 / 2.2 / tu
	z.Kd = ku * tu / 2.2 / 6.3
	return z.Values()
}

func (z *TyreusLuyben) FromKuKi(ku, ki float64) (float64, float64, float64) {
	tu := ku / 2.2 / 2.2 / ki
	return z.FromKuTu(ku, tu)
}

func (z *TyreusLuyben) FromKuKd(ku, kd float64) (float64, float64, float64) {
	tu := kd / ku * 2.2 * 6.3
	return z.FromKuTu(ku, tu)
}

func (z *TyreusLuyben) FromTuKp(tu, kp float64) (float64, float64, float64) {
	ku := kp * 2.2
	return z.FromKuTu(ku, tu)
}

func (z *TyreusLuyben) FromTuKi(tu, ki float64) (float64, float64, float64) {
	ku := ki * 2.2 * 2.2 * tu
	return z.FromKuTu(ku, tu)
}

func (z *TyreusLuyben) FromTuKd(tu, kd float64) (float64, float64, float64) {
	ku := kd / tu * 2.2 * 6.3
	return z.FromKuTu(ku, tu)
}

var _ pid.PID = &TyreusLuyben{}
