package routers

import (
	"auht_fake/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	//跨域中间件
	r.Use(middleware.CORSMiddleware())
	//通用路由
	commonRouters(r)
}
