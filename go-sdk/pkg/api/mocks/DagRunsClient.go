// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"

	"github.com/apache/airflow/go-sdk/pkg/api"
	mock "github.com/stretchr/testify/mock"
	"resty.dev/v3"
)

// NewDagRunsClient creates a new instance of DagRunsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDagRunsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DagRunsClient {
	mock := &DagRunsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// DagRunsClient is an autogenerated mock type for the DagRunsClient type
type DagRunsClient struct {
	mock.Mock
}

type DagRunsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *DagRunsClient) EXPECT() *DagRunsClient_Expecter {
	return &DagRunsClient_Expecter{mock: &_m.Mock}
}

// Clear provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) Clear(ctx context.Context, dagId string, runId string) error {
	ret := _mock.Called(ctx, dagId, runId)

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = returnFunc(ctx, dagId, runId)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// DagRunsClient_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type DagRunsClient_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
//   - ctx context.Context
//   - dagId string
//   - runId string
func (_e *DagRunsClient_Expecter) Clear(ctx interface{}, dagId interface{}, runId interface{}) *DagRunsClient_Clear_Call {
	return &DagRunsClient_Clear_Call{Call: _e.mock.On("Clear", ctx, dagId, runId)}
}

func (_c *DagRunsClient_Clear_Call) Run(run func(ctx context.Context, dagId string, runId string)) *DagRunsClient_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *DagRunsClient_Clear_Call) Return(err error) *DagRunsClient_Clear_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *DagRunsClient_Clear_Call) RunAndReturn(run func(ctx context.Context, dagId string, runId string) error) *DagRunsClient_Clear_Call {
	_c.Call.Return(run)
	return _c
}

// ClearResponse provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) ClearResponse(ctx context.Context, dagId string, runId string) (*resty.Response, error) {
	ret := _mock.Called(ctx, dagId, runId)

	if len(ret) == 0 {
		panic("no return value specified for ClearResponse")
	}

	var r0 *resty.Response
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) (*resty.Response, error)); ok {
		return returnFunc(ctx, dagId, runId)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) *resty.Response); ok {
		r0 = returnFunc(ctx, dagId, runId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resty.Response)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = returnFunc(ctx, dagId, runId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// DagRunsClient_ClearResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearResponse'
type DagRunsClient_ClearResponse_Call struct {
	*mock.Call
}

// ClearResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - dagId string
//   - runId string
func (_e *DagRunsClient_Expecter) ClearResponse(ctx interface{}, dagId interface{}, runId interface{}) *DagRunsClient_ClearResponse_Call {
	return &DagRunsClient_ClearResponse_Call{Call: _e.mock.On("ClearResponse", ctx, dagId, runId)}
}

func (_c *DagRunsClient_ClearResponse_Call) Run(run func(ctx context.Context, dagId string, runId string)) *DagRunsClient_ClearResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *DagRunsClient_ClearResponse_Call) Return(response *resty.Response, err error) *DagRunsClient_ClearResponse_Call {
	_c.Call.Return(response, err)
	return _c
}

func (_c *DagRunsClient_ClearResponse_Call) RunAndReturn(run func(ctx context.Context, dagId string, runId string) (*resty.Response, error)) *DagRunsClient_ClearResponse_Call {
	_c.Call.Return(run)
	return _c
}

// GetDrCount provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) GetDrCount(ctx context.Context, params *api.GetDrCountParams) (*int, error) {
	ret := _mock.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetDrCount")
	}

	var r0 *int
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *api.GetDrCountParams) (*int, error)); ok {
		return returnFunc(ctx, params)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *api.GetDrCountParams) *int); ok {
		r0 = returnFunc(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *api.GetDrCountParams) error); ok {
		r1 = returnFunc(ctx, params)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// DagRunsClient_GetDrCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDrCount'
type DagRunsClient_GetDrCount_Call struct {
	*mock.Call
}

// GetDrCount is a helper method to define mock.On call
//   - ctx context.Context
//   - params *api.GetDrCountParams
func (_e *DagRunsClient_Expecter) GetDrCount(ctx interface{}, params interface{}) *DagRunsClient_GetDrCount_Call {
	return &DagRunsClient_GetDrCount_Call{Call: _e.mock.On("GetDrCount", ctx, params)}
}

func (_c *DagRunsClient_GetDrCount_Call) Run(run func(ctx context.Context, params *api.GetDrCountParams)) *DagRunsClient_GetDrCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 *api.GetDrCountParams
		if args[1] != nil {
			arg1 = args[1].(*api.GetDrCountParams)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *DagRunsClient_GetDrCount_Call) Return(n *int, err error) *DagRunsClient_GetDrCount_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *DagRunsClient_GetDrCount_Call) RunAndReturn(run func(ctx context.Context, params *api.GetDrCountParams) (*int, error)) *DagRunsClient_GetDrCount_Call {
	_c.Call.Return(run)
	return _c
}

