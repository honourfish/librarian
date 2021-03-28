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
    },
    "/book/{title}": {
      "get": {
        "summary": "request a book by its title",
        "parameters": [
          {
            "type": "string",
            "name": "title",
            "in": "path",
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
      "put": {
        "summary": "request a book be updated",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          }
        }
      },
      "delete": {
        "summary": "delete a book by its title",
        "parameters": [
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Deleted"
          }
        }
      }
    },
    "/librarian/{username}/book": {
      "post": {
        "summary": "request a book be added to the library",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/book/{title}/{author}": {
      "get": {
        "summary": "request a books stock information by its title and author",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "author",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Retrieved",
            "schema": {
              "$ref": "#/definitions/BookStock"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/book/{title}/{author}/{copies}": {
      "delete": {
        "summary": "delete a book by its title/author",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "author",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "name": "copies",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Deleted"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/user": {
      "post": {
        "summary": "request a user be added to the library",
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/user/{user}": {
      "get": {
        "summary": "request a user by their username",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Retrieved",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "delete": {
        "summary": "delete a user by their username",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Deleted"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/user/{user}/checkout": {
      "put": {
        "summary": "request a book to be checked out by a user",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "title",
                "author"
              ],
              "properties": {
                "author": {
                  "type": "string"
                },
                "title": {
                  "type": "string"
                }
              }
            }
          },
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          },
          "404": {
            "description": "Not Found"
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
        "copies": {
          "type": "integer"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "BookStock": {
      "description": "A JSON object containing book stock information",
      "type": "object",
      "properties": {
        "checkedout": {
          "type": "integer"
        },
        "copies": {
          "type": "integer"
        },
        "instock": {
          "type": "boolean"
        }
      }
    },
    "User": {
      "description": "A JSON object containing user information",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "username": {
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
    },
    "/book/{title}": {
      "get": {
        "summary": "request a book by its title",
        "parameters": [
          {
            "type": "string",
            "name": "title",
            "in": "path",
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
      "put": {
        "summary": "request a book be updated",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          }
        }
      },
      "delete": {
        "summary": "delete a book by its title",
        "parameters": [
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Deleted"
          }
        }
      }
    },
    "/librarian/{username}/book": {
      "post": {
        "summary": "request a book be added to the library",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/book/{title}/{author}": {
      "get": {
        "summary": "request a books stock information by its title and author",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "author",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Retrieved",
            "schema": {
              "$ref": "#/definitions/BookStock"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/book/{title}/{author}/{copies}": {
      "delete": {
        "summary": "delete a book by its title/author",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "title",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "author",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "name": "copies",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Deleted"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/user": {
      "post": {
        "summary": "request a user be added to the library",
        "parameters": [
          {
            "name": "user",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/user/{user}": {
      "get": {
        "summary": "request a user by their username",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Retrieved",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "delete": {
        "summary": "delete a user by their username",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Deleted"
          },
          "404": {
            "description": "Not Found"
          }
        }
      }
    },
    "/librarian/{username}/user/{user}/checkout": {
      "put": {
        "summary": "request a book to be checked out by a user",
        "parameters": [
          {
            "name": "book",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "title",
                "author"
              ],
              "properties": {
                "author": {
                  "type": "string"
                },
                "title": {
                  "type": "string"
                }
              }
            }
          },
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          },
          "404": {
            "description": "Not Found"
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
        "copies": {
          "type": "integer"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "BookStock": {
      "description": "A JSON object containing book stock information",
      "type": "object",
      "properties": {
        "checkedout": {
          "type": "integer"
        },
        "copies": {
          "type": "integer"
        },
        "instock": {
          "type": "boolean"
        }
      }
    },
    "User": {
      "description": "A JSON object containing user information",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  }
}`))
}
