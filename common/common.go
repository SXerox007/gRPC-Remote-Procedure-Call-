package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetMd5(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func GetSha1String(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UserPwdEncrypt(password string, salts ...string) string {
	salt := "BL@cK_"
	if len(salts) > 0 {
		salt = salts[0]
	}
	return GetSha1String(GetSha1String(string(GetMd5(password))+salt) + GetMd5String(salt))
}

func HmacSha1ToString(k, v string) string {
	key := []byte(v)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(k))
	return hex.EncodeToString(mac.Sum(nil))
}
