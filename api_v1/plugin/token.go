package plugin

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func MakeToken(id int, SecretKey string) string {
	SID := strconv.Itoa(id)
	str := SecretKey + SID
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}