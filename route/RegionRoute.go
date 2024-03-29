package route

import (
	"goTestProj/jwt"
	"goTestProj/service"

	"github.com/gin-gonic/gin"
)

func AddRegionRouter(r *gin.RouterGroup) {
	region := r.Group("/region")
	region.GET("/", service.FindAllRegion)
	region.POST("/", jwt.JWT(), service.SetRegion)
	region.GET("/area", service.FindAllArea)
	region.POST("/area", jwt.JWT(), service.SetArea)

}
