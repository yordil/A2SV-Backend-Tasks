---

# 🚀 Task Management API

Welcome to the **Task Management API**! This API allows you to manage users and tasks  with full authentication and authorization support! 🎯

---

## 🔧 **Setup Instructions**

### 1️⃣ **Clone the Repository**
```bash
git clone <repository-url>
cd <repository-directory>
```

### 2️⃣ **Environment Variables**
Set up your `.env` file with the following:
```bash
JWT_SECRET=your_secret_key 🔑
MONGO_URI=your_mongodb_url 🌐
DB_NAME=your_database_name 📂
COLLECTION_NAME=your_task_collection_name 🗂️
USER_COLLECTION_NAME=your_user_collection_name 👥
ADMIN_CODE=your_admin_code 🛡️
```

### 3️⃣ **Run the Application**
```bash
go run main.go 🏃‍♂️
```

---

## 📜 **API Endpoints**

### 🔒 **Authentication**

#### Register User
- **URL**: `/register`
- **Method**: `POST`
- **Request**:
    ```json
    {
      "email": "user@example.com",
      "password": "password123",
      "role": "user"
    }
    ```
- **Response**:
    ```json
    {
      "message": "User created successfully"
    }
    ```

#### Login User
- **URL**: `/login`
- **Method**: `POST`
- **Request**:
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
- **Response**:
    ```json
    {
      "message": "User logged in successfully",
      "token": "<JWT_TOKEN>"
    }
    ```

### 📋 **Task Management**

#### Get All Tasks (Admin only)
- **URL**: `/api/tasks/`
- **Method**: `GET`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Response**:
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

#### Get Task by ID
- **URL**: `/api/tasks/:id`
- **Method**: `GET`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Response**:
    ```json
    {
      "task": {
        "id": "<TASK_ID>",
        "user_id": "<USER_ID>",
        "title": "Task Title",
        "description": "Task Description",
        "due_date": "2024-01-01T00:00:00Z",
        "status": "Pending"
      }
    }
    ```

#### Create Task
- **URL**: `/api/tasks/`
- **Method**: `POST`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Request**:
    ```json
    {
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2024-01-01T00:00:00Z",
      "status": "Pending"
    }
    ```
- **Response**:
    ```json
    {
      "message": "Task created successfully",
      "task": {
        "id": "<TASK_ID>",
        "user_id": "<USER_ID>",
        "title": "Task Title",
        "description": "Task Description",
        "due_date": "2024-01-01T00:00:00Z",
        "status": "Pending"
      }
    }
    ```

#### Update Task
- **URL**: `/api/tasks/:id`
- **Method**: `PUT`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Request**:
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "due_date": "2024-01-01T00:00:00Z",
      "status": "Completed"
    }
    ```
- **Response**:
    ```json
    {
      "task": {
        "id": "<TASK_ID>",
        "user_id": "<USER_ID>",
        "title": "Updated Task Title",
        "description": "Updated Task Description",
        "due_date": "2024-01-01T00:00:00Z",
        "status": "Completed"
      }
    }
    ```

#### Delete Task 🗑️
- **URL**: `/api/tasks/:id`
- **Method**: `DELETE`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Response**:
    ```json
    {
      "message": "Task Deleted"
    }
    ```

#### Get Tasks by User ID
- **URL**: `/api/tasks/mytask`
- **Method**: `GET`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Response**:
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

### 👥 **User Management**

#### Get All Users (Admin only)
- **URL**: `/api/users/getAllUser`
- **Method**: `GET`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Response**:
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

#### Delete User 🗑️
- **URL**: `/api/users/deleteUser/:id`
- **Method**: `DELETE`
- **Headers**:
    ```json
    {
      "Authorization": "Bearer <JWT_TOKEN>"
    }
    ```
- **Response**:
    ```json
    {
      "message": "User deleted successfully"
    }
    ```

---

## 📚 **Full API Documentation**
For the complete API documentation, check out [this link](https://documenter.getpostman.com/view/28611859/2sA3s1pCqr) 📖.

---
