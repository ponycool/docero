# Docero

Docero is a high-performance document conversion service developed using Go language and the Gin framework. By integrating LibreOffice, it quickly converts common office documents (such as .docx, .pptx, .xlsx, etc.) into PDF format, suitable for document archiving, online preview, and report generation scenarios.

Whether you are building an enterprise document system, a file sharing platform, or an automated reporting service, Docero can help you easily implement document format conversion with a simple RESTful API.

## Core Features

- 🚀 High Performance & Lightweight – Written in Go, providing high concurrency and low resource consumption.
- 📄 Format Conversion – Supports converting .docx, .pptx, .xlsx, .odt and other formats to PDF.
- 🌐 HTTP API Interface – Upload files and receive converted PDFs, easy to integrate.
- 🧰 Simple Deployment – Supports local execution or Docker deployment.
- 🔧 Headless Mode – Based on LibreOffice running without GUI, suitable for server environments.

## Quick Start

```
git clone https://github.com/ponycool/docero.git
cd docero
go run main.go
```

After the service starts, visit http://localhost:8528 to upload files and complete conversion.

## 🐳 Docker Deployment


## API Usage Example

```
curl -X POST -F "file=@contract.docx" http://localhost:8080/convert --output contract.pdf
```

## 🎯 Application Scenarios

- Automatically generate PDF versions of contracts, resumes, and reports
- Online document preview systems (convert to PDF first, then render)
- Document archiving and standardization processing
- Seamless integration with web frontends, mobile apps, or backend services

## 🤝 Contribution and Feedback

Welcome to submit Issues or Pull Requests! Let's work together to create the most concise and reliable document conversion tool.