package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//check role is the middle ware to check the auth.
func checkRole(ctx *gin.Context) {
	query, ok := ctx.GetQuery("action")
	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, query)
}
