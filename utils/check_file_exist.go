package utils

import "os"

// CheckFileIsExist 检查文件是否存在
func CheckFileIsExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}

	return exist
}
