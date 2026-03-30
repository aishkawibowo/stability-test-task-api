# Task Manager API – Stability Test

A simple REST API built using **Golang** and the **Fiber framework**.

This project was provided as part of a **technical assessment** to evaluate the ability to analyze an existing codebase, identify bugs, and improve the stability and correctness of the API.

The main objective of this task is to **detect issues in the original implementation and implement fixes and improvements**.

---

# Tech Stack

- Go (Golang)
- Fiber Web Framework
- REST API
- JSON

---

# Project Structure

```
stability-test-task-api
│
├── handlers
│   └── task_handler.go
│
├── models
│   └── task.go
│
├── store
│   └── task_store.go
│
├── main.go
├── go.mod
└── README.md
```

---

# How to Run the Project

### 1. Clone the repository

```bash
git clone https://github.com/aishkawibowo/stability-test-task-api.git
cd stability-test-task-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the application

```bash
go run main.go
```

The API will run at:

```
http://localhost:3000
```

---

# Available Endpoints

| Method | Endpoint | Description |
|------|------|------|
| GET | `/` | API health check |
| GET | `/tasks` | Retrieve all tasks |
| GET | `/tasks/:id` | Retrieve task by ID |
| POST | `/tasks` | Create new task |
| DELETE | `/tasks/:id` | Delete task |

---

# Example Request

### Create Task

```
POST /tasks
```

```json
{
  "title": "Learn API",
  "done": false
}
```

---

# Example Response

### GET /tasks

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "done": false
  },
  {
    "id": 2,
    "title": "Build API",
    "done": false
  }
]
```

---

# Issues Identified in the Original Code

During the code review, several problems were found in the initial implementation.

### 1️. Pointer Bug in Slice Iteration

The original code returned the address of a loop variable which could cause incorrect memory references.

**Before**

```go
for _, t := range Tasks {
    if t.ID == id {
        return &t
    }
}
```

**After**

```go
for i := range Tasks {
    if Tasks[i].ID == id {
        return &Tasks[i]
    }
}
```

---

### 2️. Missing Error Handling for Invalid ID

The original implementation ignored potential errors when converting route parameters.

**Improvement**

Return `400 Bad Request` when the ID is invalid.

---

### 3️. Incorrect HTTP Status Codes

Some endpoints returned incorrect status codes.

**Fix**

- `200 OK` for successful GET requests
- `201 Created` for successful POST requests
- `404 Not Found` when a task does not exist
- `400 Bad Request` for invalid input

---

### 4️. Missing Validation on Task Creation

The API previously allowed tasks to be created without validating the request body.

**Improvement**

Added validation to ensure `title` is provided.

---

### 5️. Delete Operation Did Not Indicate Success

Previously the delete function did not indicate whether the task existed.

**Improvement**

Delete function now returns a boolean value indicating success.

---

### 6️. Task ID Handling

Tasks were previously added without assigning an ID.

**Improvement**

Implemented simple **auto-increment ID generation**.

---

# Improvements Implemented

Besides fixing bugs, several improvements were made:

- Added root endpoint `/` for API health check
- Implemented request validation
- Improved HTTP status codes
- Improved delete logic
- Implemented auto-increment ID generation

---

# Author

**Aishka Syakirah Wibowo**

Information Systems Graduate  
UIN Sunan Ampel Surabaya
