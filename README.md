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
- **Chi router** provides a lightweight and idiomatic way to define HTTP routes.
- **Modular package structure**: Each responsibility (task handling, storage, routing) is isolated.
- **In-memory storage** for simplicity and testability (easily replaceable with a real database).
- **Status-based filtering** and **pagination** included to simulate real-world requirements.

---

## 🚀 How to Run the Service

### 1. Clone the repository:
```bash
git clone https://github.com/invocoder/task-manager.git
cd task-manager
