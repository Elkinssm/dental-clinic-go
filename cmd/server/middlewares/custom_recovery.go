package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomRecoveryMiddleware(ctx *gin.Context, err any) {

	fmt.Printf("error is %+v\n", err)

	fmt.Printf("method is %+v\n", ctx.Request.Method)

	fmt.Printf("path is %+v\n", ctx.Request.URL.Path)

	fmt.Printf("the time is %+v\n", time.Now())

	fmt.Printf("the size of the request is %+v\n", ctx.Request.ContentLength)

	ctx.AbortWithError(http.StatusInternalServerError, errors.New("internal server error"))
}
