# Ecommerce Go API

A RESTful API built in Go (1.24) for a sample e-commerce (Northwind) database. It uses:
- gin-gonic for HTTP routing
- SQLC for type-safe database queries
- pgx as the Postgres driver
- Viper for configuration management
- migrate for database migrations

> **Note:** This project demonstrates CRUD operations, search/filtering, and pagination using real-world schema.

## Features

- CRUD operations for Customers (create, retrieve, update, delete)
- Search customers by company name and city
- List customers with pagination
- Database migrations using SQL scripts
- SQLC-generated Go code for all queries (auto-generated, do not edit)
- Environment-based configuration (.env files)
- Unit tests for database queries

## Project Structure

```
.
├── Makefile            # Common tasks (run DB, migrations, codegen, tests, server)
├── app.env             # Environment variables for development
├── test.env            # Environment variables for testing
├── sqlc.yaml           # Configuration for sqlc code generation
├── main.go             # Application entry point
├── api/                # HTTP handlers and routing (Gin)
├── db/                 # Database resources
│   ├── migration/      # SQL migration scripts
│   ├── queries/        # SQL files for sqlc generation
│   └── sqlc/           # Generated Go code by sqlc (do not edit)
└── util/               # Utilities (config loading, helpers, random data)
```

## Prerequisites

- Go 1.24+
- Docker
- [migrate CLI](https://github.com/golang-migrate/migrate)
- [sqlc CLI](https://github.com/kyleconroy/sqlc)
- make

## Getting Started

1. Copy the development environment template:
   ```bash
   cp app.env .env
   ```
2. Start a Postgres container and create the database:
   ```bash
   make postgres
   make createdb
   ```
3. Apply database migrations:
   ```bash
   make migrateup
   ```
4. Generate or update SQLC code (after modifying queries):
   ```bash
   make sqlc
   ```
5. Run the API server:
   ```bash
   make server
   ```
   The server listens on the address specified in `HTTP_SERVER_ADDRESS` (default `0.0.0.0:8080`).

## API Endpoints

### Customer

| Method | Endpoint                          | Description                        |
| ------ | --------------------------------- | ---------------------------------- |
| POST   | /api/customer                     | Create a new customer              |
| GET    | /api/customer/:customer_id        | Retrieve customer by ID            |
| PUT    | /api/customer/:customer_id        | Update an existing customer        |
| DELETE | /api/customer/:customer_id        | Delete a customer                  |
| GET    | /api/customers/list?page_id=&page_size= | List customers with pagination |
| GET    | /api/customer/company?company_name=     | Search customers by company    |
| GET    | /api/customer/city?city=               | List customers by city         |

## Testing

Run all tests:
```bash
make test
```

## Code Generation

- SQL migrations: plain SQL files in `db/migration/`
- SQLC: configured via `sqlc.yaml`, SQL files in `db/queries/`
- Generated code appears under `db/sqlc/` (do not edit)

## License

*Specify license here (e.g., MIT, Apache 2.0, etc.)*