package utility

import (
	"regexp"
)

func CheckUsername(username string) bool {
	re, _ := regexp.MatchString("^(\\w){8,30}$", username)
	return re
}

func CheckPassword(password string) bool {
	re, _ := regexp.MatchString("^\\S{8,30}$", password)
	return re
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
func CheckEmail(email string) bool {
	re, _ := regexp.MatchString("^\\w+@\\w+\\.\\w+(\\.\\w+)?$", email)
	return re
}
func CheckPhone(phone string) bool {
	re, _ := regexp.MatchString("^(\\d){10}$", phone)
	return re
}

func CheckFullName(fullName string) bool {
	re, _ := regexp.MatchString("^(.)+${8,50}", fullName)
	return re
}
