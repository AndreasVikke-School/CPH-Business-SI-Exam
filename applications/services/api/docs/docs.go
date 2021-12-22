// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/book/all/": {
            "get": {
                "description": "Gets a list of all books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Book"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/checkout/{title}": {
            "get": {
                "description": "Checkouts a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.BookTitle"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/get-recs-author/{title}": {
            "get": {
                "description": "Gets a recommended list of all books from author by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.BookSimple"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/get-recs-year/{title}": {
            "get": {
                "description": "Gets a recommended list of all books from year by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.BookSimple"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/get-simple/{title}": {
            "get": {
                "description": "Gets a simplified list of books by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.BookSimple"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/get/{title}": {
            "get": {
                "description": "Gets a book by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/return/{title}": {
            "get": {
                "description": "Returns a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.BookTitle"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/search/{title}": {
            "get": {
                "description": "Gets a list of all searched books by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of book",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Book"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/book/write-csv-to-db/": {
            "get": {
                "description": "Writes a csv file to the db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.BookTitle"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/loan/all-by-user/{id}": {
            "get": {
                "description": "Gets a list of all loans by a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of all loans by user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Loan"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/loan/all/": {
            "get": {
                "description": "Gets a list of all loans",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Loan"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/loan/create/": {
            "post": {
                "description": "Creates a loan entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "parameters": [
                    {
                        "description": "Create loan",
                        "name": "LoanEntry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.LoanEntry"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/loan/get/{id}": {
            "get": {
                "description": "Gets a loan by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loan"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of loan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Loan"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/log/all-by-user/{id}": {
            "get": {
                "description": "Says a list of all logs by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Log"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/log/create/": {
            "post": {
                "description": "Creates a log entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "parameters": [
                    {
                        "description": "Create log entry",
                        "name": "LogEntry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.LogEntry"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/log/get-by-user/{userId}/{logId}": {
            "get": {
                "description": "Gets a log by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Id of log",
                        "name": "logId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Log"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/user/all/": {
            "get": {
                "description": "Gets a list of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.User"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/user/create/": {
            "post": {
                "description": "Creates a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "User to create",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/get/{id}": {
            "get": {
                "description": "Gets a user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/all/": {
            "get": {
                "description": "Gets a list of all vinyls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Vinyl"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/get-by-title/{title}": {
            "get": {
                "description": "Gets a vinyl by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of vinyl",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Vinyl"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/get-recs-artist/{title}": {
            "get": {
                "description": "Gets a recommended list of all vinyls from artist by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of vinyl",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.VinylSimple"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/get-recs-year/{title}": {
            "get": {
                "description": "Gets a recommended list of all vinyls from year by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of vinyl",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.VinylSimple"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/get-simple/{title}": {
            "get": {
                "description": "Gets a simplified list of vinyls by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of vinyl",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.VinylSimple"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/get/{id}": {
            "get": {
                "description": "Gets a vinyl by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of vinyl",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Vinyl"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/vinyl/search/{title}": {
            "get": {
                "description": "Gets a list of vinyls from search by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vinyl"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title of vinyl",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Vinyl"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Book": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "author": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "isbn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "main.BookSimple": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "isbn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "main.BookTitle": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "main.Loan": {
            "type": "object",
            "properties": {
                "entityId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "main.LoanEntry": {
            "type": "object",
            "properties": {
                "entityId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "main.Log": {
            "type": "object",
            "properties": {
                "entityId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "unix": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "main.LogEntry": {
            "type": "object",
            "properties": {
                "entityId": {
                    "type": "integer"
                },
                "unix": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "main.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "main.Vinyl": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "artist": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "main.VinylSimple": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api",
	Schemes:     []string{"http"},
	Title:       "Book & Venyl Loan Service",
	Description: "API for school project",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
