package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// EncryptAESHex mã hóa dữ liệu (password, token, ...) bằng key hex1024 (128 bytes)
func EncryptAESHex(plainText string, hexKey string) (string, error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return "", err
	}
	if len(key) != 128 {
		return "", errors.New("key must be 1024 bits (128 bytes) in hex")
	}
	block, err := aes.NewCipher(key[:32]) // AES-256 chỉ dùng 32 bytes đầu
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptAESHex giải mã dữ liệu (password, token, ...) bằng key hex1024 (128 bytes)
func DecryptAESHex(encrypted string, hexKey string) (string, error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return "", err
	}
	if len(key) != 128 {
		return "", errors.New("key must be 1024 bits (128 bytes) in hex")
	}
	ciphertext, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key[:32]) // AES-256 chỉ dùng 32 bytes đầu
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("ciphertext too short")
	}
	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
} 