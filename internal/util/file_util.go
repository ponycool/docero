package util

import "path/filepath"

// GetFileExtension 获取文件扩展名 (不带点)
func GetFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	if len(ext) > 1 {
		return ext[1:] // 移除开头的 '.'
	}
	return ""
}

// GetFileNameWithoutExtension 获取不带扩展名的文件名
func GetFileNameWithoutExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
