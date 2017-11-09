# Working on the Bridge API

The Guruvest Bridge API is a REST (over HTTP or WebSocket) API used by third-party clients to communicate with the daemon.

It consists of various components in this repository:

- `api/swagger.yaml` A Swagger definition of the API.
- `daemon/` The daemon, which serves the API.

## Swagger definition

The API is defined by the [Swagger](http://swagger.io/specification/) definition in `api/swagger.yaml`. This definition can be used to:

1. Automatically generate documentation.
2. Automatically generate the Go server and client. (A work-in-progress.)
3. Provide a machine readable version of the API for introspecting what it can do, automatically generating clients for other languages, etc.

## Updating the API documentation
The API documentation is generated entirely from `api/swagger.yaml`. If you make updates to the API, edit this file to represent the change in the documentation.

- `definitions`, which defines re-usable objects used in requests and responses
- `paths`, which defines the API endpoints (and some inline objects which don't need to be reusable)

To make an edit, first look for the endpoint you want to edit under `paths`, then make the required edits. Endpoints may reference reusable objects with `$ref`, which can be found in the `definitions` section.