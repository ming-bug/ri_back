package utils

import (
	"crypto/sha512"
	"fmt"
)

func HashPassword(password string) string {
	sum512 := sha512.Sum512([]byte(password))
	return fmt.Sprintf("%x", sum512)
}
