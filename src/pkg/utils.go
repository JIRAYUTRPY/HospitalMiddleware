package pkg

import "golang.org/x/crypto/bcrypt"

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

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
