package middleware

import (
	"belajar-go-restapi-gin-latihan2/model/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	if "API-SIMPLE" == ctx.Request.Header.Get("API-KEY") {
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponses{
			Code:   http.StatusBadRequest,
			Status: "UnAuthorized",
		})
	}

}
