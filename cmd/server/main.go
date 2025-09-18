package main

import (
	"docero/internal/config"
	"docero/internal/handler"
	"docero/internal/router"
	"docero/internal/service"
	"docero/internal/tool"
	"docero/pkg/logger"
	"docero/pkg/storage"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	// Go标准库log用于初始化slog前的错误
	stdlog := log.Default()

	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// 2. 初始化 slog 日志
	logger.InitSlogLogger(cfg)
	// 在此处设置 Gin 的默认日志输出到 slog
	// Gin 默认会使用 log.Default()，我们可以通过 Hook 来捕获它
	// 或者在 Gin 初始化时传递自定义的 Logger。更简单的是，
	// 如果你希望 Gin 自己的 Access Log 也走 slog，你可以创建一个自定义 Gin 中间件。
	// 但默认的 gin.Default() 已经集成了日志，这里先让它独立运行，
	// 我们的应用日志走 slog。

	slog.Info("Application initialization started.", "config_path", cfg.Log.Filename)
	slog.Debug("Full configuration loaded", "config", cfg) // debug 级别，只有当日志级别设置为 debug 时才显示

	gin.SetMode(cfg.Server.Mode)
	slog.Info("Gin mode set", "mode", cfg.Server.Mode)

	// 2. 依赖注入
	// Storage
	localStorage := storage.NewLocalStorage()

	// Document Converter
	libreOfficeConverter := tool.NewLibreOfficeConverter(cfg.Converter.LibreofficePath)

	// Service
	convertService := service.NewConvertService(localStorage, libreOfficeConverter)

	// Handler
	convertHandler := handler.NewConvertHandler(
		convertService,
		cfg.Converter.UploadDir,
		cfg.Converter.OutputDir,
	)

	// 3. 初始化Gin路由器
	routerEngine := gin.Default()

	// 加载HTML模板 (如果提供上传页面)
	routerEngine.LoadHTMLGlob("web/templates/*.html")

	// 注册静态文件服务路由
	// 这将允许浏览器通过 /static/style.css 访问 web/static/style.css
	routerEngine.Static("/static", "./web/static")

	// 4. 注册路由
	// 注册 Web 页面路由
	router.SetupWebRoutes(routerEngine, convertHandler)
	slog.Info("Web routes registered.")

	// 6. 注册 API 路由
	apiV1 := routerEngine.Group("/api/v1")
	router.SetupAPIRoutes(apiV1, convertHandler)
	slog.Info("API routes registered.")

	// 5. 启动服务器
	slog.Info("Server starting...", "port", cfg.Server.Port, "mode", cfg.Server.Mode)
	if err := routerEngine.Run(cfg.Server.Port); err != nil {
		slog.Error("Server failed to start", "error", err)
		stdlog.Fatalf("Server failed to start: %v", err)
	}
	slog.Info("Server stopped.")
}