// GetDrCountResponse provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) GetDrCountResponse(ctx context.Context, params *api.GetDrCountParams) (*resty.Response, error) {
	ret := _mock.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetDrCountResponse")
	}

	var r0 *resty.Response
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *api.GetDrCountParams) (*resty.Response, error)); ok {
		return returnFunc(ctx, params)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *api.GetDrCountParams) *resty.Response); ok {
		r0 = returnFunc(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resty.Response)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *api.GetDrCountParams) error); ok {
		r1 = returnFunc(ctx, params)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// DagRunsClient_GetDrCountResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDrCountResponse'
type DagRunsClient_GetDrCountResponse_Call struct {
	*mock.Call
}

// GetDrCountResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - params *api.GetDrCountParams
func (_e *DagRunsClient_Expecter) GetDrCountResponse(ctx interface{}, params interface{}) *DagRunsClient_GetDrCountResponse_Call {
	return &DagRunsClient_GetDrCountResponse_Call{Call: _e.mock.On("GetDrCountResponse", ctx, params)}
}

func (_c *DagRunsClient_GetDrCountResponse_Call) Run(run func(ctx context.Context, params *api.GetDrCountParams)) *DagRunsClient_GetDrCountResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 *api.GetDrCountParams
		if args[1] != nil {
			arg1 = args[1].(*api.GetDrCountParams)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *DagRunsClient_GetDrCountResponse_Call) Return(response *resty.Response, err error) *DagRunsClient_GetDrCountResponse_Call {
	_c.Call.Return(response, err)
	return _c
}

func (_c *DagRunsClient_GetDrCountResponse_Call) RunAndReturn(run func(ctx context.Context, params *api.GetDrCountParams) (*resty.Response, error)) *DagRunsClient_GetDrCountResponse_Call {
	_c.Call.Return(run)
	return _c
}

// GetState provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) GetState(ctx context.Context, dagId string, runId string) (*api.DagRunStateResponse, error) {
	ret := _mock.Called(ctx, dagId, runId)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 *api.DagRunStateResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) (*api.DagRunStateResponse, error)); ok {
		return returnFunc(ctx, dagId, runId)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) *api.DagRunStateResponse); ok {
		r0 = returnFunc(ctx, dagId, runId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.DagRunStateResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = returnFunc(ctx, dagId, runId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// DagRunsClient_GetState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetState'
type DagRunsClient_GetState_Call struct {
	*mock.Call
}

// GetState is a helper method to define mock.On call
//   - ctx context.Context
//   - dagId string
//   - runId string
func (_e *DagRunsClient_Expecter) GetState(ctx interface{}, dagId interface{}, runId interface{}) *DagRunsClient_GetState_Call {
	return &DagRunsClient_GetState_Call{Call: _e.mock.On("GetState", ctx, dagId, runId)}
}

func (_c *DagRunsClient_GetState_Call) Run(run func(ctx context.Context, dagId string, runId string)) *DagRunsClient_GetState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *DagRunsClient_GetState_Call) Return(dagRunStateResponse *api.DagRunStateResponse, err error) *DagRunsClient_GetState_Call {
	_c.Call.Return(dagRunStateResponse, err)
	return _c
}

func (_c *DagRunsClient_GetState_Call) RunAndReturn(run func(ctx context.Context, dagId string, runId string) (*api.DagRunStateResponse, error)) *DagRunsClient_GetState_Call {
	_c.Call.Return(run)
	return _c
}

// GetStateResponse provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) GetStateResponse(ctx context.Context, dagId string, runId string) (*resty.Response, error) {
	ret := _mock.Called(ctx, dagId, runId)

	if len(ret) == 0 {
		panic("no return value specified for GetStateResponse")
	}

	var r0 *resty.Response
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) (*resty.Response, error)); ok {
		return returnFunc(ctx, dagId, runId)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) *resty.Response); ok {
		r0 = returnFunc(ctx, dagId, runId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resty.Response)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = returnFunc(ctx, dagId, runId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// DagRunsClient_GetStateResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateResponse'
type DagRunsClient_GetStateResponse_Call struct {
	*mock.Call
}

// GetStateResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - dagId string
//   - runId string
func (_e *DagRunsClient_Expecter) GetStateResponse(ctx interface{}, dagId interface{}, runId interface{}) *DagRunsClient_GetStateResponse_Call {
	return &DagRunsClient_GetStateResponse_Call{Call: _e.mock.On("GetStateResponse", ctx, dagId, runId)}
}

func (_c *DagRunsClient_GetStateResponse_Call) Run(run func(ctx context.Context, dagId string, runId string)) *DagRunsClient_GetStateResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *DagRunsClient_GetStateResponse_Call) Return(response *resty.Response, err error) *DagRunsClient_GetStateResponse_Call {
	_c.Call.Return(response, err)
	return _c
}

func (_c *DagRunsClient_GetStateResponse_Call) RunAndReturn(run func(ctx context.Context, dagId string, runId string) (*resty.Response, error)) *DagRunsClient_GetStateResponse_Call {
	_c.Call.Return(run)
	return _c
}

// Trigger provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) Trigger(ctx context.Context, dagId string, runId string, body *api.TriggerDAGRunPayload) error {
	ret := _mock.Called(ctx, dagId, runId, body)

	if len(ret) == 0 {
		panic("no return value specified for Trigger")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, *api.TriggerDAGRunPayload) error); ok {
		r0 = returnFunc(ctx, dagId, runId, body)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// DagRunsClient_Trigger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trigger'
type DagRunsClient_Trigger_Call struct {
	*mock.Call
}

// Trigger is a helper method to define mock.On call
//   - ctx context.Context
//   - dagId string
//   - runId string
//   - body *api.TriggerDAGRunPayload
func (_e *DagRunsClient_Expecter) Trigger(ctx interface{}, dagId interface{}, runId interface{}, body interface{}) *DagRunsClient_Trigger_Call {
	return &DagRunsClient_Trigger_Call{Call: _e.mock.On("Trigger", ctx, dagId, runId, body)}
}

func (_c *DagRunsClient_Trigger_Call) Run(run func(ctx context.Context, dagId string, runId string, body *api.TriggerDAGRunPayload)) *DagRunsClient_Trigger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		var arg3 *api.TriggerDAGRunPayload
		if args[3] != nil {
			arg3 = args[3].(*api.TriggerDAGRunPayload)
		}
		run(
			arg0,
			arg1,
			arg2,
			arg3,
		)
	})
	return _c
}

