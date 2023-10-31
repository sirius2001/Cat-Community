package main

import (
	"github.com/sirius2001/Cat-Community/dao"
	"github.com/sirius2001/Cat-Community/web"
)

func main() {
	DB := dao.NewDao("")
	router := web.NewRouter()
	mananger.CreateMoudle("Cat", *DB, router.API.Group("/cat"))
}
