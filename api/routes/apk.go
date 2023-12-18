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
	auth := s.handler.Gin.Group("/apk")
	{
		auth.POST("/create", s.apkController.CreateApk)
		auth.POST("/", s.apkController.GetApk)
		auth.POST("/create/review", s.apkController.CreateReview)
		auth.POST("/reviews", s.apkController.GetReviews)
		auth.GET("/featured", s.apkController.GetFeaturedApk)
		auth.POST("/similar", s.apkController.GetSimilarApk)
		auth.POST("/similar/develop", s.apkController.GetSimilarDevelopApk)

		auth.GET("/ads", s.apkController.GetAds)
		auth.GET("/click-ads", s.apkController.ClickAds)
		auth.GET("/feed", s.apkController.GetFeed)
		auth.GET("/click-feed", s.apkController.ClickFeed)
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
