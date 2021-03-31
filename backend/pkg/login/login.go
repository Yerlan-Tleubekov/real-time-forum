package login

func CheckLogin(login string) bool {

	if len(login) == 0 {
		return false
	}

	for _, value := range login {

		if (value < 'a' || value > 'z') &&
			(value < 'A' || value > 'Z') &&
			(value < '0' || value > '9') &&
			value != '-' &&
			value != '_' {
			return false
		}
	}

	return true
}
