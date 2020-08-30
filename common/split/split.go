package split

import "strings"

var sep = map[rune]bool{
	' ':  true,
	'\n': true,
	',':  true,
	';':  true,
	'\t': true,
	'\f': true,
	'\v': true,
	'\r': true,
}

func Split(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return sep[r]
	})
}
