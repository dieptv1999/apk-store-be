package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/lib"
)

// ApkRoutes struct
type ApkRoutes struct {
	logger        lib.Logger
	handler       lib.RequestHandler
	apkController controllers.ApkController
}

// Setup user routes
func (s ApkRoutes) Setup() {
	s.logger.Info("Setting up apk")
	apkGroup := s.handler.Gin.Group("/apk")
	{
		apkGroup.POST("/create", s.apkController.CreateApk)
		apkGroup.POST("/", s.apkController.GetApk)
		apkGroup.POST("/create/review", s.apkController.CreateReview)
		apkGroup.POST("/reviews", s.apkController.GetReviews)
		apkGroup.GET("/featured", s.apkController.GetFeaturedApk)
		apkGroup.POST("/similar", s.apkController.GetSimilarApk)
		apkGroup.POST("/similar/develop", s.apkController.GetSimilarDevelopApk)
		apkGroup.POST("/version", s.apkController.CreateApkVersion)
		apkGroup.POST("/versions", s.apkController.GetApkVersion)
		apkGroup.POST("/category", s.apkController.GetApkInCategory)

		apkGroup.GET("/ads", s.apkController.GetAds)
		apkGroup.GET("/click-ads", s.apkController.ClickAds)
		apkGroup.GET("/feed", s.apkController.GetFeed)
		apkGroup.GET("/click-feed", s.apkController.ClickFeed)
	}
}

// NewApkRoutes creates new user controller
func NewApkRoutes(
	handler lib.RequestHandler,
	apkController controllers.ApkController,
	logger lib.Logger,
) ApkRoutes {
	return ApkRoutes{
		handler:       handler,
		logger:        logger,
		apkController: apkController,
	}
}
