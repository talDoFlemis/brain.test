// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/auth/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Update an Account",
                "parameters": [
                    {
                        "description": "Update Account Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.UpdateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Log In a User",
                "parameters": [
                    {
                        "description": "login Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.TokenResponse"
                        }
                    },
                    "401": {
                        "description": "Authentication Failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Authentication"
                ],
                "summary": "Delete an Account",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh an Access Token",
                "parameters": [
                    {
                        "description": "Refresh Token Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Refresh Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Expired Token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register an User",
                "parameters": [
                    {
                        "description": "Register User Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.RegisterUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/web.TokenResponse"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "web.LoginRequest": {
            "description": "Request of Login",
            "type": "object",
            "properties": {
                "password": {
                    "description": "the password of the user",
                    "type": "string"
                },
                "username": {
                    "description": "the username of the user",
                    "type": "string"
                }
            }
        },
        "web.RefreshTokenRequest": {
            "description": "Request of Refresh Token",
            "type": "object",
            "properties": {
                "refresh_token": {
                    "description": "refresh token",
                    "type": "string"
                }
            }
        },
        "web.RegisterUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "the email of the user",
                    "type": "string"
                },
                "password": {
                    "description": "the password of the user",
                    "type": "string"
                },
                "username": {
                    "description": "the username of the user",
                    "type": "string"
                }
            }
        },
        "web.TokenResponse": {
            "description": "A Token Response",
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "access token",
                    "type": "string"
                },
                "expire_at": {
                    "description": "expired at",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "refresh token",
                    "type": "string"
                }
            }
        },
        "web.UpdateAccountRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "the password of the user",
                    "type": "string"
                },
                "password": {
                    "description": "the email of the user",
                    "type": "string"
                },
                "username": {
                    "description": "the username of the user",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.69420",
	Host:             "localhost:42069",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Brain.test API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
