// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "docs": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/people/{id}": {
            "put": {
                "description": "Update a person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a person",
                "operationId": "update-person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a person",
                "operationId": "delete-person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                }
            }
        },
        "/person": {
            "get": {
                "description": "Get all persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all persons",
                "operationId": "get-person",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a person",
                "operationId": "create-person",
                "parameters": [
                    {
                        "description": "Person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                }
            }
        },
        "/person/{id}": {
            "get": {
                "description": "Get a person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a person",
                "operationId": "get-person-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                }
            }
        },
        "/region": {
            "get": {
                "description": "Get all regions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all regions",
                "operationId": "get-regions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Region"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a region",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a region",
                "operationId": "create-region",
                "parameters": [
                    {
                        "description": "Region",
                        "name": "region",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Region"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Region"
                        }
                    }
                }
            }
        },
        "/regionresidents/{id}": {
            "get": {
                "description": "Get all residents in a region",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all residents in a region",
                "operationId": "get-residents",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Region ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Person"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Person": {
            "type": "object",
            "properties": {
                "forename": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "patronymic": {
                    "type": "string"
                },
                "region_id": {
                    "type": "integer"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "main.Region": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "docs",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
