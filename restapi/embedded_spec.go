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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This API is for a coding interview challenge for Canonical's snap team.",
    "title": "snap API challenge",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/api/v0",
  "paths": {
    "/alive": {
      "get": {
        "description": "Any non-200 response means the service is not alive.",
        "tags": [
          "system"
        ],
        "summary": "Used by Caddy or other reverse proxy to determine if the service is alive.",
        "operationId": "alive",
        "responses": {
          "200": {
            "description": "Service is alive."
          }
        }
      }
    },
    "/group": {
      "post": {
        "description": "Insert all given quota-groups to the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Insert quota-groups.",
        "operationId": "groupInsert",
        "parameters": [
          {
            "description": "The names of the quota-groups to insert. Order matters. If a quota-group is referenced as a member before it was inserted, an error will occur.",
            "name": "groups",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/QuotaGroup"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The group was successfully inserted."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete all the given quota-groups from the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Delete quota-groups.",
        "operationId": "groupDelete",
        "parameters": [
          {
            "description": "The names of the quota-groups to delete.",
            "name": "groups",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The group was successfully deleted."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/group/limits": {
      "get": {
        "description": "Given an array of quota-groups, return all of their limits in a map.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the resource limits for the given quota-groups.",
        "operationId": "groupLimitRead",
        "parameters": [
          {
            "description": "The name of the quota-groups to get the resource limits for.",
            "name": "groups",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The map of quota-group names to resource limits.",
            "schema": {
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Limits"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Given a map of quota-group names to resource limits, set the resource limits for the quota-groups on the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Set the resource limits for the given quota-groups.",
        "operationId": "groupLimitWrite",
        "parameters": [
          {
            "description": "The mapping of quota-group names to resource limits.",
            "name": "groupLimitsMap",
            "in": "body",
            "required": true,
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/Limits"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The resource limits were successfully set for the quota-group."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/group/members": {
      "get": {
        "description": "Given an array of quota-group names, return a mapping of quota-group names to members.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the members of quota-groups.",
        "operationId": "groupMembersRead",
        "parameters": [
          {
            "description": "The names of the quota-groups to get members from.",
            "name": "groups",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The mapping of quota-group names its members.",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/GroupMembers"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Given a mapping of quota-group names to members, add the members to the quota-groups.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Add members to the quota-group.",
        "operationId": "groupMembersAdd",
        "parameters": [
          {
            "description": "The mapping of quota-group names to the snaps and member quota-groups to add.",
            "name": "groupMembers",
            "in": "body",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/GroupMembers"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The members of the group were successfully added."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Given a map of quota-groups to members, disassociate the given members from their associated quota-groups on the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Delete members from the quota-group.",
        "operationId": "groupMembersDelete",
        "parameters": [
          {
            "description": "The mapping of quota-group names to the snaps and member quota-groups to remove.",
            "name": "groupMembers",
            "in": "body",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/GroupMembers"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The members of the group were successfully deleted."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/group/usage": {
      "get": {
        "description": "Given an array of quota-groups, create a mapping of quota-groups to resource usage.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the quota usage information for a quota-group.",
        "operationId": "groupUsage",
        "parameters": [
          {
            "description": "The names of the quota-groups to get the usage information for.",
            "name": "groups",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The a mapping of quota-group names to usage information.",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/Usage"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "x-nullable": false
        },
        "message": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "GroupMembers": {
      "type": "object",
      "properties": {
        "memberGroups": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "snaps": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Limits": {
      "properties": {
        "maxMemory": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "QuotaGroup": {
      "properties": {
        "limits": {
          "$ref": "#/definitions/Limits"
        },
        "name": {
          "type": "string"
        },
        "snaps": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "subGroups": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Usage": {
      "additionalProperties": {
        "properties": {
          "max": {
            "type": "number"
          },
          "usage": {
            "type": "number"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "Endpoints fulfilling the challenge requirements.",
      "name": "api"
    },
    {
      "description": "Endpoints required by the system that are not public facing.",
      "name": "system"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This API is for a coding interview challenge for Canonical's snap team.",
    "title": "snap API challenge",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/api/v0",
  "paths": {
    "/alive": {
      "get": {
        "description": "Any non-200 response means the service is not alive.",
        "tags": [
          "system"
        ],
        "summary": "Used by Caddy or other reverse proxy to determine if the service is alive.",
        "operationId": "alive",
        "responses": {
          "200": {
            "description": "Service is alive."
          }
        }
      }
    },
    "/group": {
      "post": {
        "description": "Insert all given quota-groups to the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Insert quota-groups.",
        "operationId": "groupInsert",
        "parameters": [
          {
            "description": "The names of the quota-groups to insert. Order matters. If a quota-group is referenced as a member before it was inserted, an error will occur.",
            "name": "groups",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/QuotaGroup"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The group was successfully inserted."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete all the given quota-groups from the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Delete quota-groups.",
        "operationId": "groupDelete",
        "parameters": [
          {
            "description": "The names of the quota-groups to delete.",
            "name": "groups",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The group was successfully deleted."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/group/limits": {
      "get": {
        "description": "Given an array of quota-groups, return all of their limits in a map.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the resource limits for the given quota-groups.",
        "operationId": "groupLimitRead",
        "parameters": [
          {
            "description": "The name of the quota-groups to get the resource limits for.",
            "name": "groups",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The map of quota-group names to resource limits.",
            "schema": {
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Limits"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Given a map of quota-group names to resource limits, set the resource limits for the quota-groups on the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Set the resource limits for the given quota-groups.",
        "operationId": "groupLimitWrite",
        "parameters": [
          {
            "description": "The mapping of quota-group names to resource limits.",
            "name": "groupLimitsMap",
            "in": "body",
            "required": true,
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/Limits"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The resource limits were successfully set for the quota-group."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/group/members": {
      "get": {
        "description": "Given an array of quota-group names, return a mapping of quota-group names to members.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the members of quota-groups.",
        "operationId": "groupMembersRead",
        "parameters": [
          {
            "description": "The names of the quota-groups to get members from.",
            "name": "groups",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The mapping of quota-group names its members.",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/GroupMembers"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Given a mapping of quota-group names to members, add the members to the quota-groups.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Add members to the quota-group.",
        "operationId": "groupMembersAdd",
        "parameters": [
          {
            "description": "The mapping of quota-group names to the snaps and member quota-groups to add.",
            "name": "groupMembers",
            "in": "body",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/GroupMembers"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The members of the group were successfully added."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Given a map of quota-groups to members, disassociate the given members from their associated quota-groups on the backend.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Delete members from the quota-group.",
        "operationId": "groupMembersDelete",
        "parameters": [
          {
            "description": "The mapping of quota-group names to the snaps and member quota-groups to remove.",
            "name": "groupMembers",
            "in": "body",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/GroupMembers"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The members of the group were successfully deleted."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/group/usage": {
      "get": {
        "description": "Given an array of quota-groups, create a mapping of quota-groups to resource usage.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the quota usage information for a quota-group.",
        "operationId": "groupUsage",
        "parameters": [
          {
            "description": "The names of the quota-groups to get the usage information for.",
            "name": "groups",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The a mapping of quota-group names to usage information.",
            "schema": {
              "additionalProperties": {
                "$ref": "#/definitions/Usage"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "x-nullable": false
        },
        "message": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "GroupMembers": {
      "type": "object",
      "properties": {
        "memberGroups": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "snaps": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Limits": {
      "properties": {
        "maxMemory": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "QuotaGroup": {
      "properties": {
        "limits": {
          "$ref": "#/definitions/Limits"
        },
        "name": {
          "type": "string"
        },
        "snaps": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "subGroups": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Usage": {
      "additionalProperties": {
        "$ref": "#/definitions/UsageAnon"
      }
    },
    "UsageAnon": {
      "properties": {
        "max": {
          "type": "number"
        },
        "usage": {
          "type": "number"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Endpoints fulfilling the challenge requirements.",
      "name": "api"
    },
    {
      "description": "Endpoints required by the system that are not public facing.",
      "name": "system"
    }
  ]
}`))
}
