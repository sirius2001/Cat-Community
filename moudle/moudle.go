package moudle

import (
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	CatMoudle  = "Cat_Moudle"
	AreaMoudle = "Area_Moudle"
)

type Moudle interface {
	RegisterModel(api *gin.Engine, moudles map[string]Moudle)
	SetApi(rg *gin.RouterGroup)
}

type MoudleMananger struct {
	Moudles map[string]Moudle
}

var moudleManager = &MoudleMananger{}

func GetMoudleMananger() *MoudleMananger {
	sync.OnceFunc(func() {
		moudleManager = &MoudleMananger{
			Moudles: make(map[string]Moudle),
		}
	})
	return moudleManager
}
