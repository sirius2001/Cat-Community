package area

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirius2001/Cat-Community/moudle"
)

type AreaMoudel struct {
	Model Area
	Api   AreaApi
	Impl  AreaImpl
}

func NewCatModel() moudle.Moudle {
	return &AreaMoudel{}
}

func (m *AreaMoudel) RegisterModel(api *gin.Engine, moudles map[string]moudle.Moudle) {
	m.registModel(moudles)
}

func (m *CatMoudel) registModel(moudles map[string]moudle.Moudle) {
	sync.OnceFunc(func() {
		moudles["Cat"] = NewCatModel()
	})
}
