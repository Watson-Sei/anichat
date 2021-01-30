package plugin

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func MakeToken(id int, secretkey string) string {
	SID := strconv.Itoa(id)
	str := secretkey + SID
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}