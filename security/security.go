package security

import "golang.org/x/crypto/bcrypt"

/*
Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
See http://www.usenix.org/event/usenix99/provos/provos.pdf
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
