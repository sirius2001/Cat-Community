package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirius2001/Cat-Community/dao"
	"github.com/sirius2001/Cat-Community/moudle"
	"github.com/sirius2001/Cat-Community/moudle/cat"
)

type Manager struct {
	Mouldes map[string]moudle.Moudle
	Locker  sync.RWMutex
}

var mananger = Manager{Mouldes: make(map[string]moudle.Moudle)}

func GetManager() *Manager {
	mananger.Locker.Lock()
	defer mananger.Locker.Unlock()
	return &mananger
}

func (m *Manager) CreateMoudle(Name string, db dao.DB, router *gin.RouterGroup) {
	switch Name {
	case "Cat":
		cat.CreateMoudle(router, &db, GetManager().Mouldes)

	}
}
