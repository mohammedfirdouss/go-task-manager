# Task Manager CLI

A simple command-line interface (CLI) task manager application built with Go and Cobra. This application allows users to manage their tasks efficiently through a set of commands.

## Features

- Add new tasks
- List all tasks
- Mark tasks as completed
- Persistent storage of tasks in a file

## Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```
   cd cli
   ```
3. Install the dependencies:
   ```
   go mod tidy
   ```

## Installing
First, use `go get` to install the latest version of the library.

```go
go get -u github.com/spf13/cobra@latest
```
Then, import Cobra into your application:

```go
import "github.com/spf13/cobra"
```
