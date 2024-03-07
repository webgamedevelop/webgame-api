package password

import "golang.org/x/crypto/bcrypt"

var cost = bcrypt.DefaultCost

func Compare(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func Generate(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}
