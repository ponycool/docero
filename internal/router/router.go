package router

import (
	"docero/internal/handler"

	"github.com/gin-gonic/gin"
)

// SetupAPIRoutes 注册 API 版本 1 的所有路由
// 它接收一个 gin.RouterGroup 用于注册路由，以及所有需要的 handler 实例
func SetupAPIRoutes(r *gin.RouterGroup, converterHandler *handler.ConvertHandler) {
	// 注册 /api/v1/convert 相关的路由
	converterRoutes := r.Group("/convert")
	{
		converterRoutes.POST("/", converterHandler.UploadAndConvertFile)
		// 注意：下载路由通常不归属于 "convert" 组，它通常是独立的资源下载
		// 但为了简单，我们暂时放在这里，或者可以放在更顶层的 apiV1 下
	}

	// 注册 /api/v1/download 相关的路由
	downloadRoutes := r.Group("/download")
	{
		downloadRoutes.GET("/:filename", converterHandler.DownloadConvertedFile)
	}
}

// SetupWebRoutes 注册 Web 页面路由 (例如首页)
func SetupWebRoutes(r *gin.Engine, converterHandler *handler.ConvertHandler) {
	r.GET("/", converterHandler.ShowUploadPage)
}
