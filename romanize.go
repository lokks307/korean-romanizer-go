package korom

import "fmt"

var VowelMap = map[rune]string{
	'\u1161': "a",
	'\u1165': "eo",
	'\u1169': "o",
	'\u116E': "u",
	'\u1173': "eu",
	'\u1175': "i",
	'\u1162': "ae",
	'\u1166': "e",
	'\u116C': "oe",
	'\u1171': "wi",

	'\u1163': "ya",
	'\u1167': "yeo",
	'\u116D': "yo",
	'\u1172': "yu",
	'\u1164': "yae",
	'\u1168': "ye",
	'\u116A': "wa",
	'\u116B': "wae",
	'\u116F': "wo",
	'\u1170': "we",
	'\u1174': "ui",
}

var OnsetMap = map[rune]string{
	'\u0020': " ",
	'\u1100': "g",
	'\u1101': "kk",
	'\u110F': "k",
	'\u1103': "d",
	'\u1104': "tt",
	'\u1110': "t",
	'\u1107': "b",
	'\u1108': "pp",
	'\u1111': "p",
	'\u110C': "j",
	'\u110D': "jj",
	'\u110E': "ch",
	'\u1109': "s",
	'\u110A': "ss",
	'\u1112': "h",
	'\u1102': "n",
	'\u1106': "m",
	'\u1105': "r",
	'\u110B': "",
}

var CodaMap = map[rune]string{
	'\u11A8': "k",
	'\u11AE': "t",
	'\u11B8': "p",
	'\u11AB': "n",
	'\u11BC': "ng",
	'\u11B7': "m",
	'\u11AF': "l",
	0:        "",
}

var RomanizeDebug bool

func Romanize(in string) string {
	if in == "" {
		return ""
	}

	pronon := NewPronouncer(in)
	if pronon == nil {
		return ""
	}

	ret := ""
	for _, each := range pronon.PronouncedList {
		if RomanizeDebug {
			fmt.Printf("[%s] %X + %X + %X\n", string(each.Char), each.Initial, each.Medial, each.Final)
		}

		if conv, ok := OnsetMap[each.Initial]; ok {
			if conv == "r" && ret[len(ret)-1] == 'l' {
				conv = "l"
			}

			ret += conv
		}

		if conv, ok := VowelMap[each.Medial]; ok {
			ret += conv
		}

		if conv, ok := CodaMap[each.Final]; ok {
			ret += conv
		}
	}

	return ret
}
