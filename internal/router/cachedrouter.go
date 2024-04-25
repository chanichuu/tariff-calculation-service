package router

import (
	"github.com/gin-gonic/gin"
)

var cachedRouter *gin.Engine

// Return a new gin router if there is none yet
func NewRouter() *gin.Engine {
	if cachedRouter == nil {
		cachedRouter = gin.Default()
	}
	return cachedRouter
}
