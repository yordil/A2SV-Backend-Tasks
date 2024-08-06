# Task Manager API 📋

Welcome to the **Task Manager API**! This API allows you to manage user registrations, logins, and task operations with secure authentication and authorization. Here's everything you need to get started!

## Introduction 🚀

The Task Manager API provides endpoints for:

- User Registration
- User Login
- CRUD operations for tasks
- User management
- JWT-based authentication
- Role-based authorization (Admin and User)

## Functionality ✨

- **User Registration:** Register new users with email, password, and role.
- **User Login:** Authenticate users and issue JWT tokens.
- **Task Management:** Create, read, update, and delete tasks.
- **User Management:** Admins can manage all users.
- **JWT Authentication:** Secure endpoints with JSON Web Tokens.
- **Role-based Authorization:** Differentiate access for Admin and User roles.

## MongoDB Setup 🛠️

1. **Clone the repository:**

```sh
git clone <repository-url>
cd <repository-directory>
```

2. **Set up environment variables:**

Create a `.env` file in the root directory and add the following variables:

```sh
JWT_SECRET=
MONGO_URI=
DB_NAME=
COLLECTION_NAME=
USER_ROLE=
ADMIN_CODE=
```

3. **Run the application:**

```sh
go run main.go
```

## Example Test Cases 🧪

### User Registration

- **Request:**

```json
POST /register
{
  "email": "testuser@example.com",
  "password": "password123",
  "role": "user"
}
```

- **Expected Response:**

```json
{
  "message": "User created successfully"
}
```

### User Login

- **Request:**

```json
POST /login
{
  "email": "testuser@example.com",
  "password": "password123"
}
```

- **Expected Response:**

```json
{
  "message": "User logged in successfully",
  "token": "<JWT_TOKEN>"
}
```

### Get All Tasks (Admin Only)

- **Request:**

```json
GET /api/tasks/
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
```

- **Expected Response:**

```json
[
  {
    "id": "<TASK_ID>",
    "user_id": "<USER_ID>",
    "title": "Task Title",
    "description": "Task Description",
    "due_date": "2024-01-01T00:00:00Z",
    "status": "Pending"
  }
]
```

### Create Task

- **Request:**

```json
POST /api/tasks/
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
Body:
{
  "title": "New Task",
  "description": "New Task Description",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "Pending"
}
```

- **Expected Response:**

```json
{
  "message": "Task created successfully",
  "task": {
    "id": "<TASK_ID>",
    "user_id": "<USER_ID>",
    "title": "New Task",
    "description": "New Task Description",
    "due_date": "2024-01-01T00:00:00Z",
    "status": "Pending"
  }
}
```

### Update Task

- **Request:**

```json
PUT /api/tasks/<TASK_ID>
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
Body:
{
  "title": "Updated Task",
  "description": "Updated Task Description",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "Completed"
}
```

- **Expected Response:**

```json
{
  "task": {
    "id": "<TASK_ID>",
    "user_id": "<USER_ID>",
    "title": "Updated Task",
    "description": "Updated Task Description",
    "due_date": "2024-01-01T00:00:00Z",
    "status": "Completed"
  }
}
```

### Delete Task

- **Request:**

```json
DELETE /api/tasks/<TASK_ID>
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
```

- **Expected Response:**

```json
{
  "message": "Task Deleted"
}
```

### Get Tasks by User ID

- **Request:**

```json
GET /api/tasks/mytask
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
```

- **Expected Response:**

```json
{
  "tasks": [
    {
      "id": "<TASK_ID>",
      "user_id": "<USER_ID>",
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2024-01-01T00:00:00Z",
      "status": "Pending"
    }
  ]
}
```

### Get All Users (Admin Only)

- **Request:**

```json
GET /api/users/getAllUser
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
```

- **Expected Response:**

```json
{
  "users": [
    {
      "id": "<USER_ID>",
      "email": "user@example.com",
      "role": "user"
    }
  ]
}
```

### Delete User

- **Request:**

```json
DELETE /api/users/deleteUser/<USER_ID>
Headers:
{
  "Authorization": "Bearer <JWT_TOKEN>"
}
```

- **Expected Response:**

```json
{
  "message": "User deleted successfully"
}
```

## Complete API Documentation 📚

You can get the complete API documentation [here](https://documenter.getpostman.com/view/28611859/2sA3rzHraf).

---
