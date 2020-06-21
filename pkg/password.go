package pkg

import "golang.org/x/crypto/bcrypt"

// Encrypt 使用bcrypt加密
func Encrypt(source string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// Compare 验证密码
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
