package interfaces

import "github.com/gin-gonic/gin"

type Validator interface {
	ValidateAndSetPathParams(ctx *gin.Context, objectPtr any) error
}
