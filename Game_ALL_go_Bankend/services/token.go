package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-256-bit-secret") // Nên để trong biến môi trường

type Claims struct {
	UserID    uint      `json:"user_id"`
	ExpiresAt int64     `json:"exp"`
	jwt.RegisteredClaims
}

// GenerateToken tạo JWT token cho userID, thời gian sống (seconds)
func GenerateToken(userID uint, durationSeconds int64) (string, error) {
	expiresAt := time.Now().Add(time.Duration(durationSeconds) * time.Second).Unix()
	claims := &Claims{
		UserID:    userID,
		ExpiresAt: expiresAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiresAt, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken xác thực JWT token, trả về userID nếu hợp lệ
func VerifyToken(tokenString string) (uint, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return 0, errors.New("token expired")
	}
	return claims.UserID, nil
}

// EncryptToken mã hóa token bằng key hex (AES-256-GCM)
func EncryptToken(token string, hexKey string) (string, error) {
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
	ciphertext := gcm.Seal(nonce, nonce, []byte(token), nil)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptToken giải mã token bằng key hex (AES-256-GCM)
func DecryptToken(encrypted string, hexKey string) (string, error) {
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