# apierr

`apierr` is a lightweight, framework-agnostic error-handling package for Go. It simplifies API error management by standardizing error definitions, automating error code generation, and supporting consistent error responses across various frameworks.

## Features

- **Centralized Error Definitions**: Define all your errors in a single YAML file.
- **Automated Code Generation**: Generate reusable error structs, constants, and helper methods with a single command.
- **Retry-ability Support**: Specify whether errors are retryable, helping clients handle errors more effectively.
- **Framework Agnostic**: Works seamlessly with popular frameworks like Echo, Gin, and `net/http`.
- **Customizable Responses**: Define developer-facing and user-facing error messages separately.
- **Dynamic Argument Handling**: Pass and retrieve dynamic arguments for detailed error descriptions.

## Installation
To use the package, run:
```bash
go get github.com/LakshyaNegi/apierr

## License

Copyright © 2025 Lakshya Negi.

This project is under MIT license.