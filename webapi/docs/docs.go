// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "healthy",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/chat": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "create chat",
                "parameters": [
                    {
                        "description": "chat params",
                        "name": "chat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddChat"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "created chat",
                        "schema": {
                            "$ref": "#/definitions/models.Chat"
                        }
                    }
                }
            }
        },
        "/chat/:id": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "upgrade request to ws",
                "parameters": [
                    {
                        "description": "ws struct",
                        "name": "struct",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.WSChatIn"
                        }
                    }
                ],
                "responses": {
                    "101": {
                        "description": "ws struct",
                        "schema": {
                            "$ref": "#/definitions/models.WSChatOut"
                        }
                    }
                }
            }
        },
        "/chats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "get all chats",
                "responses": {
                    "200": {
                        "description": "list of chats for current user",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Chat"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "logged in user",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register user",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddUser"
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
        "/user/:username": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddChat": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Super Chat"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.AddUser": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string",
                    "example": "User"
                },
                "name": {
                    "type": "string",
                    "example": "User"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "phone": {
                    "type": "string",
                    "example": "+78005553535"
                }
            }
        },
        "models.Chat": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "1662070156"
                },
                "id": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Message"
                    }
                },
                "name": {
                    "type": "string",
                    "example": "Super Chat"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.DeleteMessage": {
            "type": "object",
            "properties": {
                "messageId": {
                    "type": "integer"
                }
            }
        },
        "models.EditMessage": {
            "type": "object",
            "properties": {
                "messageId": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.GetMessages": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "models.Login": {
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
        "models.Message": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "editedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.MessageType": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "payload": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "models.ReplyMessage": {
            "type": "object",
            "properties": {
                "replyMessageId": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.SendMessage": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string",
                    "example": "What are you taking about?"
                },
                "createdAt": {
                    "type": "string",
                    "example": "1662070156"
                },
                "lastOnline": {
                    "type": "string",
                    "example": "1662070156"
                },
                "name": {
                    "type": "string",
                    "example": "John"
                },
                "nickname": {
                    "type": "string",
                    "example": "Nickname"
                },
                "phone": {
                    "type": "string",
                    "example": "+78005553535"
                }
            }
        },
        "models.WSChatIn": {
            "type": "object",
            "properties": {
                "deleteMessage": {
                    "$ref": "#/definitions/models.DeleteMessage"
                },
                "editMessage": {
                    "$ref": "#/definitions/models.EditMessage"
                },
                "getMessages": {
                    "$ref": "#/definitions/models.GetMessages"
                },
                "messageType": {
                    "$ref": "#/definitions/models.MessageType"
                },
                "replyMessage": {
                    "$ref": "#/definitions/models.ReplyMessage"
                },
                "sendMessage": {
                    "$ref": "#/definitions/models.SendMessage"
                }
            }
        },
        "models.WSChatOut": {
            "type": "object",
            "properties": {
                "newMessages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Message"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Application Api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
