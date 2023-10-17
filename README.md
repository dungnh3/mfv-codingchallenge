# mfv-codingchallenge

## Overview
Coding Challenge

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Configuration](#configuration)
- [Contributing](#contributing)

## Getting Started
### Installation

### Prerequisites
- Golang 19+
- MySQL (or provide details if using a different database)
- Docker
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Makefile GNU](https://www.gnu.org/software/make/manual/make.html)

## Usage
### Initialize database
Use the following command to init an instance database, and create tables
```bash
make init-db
make migrate-up
```
- MySQL as our database, the schema will be stored in [migrations](migrations)
- We use [migrate](#https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md) read migrations from sources and apply them in correct order to a database.

### Starting the Server
To start the SampleService Server, use the following command:
```bash
go run main.go server
```
The service will be available at http://localhost:9090

### API Endpoints
- `GET /users/:id`: Retrieve user information by id
- `GET /users/:id/accounts`: Retrieve a list account of user
- `GET /accounts/:id`: Retrieve account information by id

- `POST /users/register`: Add a new user to the database
```json
{"name": "user name"}
```
- `POST /users/:id/accounts`: Add a new account to the database
```json
{"name": "user name"}
```

## Testing
To run mock-tests, use the following command:
```bash
make mock-test 
```

To run integration-tests, use the following command:
```bash
make test-integration
```

## Configuration

## Contributing
We welcome contributions. To contribute to this project, please follow these steps:
```bash
- Fork the repository
- Create a new branch for your feature: git checkout -b feature/your-feature-name
- Commit your changes: git commit -m 'Added a new feature'
- Push to the branch: git push origin feature/your-feature-name
- Create a pull request
```
Please follow our code of conduct and coding style guidelines.