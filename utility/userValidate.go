package utility

import (
	"regexp"
	"strings"
)

func CheckUsername(username string) bool {
	IsUserName := regexForAll("(\\w){8,}", username)
	return IsUserName
}

func CheckPassword(password string) bool {
	IsPassword := regexForAll("[^(.\\s)]{8,}", password)
	return IsPassword
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
	IsEmail := regexForAll("\\w+@\\w+\\.\\w+(\\.\\w+)?", email)
	return IsEmail
}
func CheckPhone(phone string) bool {
	IsPhone := regexForAll("(\\d){10}", phone)
	return IsPhone
}

func CheckFullName(fullname string) bool {
	isFullName := regexForAll("(\\w){8,30}", fullname)
	return isFullName
}

func regexForAll(regex string, param string) bool {
	ModifyRg := regexp.MustCompile(regex)
	if len(ModifyRg.FindAllString(param, -1)) > 1 {
		return false
	} else {
		if strings.Contains(param, " ") || len(ModifyRg.FindString(param)) < len(param) {
			return false
		} else {
			return true
		}
	}
}
