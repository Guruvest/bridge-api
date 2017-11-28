// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Guruvest Bridge API is a REST API served the Guruvest platform. It is the API that third party client uses to communicate signals to the Guruvest platform.\n\n# Errors\n\nThe API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:\n\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"message\": \"page not found\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n\n# Versioning\n\nThe API is usually changed in each release, so API calls are versioned to\nensure that clients don't break. To lock to a specific version of the API,\nyou prefix the URL with its version, for example, call ` + "`" + `/v0/info` + "`" + ` to use\nthe v0 version of the ` + "`" + `/info` + "`" + ` endpoint. If the API version specified in\nthe URL is not supported by the daemon, a HTTP ` + "`" + `400 Bad Request` + "`" + ` error message\nis returned.\n\nIf you omit the version-prefix, the current version of the API (v1) is used.\nFor example, calling ` + "`" + `/info` + "`" + ` is the same as calling ` + "`" + `/v1/info` + "`" + `. Using the\nAPI without a version-prefix is deprecated and will be removed in a future release.\n\nThe API uses an open schema model, which means server may add extra properties\nto responses. Likewise, the server will ignore any extra query parameters and\nrequest body properties. When you write clients, you need to ignore additional\nproperties in responses to ensure they do not break when talking to newer\ndaemons.\n\n\n` + "`" + `` + "`" + `` + "`" + `\n",
    "title": "Guruvest Bridge API",
    "version": "1",
    "x-logo": {
      "url": "https://www.guruvest.io/wp-content/uploads/2017/08/cropped-logo-transparent.png"
    }
  },
  "basePath": "/v1",
  "paths": {
    "/accounts": {
      "get": {
        "tags": [
          "Account Data"
        ],
        "summary": "Return the list of all the linked accounts",
        "operationId": "getLinkedAccounts",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/LinkedAccount"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/accounts/orders": {
      "get": {
        "description": "Return the list of orders across linked accounts. If the account identifier is passed, selects only the orders on the specified account",
        "tags": [
          "Account Data"
        ],
        "summary": "Return the list of orders across all the linked accounts",
        "operationId": "getOrdersByAccount",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "$ref": "#/parameters/accountIdParam"
          },
          {
            "$ref": "#/parameters/limitParam"
          },
          {
            "$ref": "#/parameters/statusParam"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Order"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/accounts/portfolios": {
      "get": {
        "description": "Return the list of portfolios across linked accounts. If the account identifier is passed, selects only the portfolios on the specified account",
        "tags": [
          "Account Data"
        ],
        "summary": "Return the list of portfolios across all the linked accounts",
        "operationId": "getPortfoliosByAccount",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "$ref": "#/parameters/accountIdParam"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Position"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/accounts/trades": {
      "get": {
        "description": "Return the history of trades across linked accounts. If account identifier is passed, selects only the trades on the specific account",
        "tags": [
          "Account Data"
        ],
        "summary": "Return the list of trades across all the linked accounts",
        "operationId": "getTradesByAccount",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "$ref": "#/parameters/accountIdParam"
          },
          {
            "$ref": "#/parameters/statusParam"
          },
          {
            "$ref": "#/parameters/limitParam"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Trade"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/info": {
      "get": {
        "tags": [
          "System Info"
        ],
        "summary": "Get system information",
        "operationId": "SystemInfo",
        "responses": {
          "200": {
            "description": "No error",
            "schema": {
              "$ref": "#/definitions/SystemInfo"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/market/exchanges": {
      "get": {
        "tags": [
          "Market Access"
        ],
        "summary": "Return the list of exchanges supported",
        "operationId": "getExchanges",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Exchange"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/market/orders": {
      "post": {
        "tags": [
          "Market Access"
        ],
        "summary": "Add a new order to the order book",
        "operationId": "addOrder",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Order object that needs to be added to the order book",
            "name": "order",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Order inserted successfully"
          },
          "400": {
            "description": "Invalid request"
          },
          "422": {
            "description": "Validation exception"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/market/orders/{order_id}": {
      "get": {
        "tags": [
          "Market Access"
        ],
        "summary": "Find an order by ID",
        "operationId": "getOrderById",
        "parameters": [
          {
            "$ref": "#/parameters/orderIdParam"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid request"
          },
          "404": {
            "description": "Order not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "put": {
        "tags": [
          "Market Access"
        ],
        "summary": "Update an existing Order",
        "operationId": "updateOrder",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "$ref": "#/parameters/orderIdParam"
          },
          {
            "description": "Order object that needs to be added to the order book",
            "name": "order",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Order updated successfully"
          },
          "400": {
            "description": "Invalid request"
          },
          "403": {
            "description": "Order already filled, can not be canceled or modified"
          },
          "404": {
            "description": "Order not found"
          },
          "422": {
            "description": "Validation exception"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "Market Access"
        ],
        "summary": "Cancel the order by ID",
        "operationId": "cancelOrder",
        "parameters": [
          {
            "$ref": "#/parameters/orderIdParam"
          }
        ],
        "responses": {
          "204": {
            "description": "Order deleted successfully"
          },
          "400": {
            "description": "Invalid request"
          },
          "403": {
            "description": "Order already filled, can not be canceled or modified"
          },
          "404": {
            "description": "Order not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/market/symbols": {
      "get": {
        "description": "Return the list of symbols for the specified trading venue. If no trading venue (exchange_id) is specified, the default one will be used",
        "tags": [
          "Market Access"
        ],
        "summary": "Return the list of symbols for the given trading venue.",
        "operationId": "getSymbols",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "$ref": "#/parameters/exchangeIdParam"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Symbol"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/market/tickers": {
      "get": {
        "description": "Return the last, high (24h), low (24h), ask, bid for specified symbol. If no trading venue (exchange_id) is specified, the default one will be used",
        "tags": [
          "Market Access"
        ],
        "summary": "Returns last, high (24h), low (24h), ask, bid for specified symbol",
        "operationId": "getTicker",
        "security": [
          {
            "bridge_auth": [
              "write:orders",
              "read:orders"
            ]
          }
        ],
        "parameters": [
          {
            "$ref": "#/parameters/symbolParam"
          },
          {
            "$ref": "#/parameters/exchangeIdParam"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Tick"
            }
          },
          "400": {
            "description": "Invalid Symbol supplied"
          },
          "404": {
            "description": "Symbol not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "description": "Represents an error.",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "description": "The error message.",
          "type": "string",
          "x-nullable": false
        }
      },
      "example": {
        "message": "Something went wrong."
      }
    },
    "Exchange": {
      "description": "The trading venue, it could be either a Broker, an exchange or a financial institution with trading capabilities",
      "properties": {
        "code": {
          "description": "The code of the exchange/broker",
          "type": "string"
        },
        "fee": {
          "description": "The fee applied to each order",
          "type": "number",
          "format": "float"
        },
        "id": {
          "description": "Guruvest internal exchange/broker identifier",
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "description": "The name of the exchange/broker",
          "type": "string"
        },
        "tradable": {
          "description": "Identify if the exchange has trading capabilities enabled or not",
          "type": "boolean"
        },
        "url": {
          "description": "The URL of the exchange/broker",
          "type": "string"
        }
      },
      "xml": {
        "name": "Exchange"
      }
    },
    "LinkedAccount": {
      "description": "Information about a linked account",
      "type": "object",
      "properties": {
        "active": {
          "description": "Identify if the account is active or not",
          "type": "boolean"
        },
        "brokerage_id": {
          "description": "Guruvest internal exchange/broker identifier",
          "type": "integer",
          "format": "int32"
        },
        "brokerage_name": {
          "description": "The name of the Venue (broker or exchange)",
          "type": "string",
          "example": "The name of the associated broker or exchange"
        },
        "id": {
          "description": "The unique Guruvest internal identifier for the linked account",
          "type": "integer",
          "format": "int32"
        },
        "is_default": {
          "description": "Identify if the account should be used as the default one",
          "type": "boolean"
        },
        "name": {
          "description": "The name of the account",
          "type": "string",
          "example": "My Bittex account"
        },
        "tradable": {
          "description": "Identify if the account is able to perform trades or not",
          "type": "boolean"
        }
      },
      "xml": {
        "name": "LinkedAccount"
      }
    },
    "Order": {
      "properties": {
        "closed_at": {
          "description": "The full date time format ISO 8601 when the order was closed",
          "type": "string",
          "format": "date-time"
        },
        "filled_quantity": {
          "type": "number",
          "format": "float"
        },
        "id": {
          "type": "string",
          "example": "d0c5340b-6d6c-49d9-b567-48c4bfca13d2"
        },
        "opened_at": {
          "description": "The full date time format ISO 8601 when the order was created",
          "type": "string",
          "format": "date-time"
        },
        "quantity": {
          "type": "number",
          "format": "float"
        },
        "rate": {
          "type": "number",
          "format": "float"
        },
        "side": {
          "type": "string",
          "enum": [
            "buy",
            "sell"
          ],
          "example": "buy or sell"
        },
        "status": {
          "type": "string",
          "enum": [
            "open",
            "pending",
            "closed"
          ]
        },
        "strategy_id": {
          "description": "The identifier of the strategy to whom the order belongs to",
          "type": "string",
          "example": "d0c5340b-6d6c-49d9-b567-48c4bfca13d2"
        },
        "symbol": {
          "description": "Symbol ID",
          "type": "string",
          "example": "A valid symbol ID"
        },
        "time_in_force": {
          "description": "[optional] GTC, GTT, IOC, or FOK (default is GTC)",
          "type": "string",
          "enum": [
            "GTC",
            "GTT",
            "IOC",
            "FOK"
          ]
        },
        "type": {
          "description": "Order Type",
          "type": "string",
          "enum": [
            "limit",
            "market",
            "stop"
          ],
          "example": "limit, market, or stop"
        }
      },
      "xml": {
        "name": "Order"
      }
    },
    "Position": {
      "properties": {
        "quantity": {
          "type": "number",
          "format": "float"
        },
        "symbol": {
          "description": "The symbol ID on the exchange of choice.",
          "type": "string",
          "example": "A valid symbol ID"
        }
      },
      "xml": {
        "name": "Position"
      }
    },
    "Symbol": {
      "properties": {
        "exchange_code": {
          "description": "The code of the exchange/broker",
          "type": "string"
        },
        "symbol": {
          "description": "The symbol identifier of the exchange of choice",
          "type": "string"
        }
      },
      "xml": {
        "name": "Symbol"
      }
    },
    "SystemInfo": {
      "description": "Information of the current system",
      "type": "object",
      "properties": {
        "message": {
          "description": "This is the info message.",
          "type": "string",
          "x-nullable": false
        }
      },
      "xml": {
        "name": "SystemInfo"
      }
    },
    "Tick": {
      "properties": {
        "ask": {
          "type": "number",
          "format": "float"
        },
        "bid": {
          "type": "number",
          "format": "float"
        },
        "current_volume": {
          "type": "number",
          "format": "float"
        },
        "exchange_code": {
          "description": "The code of the exchange/broker",
          "type": "string"
        },
        "high": {
          "type": "number",
          "format": "float"
        },
        "last": {
          "type": "number",
          "format": "float"
        },
        "low": {
          "type": "number",
          "format": "float"
        },
        "symbol": {
          "description": "The symbol of an asset",
          "type": "string",
          "example": "BTC-USD"
        },
        "timestamp": {
          "description": "The full date time format ISO 8601 from the exchange/broker",
          "type": "string",
          "format": "date-time"
        }
      },
      "xml": {
        "name": "Tick"
      }
    },
    "Trade": {
      "properties": {
        "fee": {
          "description": "The fee applied to the trade",
          "type": "number",
          "format": "float"
        },
        "filled_at": {
          "description": "The full date time format ISO 8601 when the trade was filled",
          "type": "string",
          "format": "date-time"
        },
        "filled_quantity": {
          "type": "number",
          "format": "float"
        },
        "filled_rate": {
          "type": "number",
          "format": "float"
        },
        "filled_type": {
          "description": "Whether the trade was partially filled or not",
          "type": "string",
          "enum": [
            "fill",
            "partial"
          ]
        },
        "id": {
          "type": "string",
          "example": "d0c5340b-6d6c-49d9-b567-48c4bfca13d2"
        },
        "notional_traded": {
          "type": "number",
          "format": "float"
        },
        "order_id": {
          "description": "The link to the order id",
          "type": "string",
          "example": "d0c5340b-6d6c-49d9-b567-48c4bfca13d2"
        },
        "quantity": {
          "type": "number",
          "format": "float"
        },
        "rate": {
          "type": "number",
          "format": "float"
        },
        "side": {
          "type": "string",
          "enum": [
            "buy",
            "sell"
          ],
          "example": "buy or sell"
        },
        "symbol": {
          "description": "Symbol ID",
          "type": "string",
          "example": "A valid symbol ID"
        }
      }
    }
  },
  "parameters": {
    "accountIdParam": {
      "type": "integer",
      "format": "int32",
      "description": "The linked account identifier",
      "name": "account_id",
      "in": "query"
    },
    "exchangeIdParam": {
      "type": "integer",
      "format": "int32",
      "description": "The trading venue identifier. If not specified, user's default is selected",
      "name": "exchange_id",
      "in": "query"
    },
    "limitParam": {
      "type": "integer",
      "format": "int32",
      "description": "maximum number of results to return",
      "name": "limit",
      "in": "query"
    },
    "orderIdParam": {
      "type": "string",
      "description": "The order identifier",
      "name": "order_id",
      "in": "path",
      "required": true
    },
    "statusParam": {
      "enum": [
        "open",
        "pending",
        "closed"
      ],
      "type": "string",
      "description": "the status of the orders to be retrieved - open, pending, closed",
      "name": "status",
      "in": "query"
    },
    "symbolParam": {
      "type": "string",
      "description": "The symbol identifier",
      "name": "symbol",
      "in": "query",
      "required": true
    }
  },
  "securityDefinitions": {
    "api_key": {
      "description": "API key parameter",
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "bridge_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://bridge.guruvest.io/api/oauth",
      "scopes": {
        "read:orders": "read your orders",
        "write:orders": "send orders to your account"
      }
    }
  },
  "tags": [
    {
      "description": "Gather information on the Bridge system.\n",
      "name": "System Info"
    },
    {
      "description": "Gather information about your Guruvest account.\n",
      "name": "Account Data"
    },
    {
      "description": "Access to execution methods and market data.\n",
      "name": "Market Access"
    }
  ]
}`))
}
