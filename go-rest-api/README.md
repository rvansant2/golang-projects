# Go REST API
## Required
Go Version 1.10.x - 1.11.x (not tested with version 1.13) or Docker
## Description
This is a simple example of a RESTful API following RESTful principles - see below.

## [Guiding Principles of REST](https://restfulapi.net/)
1. Client–server – By separating the user interface concerns from the data storage concerns, we improve the portability of the user interface across multiple platforms and improve scalability by simplifying the server components.
2. Stateless – Each request from client to server must contain all of the information necessary to understand the request, and cannot take advantage of any stored context on the server. Session state is therefore kept entirely on the client.
3. Cacheable – Cache constraints require that the data within a response to a request be implicitly or explicitly labeled as cacheable or non-cacheable. If a response is cacheable, then a client cache is given the right to reuse that response data for later, equivalent requests.
4. Uniform interface – By applying the software engineering principle of generality to the component interface, the overall system architecture is simplified and the visibility of interactions is improved. In order to obtain a uniform interface, multiple architectural constraints are needed to guide the behavior of components. REST is defined by four interface constraints: identification of resources; manipulation of resources through representations; self-descriptive messages; and, hypermedia as the engine of application state.
5. Layered system – The layered system style allows an architecture to be composed of hierarchical layers by constraining component behavior such that each component cannot “see” beyond the immediate layer with which they are interacting.
6. Code on demand (optional) – REST allows client functionality to be extended by downloading and executing code in the form of applets or scripts. This simplifies clients by reducing the number of features required to be pre-implemented.

### Motivation
- **Originally coded in March 2018**
- A desire to code and play with the Go Language
- Always liked coding APIs
- Simple REST API to save cooking recipes
- Simple security of an API key and protected routing of API endpoints

### Test
- Run command via `go run test`

### How to Run
- Run API `go run main.go` with a local install of Go or run `go build` or use Docker with `docker-compose up -d` (optional flag for development).
- Use Postman or Paw HTTP tool and set the header to have a key/value of `api-key` and value of `12345`.
- Go to `http://localhost:8080/recipes` to get a list of current recipes or if you're running the API via Docker go to `http://localhost/recipes`

#### TODO
- Multi-stage Docker builds
- Record duplication
- Swagger documentation for REST endpoints
- Refactor...