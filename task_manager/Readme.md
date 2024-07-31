# Task Manager API 📋

## Overview

Welcome to the Task Manager API! This API allows you to manage tasks efficiently, including creating, retrieving, updating, and deleting tasks.

## Getting Started 🚀

To start using the Task Manager API, follow these steps:

1. **Set Up Your Environment**: Install the required dependencies.
2. **Configure Your Application**: Ensure your app is set up to use the API.
3. **Understand the Request and Response Format**: The API returns responses in JSON format.

No authentication is required for this API.

## API Endpoints 📡

### Get All Tasks

#### Request

- **Method**: GET
- **Endpoint**: `/api/tasks`

#### Response

- **Status Code**: 200 OK
- **Body**:

  ```json
  {
  	"tasks": [
  		{
  			"id": "1",
  			"title": "Task 1",
  			"description": "First task",
  			"due_date": "2024-07-31T00:00:00Z",
  			"status": "Pending"
  		},
  		{
  			"id": "2",
  			"title": "Task 2",
  			"description": "Second task",
  			"due_date": "2024-08-01T00:00:00Z",
  			"status": "In Progress"
  		}
  	]
  }
  ```

- **Status Code**: 404 Not Found
- **Body**:
  ```json
  {
  	"message": "No tasks found"
  }
  ```

### Get Task by ID

#### Request

- **Method**: GET
- **Endpoint**: `/api/tasks/:id`

#### Response

- **Status Code**: 200 OK
- **Body**:

  ```json
  {
  	"id": "1",
  	"title": "Task 1",
  	"description": "First task",
  	"due_date": "2024-07-31T00:00:00Z",
  	"status": "Pending"
  }
  ```

- **Status Code**: 404 Not Found
- **Body**:
  ```json
  {
  	"message": "Task Not Found"
  }
  ```

### Create a New Task 📝

#### Request

- **Method**: POST
- **Endpoint**: `/api/tasks`
- **Body**:
  ```json
  {
  	"title": "New Task",
  	"description": "Description of the new task",
  	"due_date": "2024-08-05T00:00:00Z",
  	"status": "Pending"
  }
  ```

#### Response

- **Status Code**: 201 Created
- **Body**:

  ```json
  {
  	"id": "4",
  	"title": "New Task",
  	"description": "Description of the new task",
  	"due_date": "2024-08-05T00:00:00Z",
  	"status": "Pending"
  }
  ```

- **Status Code**: 400 Bad Request
- **Body**:
  ```json
  {
  	"message": "Error message explaining why the request was bad"
  }
  ```

### Delete a Task 🗑️

#### Request

- **Method**: DELETE
- **Endpoint**: `/api/tasks/:id`

#### Response

- **Status Code**: 200 OK
- **Body**:

  ```json
  {
  	"message": "Task deleted successfully",
  	"task": {
  		"id": "3",
  		"title": "Task 3",
  		"description": "Third task",
  		"due_date": "2024-08-02T00:00:00Z",
  		"status": "Completed"
  	}
  }
  ```

- **Status Code**: 404 Not Found
- **Body**:
  ```json
  {
  	"message": "The task Not Found"
  }
  ```

### Update a Task ✏️

#### Request

- **Method**: PUT
- **Endpoint**: `/api/tasks/:id`
- **Body**:
  ```json
  {
  	"title": "Updated Task Title",
  	"description": "Updated task description",
  	"due_date": "2024-08-06T00:00:00Z",
  	"status": "Completed"
  }
  ```

#### Response

- **Status Code**: 200 OK
- **Body**:

  ```json
  {
  	"id": "2",
  	"title": "Updated Task Title",
  	"description": "Updated task description",
  	"due_date": "2024-08-06T00:00:00Z",
  	"status": "Completed"
  }
  ```

- **Status Code**: 400 Bad Request
- **Body**:
  ```json
  {
  	"message": "Error message explaining why the request was bad"
  }
  ```

## Error Responses 🚫

- **404 Not Found**: The requested resource could not be found.
  ```json
  {
  	"message": "The route Not Found"
  }
  ```

## Examples 📚

### Get All Tasks

```bash
curl -X GET http://localhost:8080/api/tasks
```

Response:

```json
{
	"tasks": [
		{
			"id": "1",
			"title": "Task 1",
			"description": "First task",
			"due_date": "2024-07-31T00:00:00Z",
			"status": "Pending"
		},
		{
			"id": "2",
			"title": "Task 2",
			"description": "Second task",
			"due_date": "2024-08-01T00:00:00Z",
			"status": "In Progress"
		}
	]
}
```

### Get Task by ID

```bash
curl -X GET http://localhost:8080/api/tasks/1
```

Response:

```json
{
	"id": "1",
	"title": "Task 1",
	"description": "First task",
	"due_date": "2024-07-31T00:00:00Z",
	"status": "Pending"
}
```

### Create a New Task 📝

```bash
curl -X POST http://localhost:8080/api/tasks -H "Content-Type: application/json" -d '{
  "title": "New Task",
  "description": "Description of the new task",
  "due_date": "2024-08-05T00:00:00Z",
  "status": "Pending"
}'
```

Response:

```json
{
	"id": "4",
	"title": "New Task",
	"description": "Description of the new task",
	"due_date": "2024-08-05T00:00:00Z",
	"status": "Pending"
}
```

### Delete a Task 🗑️

```bash
curl -X DELETE http://localhost:8080/api/tasks/3
```

Response:

```json
{
	"message": "Task deleted successfully",
	"task": {
		"id": "3",
		"title": "Task 3",
		"description": "Third task",
		"due_date": "2024-08-02T00:00:00Z",
		"status": "Completed"
	}
}
```

### Update a Task ✏️

```bash
curl -X PUT http://localhost:8080/api/tasks/2 -H "Content-Type: application/json" -d '{
  "title": "Updated Task Title",
  "description": "Updated task description",
  "due_date": "2024-08-06T00:00:00Z",
  "status": "Completed"
}'
```

Response:

```json
{
	"id": "2",
	"title": "Updated Task Title",
	"description": "Updated task description",
	"due_date": "2024-08-06T00:00:00Z",
	"status": "Completed"
}
```

---
