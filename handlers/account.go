package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AccountGetAssets(ctx *gin.Context) {
    logrus.Infof("get asserts %v", GetDB())
}

func AccountGetLogs(ctx *gin.Context) {

}

func AccountTransfer(ctx *gin.Context) {
    
}

func GetPools(ctx *gin.Context) {
}
