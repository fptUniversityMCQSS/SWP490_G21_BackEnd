package ultity

func CheckUsername(username string) bool {

	return true
}

func CheckPassword(password string) bool {
	if password == "" {
		return false
	}
	return true
}

func CheckRole(role string) bool {
	if role == "admin" {
		return true
	} else if role == "staff" {
		return true
	} else if role == "user" {
		return true
	} else {
		return false
	}
}
