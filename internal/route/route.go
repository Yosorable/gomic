package route

import (
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/Yosorable/gomic/assets"
	"github.com/Yosorable/gomic/internal/global"

	"github.com/gin-gonic/gin"
)

func CreateRoute() (*gin.Engine, error) {
	fSys, err := fs.Sub(assets.WebStaticFiles, "web")
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	r.StaticFS("/web", http.FS(fSys))
	r.StaticFS("/media", http.Dir(global.CONFIG.Media))
	r.StaticFS("/thumb", http.Dir(filepath.Join(global.CONFIG.Data, "thumb")))
	r.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/web"
		r.HandleContext(ctx)
	})

	setUpApi(r.Group("/api"))

	return r, nil
}
