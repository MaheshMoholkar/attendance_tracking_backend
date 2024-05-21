package types

import "regexp"

const (
	bcryptCost    = 12
	minFirstName  = 2
	minLastName   = 2
	minPassword   = 6
	minClassName  = 2
	exactDivision = 1
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
