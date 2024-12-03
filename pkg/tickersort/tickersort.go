package tickersort

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func TabOrSpace(r rune) bool {
	return r=='\t' || r==' '
}
