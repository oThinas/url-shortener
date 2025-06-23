package store

import "math/rand/v2"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode() string {
	const length = 8
	randomBytes := make([]byte, length)
	for i := range length {
		randomBytes[i] = charset[rand.IntN(len(charset))]
	}

	return string(randomBytes)
}
