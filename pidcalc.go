package pidcalc

import (
	"fmt"

	"github.com/mikesmitty/pidcalc/pkg/noovershoot"
	"github.com/mikesmitty/pidcalc/pkg/pessenintegral"
	"github.com/mikesmitty/pidcalc/pkg/pid"
	"github.com/mikesmitty/pidcalc/pkg/someovershoot"
	"github.com/mikesmitty/pidcalc/pkg/tyreusluyben"
	"github.com/mikesmitty/pidcalc/pkg/zieglernichols"
)

const (
	Classic        = ZieglerNichols
	ZieglerNichols = iota
	PessenIntegral
	SomeOvershoot
	NoOvershoot
	TyreusLuyben
)

func Calculate(ku, tu, kp, ki, kd float64, algorithm int) (float64, float64, float64, error) {
	switch {
	case ku == 0 && tu == 0:
		return kp, ki, kd, nil
	case tu == 0 && ki == 0 && kd == 0:
		return 0, 0, 0, fmt.Errorf("tu cannot be calculated without ki or kd")
	}
	var p pid.PID
	switch algorithm {
	case Classic:
		p = new(zieglernichols.ZieglerNichols)
	case PessenIntegral:
		p = new(pessenintegral.PessenIntegral)
	case SomeOvershoot:
		p = new(someovershoot.SomeOvershoot)
	case NoOvershoot:
		p = new(noovershoot.NoOvershoot)
	case TyreusLuyben:
		p = new(tyreusluyben.TyreusLuyben)
	default:
		return 0, 0, 0, fmt.Errorf("unknown PID algorithm")
	}

	switch {
	case tu == 0:
		switch {
		case ki != 0:
			kp, ki, kd = p.FromKuKi(ku, ki)
		case kd != 0:
			kp, ki, kd = p.FromKuKd(ku, kd)
		}
	case ku == 0:
		switch {
		case kp != 0:
			kp, ki, kd = p.FromTuKp(ku, ki)
		case ki != 0:
			kp, ki, kd = p.FromTuKi(ku, ki)
		case kd != 0:
			kp, ki, kd = p.FromTuKd(ku, ki)
		}
	default:
		kp, ki, kd = p.FromKuTu(ku, tu)
	}
	return kp, ki, kd, nil
}
