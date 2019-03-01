package core

import (
	"github.com/gin-gonic/gin"
)

type Storage interface {
	Close() error
	Find(ctx *gin.Context)
}
