# Golang Sample App Blueprint

This project is a blueprint for starting a new Golang project from scratch. It can help you kick off a new Golang project with a solid foundation.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Running the Project](#running-the-project)
- [Project Structure](#project-structure)
- [Dependency Injection: Adding Custom Dependencies](#dependency-injection-adding-custom-dependencies)
  - [Step  1: Define Your Dependency](#step-1-define-your-dependency)
  - [Step  2: Add the Dependency to the Container](#step-2-add-the-dependency-to-the-container)
  - [Step  3: Inject Your Dependency](#step-3-inject-your-dependency)
- [Setting Up Routes](#setting-up-routes)
  - [Defining Routes](#defining-routes)
  - [Adding Additional Routes](#adding-additional-routes)
  - [Testing Routes](#testing-routes)
  - [Key Features of Chi](#key-features-of-chi)
- [Contributing](#contributing)
## Features

- **Dependency Injection**: The project uses the [Uber Dig](https://github.com/uber-go/dig) project for dependency injection, which helps manage dependencies in a clean and maintainable way.
- **CLI Applications**: [Cobra](https://github.com/spf13/cobra) is integrated for creating powerful, modern CLI applications. Cobra provides a simple interface for creating powerful modern CLI applications and snake case command handling.
- **HTTP Router**: The [Chi](https://github.com/go-chi/chi) router is used for HTTP request routing, offering a lightweight, idiomatic, and composable router for building Go HTTP services.
- **Configuration Management**: [Viper](https://github.com/spf13/viper) is employed as a complete configuration solution, allowing you to manage your application's configuration with ease.

## Getting Started

To get started with this project, you'll need to have Go installed on your machine. Once you have Go set up, you can clone this repository and start building your application.

### Prerequisites

- Ensure you have Go version  1.22 or higher installed on your system. You can check your Go version by running `go version` in your terminal.
- Familiarize yourself with Go modules, as this project uses Go modules for dependency management.

### Installation

1. **Clone the Repository**: First, clone this repository to your local machine.
```
git clone https://github.com/majidypd/goblueprint.git
cd goblueprint
```

2. **Install Dependencies**: This project uses Go modules for dependency management. Run the following command to download all necessary dependencies:
```
go mod download
```

## Running the Project
To run the project, execute the following command from the project's root directory:
```
go run main.go serv 
```
This command will compile and run the project. Ensure that the go.mod file is present in the project's root directory for the dependencies to be correctly managed.

## Project Structure
```
goblueprint/
├── cmd/                  
│   └── di.go             # Implements dependency injection
│   └── http_server.go    # Executes the HTTP server command
│   └── root.go           # Serves as the root directory for all commands
├── config/               # Configuration files and utilities
│   └── config.go         # Configuration loading and parsing
├── internal/             # Internal packages
│   ├── app/              # Contains business logic and services
│   │   └── health.go     # Example of a health check controller
│   ├── database/         # Houses database migrations
│   │   └── migration.go  # Example of a database migration script
│   ├── middleware/       # Contains API middleware
│   │   └── sample.go     # Example middleware implementation
│   ├── repository/       # Repository layer for data access
│   │   └── sample.go     # Example repository implementation
│   ├── router/           # Defines API routes
│   │   └── routes.go     # API route definitions
│   └── utils/            # Includes utility functions and helpers
│       └── utils.go      # Example utility function
├── go.mod                # Go module definition
└── README.md             # Project README
└── main.go               # Application entry points
```

## Dependency Injection: Adding Custom Dependencies
goblueprint utilizes the [Uber Dig](https://github.com/uber-go/dig) package for dependency injection, allowing for a clean and flexible way to manage dependencies across your application. This section will guide you on how to add your own dependencies to the dig container.

### Step 1: Define Your Dependency
First, you need to define the dependency you want to add. This involves creating a struct or function that your application requires. For example, let's say you want to add a custom `Database` service:
```go
package database

type Database struct {
    // Database connection details
}

func NewDatabase() *Database {
    return &Database{}
}
```
### Step 2: Add the Dependency to the Container
Next, you need to add your new dependency to the `dig` container. This is done in the `NewDig` function within the `cmd` package. Here's how you can modify the `NewDig` function to include your `Database` service:
```go
package cmd

import (
    "go.uber.org/dig"

    "github.com/majidypd/goblueprint/database"
)

// NewDig service provider...
func NewDig() *dig.Container {
    container := dig.New()
    if err := container.Provide(database.NewDatabase); err != nil {
        panic(err)
    }
    // Add other dependencies as needed...
    return container
}
```

### Step 3: Inject Your Dependency
Finally, you can inject your new dependency into any part of your application that requires it. For example, if you have a service that needs to access the database, you can do so like this:
```go
package service

import (
    "context"
    "github.com/majidypd/goblueprint/database"
)

type MyService struct {
    DB *database.Database
}

func NewMyService(db *database.Database) *MyService {
    return &MyService{DB: db}
}

func (s *MyService) DoSomething(ctx context.Context) {
    // Use s.DB to interact with the database...
}
```
And then, ensure your `NewMyService` function is provided to the `dig` container in the `NewDig` function:
```go
if err := container.Provide(service.NewMyService); err != nil {
    panic(err)
}

```
By following these steps, you can easily add and manage your own dependencies within the `goblueprint` project using the dig package for dependency injection.

## Setting Up Routes
`goblueprint` leverages the [Chi](https://github.com/go-chi/chi) router to manage HTTP request routing. This section explains how to set up and utilize the router within the project.

### Defining Routes
To define routes, you use the `NewRouteHandler` function in the `routes` package. This function initializes a new `ApiRouteHandler` instance and sets up the router with the specified routes. Here's an example of how to define a route for a health check endpoint:

```go
package routes

import (
    "github.com/go-chi/chi"
    "github.com/majidypd/goblueprint/internal/app"
)

type ApiRouteHandler struct {
    Router *chi.Mux
}

// NewRouteHandler constructor of RouteHandler
func NewRouteHandler(healthHandler *app.HealthHandler) *ApiRouteHandler {
    apiRouteHandler := &ApiRouteHandler{
        Router: chi.NewRouter(),
    }

    apiRouteHandler.Router.Get("/health", healthHandler.Health)
    return apiRouteHandler
}
```
In this example, the `NewRouteHandler` function takes a `HealthHandler` as a parameter and uses it to define a route for the `/health` endpoint. The `Health` method of `HealthHandler` is set as the handler for this route.

### Adding Additional Routes
To add more routes, you can extend the `NewRouteHandler` function to accept additional handlers and attach them to the router. For instance, if you have a `SampleHandler` with a `Foo` method that should handle requests to `/sample`, you can modify the `NewRouteHandler` function like this:

```go
func NewRouteHandler(healthHandler *app.HealthHandler, sampleHandler *app.SampleHandler) *ApiRouteHandler {
    apiRouteHandler := &ApiRouteHandler{
        Router: chi.NewRouter(),
    }

    apiRouteHandler.Router.Get("/health", healthHandler.Health)
    apiRouteHandler.Router.Get("/sample", sampleHandler.Foo) // Attach the new handler to the router

    return apiRouteHandler
}
```
### Testing Routes
To ensure your routes work as expected, you should write tests that simulate HTTP requests to your endpoints and verify the responses. This can be done using Go's `httptest` package, which allows you to create a test HTTP server and make requests to it.
### Key Features of Chi
- Lightweight: Chi is designed to be small and efficient, with less than 1000 lines of code.
- Fast: It's optimized for performance, ensuring your application responds quickly to requests.
- Composable: You can easily compose routes, middlewares, and sub-routers to build complex routing structures.
- Middleware Support: Chi allows you to add middleware to your routes, enabling features like logging, authentication, and more.

For more advanced routing techniques, such as grouping routes, mounting subrouters, and handling wildcards, refer to the [chi documentation](https://github.com/go-chi/chi)

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to improve this blueprint.
