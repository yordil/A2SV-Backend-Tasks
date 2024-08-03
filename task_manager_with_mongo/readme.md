# Task Manager API 📝

Welcome to the Task Manager API! This API allows you to manage tasks efficiently by providing endpoints to create, retrieve, update, and delete tasks. It is integrated with MongoDB for persistent data storage.

## Features ✨

- **Create Tasks**: Add new tasks with a title, description, due date, and status.
- **Retrieve Tasks**: Fetch all tasks or a specific task by its ID.
- **Update Tasks**: Modify the details of an existing task.
- **Delete Tasks**: Remove tasks that are no longer needed.

## Prerequisites 📋

### Environment Variables

To connect to MongoDB, set up the following environment variables:

- `MONGODB_URI`: The URI of your MongoDB instance.
- `DATABASE_NAME`: The name of your database.
- `TASKS_COLLECTION`: The name of the tasks collection.

## Example Requests and Responses 📬

### Get All Tasks 📋

#### Request

```bash
curl -X GET http://localhost:8080/api/tasks
```

#### Response

```json
{
	"tasks": [
		{
			"id": "64b7f7e00b8e8f3f3d9b0e1b",
			"title": "Task 1",
			"description": "First task",
			"due_date": "2024-07-31T00:00:00Z",
			"status": "Pending"
		},
		{
			"id": "64b7f7e10b8e8f3f3d9b0e1c",
			"title": "Task 2",
			"description": "Second task",
			"due_date": "2024-08-01T00:00:00Z",
			"status": "In Progress"
		}
	]
}
```

### Get Task by ID 🔍

#### Request

```bash
curl -X GET http://localhost:8080/api/tasks/64b7f7e00b8e8f3f3d9b0e1b
```

#### Response

```json
{
	"task": {
		"id": "64b7f7e00b8e8f3f3d9b0e1b",
		"title": "Task 1",
		"description": "First task",
		"due_date": "2024-07-31T00:00:00Z",
		"status": "Pending"
	}
}
```

### Create a New Task ➕

#### Request

```bash
curl -X POST http://localhost:8080/api/tasks -H "Content-Type: application/json" -d '{
  "title": "New Task",
  "description": "Description of the new task",
  "due_date": "2024-08-05T00:00:00Z",
  "status": "Pending"
}'
```

#### Response

```json
{
	"message": "Task created successfully",
	"task": {
		"id": "64b7f7e20b8e8f3f3d9b0e1d",
		"title": "New Task",
		"description": "Description of the new task",
		"due_date": "2024-08-05T00:00:00Z",
		"status": "Pending"
	}
}
```

### Update a Task ♻️

#### Request

```bash
curl -X PUT http://localhost:8080/api/tasks/64b7f7e10b8e8f3f3d9b0e1c -H "Content-Type: application/json" -d '{
  "title": "Updated Task Title",
  "description": "Updated task description",
  "due_date": "2024-08-06T00:00:00Z",
  "status": "Completed"
}'
```

#### Response

```json
{
	"task": {
		"id": "64b7f7e10b8e8f3f3d9b0e1c",
		"title": "Updated Task Title",
		"description": "Updated task description",
		"due_date": "2024-08-06T00:00:00Z",
		"status": "Completed"
	}
}
```

### Delete a Task ❌

#### Request

```bash
curl -X DELETE http://localhost:8080/api/tasks/64b7f7e00b8e8f3f3d9b0e1b
```

#### Response

```json
{
	"message": "Task deleted successfully"
}
```

## Complete API Documentation 📚

For more detailed information and additional examples, you can find the complete API documentation [here](https://documenter.getpostman.com/view/28611859/2sA3rwMZW8).