func (_c *DagRunsClient_Trigger_Call) Return(err error) *DagRunsClient_Trigger_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *DagRunsClient_Trigger_Call) RunAndReturn(run func(ctx context.Context, dagId string, runId string, body *api.TriggerDAGRunPayload) error) *DagRunsClient_Trigger_Call {
	_c.Call.Return(run)
	return _c
}

// TriggerResponse provides a mock function for the type DagRunsClient
func (_mock *DagRunsClient) TriggerResponse(ctx context.Context, dagId string, runId string, body *api.TriggerDAGRunPayload) (*resty.Response, error) {
	ret := _mock.Called(ctx, dagId, runId, body)

	if len(ret) == 0 {
		panic("no return value specified for TriggerResponse")
	}

	var r0 *resty.Response
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, *api.TriggerDAGRunPayload) (*resty.Response, error)); ok {
		return returnFunc(ctx, dagId, runId, body)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, *api.TriggerDAGRunPayload) *resty.Response); ok {
		r0 = returnFunc(ctx, dagId, runId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resty.Response)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string, *api.TriggerDAGRunPayload) error); ok {
		r1 = returnFunc(ctx, dagId, runId, body)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// DagRunsClient_TriggerResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TriggerResponse'
type DagRunsClient_TriggerResponse_Call struct {
	*mock.Call
}

// TriggerResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - dagId string
//   - runId string
//   - body *api.TriggerDAGRunPayload
func (_e *DagRunsClient_Expecter) TriggerResponse(ctx interface{}, dagId interface{}, runId interface{}, body interface{}) *DagRunsClient_TriggerResponse_Call {
	return &DagRunsClient_TriggerResponse_Call{Call: _e.mock.On("TriggerResponse", ctx, dagId, runId, body)}
}

func (_c *DagRunsClient_TriggerResponse_Call) Run(run func(ctx context.Context, dagId string, runId string, body *api.TriggerDAGRunPayload)) *DagRunsClient_TriggerResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		var arg3 *api.TriggerDAGRunPayload
		if args[3] != nil {
			arg3 = args[3].(*api.TriggerDAGRunPayload)
		}
		run(
			arg0,
			arg1,
			arg2,
			arg3,
		)
	})
	return _c
}

func (_c *DagRunsClient_TriggerResponse_Call) Return(response *resty.Response, err error) *DagRunsClient_TriggerResponse_Call {
	_c.Call.Return(response, err)
	return _c
}

func (_c *DagRunsClient_TriggerResponse_Call) RunAndReturn(run func(ctx context.Context, dagId string, runId string, body *api.TriggerDAGRunPayload) (*resty.Response, error)) *DagRunsClient_TriggerResponse_Call {
	_c.Call.Return(run)
	return _c
}
