package route

import (
	"goTestProj/jwt"
	"goTestProj/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	user.GET("/", service.FindAllUser)
	user.GET("/underJurisdiction/:id", service.FindAllJurisdiction)
	user.POST("/", jwt.JWT(), service.PostUser)
	user.DELETE("/:id", jwt.JWT(), service.DeleteUser)
	user.PUT("/:id", jwt.JWT(), service.PutUser)

	user.GET("/region/:regionCode", service.FindUserByRegion)
	user.GET("/area/:areaCode", service.FindUserByArea)

}
