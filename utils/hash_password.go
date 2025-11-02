package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hashed []byte, err error) {
	hashed, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return
}
