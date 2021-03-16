package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"time"
)
type StringUtil struct {

}
var keyString = "a very very very  mediazen.co.kr"// 32 bytes로 만들어 주세요

func (sutil StringUtil)Encrypt(instr string) (string, error) {
	key := []byte(keyString)
	now := time.Now()
	strnow := now.Format("20060102")

	plaintext := []byte( instr+ "|"+strnow)
	log.Printf("Encrypt %s\n", plaintext)
	ciphertext, err := encrypt(key, plaintext)
	log.Printf("Encrypt %0x\n", ciphertext)
	if err != nil {
		return "", err
	}
	encodedStr := base64.StdEncoding.EncodeToString(ciphertext)
	log.Printf("Encrypt base 64 : %s\n", encodedStr)
	return encodedStr, nil
}
/**
 decrypt 후 string 반환
 */
func (sutil StringUtil)Decrypt(ciphertext string) (string, error) {
	log.Printf("Decrypt base 64 : %s\n", ciphertext)
	decodedStrAsByteSlice, err := base64.StdEncoding.DecodeString(ciphertext)
	log.Printf("Decrypt %0x\n", decodedStrAsByteSlice)
	key := []byte(keyString)
	result, err := decrypt(key, decodedStrAsByteSlice)
	if err != nil {
		return "", err
	}
	log.Printf("Decrypt string %s\n", result)
	return string(result), nil
}


func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}