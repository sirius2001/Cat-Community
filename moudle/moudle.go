package moudle

import (
	"github.com/gin-gonic/gin"
)

const (
	CatMoudle  = "Cat_Moudle"
	AreaMoudle = "Area_Moudle"
)

type Moudle interface {
	RegisterMoudle(moudles map[string]Moudle)
	SetApi(rg *gin.RouterGroup)
}
