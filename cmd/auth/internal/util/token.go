package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// функция генерации токена
func GenerateToken() string {

	randomNumber := rand.Int()

	randomString := fmt.Sprintf("%d", randomNumber)

	hash := md5.New()
	hash.Write([]byte(randomString))
	hashValue := hex.EncodeToString(hash.Sum(nil))

	return hashValue
}
