// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "task7/domain"

	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *TaskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Task) (domain.Task, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(domain.Task) domain.Task); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(domain.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: id
func (_m *TaskRepository) DeleteTask(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTask provides a mock function with given fields:
func (_m *TaskRepository) GetAllTask() ([]domain.Task, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllTask")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: id, role
func (_m *TaskRepository) GetTaskByID(id string, role string) (domain.Task, error) {
	ret := _m.Called(id, role)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (domain.Task, error)); ok {
		return rf(id, role)
	}
	if rf, ok := ret.Get(0).(func(string, string) domain.Task); ok {
		r0 = rf(id, role)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByUserID provides a mock function with given fields: userID
func (_m *TaskRepository) GetTaskByUserID(userID string) ([]domain.Task, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByUserID")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Task, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Task); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: id, task
func (_m *TaskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	ret := _m.Called(id, task)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.Task) (domain.Task, error)); ok {
		return rf(id, task)
	}
	if rf, ok := ret.Get(0).(func(string, domain.Task) domain.Task); ok {
		r0 = rf(id, task)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(string, domain.Task) error); ok {
		r1 = rf(id, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserIdGetter provides a mock function with given fields: id, taskid
func (_m *TaskRepository) UserIdGetter(id string, taskid string) bool {
	ret := _m.Called(id, taskid)

	if len(ret) == 0 {
		panic("no return value specified for UserIdGetter")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(id, taskid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
