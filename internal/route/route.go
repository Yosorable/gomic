package route

import (
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"
	"strings"

	"github.com/Yosorable/gomic/assets"
	"github.com/Yosorable/gomic/internal/controller"
	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute() (*gin.Engine, error) {
	fSys, err := fs.Sub(assets.WebStaticFiles, "web")
	if err != nil {
		return nil, err
	}

	if global.CONFIG.Secret == "" {
		global.CONFIG.Secret = fmt.Sprintf("%v", rand.Float64())
	}

	r := gin.Default()

	r.GET("/*any", func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if strings.HasPrefix(path, "/media") {
			id := strings.TrimPrefix(path, "/media/")
			controller.FileController.GetMediaByID(ctx, id)
			return
		} else if strings.HasPrefix(path, "/thumb") {
			id := strings.TrimPrefix(path, "/thumb/")
			controller.FileController.GetThumbByID(ctx, id)
			return
		}
		ctx.FileFromFS(path, http.FS(fSys))
	})

	auth := r.Group("/auth")
	{
		ctr := controller.AuthController
		auth.POST("/login", ctr.Login)
		auth.POST("/user", ctr.User)
	}

	r.Use(middleware.JWTAuthMiddleware())

	setUpApi(r.Group("/api"))

	return r, nil
}
