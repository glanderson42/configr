# Configr

`configr` is a lightweight Go package for parsing and managing application configurations. It provides a simple, flexible way to set configuration values from environment variables with support for default values and type safety. Use `configr` to ensure that your applications have all required configuration settings in place, with minimal boilerplate code.

## Features

- Parse configuration values from environment variables
- Support for default values if environment variables are not set
- Type-safe handling for common data types (int, string, bool, float, etc.)
- Validation for required configuration fields

## Installation

To install the `configr` package, use the following `go get` command:

```bash
go get github.com/glanderson42/configr
