# Docero

Docero 是一款使用 Go 语言和 Gin 框架开发的高性能文档转换服务。它通过集成 LibreOffice，将常见的办公文档（如
.docx、.pptx、.xlsx 等）快速转换为 PDF 格式，适用于文档归档、在线预览、报表生成等场景。

无论你正在构建企业文档系统、文件共享平台，还是自动化报告服务，Docero 都能以极简的 RESTful API 帮你轻松实现文档格式转换。

## 核心特性

- 🚀 高性能轻量 – 使用 Go 编写，高并发、低资源占用。
- 📄 格式转换 – 支持 .docx、.pptx、.xlsx、.odt 等格式转 PDF。
- 🌐 HTTP API 接口 – 上传文件，返回转换后的 PDF，易于集成。
- 🧰 部署简单 – 支持本地运行或 Docker 部署。
- 🔧 无头模式 – 基于 LibreOffice 无界面运行，适合服务器环境。

## 快速开始

```
git clone https://github.com/ponycool/docero.git
cd docero
go run main.go
```

服务启动后，访问 http://localhost:8528 即可上传文件并完成转换。

## 🐳 Docker 部署


## API 使用示例

```
curl -X POST -F "file=@合同.docx" http://localhost:8080/convert --output 合同.pdf
```

## 🎯 适用场景

自动生成合同、简历、报告的 PDF 版本
在线文档预览系统（先转 PDF 再渲染）
文件归档与标准化处理
与 Web 前端、移动端或后端服务无缝集成

## 🤝 贡献与反馈

欢迎提交 Issue 或 Pull Request！让我们一起打造最简洁、最可靠的中文文档转换工具。