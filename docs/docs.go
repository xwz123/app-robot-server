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
        "/accounts/code": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "账号注册获取验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "邮件发送成功",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            }
        },
        "/accounts/email": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "邮箱是否已注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "邮箱可用",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "邮箱已注册",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "返回失败状态",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            }
        },
        "/accounts/login": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "账号密码登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "account",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回token和userId",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.LoginResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "403": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            }
        },
        "/accounts/password": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "发送找回密码邮件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "邮件发送成功",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "邮箱校验失败",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "返回失败状态",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "userID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "旧密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "newPassword",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "修改成功",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "修改失败",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "token验证失败",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "邮件验证重置密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the key",
                        "name": "key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "new password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "邮件发送成功",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "重置失败",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "403": {
                        "description": "验证失败",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "返回失败状态",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            }
        },
        "/accounts/register": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "账号注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名(仅字母数字组合)",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码（大小字母数字特殊字符满足三种）",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "返回成功状态",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "返回失败状态",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "返回失败状态",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            }
        },
        "/accounts/username": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "用户名是否可用",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户名可用",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "用户名已注册",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BaseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.LoginResult": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "access-token",
            "in": "header"
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Swagger app-robot-server API",
	Description: "plugin maintenance server api doc",
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
	swag.Register(swag.Name, &s{})
}
