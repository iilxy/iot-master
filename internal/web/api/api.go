package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/model"
	"net/http"
)

func catchError(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//runtime.Stack()
			//debug.Stack()
			switch err.(type) {
			case error:
				replyError(ctx, err.(error))
			case string:
				replyFail(ctx, err.(string))
			default:
				ctx.JSON(http.StatusOK, gin.H{"error": err})
			}
		}
	}()
	ctx.Next()

	//TODO 内容如果为空，返回404

}

func mustLogin(ctx *gin.Context) {
	//检查Token
	token, has := ctx.GetQuery("token")
	if has {
		user, ok := tokens.Load(token)
		if ok {
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
			ctx.Abort()
		}
		return
	}

	//检查Session
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
		ctx.Abort()
	}
}

func RegisterRoutes(app *gin.RouterGroup) {
	//错误恢复，并返回至前端
	app.Use(catchError)

	app.GET("/info", info)

	app.GET("/auth", auth)
	app.POST("/login", login)

	//检查 session，必须登录
	app.Use(mustLogin)

	app.GET("/logout", logout)
	app.POST("/password", password)

	//修改配置
	app.GET("/config", loadConfig)
	app.POST("/config", saveConfig)

	//用户接口
	app.GET("/user/me", userMe)
	app.POST("/user/search", createCurdApiSearch[model.User]())
	app.GET("/user/list", createCurdApiList[model.User]())
	app.POST("/user/create", parseParamId, createCurdApiCreate[model.User](nil, nil))
	app.GET("/user/:id", parseParamId, createCurdApiGet[model.User]())
	app.POST("/user/:id", parseParamId, createCurdApiModify[model.User](nil, nil, "username", "nickname", "disabled"))
	app.GET("/user/:id/delete", parseParamId, createCurdApiDelete[model.User](nil, nil))
	app.GET("/user/:id/password", parseParamId, userPassword)
	app.GET("/user/:id/enable", parseParamId, createCurdApiDisable[model.User](false, nil, nil))
	app.GET("/user/:id/disable", parseParamId, createCurdApiDisable[model.User](true, nil, nil))

	//项目接口
	app.POST("/project/search", createCurdApiSearch[model.Project]())
	app.GET("/project/list", createCurdApiList[model.Project]())
	app.POST("/project/create", createCurdApiCreate[model.Project](nil, afterProjectCreate))
	app.GET("/project/:id", parseParamId, createCurdApiGet[model.Project]())
	app.POST("/project/:id", parseParamId, createCurdApiModify[model.Project](nil, afterProjectUpdate,
		"name", "devices", "template_id", "hmi", "aggregators", "jobs", "alarms", "strategies", "context", "disabled"))
	app.GET("/project/:id/delete", parseParamId, createCurdApiDelete[model.Project](nil, afterProjectDelete))

	app.GET("/project/:id/values", parseParamId, projectValues)
	app.POST("/project/:id/assign", parseParamId, projectAssign)
	//app.GET("/project/:id/targets", parseParamId, projectTargets)

	//设备接口
	app.POST("/device/search", createCurdApiSearch[model.Device]())
	app.GET("/device/list", createCurdApiList[model.Device]())
	app.POST("/device/create", createCurdApiCreate[model.Device](nil, afterDeviceCreate))
	app.GET("/device/:id", parseParamId, createCurdApiGet[model.Device]())
	app.POST("/device/:id", parseParamId, createCurdApiModify[model.Device](nil, afterDeviceUpdate,
		"name", "tunnel_id", "product_id", "station",
		"hmi", "tags", "points", "pollers", "calculators", "validators", "commands",
		"context", "disabled",
	))
	app.GET("/device/:id/delete", parseParamId, createCurdApiDelete[model.Device](nil, afterDeviceDelete))

	app.GET("/device/:id/values", parseParamId, deviceValues)
	app.POST("/device/:id/assign", parseParamId, deviceAssign)
	app.GET("/device/:id/refresh", parseParamId, deviceRefresh)

	//元件接口
	app.POST("/product/search", createCurdApiSearch[model.Product]())
	app.GET("/product/list", createCurdApiList[model.Product]())
	app.POST("/product/create", createCurdApiCreate[model.Product](generateUUID, nil))
	app.GET("/product/:id", parseParamStringId, createCurdApiGet[model.Product]())
	app.POST("/product/:id", parseParamStringId, createCurdApiModify[model.Product](nil, nil,
		"name", "manufacturer", "info",
		"hmi", "tags", "points", "pollers", "calculators", "validators", "commands",
		"context", "disabled",
	))
	app.GET("/product/:id/delete", parseParamStringId, createCurdApiDelete[model.Product](nil, nil))

	//服务器接口
	app.POST("/server/search", createCurdApiSearch[model.Server]())
	app.GET("/server/list", createCurdApiList[model.Server]())
	app.POST("/server/create", createCurdApiCreate[model.Server](nil, afterServerCreate))
	app.GET("/server/:id", parseParamId, createCurdApiGet[model.Server]())
	app.POST("/server/:id", parseParamId, createCurdApiModify[model.Server](nil, afterServerUpdate,
		"name", "type", "addr",
		"retry", "register", "heartbeat", "protocol", "devices", "disabled"))
	app.GET("/server/:id/delete", parseParamId, createCurdApiDelete[model.Server](nil, afterServerDelete))

	//通道接口
	app.POST("/tunnel/search", createCurdApiSearch[model.Tunnel]())
	app.GET("/tunnel/list", createCurdApiList[model.Tunnel]())
	app.POST("/tunnel/create", createCurdApiCreate[model.Tunnel](nil, afterTunnelCreate))
	app.GET("/tunnel/:id", parseParamId, createCurdApiGet[model.Tunnel]())
	app.POST("/tunnel/:id", parseParamId, createCurdApiModify[model.Tunnel](nil, nil,
		"name", "type", "addr", "retry", "heartbeat", "serial", "protocol", "disabled"))
	app.GET("/tunnel/:id/delete", parseParamId, createCurdApiDelete[model.Tunnel](nil, afterTunnelDelete))

	//系统接口
	app.GET("/system/cpu-info", cpuInfo)
	app.GET("/system/cpu", cpuStats)
	app.GET("/system/memory", memStats)
	app.GET("/system/disk", diskStats)
	app.GET("/system/machine", machineInfo)

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	app.Use(func(ctx *gin.Context) {
		replyFail(ctx, "Not found")
		ctx.Abort()
	})
}

func replyList(ctx *gin.Context, data interface{}, total int64) {
	ctx.JSON(http.StatusOK, gin.H{"data": data, "total": total})
}

func replyOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func replyFail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{"error": err})
}

func replyError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func nop(ctx *gin.Context) {
	ctx.String(http.StatusForbidden, "Unsupported")
}