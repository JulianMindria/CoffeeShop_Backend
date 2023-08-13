package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func VerifyPassword(hashedPass, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err
}
