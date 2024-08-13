package domain



type TaskResponse struct {
	All_Task []Task
}

type SingleTaskResponse struct {
	Single_Task Task
}

type TaskSuccessResponse struct {
	Message string
	Status int
	Data Task
}