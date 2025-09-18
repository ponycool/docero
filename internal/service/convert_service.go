package service

import (
	"docero/internal/tool"
	"docero/pkg/storage"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

// ConvertService 定义文档转换服务接口
type ConvertService interface {
	UploadAndConvert(fileHeader *multipart.FileHeader, uploadDir, outputDir string) (string, error)
}

// convertService 实现了ConvertService接口
type convertService struct {
	fileStorage       storage.FileStorage
	documentConverter tool.DocumentConverter
}

func NewConvertService(fs storage.FileStorage, dc tool.DocumentConverter) ConvertService {
	return &convertService{
		fileStorage:       fs,
		documentConverter: dc,
	}
}

// UploadAndConvert 处理文件上传和转换的整个流程
func (s *convertService) UploadAndConvert(fileHeader *multipart.FileHeader, uploadDir, outputDir string) (string, error) {
	// 1. 保存上传的文件
	uploadedFilePath, err := s.fileStorage.SaveFile(fileHeader, uploadDir)
	if err != nil {
		return "", fmt.Errorf("failed to save uploaded file: %w", err)
	}
	// 确保在函数结束时删除上传的临时文件
	defer func() {
		_ = s.fileStorage.DeleteFile(filepath.Base(uploadedFilePath), uploadDir)
	}()

	// 2. 调用文档转换工具进行转换
	convertedFilePath, err := s.documentConverter.ConvertToPDF(uploadedFilePath, outputDir)
	if err != nil {
		return "", fmt.Errorf("document conversion failed: %w", err)
	}

	return convertedFilePath, nil
}
