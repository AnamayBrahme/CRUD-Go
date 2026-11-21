\
# CRUD-Go

A small CRUD application built with Go. The project is being developed to practise Go fundamentals, basic HTTP handling, and connecting a simple single-page website to a Go server.

## Project Status

**Current stage:** Sprint 1 — Go Fundamentals and In-Memory CRUD Logic.

The current version runs in the terminal. The HTTP API and single-page website are planned for Sprint 2, and testing plus final documentation are planned for Sprint 3.

## Project Goals

- Practise basic Go syntax and programming concepts.
- Represent service information using a Go struct.
- Store service records temporarily in a Go map.
- Implement Create, Read, Update, and Delete operations.
- Add basic HTTP endpoints and a single-page website in the next sprint.

## Current Features

- Define a service with an ID, name, environment, and status.
- Store services in an in-memory map.
- Create a service.
- List all services.
- Update a service's status.
- Delete a service.
- Demonstrate the CRUD operations from `main.go`.

> **Note:** Data is stored only in memory. It is lost when the application stops.

## Technologies

- Go
- Go standard library (`fmt`)
- Git and GitHub

Planned for Sprint 2: `net/http`, `encoding/json`, HTML, CSS, and JavaScript.

## Project Structure

```text
CRUD-Go/
├── main.go       # Application entry point and example execution
├── service.go    # Service struct and CRUD functions
├── go.mod        # Go module definition
└── README.md     # Project documentation
```

## Requirements

- Go installed
- Git installed
- A code editor such as Visual Studio Code

## Setup

Clone the repository:

```bash
git clone https://github.com/AnamayBrahme/CRUD-Go.git
cd CRUD-Go
```

If you are setting up the project locally for the first time instead of cloning it, initialize the module using the module path in `go.mod`.

## Run the Application

From the project root directory:

```bash
go run .
```

The program creates example services, lists them, updates one service, deletes another service, and displays the remaining services.

## Format and Build

Format the Go source files:

```bash
gofmt -w .
```

Build the application:

```bash
go build .
```

## Current CRUD Operations

| Operation | Function | Description |
|---|---|---|
| Create | `createService` | Adds a service to the map using its ID. |
| Read | `listServices` | Displays all services stored in the map. |
| Update | `updateService` | Changes the status of an existing service. |
| Delete | `deleteService` | Removes a service from the map. |

## Go Concepts Practised

- Variables and constants: `var`, `const`, and `:=`
- Basic data types: `string`, `int`, `float64`, and `bool`
- Conditional statements: `if`, `else if`, and `else`
- Loops: `for` and `range`
- Functions and parameters
- Structs
- Maps
- Slices (introductory practice)
- Packages and imports
- Multiple Go files in the same package

Pointers, function return values, and basic error returns are part of the remaining Sprint 1 learning tasks.

## Development Plan

### Sprint 1 — Environment Setup and Go Fundamentals

Set up the Go project, practise the language fundamentals, and implement in-memory CRUD operations.

### Sprint 2 — Basic HTTP Server and Single-Page Website

Add basic HTTP handlers for Create, Read, Update, and Delete. Connect them to a single HTML page using JavaScript `fetch()`.

### Sprint 3 — Testing and Documentation

Add basic Go tests, verify the application, complete the README, and prepare a project demonstration.

## Scope

This is intentionally a beginner-level learning project. It will use Go's standard library and simple in-memory storage. It does not aim to be a production-ready service management platform.

No database, external web framework, Docker, Kubernetes, authentication, or complex architecture is planned.

## Author

Anamay Brahme
```
