# Golang Project

This is a Golang project that utilizes several packages for different purposes, including validation, authentication, database management, environment configuration, and more. Below is a brief overview of the key dependencies used in this project:

## Dependencies

### 1. [github.com/go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
A package that provides a simple and flexible way to validate data in Go applications. It supports built-in validators, custom validation functions, and chaining of validation rules.

### 2. [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
This package is used for working with JSON Web Tokens (JWTs) in Go. It provides functionality to create, parse, and validate JWT tokens, typically for user authentication and authorization.

### 3. [github.com/google/uuid](https://github.com/google/uuid)
A Go package for generating and handling UUIDs (Universally Unique Identifiers). This is useful for generating unique identifiers for entities, such as database records.

### 4. [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)
A collection of cryptographic algorithms and utilities that are not part of the Go standard library. This includes hashing, encryption, and key management functions for secure data handling.

### 5. [gorm.io/driver/postgres](https://gorm.io/docs/connecting_to_the_database.html)
The PostgreSQL driver for GORM, a popular ORM (Object-Relational Mapping) library for Go. It simplifies database interactions by allowing you to define models in Go and automatically mapping them to database tables.

### 6. [gorm.io/gorm](https://gorm.io/)
GORM is an ORM framework for Golang, providing an easy way to work with databases using Go structs. It supports complex queries, relationships, migrations, and more.

### 7. [github.com/spf13/cobra](https://github.com/spf13/cobra)
A powerful CLI library for Go applications. Cobra makes it easy to build command-line applications by supporting commands, flags, and argument parsing, allowing you to build sophisticated CLI tools.

### 8. [github.com/spf13/viper](https://github.com/spf13/viper)
Viper is a configuration management package that works well with JSON, TOML, YAML, HCL, and other formats. It allows easy reading and parsing of configuration files, environment variables, and command-line flags.

### 9. [github.com/subosito/gotenv](https://github.com/subosito/gotenv)
A simple package for loading environment variables from `.env` files into the Go application. This is especially useful for managing environment-specific configuration without hardcoding values.

## Setup

1. Clone this repository.
2. Install the required dependencies using `go get`.
3. Copy the `.env-example` file to `.env` and fill in the required environment variables.
4. Configure your environment variables using `.env` or by setting them manually.

## Running the Project

1. Use the Cobra package and commands located in the `cmd/root.go` directory to interact with the project.
2. To run the project, execute the following command in your terminal:
   
   ```bash
   go run cmd/root.go runapp
