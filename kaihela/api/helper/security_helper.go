package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

func DecryptData(data string, encryptKey string) (error, string) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err, nil
	}
	sRawText := string(rawDecodedText)
	iv := sRawText[:16]
	return Ase256Decode(sRawText[16:], encryptKey, iv)
}

func Ase256Encode(plaintext string, key string, iv string, blockSize int) (error, string) {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return err, ""
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return nil, hex.EncodeToString(ciphertext)
}

func Ase256Decode(cipherText string, encKey string, iv string) (err error, decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return err, ""
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return err, ""
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return nil, string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
