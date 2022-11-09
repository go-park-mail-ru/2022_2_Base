// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/avatar": {
            "post": {
                "description": "sets user's avatar",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Set user's avatar",
                "operationId": "setAvatar",
                "parameters": [
                    {
                        "type": "file",
                        "description": "user's avatar",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/cart": {
            "get": {
                "description": "gets user's cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "gets user's cart",
                "operationId": "GetCart",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "updates user's cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "updates user's cart",
                "operationId": "UpdateCart",
                "parameters": [
                    {
                        "description": "ProductCart items",
                        "name": "items",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProductCart"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/deletefromcart": {
            "post": {
                "description": "Deletes Item From cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Deletes Item From cart",
                "operationId": "DeleteItemFromCart",
                "parameters": [
                    {
                        "description": "ProductCart item",
                        "name": "items",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProductCartItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/insertintocart": {
            "post": {
                "description": "Adds item to cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Adds item to cart",
                "operationId": "AddItemToCart",
                "parameters": [
                    {
                        "description": "ProductCart item",
                        "name": "items",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProductCartItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Log in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Logs in and returns the authentication  cookie",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "UserDB params",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/logout": {
            "delete": {
                "description": "Logs out user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Logs out user",
                "operationId": "logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/makeorder": {
            "post": {
                "description": "makes user's order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "makes user's order",
                "operationId": "MakeOrder",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Gets products for main page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Gets products for main page",
                "operationId": "getMain",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "description": "gets user by username in cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get current user",
                "operationId": "getUser",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserProfile"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "changes user parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "changes user parameters",
                "operationId": "changeUserParameters",
                "parameters": [
                    {
                        "description": "UserProfile params",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/session": {
            "get": {
                "description": "Checks if user has active session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Checks if user has active session",
                "operationId": "session",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Access token is missing or invalid",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Sign up user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Signs up and returns the authentication  cookie",
                "operationId": "signup",
                "parameters": [
                    {
                        "description": "UserDB params",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreateParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict - UserDB already exists",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Adress": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "house": {
                    "type": "string"
                },
                "primary": {
                    "type": "boolean"
                },
                "street": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "model.Order": {
            "type": "object",
            "properties": {
                "adress": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderItem"
                    }
                },
                "orderstatus": {
                    "type": "string"
                },
                "paymentstatus": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "model.OrderItem": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "item": {
                    "$ref": "#/definitions/model.Product"
                }
            }
        },
        "model.PaymentMethod": {
            "type": "object",
            "properties": {
                "expirydate": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "priority": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imgsrc": {
                    "type": "string"
                },
                "lowprice": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "model.ProductCart": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.ProductCartItem": {
            "type": "object",
            "properties": {
                "itemid": {
                    "type": "integer"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "body": {}
            }
        },
        "model.UserCreateParams": {
            "type": "object",
            "properties": {
                "email": {
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
        "model.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserProfile": {
            "type": "object",
            "properties": {
                "adress": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Adress"
                    }
                },
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "paymentmethods": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PaymentMethod"
                    }
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Reozon API",
	Description:      "Reazon back server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
