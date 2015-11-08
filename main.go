package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"runtime"
)

func main() {

	if err := ItemPriceCacheInit(); err != nil {
		fmt.Println("Cache Initialization Error: " + err.Error())
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	m := martini.Classic()
	m.Use(render.Renderer())
	//m.Use(loggerMiddleware())

	m.Get("/item-price-service/", GetPriceHandler)

	m.RunOnAddr(":3000")
	m.Run()
}
