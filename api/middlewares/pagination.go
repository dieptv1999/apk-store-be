package middlewares

import (
	"github.com/dipeshdulal/clean-gin/lib"
	pagination "github.com/webstradev/gin-pagination"
)

// PaginationMiddleware middleware for cors
type PaginationMiddleware struct {
	handler lib.RequestHandler
	logger  lib.Logger
	env     lib.Env
}

// NewPaginationMiddleware creates new cors middleware
func NewPaginationMiddleware(handler lib.RequestHandler, logger lib.Logger, env lib.Env) PaginationMiddleware {
	return PaginationMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m PaginationMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	m.handler.Gin.Use(pagination.New("page", "size", "0", "10", 10, 1000))
}
