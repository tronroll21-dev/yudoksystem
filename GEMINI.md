# GEMINI.md

## Project Overview

This project is a web application for managing daily sales reports, named "yudoksystem". It is built with a Go backend and a frontend that uses Alpine.js and Tailwind CSS. The application allows users to view, edit, and create daily sales reports, as well as upload sales data from TSV files.

The backend is a Go application that uses the Gin web framework to provide a RESTful API. It connects to a MySQL database to store and retrieve sales data. The database schema uses a mix of Japanese and English names for tables and columns.

The frontend is a single-page application that uses Alpine.js for interactivity and Tailwind CSS for styling. The user interface is in Japanese.

## Building and Running

### Prerequisites

- Go (version 1.20 or higher)
- MySQL
- Task (a task runner for Go)
- Node.js and npm (for browser-sync)

### Environment Variables

The application requires a `.env` file with the following variables:

```
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_NAME=your_database_name
```

### Development

To run the application in a development environment, you can use the `dev` task defined in the `Taskfile.yml` file. This will start the Go backend with `fresh` for live reloading, and `browser-sync` to proxy the backend and watch for changes in the asset files.

```bash
task dev
```

The application will be available at `http://localhost:3000`.

### Building

To build the application, you can use the `go build` command:

```bash
go build -o yudoksystem main.go
```

This will create an executable file named `yudoksystem` in the root directory.

## Development Conventions

### Code Style

The Go code follows standard Go conventions. The frontend code uses Alpine.js and Tailwind CSS.
Tailwind CSS is to be used for any html pages in the assets folder.
Alpine.js is to be used to add interactivity to any html pages in the assets folder.
Tailwind CSS is not be used for the templates in the templates, which are meant to eventually be rendered into PDF using wkhtmltopdf. Use vanilla CSS or Bulma for the templates.
Keep comments to a minimum. Only comment non-trivial code.
Any code for interactions with the database should be in the models package.
The application user interacts with the application using http requests, which invoke handlers in the controllers package, which in turn invoke methods in the models package.

### Database

The application uses a MySQL database. The database schema is not defined in the codebase, but the `models` package provides functions for interacting with the database. The table and column names are in Japanese.

### API

The backend provides a RESTful API for interacting with the sales data. The API endpoints are defined in `main.go` and the handlers are in the `controllers` package.

### Frontend

The frontend is a single-page application that uses Alpine.js for interactivity and Tailwind CSS for styling. The main page is `assets/index.html`.

### Dependencies

The Go dependencies are managed with Go modules and are listed in the `go.mod` file. The frontend dependencies are included via CDNs in the `assets/index.html` file.
