package email

import "regexp"

// detected login value its nickname or email
func IsLoginAnEmail(login string) bool {
	for _, value := range login {
		if value == '@' {
			return true
		}
	}

	return false
}

func CheckValidEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	emailLength := len(email)

	if emailLength < 3 || emailLength > 254 {
		return false
	}

	emailMatchResult := emailRegex.MatchString(email)

	return emailMatchResult
}
