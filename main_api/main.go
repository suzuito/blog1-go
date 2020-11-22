package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgin"
)

func main() {
	app := application.NewApplication()
	root := gin.New()
	bgin.SetUpRoot(root, app)
	if err := root.Run(); err != nil {
		panic(err)
	}
}
