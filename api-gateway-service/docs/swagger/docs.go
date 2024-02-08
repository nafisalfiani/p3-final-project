// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/admin/scheduler/trigger": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    },
                    {
                        "XDateTimes": []
                    }
                ],
                "description": "Trigger Scheduler",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scheduler"
                ],
                "summary": "Trigger Scheduler",
                "parameters": [
                    {
                        "description": "Parameter for triggering scheduler",
                        "name": "trigger_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.TriggerSchedulerParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    }
                }
            }
        },
        "/auth/v1/login": {
            "post": {
                "description": "This endpoint will sign in user with username and password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign In With Password",
                "parameters": [
                    {
                        "description": "Input Username And Password",
                        "name": "login_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    }
                }
            }
        },
        "/auth/v1/register": {
            "post": {
                "description": "This endpoint will register new user as member",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Input Username And Password",
                        "name": "register_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    }
                }
            }
        },
        "/auth/v1/verify-email/{id}": {
            "post": {
                "description": "This endpoint will mark user email as verified",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify user email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.HTTPResp"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "This endpoint will hit the server",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Ping"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.HTTPMessage": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entity.HTTPResp": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "$ref": "#/definitions/entity.HTTPMessage"
                },
                "metadata": {
                    "$ref": "#/definitions/entity.Meta"
                },
                "pagination": {
                    "$ref": "#/definitions/entity.Pagination"
                }
            }
        },
        "entity.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.Meta": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/entity.MetaError"
                },
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "time_elapsed": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "entity.MetaError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.Pagination": {
            "type": "object",
            "properties": {
                "current_elements": {
                    "type": "integer"
                },
                "current_page": {
                    "type": "integer"
                },
                "cursor_end": {
                    "type": "string"
                },
                "cursor_start": {
                    "type": "string"
                },
                "sort_by": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "total_elements": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "entity.Ping": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "entity.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.TriggerSchedulerParams": {
            "type": "object",
            "properties": {
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
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
