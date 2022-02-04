package verbalid

import (
	"log"
	"strings"
)

const idChars = "3479ACDEFHJKMNPRTUWXY"
const charlen = len(idChars)

func IntToVerbalID(i int) string {
	rv := ""
	for i > 0 {
		r := i / charlen
		d := i - (r * charlen)
		rv = string(idChars[d]) + rv
		i = r
	}
	return rv
}

func VerbalIDToInt(s string) int {
	s = strings.ToUpper(s)
	var rv int
	pos := 0
	for i := len(s) - 1; i >= 0; i-- {
		var char int = -1
		for c := 0; c < charlen; c++ {
			if s[i] == idChars[c] {
				char = c
				break
			}
		}
		if char == -1 {
			log.Panicf("Invalid character %s", string(s[i]))
		}
		rv += char * intPow(charlen, pos)
		pos++
	}
	return rv
}

func intPow(x, y int) int {
	rv := 1
	for i := 0; i < y; i++ {
		rv *= x
	}
	return rv
}
