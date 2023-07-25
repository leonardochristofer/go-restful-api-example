package main

import (
	"github.com/leonardochristofer/go-restful-api-example/controller"
	"github.com/leonardochristofer/go-restful-api-example/model"
)

func main() {
	model.Init()
	controller.Start()
}
