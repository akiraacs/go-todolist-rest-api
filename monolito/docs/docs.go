// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Adrian Coradini",
            "email": "adrian.coradini3141592@hotmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks": {
            "post": {
                "description": "Adiciona uma nova tarefa à lista de tarefas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Adiciona uma nova tarefa",
                "parameters": [
                    {
                        "description": "Task data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Task"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Task": {
            "type": "object",
            "required": [
                "status",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/main.taskStatus"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "main.taskStatus": {
            "type": "string",
            "enum": [
                "Pending",
                "In Progress",
                "Completed"
            ],
            "x-enum-varnames": [
                "StatusPending",
                "StatusInProgress",
                "StatusCompleted"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "REST API To-Do List",
	Description:      "REST API para consultar, criar, atualizar e deletas tarefas",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
