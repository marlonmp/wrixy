package service

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) string {

	buf, _ := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(buf)
}

func matchPassword(hash, password string) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err != nil
}
