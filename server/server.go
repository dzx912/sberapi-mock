package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/gin-gonic/gin"
	"github.com/rige1/sberapi-mock/config"
	"github.com/rige1/sberapi-mock/oapi"

	legacyrouter "github.com/getkin/kin-openapi/routers/legacy"
	log "github.com/sirupsen/logrus"
)

type MockServer struct {
	addr       string
	cert       string
	key        string
	clientCert string
	tlsConfig  *tls.Config
	e          *gin.Engine
}

func NewMockServerDefault(config config.Config) (*MockServer, error) {
	return NewMockServer(config, nil)
}

func NewMockServer(config config.Config, engine *gin.Engine) (*MockServer, error) {
	var e *gin.Engine

	if engine == nil {
		e = gin.Default()
	} else {
		e = engine
	}

	if err := initRouters(e, config.Validate); err != nil {
		log.Error(err)
		return nil, err
	}

	var tlsConfig *tls.Config = nil

	if config.ClientCert != "" {
		clientCert, err := ioutil.ReadFile(config.ClientCert)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(clientCert)

		tlsConfig = &tls.Config{
			ClientCAs:  certPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
	}

	return &MockServer{
		addr:       fmt.Sprintf(":%d", config.Port),
		clientCert: config.Cert,
		cert:       config.Cert,
		key:        config.Key,
		tlsConfig:  tlsConfig,
		e:          e,
	}, nil
}

func (m *MockServer) Run() error {

	server := &http.Server{
		Addr:    m.addr,
		Handler: m.e,
	}

	if m.tlsConfig != nil {
		server.TLSConfig = m.tlsConfig
	}

	if m.cert != "" && m.key != "" {
		log.Infof("Listening and serving HTTP on %s", server.Addr)
		return server.ListenAndServeTLS(m.cert, m.key)
	} else {
		log.Infof("Listening and serving HTTP on %s", server.Addr)
		return server.ListenAndServe()
	}
}

func initRouters(engine *gin.Engine, validate bool) error {
	docs, err := oapi.LoadFromEmbedFS()

	if err != nil {
		log.Error(err)
	}

	generatedHanders := ""

	for _, doc := range docs {
		router, err := legacyrouter.NewRouter(doc)

		if err != nil {
			log.Error(err)
			return err
		}

		for method, methodRef := range doc.Paths {

			g := oapi.NewGenerator(methodRef.Post.Responses.Get(200).Value.Content.Get("application/json").Schema)

			engine.POST(method, buildPostHandler(router, validate, g))

			generatedHanders = generatedHanders + "(POST) " + method + "\n"
		}
	}

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, generatedHanders)
	})

	return nil
}

var nilAuthenticationFn openapi3filter.AuthenticationFunc = func(c context.Context, ai *openapi3filter.AuthenticationInput) error { return nil }

func buildPostHandler(router routers.Router, validate bool, generator *oapi.Generator) gin.HandlerFunc {
	return func(c *gin.Context) {

		if validate {
			route, pathParams, err := router.FindRoute(c.Request)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})

				return
			}

			requestValidationInput := &openapi3filter.RequestValidationInput{
				Request:    c.Request,
				PathParams: pathParams,
				Route:      route,
				Options: &openapi3filter.Options{
					AuthenticationFunc: nilAuthenticationFn,
				},
			}

			if err := openapi3filter.ValidateRequest(c.Request.Context(), requestValidationInput); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

				return
			}
		}

		body, err := generator.Generate()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, body)
	}
}
