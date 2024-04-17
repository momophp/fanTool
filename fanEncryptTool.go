package fanTool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// GenerateEncryptKey 生成加密密钥
func GenerateEncryptKey() string {
	encryptKey := make([]byte, 32)
	_, _ = rand.Read(encryptKey)
	return base64.StdEncoding.EncodeToString(encryptKey)
}

// Encrypt 加密
func Encrypt(data, key string) (string, error) {
	KeyBt, _ := base64.StdEncoding.DecodeString(key)
	block, err := aes.NewCipher(KeyBt)
	if err != nil {
		return "", err
	}
	//补码
	dataByte := []byte(data)
	blockSize := block.BlockSize()
	padding := blockSize - len(dataByte)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	dataByte = append(dataByte, padText...)
	ciphertext := make([]byte, aes.BlockSize+len(dataByte))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext[aes.BlockSize:], dataByte)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密
func Decrypt(data, key string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	KeyBt, _ := base64.StdEncoding.DecodeString(key)
	block, err := aes.NewCipher(KeyBt)
	if err != nil {
		return "", err
	}
	iv := decoded[:aes.BlockSize]
	ciphertext := decoded[aes.BlockSize:]
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(ciphertext, ciphertext)
	//去除补码
	length := len(ciphertext)
	unPadding := int(ciphertext[length-1])
	plaintext := ciphertext[:(length - unPadding)]
	return string(plaintext), nil
}
