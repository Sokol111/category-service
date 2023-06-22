// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "github.com/Sokol111/category-service/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockCategoryRepository is an autogenerated mock type for the CategoryRepository type
type MockCategoryRepository struct {
	mock.Mock
}

type MockCategoryRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCategoryRepository) EXPECT() *MockCategoryRepository_Expecter {
	return &MockCategoryRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: category
func (_m *MockCategoryRepository) Create(category model.Category) (model.Category, error) {
	ret := _m.Called(category)

	var r0 model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Category) (model.Category, error)); ok {
		return rf(category)
	}
	if rf, ok := ret.Get(0).(func(model.Category) model.Category); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	if rf, ok := ret.Get(1).(func(model.Category) error); ok {
		r1 = rf(category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCategoryRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockCategoryRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - category model.Category
func (_e *MockCategoryRepository_Expecter) Create(category interface{}) *MockCategoryRepository_Create_Call {
	return &MockCategoryRepository_Create_Call{Call: _e.mock.On("Create", category)}
}

func (_c *MockCategoryRepository_Create_Call) Run(run func(category model.Category)) *MockCategoryRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Category))
	})
	return _c
}

func (_c *MockCategoryRepository_Create_Call) Return(_a0 model.Category, _a1 error) *MockCategoryRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCategoryRepository_Create_Call) RunAndReturn(run func(model.Category) (model.Category, error)) *MockCategoryRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *MockCategoryRepository) GetById(id string) (model.Category, error) {
	ret := _m.Called(id)

	var r0 model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Category, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) model.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCategoryRepository_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockCategoryRepository_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id string
func (_e *MockCategoryRepository_Expecter) GetById(id interface{}) *MockCategoryRepository_GetById_Call {
	return &MockCategoryRepository_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *MockCategoryRepository_GetById_Call) Run(run func(id string)) *MockCategoryRepository_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCategoryRepository_GetById_Call) Return(_a0 model.Category, _a1 error) *MockCategoryRepository_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCategoryRepository_GetById_Call) RunAndReturn(run func(string) (model.Category, error)) *MockCategoryRepository_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: name
func (_m *MockCategoryRepository) GetByName(name string) (model.Category, error) {
	ret := _m.Called(name)

	var r0 model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Category, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) model.Category); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCategoryRepository_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockCategoryRepository_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - name string
func (_e *MockCategoryRepository_Expecter) GetByName(name interface{}) *MockCategoryRepository_GetByName_Call {
	return &MockCategoryRepository_GetByName_Call{Call: _e.mock.On("GetByName", name)}
}

func (_c *MockCategoryRepository_GetByName_Call) Run(run func(name string)) *MockCategoryRepository_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCategoryRepository_GetByName_Call) Return(_a0 model.Category, _a1 error) *MockCategoryRepository_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCategoryRepository_GetByName_Call) RunAndReturn(run func(string) (model.Category, error)) *MockCategoryRepository_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetCategories provides a mock function with given fields:
func (_m *MockCategoryRepository) GetCategories() ([]model.Category, error) {
	ret := _m.Called()

	var r0 []model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Category, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Category)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCategoryRepository_GetCategories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCategories'
type MockCategoryRepository_GetCategories_Call struct {
	*mock.Call
}

// GetCategories is a helper method to define mock.On call
func (_e *MockCategoryRepository_Expecter) GetCategories() *MockCategoryRepository_GetCategories_Call {
	return &MockCategoryRepository_GetCategories_Call{Call: _e.mock.On("GetCategories")}
}

func (_c *MockCategoryRepository_GetCategories_Call) Run(run func()) *MockCategoryRepository_GetCategories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCategoryRepository_GetCategories_Call) Return(_a0 []model.Category, _a1 error) *MockCategoryRepository_GetCategories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCategoryRepository_GetCategories_Call) RunAndReturn(run func() ([]model.Category, error)) *MockCategoryRepository_GetCategories_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: category
func (_m *MockCategoryRepository) Update(category model.Category) (model.Category, error) {
	ret := _m.Called(category)

	var r0 model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Category) (model.Category, error)); ok {
		return rf(category)
	}
	if rf, ok := ret.Get(0).(func(model.Category) model.Category); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	if rf, ok := ret.Get(1).(func(model.Category) error); ok {
		r1 = rf(category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCategoryRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockCategoryRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - category model.Category
func (_e *MockCategoryRepository_Expecter) Update(category interface{}) *MockCategoryRepository_Update_Call {
	return &MockCategoryRepository_Update_Call{Call: _e.mock.On("Update", category)}
}

func (_c *MockCategoryRepository_Update_Call) Run(run func(category model.Category)) *MockCategoryRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Category))
	})
	return _c
}

func (_c *MockCategoryRepository_Update_Call) Return(_a0 model.Category, _a1 error) *MockCategoryRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCategoryRepository_Update_Call) RunAndReturn(run func(model.Category) (model.Category, error)) *MockCategoryRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCategoryRepository creates a new instance of MockCategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCategoryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCategoryRepository {
	mock := &MockCategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
