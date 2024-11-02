package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for the healthcheck.
func Handler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
