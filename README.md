# Bank Lite API

## Overview

Bank Lite API is a simple banking application that allows users to manage accounts and transactions. This project is built using Go and provides endpoints for creating accounts, making transactions, and transferring funds between accounts.

## Prerequisites

- Go 1.19 or later
- Docker (for running the application using Docker)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/achepkin/banklite.git
    cd banklite
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Running the Application

### Using Makefile

To run the application using the Makefile, use the following commands:

- Start the application:
    ```sh
    make up
    ```

## API Documentation

The API documentation is available in the `docs/banklite.postman_collection.json` file. You can import this file into Postman to explore the available endpoints.

## Testing

To run the tests, use the following command:
```sh
go test ./...