package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rige1/sberapi-mock/oapi"

	log "github.com/sirupsen/logrus"
)

type MockServer struct {
	engine *gin.Engine
}

func NewMockServerDefault() (*MockServer, error) {
	return NewMockServer(nil)
}

func NewMockServer(engine *gin.Engine) (*MockServer, error) {
	var e *gin.Engine

	if engine == nil {
		e = gin.Default()
	} else {
		e = engine
	}

	if err := initRouters(e); err != nil {
		log.Error(err)
		return nil, err
	}

	return &MockServer{engine: e}, nil
}


func (m *MockServer) Run(args ...string) error {
	return m.engine.Run(args...)
}

func initRouters(engine *gin.Engine) error {
	docs, err := oapi.LoadFromEmbedFS()

	if err != nil {
		log.Error(err)
	}

	for _, doc := range docs {
		for method, methodRef := range doc.Paths {

			g := &oapi.Generator { Schema: methodRef.Post.Responses.Get(200).Value.Content.Get("application/json").Schema }

			engine.POST(method, buildPostHandler(g))
		}
	}

	return nil
}

func buildPostHandler(generator *oapi.Generator) gin.HandlerFunc {
	return func(c *gin.Context) {
			body, err := generator.Generate()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			}

			c.JSON(http.StatusOK, body)
		}
}
