# Jump Interview Application

This is a Go application developed as part of the Jump interview process.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Database Setup](#database-setup)
  - [Running the Application](#running-the-application)
- [Running Tests](#running-tests)
- [Possible Enhancements](#possible-enhancements)

## Overview

The Jump Interview Application is a Go-based web application that interacts with a PostgreSQL database to manage users, invoices, and transactions. It provides APIs to retrieve user information, create invoices, and accept transactions.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go (1.16 or higher)
- Docker (for running PostgreSQL in a container)
- PostgreSQL (if not using Docker)

## Getting Started

### Database Setup

1. **Using Docker (Recommended)**

   Build the Docker image for the PostgreSQL database and run it with the following commands:

   ```bash
   git clone https://github.com/Freelance-launchpad/backend-interviews.git
   cd backend-interviews/database
   docker build . -t jump-database
   docker run -p 5432:5432 jump-database
   ```

   The database will be accessible on port 5432.

2. **Without Docker**

   If you have a PostgreSQL instance running locally, update the database connection configuration in the `db.ConnectDB` function in the `db` package.

### Running the Application

1. Clone this repository:

   ```bash
   git clone https://https://github.com/darkcryptodev/jump_interview.git
   cd jump_interview
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Start the Go application:

   ```bash
   go run api/main.go
   ```

   The application will be accessible at `http://localhost:8080`.

## Running Tests

Run the tests using the following command:

```bash
go test ./...
```

## Possible Enhancements

- Add more test cases to ensure robust test coverage.
- Add unicity fields for invoice, to ensure unicity.
- Improve error handling and logging for better debugging.
- Implement pagination for listing users, invoices, and transactions.
- Add indexes for big tables.
- Use SERIAL instead of creating custom SEQUENCEs.
- Provide API documentation using tools like Swagger.
- Use a database migration tool for managing database schema changes.
- Implement a frontend interface to interact with the APIs.
- Containerize the Go application for easier deployment.
