package cat

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirius2001/Cat-Community/dao"
	"github.com/sirius2001/Cat-Community/moudle"
)

type CatMoudel struct {
	Model Cat
	Api   *CatApi
	Impl  CatImpl
}

func CreateMoudle(api *gin.RouterGroup, db *dao.DB, moudles map[string]moudle.Moudle) moudle.Moudle {
	catMoudle := CatMoudel{
		Impl: CatImpl{DB: db},
	}
	catMoudle.SetApi(api.Group("cat", nil))
	return &catMoudle
}

func (m *CatMoudel) RegisterMoudle(moudles map[string]moudle.Moudle) {
	sync.OnceFunc(func() {
		moudles["Cat"] = m
	})
}

func (m *CatMoudel) SetApi(rg *gin.RouterGroup) {
	m.Api = NewApi(rg)
}
