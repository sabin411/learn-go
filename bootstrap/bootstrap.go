package bootstrap

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
){
	appStop := func(context.Context) error {
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				r := gin.Default()

				r.GET("/",  func (ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{
						"message": "_____ Running boilerplate 📺 _____",
					})
				})
				r.Run(":8000")
			}()
			return nil
	},
		OnStop: appStop,
	})
}