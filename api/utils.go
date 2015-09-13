package api

func in(elm string, array []string) bool {
	for _, e := range array {
		if e == elm {
			return true
		}
	}
	return false
}
