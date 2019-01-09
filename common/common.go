package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

func GetRootDir() string {
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = "."
	}
	return file
}

func WritePidToFile(filename string) error {
	return ioutil.WriteFile(GetRootDir()+"/var/"+filename+".pid", []byte(strconv.Itoa(os.Getpid())), 0666)
}

func RemovePidFile(filename string) error {
	return os.Remove(GetRootDir() + "/var/" + filename + ".pid")
}

func GetFileMd5(picPath string) (string, error) {

	file, err := os.Open(picPath)
	if err != nil {
		return "", err
	}
	md5h := md5.New()
	_, err = io.Copy(md5h, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
