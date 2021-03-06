// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-04-14 21:03:30.466435871 +0530 IST m=+0.024908258

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/events": {
            "get": {
                "description": "get events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a events",
                "responses": {
                    "200": {},
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/events/{id}": {
            "get": {
                "description": "get event by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a event",
                "operationId": "int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
