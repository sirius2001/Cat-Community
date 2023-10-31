package main

import (
	"log/slog"

	"github.com/sirius2001/Cat-Community/dao"
	"github.com/sirius2001/Cat-Community/web"
)

func main() {
	DB := dao.NewDao("")
	if DB == nil {
		slog.Error("connect to db err :")
		return
	}

	router := web.NewRouter()
	GetManager().CreateMoudle("Cat", *DB, router.API.Group("/cat"))
}
