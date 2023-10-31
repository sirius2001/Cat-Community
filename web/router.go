package web

import "github.com/gin-gonic/gin"

type Router struct {
	API *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		API: gin.Default(),
	}
}
