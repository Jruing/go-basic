package core

import "math/rand"

func GenerateRandomString() string {
	var chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 16)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
