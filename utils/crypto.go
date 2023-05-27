package utils

import "golang.org/x/crypto/bcrypt"

// 密码加密
func Encryption(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// 解密
func Decrypt(pwd, hash []byte) error {
	return bcrypt.CompareHashAndPassword(pwd, hash)
}
