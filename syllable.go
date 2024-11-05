package korom

import (
	"fmt"
)

type Syllable struct {
	Char    rune
	Initial rune
	Medial  rune
	Final   rune
}

func (m Syllable) Print() {
	fmt.Printf("%s['%s'+'%s'+'%s']\n", string(m.Char), string(m.Initial), string(m.Medial), string(m.Final))
	fmt.Printf("%X %X %X\n", m.Initial, m.Medial, m.Final)
}

func NewSyllable(in rune) *Syllable {
	if in < 'ㄱ' || in > '힣' {
		return nil
	}

	if in < '가' {
		if in >= '\u1100' && in <= '\u1112' {
			return &Syllable{
				Initial: in,
			}
		}

		if in >= '\u1161' && in <= '\u1175' {
			return &Syllable{
				Medial: in,
			}
		}

		if in >= '\u11A8' && in <= '\u11C2' {
			return &Syllable{
				Final: in,
			}
		}
	}

	calc := in - '\uAC00'

	final := rune(0)
	if (calc % 28) > 0 {
		final = '\u11A7' + (calc % 28)
	}

	calc = (calc / 28)
	medial := '\u1161' + (calc % 21)
	initial := '\u1100' + (calc / 21)

	return &Syllable{
		Char:    in,
		Initial: initial,
		Medial:  medial,
		Final:   final,
	}
}
