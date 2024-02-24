# Golang Sample App Blueprint

This project is a blueprint for starting a new Golang project from scratch. It can help you kick off a new Golang project with a solid foundation.

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


This command will compile and run the project. Ensure that the go.mod file is present in the project's root directory for the dependencies to be correctly managed.
## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to improve this blueprint.