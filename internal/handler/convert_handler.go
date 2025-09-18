package handler

import (
	"docero/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// ConvertHandler 处理文档转换相关的HTTP请求
type ConvertHandler struct {
	convertService service.ConvertService
	uploadDir      string
	outputDir      string
}

// NewConvertHandler 创建一个新的ConvertHandler
func NewConvertHandler(s service.ConvertService, uploadDir, outputDir string) *ConvertHandler {
	return &ConvertHandler{
		convertService: s,
		uploadDir:      uploadDir,
		outputDir:      outputDir,
	}
}

// UploadAndConvertFile 处理文件上传和转换请求
func (h *ConvertHandler) UploadAndConvertFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file") // "file" 是前端表单中input name
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form: " + err.Error()})
		return
	}

	// 限制文件类型
	allowedExtensions := []string{".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".odt", ".ods", ".odp", ".txt", ".rtf"}
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	isValidExtension := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isValidExtension = true
			break
		}
	}
	if !isValidExtension {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type. Allowed: " + strings.Join(allowedExtensions, ", ")})
		return
	}

	convertedFilePath, err := h.convertService.UploadAndConvert(fileHeader, h.uploadDir, h.outputDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Document conversion failed: " + err.Error()})
		return
	}

	// 返回转换后的文件信息，或直接提供下载链接
	convertedFileName := filepath.Base(convertedFilePath)
	c.JSON(http.StatusOK, gin.H{
		"message":            "File converted successfully",
		"original_filename":  fileHeader.Filename,
		"converted_filename": convertedFileName,
		"download_url":       "/api/v1/download/" + convertedFileName, // 提供下载链接
	})
}

// DownloadConvertedFile 处理下载转换后文件的请求
func (h *ConvertHandler) DownloadConvertedFile(c *gin.Context) {
	filename := c.Param("filename")

	// 1. 获取 outputDir 的绝对路径，用于安全检查
	absOutputDir, err := filepath.Abs(h.outputDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve output directory"})
		return
	}

	// 2. 拼接完整的待下载文件路径
	// 使用 filepath.Join 拼接 filename，确保安全路径处理
	potentialFilePath := filepath.Join(absOutputDir, filename)

	// 3. 将 potentialFilePath 也规范化为绝对路径，进行严格的安全检查
	absFilePath, err := filepath.Abs(potentialFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve file path"})
		return
	}

	// 4. 最重要的安全检查：确保请求的文件路径是在指定的输出目录下
	// 这一步防止用户通过 ../../ 等方式访问到不应该访问的文件
	if !strings.HasPrefix(absFilePath, absOutputDir) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path detected"})
		return
	}

	// 5. 检查文件是否存在
	if _, err := os.Stat(absFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 6. 提供文件下载
	// c.FileAttachment 会自动设置正确的Content-Type和Content-Disposition头
	c.FileAttachment(absFilePath, filename)
}

// ShowUploadPage 显示文件上传页面
func (h *ConvertHandler) ShowUploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}
