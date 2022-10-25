package handlers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RespError struct {
    Err error
    Code int
}

func (e *RespError) Error() string {
    return e.Err.Error()
} 


type ErrorResp struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

func HandleNoRoute() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.AbortWithStatusJSON(
            404, ErrorResp{
                Code: 404,
                Message: "404 Not Found",
        })
    }
}

func HandleErrors() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                msg := fmt.Sprintf("%s", r)
                ctx.AbortWithStatusJSON(
                    500, ErrorResp{
                    Code: 500,
                    Message: msg,
                })
            }
        }()
        ctx.Next()
        errs := ctx.Errors
        if errs != nil {
            if len(errs) <= 0 {
                return
            }
            firstError := errs[0].Err
            var parsedErr *RespError
            errCode := 0
            errMsg := firstError.Error()
            if ok := errors.As(firstError, &parsedErr); ok {
                errCode = parsedErr.Code
                errMsg = parsedErr.Error()
            }
            ctx.AbortWithStatusJSON(
                400, ErrorResp{
                Code: errCode,
                Message: errMsg,
            })
        }
    }
}
