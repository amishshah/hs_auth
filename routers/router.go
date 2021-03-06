package routers

import (
	"github.com/unicsmcr/hs_auth/routers/api/models"
	v1 "github.com/unicsmcr/hs_auth/routers/api/v1"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// MainRouter is router to connect all routers used by the app
type MainRouter interface {
	models.Router
}

type mainRouter struct {
	models.BaseRouter
	logger *zap.Logger
	apiV1  v1.APIV1Router
}

// NewMainRouter creates a new MainRouter
func NewMainRouter(logger *zap.Logger, apiV1Router v1.APIV1Router) MainRouter {
	return &mainRouter{
		logger: logger,
		apiV1:  apiV1Router,
	}
}

// RegisterRoutes registers all of the app's routes
func (r *mainRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", r.Heartbeat)

	apiV1Group := routerGroup.Group("/api/v1")
	r.apiV1.RegisterRoutes(apiV1Group)
}
