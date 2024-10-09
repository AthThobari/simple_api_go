package token

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() string {
	b := make([]byte, 18)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}