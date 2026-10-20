package main

func rotateRunes(s string, shift int) string {
	if s == "" {
		return s
	}
	r := []rune(s)

	l := len(r)
	shift = shift % l
	if shift < 0 {
		shift = l + shift
	}

	return string(append(r[shift:], r[:shift]...))
}
