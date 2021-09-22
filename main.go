package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
