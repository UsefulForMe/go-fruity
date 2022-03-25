package utils

import "strings"

func InternationPhoneToNational(phone string) string {
	if strings.HasPrefix(phone, "+") {
		return strings.Replace(phone, "+84", "0", 1)
	}
	return phone
}
