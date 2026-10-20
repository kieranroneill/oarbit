<div align="center">

[![License: CC0-1.0](https://img.shields.io/github/license/kieranroneill/moody-376-assistant)][license]

</div>

<h1 align="center">
  Oarbit
</h1>

<p align="center">
  Oarbit is an onboard knowledge and maintenance copilot that centralizes technical documents, service logs, and operational notes for your vessel.
</p>

---

### Table of contents

* [1. Overview](#-1-overview)
  - [1.1. Project structure](#11-project-structure)
* [2. Usage](#-2-usage)
  - [2.1. Requirements](#21-requirements)
  - [2.2. Getting Started](#22-getting-started)
* [3. Development](#-3-development)
  - [3.1. With Docker (Recommended)](#31-with-docker-recommended)
    - [3.1.1. Requirements](#311-requirements)
    - [3.1.2. Start Docker](#312-start-docker)
* [4. Appendix](#-4-appendix)
  - [4.1. Useful commands](#41-useful-commands)
  - [4.2. Database migrations](#42-database-migrations)
* [5. How to contribute](#-5-how-to-contribute)
* [6. License](#-6-license)

## 🔭 1. Overview

### 1.1. Project structure

The project structure is based on the layout outlined in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

```text
.
├─ cmd/
│   ├── api/
│   │   └── main.go                         <-- API entrypoint
│   └── ...
├─ build/
│   ├── package/                            <-- Docker image files and entrypoint scripts
│   │   ├── <service-name>/
│   │   │   ├── Dockerfile
│   │   │   └── Dockerfile.development
│   │   └── ...
│   └── ...
├─ configs/                                 <-- Configuration files and environment variables
│   ├── <service-name>/
│   │   ├── env.dev
│   │   ├── env.example
│   │   ├── env.example
│   │   └── ...
│   └── ...
├─ deployments/                             <-- Container orchestration configurations
│   ├── compose.yml
│   ├── compose.development.yml
│   └── compose.test.yml
├─ internal/                                <-- Internal packages - uses a functional layering approach
│   ├── <function-name>/
│   │   ├── constants.go
│   │   ├── dtos.go
│   │   ├── models.go
│   │   ├── repositories.go
│   │   ├── routes.go
│   │   ├── services.go
│   │   └── ...
│   └── ...
├─ migrations/                              <-- Goose migrations
│   ├── 19700101000000_create_a_table.sql
│   └── ...
├─ web/
│   ├── main.ts                             <-- Web UI entrypoint
│   └── ...
├── commitlint.config.ts                    <-- Commitlint configuration
├── LICENSE                                 <-- Project license
├── Makefile                                <-- Make commands
├── package.json                            <-- JavaScript dependencies for tools and web UI
├── prettier.config.ts                      <-- Prettier configuration for formatting JS/TS files
├── README.md
├── VERSION
└── ...
```

<sup>[Back to top ^][table-of-contents]</sup>

## 🪄 2. Usage

#### 2.1. Requirements

- [Docker](https://docs.docker.com/engine/install/)

<sup>[Back to top ^][table-of-contents]</sup>

#### 2.2. Getting Started

1. Add a file the relevant `.env` files to the `configs/` directory for each of the services.

2. Using Docker compose, you can run the orchestration file using:

```bash
$ make start
```

<sup>[Back to top ^][table-of-contents]</sup>

## 🛠️ 3. Development

### 3.1. With Docker (Recommended)

#### 3.1.1. Requirements

- [Docker](https://docs.docker.com/engine/install/)

<sup>[Back to top ^][table-of-contents]</sup>

#### 3.1.2. Start Docker

1. Using Docker compose, you can run the orchestration file using:

```bash
$ make dev
```

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 4. Appendix

### 4.1. Useful commands

| Command                    | Description                                                         |
|----------------------------|---------------------------------------------------------------------|
| `make dev`                 | Runs the platform in development mode via Docker.                   |
| `make format`              | Formats all soruce code files.                                      |
| `make format_go`           | Formats all Go soruce code files.                                   |
| `make format_js`           | Formats all JS/TS soruce code files.                                |
| `make install`             | Installs all the dependencies - including development dependencies. |
| `make install_go_deps`     | Installs the Go dependencies.                                       |
| `make install_go_dev_deps` | Installs the Go dev dependencies.                                   |
| `make install_js_deps`     | Installs the JS/TS dependencies.                                    |
| `make migrations_down`     | Downgrades the database migrations by a version.                    |
| `make migrations_down_dev` | Downgrades the database migrations by a version (development only). |
| `make migrations_up`       | Upgrades the database migrations.                                   |
| `make migrations_up_dev`   | Upgrades the database migrations (development only).                |
| `make start`               | Runs the platform in production mode via Docker.                    |
| `make test`                | Runs all tests.                                                     |
| `make test_go_unit`        | Runs Go unit tests.                                                 |
| `make test_unit`           | Runs all unit tests.                                                |

<sup>[Back to top ^][table-of-contents]</sup>

### 4.2. Database migrations

Database migrations are managed with [Goose](https://github.com/pressly/goose) and run through Docker, so they use the same environment as the API service.

Use the following commands to create, apply, or roll back migrations (assuming you have installed Goose).

To create a new migration to the `migrations/` directory, use:

```shell
goose create add_boat_profiles_table sql
```

This creates a SQL file to as `migrations/YYYYMMDDHHmmss_add_boat_profiles_table.sql`. Add the necessary SQL statements to the generated migration file, then use the following command to apply the migration:
```shell
make migrations_up
```

To downgrade/rollback to the previous version, use:
```shell
make migrations_down
```

For development, the matching commands are:

```shell
make migrations_up_dev
make migrations_down_dev
```

- `migrations_up` applies all pending migrations in the production Compose setup.
- `migrations_down` rolls back the latest migration in the production Compose setup.
- `migrations_up_dev` applies all pending migrations in the development Compose setup.
- `migrations_down_dev` rolls back the latest migration in the development Compose setup.

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 5. How to contribute

Please read the [**contributing guide**](https://github.com/kieranroneill/moody-376-assistant/blob/main/CONTRIBUTING.md) to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 6. License

Please refer to the [LICENSE][license] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- links -->

[license]: https://github.com/kieranroneill/moody-376-assistant/blob/main/LICENSE
[table-of-contents]: #table-of-contents
