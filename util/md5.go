// Package util MD5 加密算法
package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 密码加密小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil) // 求和
	// 转成16进制的字符串
	return hex.EncodeToString(cipherStr)
}

// MD5Encode 密码转大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// ValidatePasswd 验证密码
func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}

// MakePasswd 创建加密密码
func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}
