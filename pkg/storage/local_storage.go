package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// FileStorage 定义文件存储接口
type FileStorage interface {
	SaveFile(fileHeader *multipart.FileHeader, destDir string) (string, error)
	GetFilePath(filename string, dir string) string
	DeleteFile(filename string, dir string) error
}

// LocalStorage 是基于本地文件系统的FileStorage实现
type LocalStorage struct{}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

// SaveFile 保存上传的文件到指定目录
func (s *LocalStorage) SaveFile(fileHeader *multipart.FileHeader, destDir string) (string, error) {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create destination directory %s: %w", destDir, err)
	}

	filename := filepath.Base(fileHeader.Filename)
	filePath := filepath.Join(destDir, filename)

	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	return filePath, nil
}

// GetFilePath 返回文件的完整路径
func (s *LocalStorage) GetFilePath(filename string, dir string) string {
	return filepath.Join(dir, filename)
}

// DeleteFile 删除指定目录下的文件
func (s *LocalStorage) DeleteFile(filename string, dir string) error {
	filePath := filepath.Join(dir, filename)
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file %s: %w", filePath, err)
	}
	return nil
}
