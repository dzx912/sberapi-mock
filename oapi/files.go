package oapi

import (
	"embed"
	"io/fs"

	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

//go:embed schemas
var openapiFS embed.FS

func LoadFromEmbedFS() ([]*openapi3.T, error) {
	result := make([]*openapi3.T, 0)

	err := fs.WalkDir(openapiFS, ".", func(path string, file fs.DirEntry, err error) error {

		if err != nil {
			log.Errorf("error while walk %s %s", path, err)
		}
		if !file.IsDir() {
			data, _ := openapiFS.ReadFile(path)

			doc, err := openapi3.NewLoader().LoadFromData(data)

			if err != nil {
				log.Errorf("error while open document %s %s", path, err)
				return nil
			}

			result = append(result, doc)
		}

		return nil
	})

	if err != nil {
		return nil, err 
	}

	return result, nil
}