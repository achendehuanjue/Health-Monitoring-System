package router

import "github.com/gin-gonic/gin"

func LoadRouters(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		BrandRouter(v1)
	}

}
