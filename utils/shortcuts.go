package utils

func GetNewUUID(save ...bool) string {
	if len(save) > 0 && len(save) < 2 {
		s, _ := Generate()
		return s
	} else {
		s, _ := Generate()
		return s
	}
}
