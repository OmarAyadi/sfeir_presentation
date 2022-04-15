package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sfeir/storage/dto"
)

func PublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/health", Health())
	routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Health godoc
// @Summary Health check
// @Schemes
// @Description Do health check
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthReturn
// @Router /health [get]
func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		health := dto.HealthReturn{
			Status: "ok",
		}

		c.JSON(200, health)
	}
}
