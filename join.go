package piscine

func Join(strs []string, sep string) string {
	concat := strs[0]
	for _, val := range strs[1:] {
		concat = concat + sep + val
	}
	return concat
}
