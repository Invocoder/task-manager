# Task Manager Service

## 📌 Problem Breakdown

This service is a simple **Task Manager API** designed to manage tasks with basic CRUD functionality. The goal is to demonstrate a modular, scalable service that adheres to microservices architecture principles.

### ✅ Features:
- Create a new task
- Retrieve tasks by status (with pagination)
- Update an existing task
- Delete a task

### 💡 Design Decisions:
- **Go (Golang)** is used for high performance and simplicity in building REST APIs.
- **Status-based filtering** and **pagination** included to simulate real-world requirements.

---

## 🚀 How to Run the Service

### 1. Clone the repository:
```bash
git clone https://github.com/invocoder/task-manager.git
cd task-manager

## 📘 API Documentation

The Task Management API allows you to create, retrieve, update, and delete tasks.

### 🔗 Swagger Documentation
Access full interactive API documentation via Swagger UI:

- [http://localhost:8082/swagger/index.html](http://localhost:8082/swagger/index.html)

---

### 📚 Endpoints

#### ✅ Create a Task

- **URL**: `/api/tasks`
- **Method**: `POST`
- **Request Body**:
```json
{
  "title": "Buy groceries",
  "status": "pending"
}

{
  "id": 1,
  "title": "Buy groceries",
  "status": "pending"
}

curl -X POST http://localhost:8082/api/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Buy groceries", "status":"pending"}'

[
  {
    "id": 1,
    "title": "Buy groceries",
    "status": "pending"
  }
]

{
  "title": "Buy vegetables",
  "status": "completed"
}

{
  "message": "Task deleted successfully"
}
