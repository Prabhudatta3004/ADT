# Advanced Distributed Table (ADT)

## Overview

The **Advanced Distributed Table (ADT)** project is a key-value store designed for extensibility and integration with distributed systems. It supports basic CRUD (Create, Read, Update, Delete) operations with in-memory storage, data persistence using a file, and configuration management through a simple YAML configuration file. The ADT system also offers a command-line interface (CLI) for easy interaction and testing.

## Features

- **CRUD Operations**: Add, retrieve, update, and delete key-value pairs using simple commands.
- **Data Persistence**: Data is saved to a file and loaded automatically when the system starts, ensuring data retention across sessions.
- **Configuration Management**: Easily configurable using a `config.yaml` file and environment variables through the Viper library.
- **User-Friendly CLI**: Interact with the ADT system using intuitive commands.
- **Thread-Safe Operations**: Ensures data consistency and thread safety using mutexes.

## Project Structure

```
adt-project/
├── main.go                   # Main entry point with a CLI interface for interacting with the key-value store
├── go.mod                    # Go module file for dependency management
├── internal/
│   ├── storage/
│   │   ├── storage.go        # Core CRUD operations for the key-value store with data persistence
│   │   └── storage_test.go   # Unit tests for storage operations (not yet implemented)
├── config/
│   └── config.go             # Configuration management using Viper
├── utils/
│   └── logger.go             # Basic logging utility for consistent logging
├── config.yaml               # Configuration file for data path
├── README.md                 # Project documentation (this file)
└── .gitignore                # Specifies files and directories to be ignored by git
```

## Prerequisites

- [Go](https://golang.org/) installed (v1.16 or higher recommended)
- [Viper](https://github.com/spf13/viper) library for configuration management

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/prabhudatta3004/adt-project.git
   cd adt-project
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configuration**:
   Create a `config.yaml` file in the root directory with the following content:
   ```yaml
   data_path: "./data/adt_data.json"
   ```

## Usage

### Running the Application

To start the application and interact with the key-value store using the CLI, run:

```bash
go run main.go
```

### Supported CLI Commands

- **Set a key-value pair**:
  ```
  set <key> <value>
  ```
  Example:
  ```
  Enter command (set/get/delete/exit): set name John Doe
  Key 'name' set successfully.
  ```
- **Get the value associated with a key**:
  ```
  get <key>
  ```
  Example:
  ```
  Enter command (set/get/delete/exit): get name
  Value for key 'name': John Doe
  ```
- **Delete a key-value pair**:
  ```
  delete <key>
  ```
  Example:
  ```
  Enter command (set/get/delete/exit): delete name
  Key 'name' deleted successfully.
  ```
- **Exit the CLI**:
  ```
  exit
  ```

### Configuration

The application reads configurations from `config.yaml` or environment variables using the Viper library. Below is a sample `config.yaml` file:

```yaml
data_path: "./data/adt_data.json"
```

- `data_path`: Specifies the file path where data is saved. You can change this path to customize where data is stored.

## Code Modules

### `main.go`

Provides an interactive CLI for interacting with the key-value store. Users can input commands to perform CRUD operations.

### `internal/storage/storage.go`

Implements core CRUD operations for an in-memory key-value store with file-based persistence. It includes functions for setting, getting, and deleting key-value pairs, and saving/loading data from a specified file.

### `config/config.go`

Manages application configuration using the Viper library, supporting both file-based and environment variable configurations.

### `utils/logger.go`

Provides basic logging functionality for the application, enabling consistent and formatted logging throughout the project.

## Testing

Unit tests for the key-value store are located in `internal/storage/storage_test.go`. To run tests, use:

```bash
go test ./internal/storage
```

## Troubleshooting

- **Error: `Data file does not exist, initializing new data.`**
  - This message appears when the application is run for the first time and there is no existing data file. This is normal behavior and indicates a new data store is being initialized.

- **Unexpected issues with file paths**
  - Ensure that `data_path` in your `config.yaml` specifies a valid file path.

## Future Enhancements

- **Enhanced Data Structures**: Support for nested and complex data types.
- **Networked Access**: REST API or gRPC endpoints for distributed usage.
- **Advanced Consistency Models**: Distributed data consistency features for multi-node deployments.
- **Performance Optimizations**: Caching mechanisms for frequently accessed data.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

