// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-13 01:16:02.939174 +0300 MSK m=+0.044950722

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
        "/api/session": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "check session of current user",
                "operationId": "is-session-valid",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/upload_avatar": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "upload avatar on server",
                "operationId": "upload-avatar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.AvatarLinkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Current user info",
                "operationId": "get-user-from-sesion",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UserInfoResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update current profile (only avatar, only password or both)",
                "operationId": "update-profile",
                "parameters": [
                    {
                        "description": "user data to update",
                        "name": "AuthData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.ProfileUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.ProfileUpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create account in our perfect game",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "user data to create",
                        "name": "AuthData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.SingUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/count": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get count of registered users",
                "operationId": "get-users-count",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UsersCountInfoResponse"
                        }
                    }
                }
            }
        },
        "/api/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get user by id",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UserInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "summary": "on hello returning world",
                "operationId": "hello-world",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/session": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Sign in with your account with email and password, set session cookie",
                "operationId": "post-session",
                "parameters": [
                    {
                        "description": "user auth data",
                        "name": "AuthData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.signInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Sign out from your account, expire cookie",
                "operationId": "delete-session",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            }
        },
        "/user/score": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "return slice of Scores (Nickname + score)",
                "operationId": "get-leaderboard",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "default: 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "default: 10",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.GetLeaderboardResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AvatarLinkResponse": {
            "type": "object",
            "properties": {
                "avatar_link": {
                    "type": "string"
                }
            }
        },
        "controllers.GetLeaderboardResponse": {
            "type": "object",
            "properties": {
                "scores": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.Scores"
                    }
                }
            }
        },
        "controllers.ProfileUpdateRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "controllers.ProfileUpdateResponse": {
            "type": "object",
            "properties": {
                "avatar_link": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "controllers.SingUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.UserInfoResponse": {
            "type": "object",
            "properties": {
                "avatar_link": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "controllers.UsersCountInfoResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "controllers.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controllers.signInRequest": {
            "type": "object",
            "properties": {
                "loginOrEmail": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.Scores": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "score": {
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
