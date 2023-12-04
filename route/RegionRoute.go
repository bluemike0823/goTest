package route

import (
	"goTestProj/service"

	"github.com/gin-gonic/gin"
)

func AddRegionRouter(r *gin.RouterGroup) {
	region := r.Group("/region")
	region.GET("/", service.FindAllRegion)
	region.POST("/", service.SetRegion)
	region.GET("/area", service.FindAllArea)
	region.POST("/area", service.SetArea)

}
