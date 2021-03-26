// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

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
    "title": "Library",
    "version": "0.0.1"
  },
  "paths": {
    "/book": {
      "get": {
        "summary": "request a book by its title",
        "parameters": [
          {
            "type": "string",
            "name": "title",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Retrieved",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        }
      },
      "post": {
        "summary": "request a book be added to the library",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          }
        }
      }
    }
  },
  "definitions": {
    "Book": {
      "description": "A JSON object containing book information",
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
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
    "title": "Library",
    "version": "0.0.1"
  },
  "paths": {
    "/book": {
      "get": {
        "summary": "request a book by its title",
        "parameters": [
          {
            "type": "string",
            "name": "title",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Retrieved",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        }
      },
      "post": {
        "summary": "request a book be added to the library",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          }
        }
      }
    }
  },
  "definitions": {
    "Book": {
      "description": "A JSON object containing book information",
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}`))
}
