package main

import (
	"encoding/hex"
	"crypto/sha256"
)


func HashString(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	md := hash.Sum(nil)
	hashedStr := hex.EncodeToString(md)

	return hashedStr
}
