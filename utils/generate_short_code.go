package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func Generate_Short_Code() (string, error){
	short_code := make([]byte, 6)
	_, err := rand.Read(short_code)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(short_code), nil
}