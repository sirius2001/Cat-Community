package cat

import "github.com/gin-gonic/gin"

type CatApi struct {
	RG *gin.RouterGroup
}

func NewApi(RG *gin.RouterGroup) *CatApi {
	return &CatApi{
		RG: RG,
	}
}
