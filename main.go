package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

func main() {
	doc, _ := openapi3.NewLoader().LoadFromFile("openapi/qr.v1.yaml")
	g := &Generator { Schema: doc.Paths.Find("/creation").Post.Responses.Get(200).Value.Content.Get("application/json").Schema }

	r := gin.Default()

	r.POST("/creation", func(c *gin.Context) {

		body, err := g.Generate()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(200, body)
	})

	r.Run(":8080")
}