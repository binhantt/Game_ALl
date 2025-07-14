package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"sync"
	"time"
)

var (
	emailCodeMap = make(map[string]string)
	emailCodeExpiry = make(map[string]time.Time)
	mu sync.Mutex
)

const (
	codeLength = 6
	codeExpiry = 10 * time.Minute
)

// GenerateVerificationCode sinh mã xác thực ngẫu nhiên
func GenerateVerificationCode() string {
	max := big.NewInt(10)
	code := ""
	for i := 0; i < codeLength; i++ {
		n, _ := rand.Int(rand.Reader, max)
		code += fmt.Sprintf("%d", n.Int64())
	}
	return code
}

// SendVerificationCode gửi mã xác thực tới email và lưu vào map
func SendVerificationCode(email string) error {
	code := GenerateVerificationCode()
	mu.Lock()
	emailCodeMap[email] = code
	emailCodeExpiry[email] = time.Now().Add(codeExpiry)
	mu.Unlock()

	// Gửi email (cấu hình SMTP ở đây, demo chỉ in ra log)
	// Thực tế: dùng smtp.SendMail hoặc thư viện khác
	fmt.Printf("[DEMO] Send code %s to email %s\n", code, email)
	return nil
}

// VerifyCode kiểm tra mã xác thực
func VerifyCode(email, code string) bool {
	mu.Lock()
	defer mu.Unlock()
	c, ok := emailCodeMap[email]
	exp, expOk := emailCodeExpiry[email]
	if !ok || !expOk || time.Now().After(exp) {
		return false
	}
	if c != code {
		return false
	}
	// Xác thực xong thì xóa code
	delete(emailCodeMap, email)
	delete(emailCodeExpiry, email)
	return true
} 