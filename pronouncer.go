package korom

import (
	"fmt"

	"github.com/lokks307/go-util/mt"
)

type Pronouncer struct {
	OrginList      []*Syllable
	PronouncedList []*Syllable
}

var DOUBLE_CONSONANT_FINAL = []rune{
	'\u11AA', // ㄳ
	'\u11AC', // ㄵ
	'\u11AD', // ㄶ
	'\u11B0', // ㄺ
	'\u11B1', // ㄻ
	'\u11B2', // ㄼ
	'\u11B3', // ㄽ
	'\u11B4', // ㄾ
	'\u11B5', // ㄿ
	'\u11B6', // ㅀ
	'\u11B9', // ㅄ
	'\u11BB', // ㅆ
}

func NewPronouncer(text string) *Pronouncer {
	proc := Pronouncer{}

	runed := []rune(text)
	for _, eachChar := range runed {
		out := NewSyllable(eachChar)

		if out == nil {
			out = &Syllable{
				Char:    ' ',
				Initial: ' ',
				Medial:  0,
				Final:   0,
			}
		}

		proc.OrginList = append(proc.OrginList, out)
	}

	for i, each := range proc.OrginList {
		isFinalSyllable := ((i + 1) == len(proc.OrginList))
		finalBeforeC := false
		// finalBeforeV := false

		if !isFinalSyllable {
			finalBeforeC = (each.Final > 0 && !mt.IsIn(proc.OrginList[i+1].Initial, 0, '\u110B'))
			// finalBeforeV = (each.Final > 0 && mt.IsIn(proc.OrginList[i+1].Initial, 0, '\u110B'))
		}

		copied := Syllable{
			Char:    each.Char,
			Initial: each.Initial,
			Medial:  each.Medial,
			Final:   each.Final,
		}

		if isFinalSyllable || finalBeforeC {
			switch copied.Final {
			case '\u11A9', '\u11BF', '\u11AA', '\u11B0': // 종성 'ㄲ','ㅋ','ㄳ','ㄺ' -> 'ㄱ'
				copied.Final = '\u11A8'
			case '\u11BA', '\u11BB', '\u11BD', '\u11BE', '\u11C0': // 종성 'ㅅ','ㅆ','ㅈ','ㅊ','ㅌ' -> 'ㄷ'
				copied.Final = '\u11AE'
			case '\u11C1', '\u11B9', '\u11B5': // 종성 'ㅍ','ㅄ','ㄿ' -> 'ㅂ'
				copied.Final = '\u11B8'
			case '\u11AC': // 종성 'ㄵ' -> 'ㄴ'
				copied.Final = '\u11AB'
			case '\u11B2', '\u11B3', '\u11B4': // 종성 'ㄼ','ㄽ','ㄾ' -> 'ㄹ'
				copied.Final = '\u11AF'
			case '\u11B1': // 종성 'ㄻ' -> 'ㅁ'
				copied.Final = '\u11B7'
			}
		}

		if mt.IsIn(each.Final, '\u11C2', '\u11AD', '\u11B6') { // 'ㅎ','ㄶ','ㅀ'
			if isFinalSyllable && each.Final == '\u11C2' {
				copied.Final = rune(0)
			} else {
				// 종성 'ㄶ' -> 'ㄴ'
				if each.Final == '\u11AD' {
					copied.Final = '\u11AB'
				}

				// 종성 'ㅀ' -> 'ㄹ'
				if each.Final == '\u11B6' {
					copied.Final = '\u11AF'

					// // 종성 'ㅀ' + 다음 초성 'ㄹ' -> 종성 'ㄹ'
					// if proc.OrginList[i+1].Initial == '\u1105' {
					// 	copied.Final = '\u11AF'
					// 	// 다음 초성 'ㄱ','ㄷ','ㅅ','ㅈ','ㅇ' -> 종성 '
					// } else if mt.IsIn(proc.OrginList[i+1].Initial, '\u1100', '\u1103', '\u1109', '\u110B', '\u110C') {
					// 	copied.Final = '\u11AF'
					// }
				}

				// 종성 'ㅎ'
				if each.Final == '\u11C2' {
					// 종성 'ㅎ' + 다음 초성 'ㄴ' -> 종성 'ㄴ'
					if proc.OrginList[i+1].Initial == '\u1102' {
						copied.Final = '\u11AB'
					} else {
						copied.Final = rune(0)
					}
				}
			}
		}

		if i-1 >= 0 {
			// 앞 음절의 종성이 'ㅎ','ㄶ','ㅀ'일 때
			if mt.IsIn(proc.OrginList[i-1].Final, '\u11C2', '\u11AD', '\u11B6') {
				// 현 음절 초성 'ㄱ','ㄷ','ㅈ','ㅅ' -> 'ㅋ','ㅌ','ㅊ','ㅆ'
				switch each.Initial {
				case '\u1100': // 'ㄱ'
					copied.Initial = '\u110F' // 'ㅋ'
				case '\u1103': // 'ㄷ'
					copied.Initial = '\u1110' // 'ㅌ'
				case '\u110C': // 'ㅈ'
					copied.Initial = '\u110E'
				case '\u1109':
					copied.Initial = '\u110A'
				}
			}

			if each.Initial == '\u110B' {
				if mt.IsIn(proc.OrginList[i-1].Final, DOUBLE_CONSONANT_FINAL...) {
					switch proc.OrginList[i-1].Final {
					case '\u11AA', '\u11B9', '\u11BB': // ㄳ,ㅄ,ㅆ
						copied.Initial = '\u110A' // ㅆ
					case '\u11AC': // ㄵ
						fmt.Println("222 char=", string(copied.Char))
						copied.Initial = '\u110C' // ㅈ
					case '\u11AD': // ㄶ, ㅀ
						copied.Initial = '\u1105'
						proc.PronouncedList[i-1].Final = rune(0)
						// copied.Initial = '\u1112' // ㅎ
					case '\u11B6': // ㅀ
						copied.Initial = '\u1105'
						proc.PronouncedList[i-1].Final = rune(0)
					case '\u11B0': // ㄺ
						copied.Initial = '\u1100' // ㄱ
					case '\u11B1': // ㄻ
						copied.Initial = '\u1106' // ㅁ
					case '\u11B2': // ㄼ
						copied.Initial = '\u1107' // ㅂ
					case '\u11B3': // ㄽ
						copied.Initial = '\u110A' // ㅆ
					case '\u11B4': // ㄾ
						copied.Initial = '\u1110' // ㅌ
					case '\u11B5': // ㄿ
						copied.Initial = '\u1111' // ㅍ
					}
				} else if proc.OrginList[i-1].Final > 0 && proc.OrginList[i-1].Final != '\u11BC' {
					fmt.Println("333 char=", string(copied.Char))

					switch proc.OrginList[i-1].Final {
					case '\u11A8':
						copied.Initial = '\u1100'
					case '\u11A9':
						copied.Initial = '\u1101'
					case '\u11AB':
						copied.Initial = '\u1102'
					case '\u11AE':
						copied.Initial = '\u1103'
					case '\u11AF':
						copied.Initial = '\u1105'
					case '\u11B7':
						copied.Initial = '\u1106'
					case '\u11B8':
						copied.Initial = '\u1107'
					case '\u11BA':
						copied.Initial = '\u1109'
					case '\u11BB':
						copied.Initial = '\u110A'
					case '\u11BD':
						copied.Initial = '\u110C'
					case '\u11BE':
						copied.Initial = '\u110E'
					case '\u11BF':
						copied.Initial = '\u110F'
					case '\u11C0':
						copied.Initial = '\u1110'
					case '\u11C1':
						copied.Initial = '\u1111'
					}

					proc.PronouncedList[i-1].Final = rune(0)
				}
			}
		}

		// 종성 겹받침 + 다음 음절 초성 'ㅇ' ->
		if mt.IsIn(each.Final, DOUBLE_CONSONANT_FINAL...) && proc.OrginList[i+1].Initial == '\u110B' {
			switch each.Final {
			case '\u11AA': // ㄳ
				copied.Final = '\u11A8' // ㄱ
			case '\u11AC', '\u11AD': // ㄵ,ㄶ
				copied.Final = '\u11AB' // ㄴ
			case '\u11B0', '\u11B1', '\u11B2', '\u11B3', '\u11B4', '\u11B5', '\u11B6': // ㄺ,ㄻ,ㄼ,ㄽ,ㄾ,ㄿ,ㅀ
				copied.Final = '\u11AF' // ㄹ
			case '\u11B9': // ㅄ
				fmt.Println("111 char=", string(copied.Char))
				copied.Final = '\u11B8' // ㅂ
			case '\u11BB': // ㅆ
				copied.Final = '\u11BA' // ㅅ
			}
		}

		// if !isFinalSyllable && finalBeforeV {
		// 	if each.Final != '\u11BC' {
		// 		fmt.Println("222 char=", string(copied.Char))
		// 		copied.Final = rune(0)
		// 	}
		// }

		if !isFinalSyllable {
			// 종성 'ㄴ' + 다음 음절 초성 'ㄹ' -> 종성 'ㄹ'
			if each.Final == '\u11AB' && proc.OrginList[i+1].Initial == '\u1105' {
				copied.Final = '\u11AF'
			}
		}

		proc.PronouncedList = append(proc.PronouncedList, &copied)
	}

	return &proc
}
