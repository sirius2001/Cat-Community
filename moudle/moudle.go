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

func (m *MoudleMananger) CreateModel(moudleName string) {
	switch moudleName {
	case CatMoudle:

	}

}
