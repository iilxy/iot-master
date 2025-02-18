package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal/args"
	"runtime"
)

func info(ctx *gin.Context) {
	replyOk(ctx, gin.H{
		"version": args.Version,
		"build":   args.BuildTime,
		"git":     args.GitHash,
		"gin":     gin.Version,
		"runtime": runtime.Version(),
	})
}
