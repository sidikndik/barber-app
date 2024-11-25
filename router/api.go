package router

import (
	"barber-app/infra"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupReouter(ctx infra.Context) {
	r := gin.Default()

	fmt.Println("server start on port ", ctx.Config.Port)
	r.Run()
}
