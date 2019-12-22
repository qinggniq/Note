package ast

type UnitType int

const (
	SingalType = iota
	TimesType
	TimesCombType
	SingalCombType
)

type Unit struct {
	Times int
	Other *Unit
	Mine  *Unit
	Cur   string
	Ty    UnitType
}

func (u *Unit) Eval() string {
	var res string
	switch u.Ty {
	case SingalType:
		res = u.Cur
	case TimesType:
		for i := 0; i < u.Times; i++ {
			res += u.Mine.Eval()
		}
	case SingalCombType:
		res = u.Cur
		res += u.Other.Eval()
	case TimesCombType:
		for i := 0; i < u.Times; i++ {
			res += u.Mine.Eval()
		}
		res += u.Other.Eval()
	}
	return res
}
