package server

import (
	"bytes"
	"io"
	"os"
)

var (
	swagger string
)

//SwaggerDoc ...
type SwaggerDoc struct{}

//NewSwaggerDoc ...
func NewSwaggerDoc() *SwaggerDoc {
	return &SwaggerDoc{}
}

//ReadDoc ...
func (s *SwaggerDoc) ReadDoc() string {
	if swagger != "" {
		return swagger
	}

	buf := bytes.NewBuffer(nil)

	f, err := os.Open("swagger.json")

	if err != nil {
		panic(err)
	}

	io.Copy(buf, f)
	f.Close()

	swagger = string(buf.Bytes())

	return swagger
}
