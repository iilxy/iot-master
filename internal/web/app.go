package web

import (
	"context"
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal/web/api"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"mime"
	"net/http"
	"path"
	"strings"
	"time"
)

func init() {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		log.Error(err)
	}
}

//go:embed all:www
var wwwFiles embed.FS

var server *http.Server

func Serve(addr string) {
	gin.SetMode(gin.ReleaseMode)

	//GIN初始化
	//app := gin.Default()
	app := gin.New()
	app.Use(gin.Recovery())

	//if cfg.Debug {
	app.Use(gin.Logger())

	//跨域问题
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	app.Use(cors.New(cfg))

	//启用session
	app.Use(sessions.Sessions("iot-master", cookie.NewStore([]byte("iot-master"))))

	//开启压缩
	//if cfg.Compress {
	app.Use(gzip.Gzip(gzip.DefaultCompression)) //gzip.WithExcludedPathsRegexs([]string{".*"})

	//注册前端接口
	api.RegisterRoutes(app.Group("/api"))

	app.Any("/service/:name/*path", api.ServiceProxy)

	//前端静态文件
	//app.StaticFS("/www", http.FS(wwwFiles))

	wwwFS := http.FS(wwwFiles)
	app.Use(func(c *gin.Context) {

		if c.Request.Method == http.MethodGet {
			//支持前端框架的无“#”路由
			fn := path.Join("www", c.Request.URL.Path) //删除查询参数
			f, err := wwwFS.Open(fn)
			if err == nil {
				defer f.Close()
				stat, err := f.Stat()
				if err != nil {
					c.Next() //500错误
					return
				}
				if !stat.IsDir() {
					http.ServeContent(c.Writer, c.Request, fn, stat.ModTime(), f)
					return
				}
			}

			//默认首页
			f, err = wwwFS.Open("www/index.html")
			if err != nil {
				c.Next()
				return
			}
			defer f.Close()

			fn += ".html" //避免DetectContentType
			http.ServeContent(c.Writer, c.Request, fn, time.Now(), f)
		}
	})

	//监听HTTP
	//if err := app.Run(cfg.Addr); err != nil {
	//	log.Fatal("HTTP 服务启动错误", err)
	//}

	log.Println("Web服务启动", addr)
	server = &http.Server{
		Addr:    resolvePort(addr),
		Handler: app,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Web服务启动错误", err)
	}
}

func Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}

func resolvePort(addr string) string {
	if strings.IndexByte(addr, ':') == -1 {
		return ":" + addr
	}
	return addr
}
