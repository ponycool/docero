package tool

import (
	"docero/internal/util"
	"fmt"
	"os/exec"
	"path/filepath"
)

// DocumentConverter 定义文档转换工具接口
type DocumentConverter interface {
	ConvertToPDF(inputPath string, outputPath string) (string, error)
	// Add other conversion methods as needed (e.g., ConvertToHTML)
}

// LibreOfficeConverter 是基于LibreOffice命令行的实现
type LibreOfficeConverter struct {
	sofficePath string
}

func NewLibreOfficeConverter(path string) *LibreOfficeConverter {
	return &LibreOfficeConverter{sofficePath: path}
}

// ConvertToPDF 将输入文件转换为PDF
// inputPath: 待转换文件的完整路径
// outputPath: 转换后PDF文件的输出目录
// 返回值: 转换后的PDF文件的完整路径, 错误
func (c *LibreOfficeConverter) ConvertToPDF(inputPath string, outputDir string) (string, error) {
	// 获取不带扩展名的文件名
	baseName := util.GetFileNameWithoutExtension(filepath.Base(inputPath))
	outputFileName := baseName + ".pdf"
	convertedFilePath := filepath.Join(outputDir, outputFileName)

	// LibreOffice命令: soffice --headless --convert-to pdf input.docx --outdir output_dir
	cmd := exec.Command(
		c.sofficePath,
		"--headless",
		"--convert-to", "pdf",
		inputPath,
		"--outdir", outputDir,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("LibreOffice conversion failed: %w, output: %s", err, string(output))
	}

	// 检查转换后的文件是否存在
	// LibreOffice在成功转换时，默认会生成一个同名但不同后缀的文件在outdir中
	// 所以我们检查的是 convertedFilePath
	if _, err := exec.Command("test", "-f", convertedFilePath).CombinedOutput(); err != nil {
		return "", fmt.Errorf("converted PDF file %s not found, LibreOffice output: %s", convertedFilePath, string(output))
	}

	return convertedFilePath, nil
}
