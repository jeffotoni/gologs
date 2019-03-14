// Back-End in Go server
// @jeffotoni
// 2019-01-09

package crypt

//
//
//
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	// "strings"
)

var (
	HASH_SALT      = "jeff102912..#golang@#455x848484"
	SHA1_SALT      = "doglas#145$1000.gojefgo#4love.x"
	letters        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	UKK_KEY_CIPHER = []byte("LURYRYXEUDU393CKXKKS3KXUCE847302")
)

func createHash(key string) []byte {
	hasher := md5.New()
	io.WriteString(hasher, key)
	hash := hasher.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash)
	// base 64
	return dst
}

func Decrypt(text, passphrase string) (strText string) {
	strText = Decode64String(text)
	data := []byte(strText)
	key := createHash(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Decrypt:: ", err)
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Cipher:: ", err)
		return
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println("Open: ", err)
		return
	}

	return string(plaintext)
}

func Encrypt(text, passphrase string) (cryptStr string) {
	data := []byte(text)
	block, _ := aes.NewCipher(createHash(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Encrypt:: ", err)
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		log.Println(err)
		return
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	cryptStr = Encode64Byte(ciphertext)
	return cryptStr
}

func Encode64String(content string) string {
	if len(content) > 0 {
		return base64.StdEncoding.EncodeToString([]byte(content))
	}
	return ""
}

func Encode64Byte(content []byte) string {
	if len(string(content)) > 0 {
		return base64.StdEncoding.EncodeToString(content)
	}
	return ""
}

func Decode64String(encoded string) string {
	if len(encoded) > 0 {
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			log.Println("base64 decode:: ", err)
			return ""
		}
		return (string(decoded))
	}
	return ""
}
