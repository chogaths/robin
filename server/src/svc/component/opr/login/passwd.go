package opr_login

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncodePassword(orgpassword string) string {
	enc := sha1.New()
	enc.Write([]byte(orgpassword))
	return hex.EncodeToString(enc.Sum(nil))
}
