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
            "name": "webgamedevelop",
            "url": "http://www.swagger.io/support",
            "email": "webgamedevelop@163.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/refresh_token": {
            "get": {
                "description": "refresh token",
                "tags": [
                    "user"
                ],
                "summary": "refresh token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginFailedResponse"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "sign in",
                "tags": [
                    "user"
                ],
                "summary": "sign in",
                "parameters": [
                    {
                        "description": "login request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginFailedResponse"
                        }
                    }
                }
            }
        },
        "/user/signout": {
            "get": {
                "description": "sign out",
                "tags": [
                    "user"
                ],
                "summary": "sign out",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginFailedResponse"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "sign up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "sign up",
                "parameters": [
                    {
                        "description": "sign up request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.simpleResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.simpleResponse"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "description": "update user info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user info",
                "parameters": [
                    {
                        "description": "update user info request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.simpleResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.simpleResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "required": [
                "confirmPassword",
                "email",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 50
                },
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 16
                },
                "phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 11
                }
            }
        },
        "models.UserLoginRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 16
                }
            }
        },
        "models.UserUpdateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 50
                },
                "phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 11
                }
            }
        },
        "v1.LoginFailedResponse": {
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
        "v1.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "expire": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "v1.simpleResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "extend": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "webgame-api",
	Description:      "webgame-api docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
