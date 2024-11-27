package main

import (
	"barber-app/infra"
	"barber-app/router"
)

func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("errer initialis context")
	}

	router.SetupReouter(ctx)
}
