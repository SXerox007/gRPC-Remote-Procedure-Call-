// simple forensic image file use stlib Golang
// test gif, png, jpg

// package main

// import (
// 	"fmt"
// 	"image"
// 	_ "image/gif"
// 	_ "image/jpeg"
// 	_ "image/png"
// 	"os"
// )

// func main() {
// 	if len(os.Args) == 0 {
// 		fmt.Printf("Use: %v fileUnknown fileUnknown fileUnknown ...\n", os.Args[0])
// 	} else {
// 		for _, file := range os.Args[1:] {
// 			f, err := os.Open(file)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "file: %v\n", err)
// 				continue
// 			}
// 			defer f.Close()

// 			_, kind, err := image.Decode(f)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "%v: %v\n", f.Name(), err)
// 				continue
// 			}
// 			fmt.Fprintln(os.Stderr, f.Name(), "Format =", kind)
// 		}
// 	}
// }

// Example 2
// Schedular example in golang

package main

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 16
const MINUTE_TO_TICK int = 47
const SECOND_TO_TICK int = 00

type jobTicker struct {
	t *time.Timer
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
	fmt.Println("new tick here")
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
	fmt.Println("next tick here")
	jt.t.Reset(getNextTickDuration())
}

func main() {
	jt := NewJobTicker()
	for {
		<-jt.t.C
		fmt.Println(time.Now(), "- just ticked")
		jt.updateJobTicker()
	}
}

// // Example-3
// // AES

// package main

// import (
// 	"bytes"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/x509"
// 	"encoding/base64"
// 	"encoding/pem"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	// input := []byte("this is a test")
// 	// iv := []byte("532b6195636c6127")[:aes.BlockSize]
// 	// key := []byte("532b6195636c61279a010000")

// 	// fmt.Println("Input:     ", input)
// 	// fmt.Println("Key:       ", key)
// 	// fmt.Println("IV:        ", iv)

// 	// encrypted := make([]byte, len(input))
// 	// EncryptAES(encrypted, input, key, iv)

// 	// fmt.Println("Output:    ", encrypted)

// 	var test string = `"{"name":"John", "age":32, "city":"New Yorkerere"}"`
// 	key := make([]byte, 8)
// 	_, err := rand.Read(key)
// 	if err != nil {
// 		// handle error here
// 	}
// 	fmt.Println(Encrypt(fmt.Sprintf("%x", key), test))

// 	//LoadX509KeyPair("/Users/sumitthakur/Downloads/m2psolutions_pub.der", "/Users/sumitthakur/Downloads/m2psolutions_pub.der")

// 	//GetRSAPublicKeyFromCertificate()
// }

// func EncryptAES(dst, src, key, iv []byte) error {
// 	aesBlockEncryptor, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		return err
// 	}
// 	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
// 	aesEncrypter.XORKeyStream(dst, src)
// 	return nil
// }

// func DecryptAES(dst, src, key, iv []byte) error {
// 	aesBlockEncryptor, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		return err
// 	}
// 	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
// 	aesEncrypter.XORKeyStream(dst, src)
// 	return nil
// }

// func Encrypt(keyStr string, textStr string) string {
// 	key := []byte(keyStr)
// 	plaintext := PKCS5Padding([]byte(textStr))

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// The IV needs to be unique, but not secure. Therefore it's common to
// 	// include it at the beginning of the ciphertext.
// 	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
// 	iv := ciphertext[:aes.BlockSize]
// 	iv = key[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		panic(err)
// 	}
// 	fmt.Sprintf("%x", iv)

// 	encrypter := cipher.NewCBCEncrypter(block, iv)
// 	out := make([]byte, len(plaintext))

// 	encrypter.CryptBlocks(out, plaintext)
// 	//fmt.Println("#Encrypt the data using AES 256, mode CBC, Padding PKCS7 with the MD5 as the `key` (\""+textStr+"\", \""+keyStr+"\"):\n\t", base64.URLEncoding.EncodeToString(out),"\n")

// 	// convert to base64
// 	return base64.URLEncoding.EncodeToString(out)
// }

// /**
//  *	PKCS7补码
//  *	这里可以参考下http://blog.studygolang.com/167.html
//  */
// func PKCS7Padding(data []byte) []byte {
// 	blockSize := 16
// 	padding := blockSize - len(data)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(data, padtext...)
// }

// func PKCS5Padding(data []byte) []byte {
// 	blockSize := 8
// 	padding := blockSize - len(data)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(data, padtext...)
// }

// // load the x509
// func LoadX509KeyPair(certFile, keyFile string) (*x509.Certificate, *rsa.PrivateKey) {
// 	cf, e := ioutil.ReadFile(certFile)
// 	if e != nil {
// 		fmt.Println("cfload:", e.Error())
// 		os.Exit(1)
// 	}

// 	kf, e := ioutil.ReadFile(keyFile)
// 	if e != nil {
// 		fmt.Println("kfload:", e.Error())
// 		os.Exit(1)
// 	}
// 	cpb, cr := pem.Decode(cf)
// 	fmt.Println("Pem Decode: ", string(cr))
// 	kpb, kr := pem.Decode(kf)
// 	fmt.Println("KeyFile Decode: ", string(kr))
// 	crt, e := x509.ParseCertificate(cpb.Bytes)

// 	if e != nil {
// 		fmt.Println("parsex509:", e.Error())
// 		os.Exit(1)
// 	}
// 	key, e := x509.ParsePKCS1PrivateKey(kpb.Bytes)
// 	if e != nil {
// 		fmt.Println("parsekey:", e.Error())
// 		os.Exit(1)
// 	}
// 	return crt, key
// }
