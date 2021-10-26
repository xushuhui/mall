package main

import (
	"log"
	"mall-go/internal/router"
	"mall-go/pkg/core"
)

func main() {

	app := router.InitRouter()
	core.InitValidate()
	log.Fatal(app.Listen(":3000"))

}
