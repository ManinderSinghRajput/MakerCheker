# Maker-Checker Service

## Overview
The Maker-Checker Service is a simplified implementation of an approval workflow for sending messages. A message must be validated and approved by a user before being sent to the recipient. If the message is rejected, it will not be sent.

This service is implemented in Go and exposes RESTful APIs to manage the message approval process.

## Features
- Submit a new message for approval.
- Approve or reject a submitted message.
- Retrieve the status of a message.

## Directory Structure
```
maker_checker_service/
├── cmd/
│   └── main.go
├── internal/
│   └── handler/
│       └── handler.go
├── pkg/
│   └── model/
│       └── message.go
└── .gitignore
└── go.mod
└── go.sum
└── Makefile
└── README.md
```

## Prerequisites
- Go 1.20 or later
- cURL for testing the API

## Running the Service

1. Clone the repository and navigate to the project directory:
   ```bash
   git clone <repository-url>
   cd maker_checker_service
   ```

2. Build and run the application using the provided Makefile:
   ```bash
   make build
   ./bin/maker_checker_service
   ```
   Or run directly:
   ```bash
   make run
   ```

The service will start at `http://localhost:8080`.

## API Endpoints

### 1. **Submit a Message**
**Endpoint:** `POST /messages`

**Request Body:**
```json
{
  "message": "Hello, this is a test message!",
  "recipient": "test@example.com"
}
```

**cURL Command:**
```bash
curl -X POST http://localhost:8080/messages \
-H "Content-Type: application/json" \
-d '{
  "message": "Hello, this is a test message!",
  "recipient": "test@example.com"
}'
```

**Response:**
```json
{
  "id": "<message_id>",
  "message": "Hello, this is a test message!",
  "recipient": "test@example.com",
  "status": "pending"
}
```

### 2. **Approve or Reject a Message**
**Endpoint:** `POST /messages/{id}/status`

**Request Body:**
```json
{
  "status": "approved"
}
```
Or:
```json
{
  "status": "rejected"
}
```

**cURL Commands:**
#### Approve:
```bash
curl -X POST http://localhost:8080/messages/<message_id>/status \
-H "Content-Type: application/json" \
-d '{
  "status": "approved"
}'
```
#### Reject:
```bash
curl -X POST http://localhost:8080/messages/<message_id>/status \
-H "Content-Type: application/json" \
-d '{
  "status": "rejected"
}'
```

**Response:**
```json
{
  "id": "<message_id>",
  "message": "Hello, this is a test message!",
  "recipient": "test@example.com",
  "status": "approved"
}
```

### 3. **Get Message Status**
**Endpoint:** `GET /messages/{id}`

**cURL Command:**
```bash
curl -X GET http://localhost:8080/messages/<message_id>
```

**Response:**
```json
{
  "id": "<message_id>",
  "message": "Hello, this is a test message!",
  "recipient": "test@example.com",
  "status": "pending"
}
```

## Makefile Targets
- `build`: Build the application binary.
- `fmt`: Format the code.
- `run`: Run the application.
- `tidy`: Clean up module dependencies.
- `vendor`: Create a vendor directory for dependencies.
- `clean`: Remove build artifacts.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments
- [Gorilla Mux](https://github.com/gorilla/mux) for HTTP request routing.
- [Google UUID](https://github.com/google/uuid) for UUID generation.
