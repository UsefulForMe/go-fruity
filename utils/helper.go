package utils

import (
	"io/ioutil"
	"mime/multipart"
	"strings"
)

func InternationPhoneToNational(phone string) string {
	if strings.HasPrefix(phone, "+") {
		return strings.Replace(phone, "+84", "0", 1)
	}
	return phone
}

func MultipartFileToByte(multipartFile multipart.FileHeader) ([]byte, error) {
	fileContent, _ := multipartFile.Open()

	byteContainer, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return nil, err
	}

	return byteContainer, nil
}
