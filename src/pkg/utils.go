package pkg

func CheckLang(lang string) bool {
	return lang == "th" || lang == "en"
}

func IsTH(lang string) bool {
	return lang == "th"
}

func IsEN(lang string) bool {
	return lang == "en"
}

func IsNullReturnString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
