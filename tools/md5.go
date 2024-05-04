package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 生成MD5哈希
func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
