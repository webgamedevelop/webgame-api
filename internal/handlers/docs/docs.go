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
        "/secret/create": {
            "post": {
                "description": "create image pull secret",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "secret"
                ],
                "summary": "create image pull secret",
                "parameters": [
                    {
                        "description": "secret creation request",
                        "name": "secret",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ImagePullSecret"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.detailResponse-models_ImagePullSecret"
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
        "/secret/detail": {
            "get": {
                "description": "details of the secret",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "secret"
                ],
                "summary": "details of the secret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "secret id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.detailResponse-models_ImagePullSecret"
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
        "/secret/list": {
            "get": {
                "description": "list of the secret",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "secret"
                ],
                "summary": "list of the secret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "column name to order by",
                        "name": "column",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.listResponse-array_models_ImagePullSecret-models_ImagePullSecret"
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
        "/secret/update": {
            "post": {
                "description": "update image pull secret",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "secret"
                ],
                "summary": "update image pull secret",
                "parameters": [
                    {
                        "description": "secret update request",
                        "name": "secret",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ImagePullSecret"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.detailResponse-models_ImagePullSecret"
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
        "/user/password": {
            "post": {
                "description": "change password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "change password",
                "parameters": [
                    {
                        "description": "change password request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserChangePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginFailedResponse"
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
                            "$ref": "#/definitions/v1.detailResponse-models_User"
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
                            "$ref": "#/definitions/v1.detailResponse-models_User"
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
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.ImagePullSecret": {
            "type": "object",
            "required": [
                "dockerPassword",
                "dockerServer",
                "dockerUsername",
                "name",
                "secretName",
                "secretNamespace"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "dockerEmail": {
                    "description": "Email for Docker registry",
                    "type": "string",
                    "maxLength": 100
                },
                "dockerPassword": {
                    "description": "Password for Docker registry authentication",
                    "type": "string",
                    "maxLength": 100
                },
                "dockerServer": {
                    "description": "Server location for Docker registry, default https://index.docker.io/v1/",
                    "type": "string",
                    "maxLength": 100
                },
                "dockerUsername": {
                    "description": "Username for Docker registry authentication",
                    "type": "string",
                    "maxLength": 50
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "Display name",
                    "type": "string",
                    "maxLength": 50
                },
                "secretName": {
                    "description": "K8S secret resource name",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "secretNamespace": {
                    "description": "K8S namespace",
                    "type": "string",
                    "maxLength": 60,
                    "minLength": 3
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
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
        "models.UserChangePasswordRequest": {
            "type": "object",
            "required": [
                "confirmPassword",
                "password"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 16
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
        "v1.detailResponse-models_ImagePullSecret": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.ImagePullSecret"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.detailResponse-models_User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/models.User"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.listResponse-array_models_ImagePullSecret-models_ImagePullSecret": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ImagePullSecret"
                    }
                },
                "len": {
                    "type": "integer"
                },
                "message": {
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
                "data": {},
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
