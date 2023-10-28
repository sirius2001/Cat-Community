package cat

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirius2001/Cat-Community/moudle"
)

type CatMoudel struct {
	Model Cat
	Api   CatApi
	Impl  CatImpl
}

func NewCatModel() moudle.Moudle {
	return &CatMoudel{}
}

func (m *CatMoudel) RegisterModel(api *gin.Engine, moudles map[string]moudle.Moudle) {
	m.registModel(moudles)
}

func (m *CatMoudel) registModel(moudles map[string]moudle.Moudle) {
	sync.OnceFunc(func() {
		moudles["Cat"] = m
	})
}
