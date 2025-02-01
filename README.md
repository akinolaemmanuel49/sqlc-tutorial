# sqlc-tutorial

This is a command-line interface (CLI) application written in Go that performs CRUD (Create, Read, Update, Delete) operations on a PostgreSQL database. It uses the `pgx` library for database connectivity and `sqlc` for generating type-safe SQL queries.

## Features

- **Create Authors**: Add new authors to the database with a name and bio.
- **Read Author**: Fetch a single author by ID.
- **Read Authors**: Fetch a list of all authors.
- **Update Author**: Modify an author's name or bio.
- **Delete Author**: Remove an author from the database by ID.
- **Readable Output**: Display author details in a clean, formatted table.

## Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.20 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) (version 12 or higher)
- [sqlc](https://sqlc.dev/) (for generating type-safe SQL queries)
- [golang-migrate](https://github.com/golang-migrate/migrate) (For instructions on how to install golang-migrate cli on your machine [Visit](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate))


## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/akinolaemmanuel49/sqlc-tutorial.git
cd sqlc-tutorial
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory with the following content:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your-db-username
DB_PASSWORD=your-db-password
DB_NAME=tutorial
```

### 3. Set Up the Database

1. Create a PostgreSQL database:
```bash
createdb tutorial
```

2. Create a migrations directory:
```bash
mkdir -p migrations
```

3. Add your first migration file:
```bash
migrate create -ext sql -dir migrations -seq create_authors_table
```
This will create two files in the migrations directory:
  - 000001_create_authors_table.up.sql (for applying the migration)
  - 000001_create_authors_table.down.sql (for reverting the migration)

4. Write SQL migrations:
Edit the generated files to define your schema changes.
- 000001_create_authors_table.up.sql:
  ```sql
  CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    bio  TEXT
  );
  ```
- 000001_create_authors_table.down.sql:
  ```sql
  DROP TABLE authors;
  ```

5. Applying migrations:
- Upgrade migration:
    ```bash
    migrate -path migrations -database "postgres://<DB_USER>:<DB_PASSWORD>@<DB_HOST>:<DB_PORT>/<DB_NAME>?sslmode=disable" up
    ```
- Downgrade migration:
    ```bash
    migrate -path migrations -database "postgres://<DB_USER>:<DB_PASSWORD>@<DB_HOST>:<DB_PORT>/<DB_NAME>?sslmode=disable" down
    ```
Migrations directory should now look like this:
```bash
migrations/
├── 0001_create_authors_table.up.sql
└── 0001_create_authors_table.down.sql
```

Replace the placeholders with your actual database credentials.

### 4. Generate SQLC Code

Run `sqlc` to generate the Go code for your SQL queries:

```bash
sqlc generate
```

This will create the necessary Go files in the `tutorial` package.

### 5. Build and Run the Application

Build the application:

```bash
go build -o sqlc-tutorial
```

Run the application:

```bash
./sqlc-tutorial
```

## Usage

The application supports the following command-line flags:

| Flag   | Description                          | Example                              |
|--------|--------------------------------------|--------------------------------------|
| `-c`   | Create a new author                  | `./sqlc-tutorial -c --name "John Doe" --bio "A passionate writer"` |
| `-r`   | Read authors (all or by ID)          | `./sqlc-tutorial -r` or `./sqlc-tutorial -r --id 1` |
| `-u`   | Update an author's name or bio       | `./sqlc-tutorial -u --id 1 --name "Jane Doe"` or `./sqlc-tutorial -u --id 1 --bio "A phenomenal writer"` |
| `-d`   | Delete an author by ID               | `./sqlc-tutorial -d --id 1`             |
| `-id`  | Specify the author ID                | `./sqlc-tutorial -r --id 1`             |
| `-name`| Specify the author name              | `./sqlc-tutorial -c --name "John Doe"`  |
| `-bio` | Specify the author bio               | `./sqlc-tutorial -c --name --bio "A writer"`   |

### Examples

1. **Create an Author**:
   ```bash
   ./sqlc-tutorial -c -name "John Doe" -bio "A passionate writer"
   ```

2. **Read All Authors**:
   ```bash
   ./sqlc-tutorial -r
   ```

3. **Read a Specific Author**:
   ```bash
   ./sqlc-tutorial -r -id 1
   ```

4. **Update an Author's Name**:
   ```bash
   ./sqlc-tutorial -u -id 1 -name "Jane Doe"
   ```

5. **Update an Author's Bio**:
   ```bash
   ./sqlc-tutorial -u -id 1 -bio "An avid reader and writer"
   ```

6. **Delete an Author**:
   ```bash
   ./sqlc-tutorial -d -id 1
   ```

## Project Structure

```
.
├── main.go                                    # Main application entry point
├── go.mod                                     # Go module file
├── go.sum                                     # Go checksum file
├── .env                                       # Environment variables
├── sqlc.yaml                                  # SQLC configuration file
├── schema.sql                                 # SQLC schema file
├── query.sql                                  # SQL queries for SQLC
├── tutorial/                                  # Generated SQLC code
│   ├── db.go
│   ├── models.go
│   └── query.sql.go
├── migrations/                                # Database migrations
│   ├── 000001_create_authors_table.down.sql
│   └── 000001_create_authors_table.up.sql
└── README.md                                  # This file
```

## Dependencies

- [pgx](https://github.com/jackc/pgx): PostgreSQL driver and toolkit for Go.
- [sqlc](https://sqlc.dev/): SQL compiler that generates type-safe Go code from SQL.
- [godotenv](https://github.com/joho/godotenv): Load environment variables from a `.env` file.
- [tablewriter](https://github.com/olekukonko/tablewriter): Generate formatted tables for CLI output.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
