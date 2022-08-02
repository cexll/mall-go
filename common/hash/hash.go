package hash

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) (string, error) {
	bts, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bts), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
