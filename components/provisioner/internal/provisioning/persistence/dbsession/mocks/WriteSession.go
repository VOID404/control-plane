// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	apperrors "github.com/kyma-project/control-plane/components/provisioner/internal/apperrors"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-project/control-plane/components/provisioner/internal/model"

	time "time"
)

// WriteSession is an autogenerated mock type for the WriteSession type
type WriteSession struct {
	mock.Mock
}

// DeleteCluster provides a mock function with given fields: runtimeID
func (_m *WriteSession) DeleteCluster(runtimeID string) apperrors.AppError {
	ret := _m.Called(runtimeID)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(runtimeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// InsertAdministrators provides a mock function with given fields: clusterId, administrators
func (_m *WriteSession) InsertAdministrators(clusterId string, administrators []string) apperrors.AppError {
	ret := _m.Called(clusterId, administrators)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, []string) apperrors.AppError); ok {
		r0 = rf(clusterId, administrators)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// InsertCluster provides a mock function with given fields: cluster
func (_m *WriteSession) InsertCluster(cluster model.Cluster) apperrors.AppError {
	ret := _m.Called(cluster)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(model.Cluster) apperrors.AppError); ok {
		r0 = rf(cluster)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// InsertGardenerConfig provides a mock function with given fields: config
func (_m *WriteSession) InsertGardenerConfig(config model.GardenerConfig) apperrors.AppError {
	ret := _m.Called(config)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(model.GardenerConfig) apperrors.AppError); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// InsertOperation provides a mock function with given fields: operation
func (_m *WriteSession) InsertOperation(operation model.Operation) apperrors.AppError {
	ret := _m.Called(operation)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(model.Operation) apperrors.AppError); ok {
		r0 = rf(operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// MarkClusterAsDeleted provides a mock function with given fields: runtimeID
func (_m *WriteSession) MarkClusterAsDeleted(runtimeID string) apperrors.AppError {
	ret := _m.Called(runtimeID)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(runtimeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// TransitionOperation provides a mock function with given fields: operationID, message, stage, transitionTime
func (_m *WriteSession) TransitionOperation(operationID string, message string, stage model.OperationStage, transitionTime time.Time) apperrors.AppError {
	ret := _m.Called(operationID, message, stage, transitionTime)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, model.OperationStage, time.Time) apperrors.AppError); ok {
		r0 = rf(operationID, message, stage, transitionTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateGardenerClusterConfig provides a mock function with given fields: config
func (_m *WriteSession) UpdateGardenerClusterConfig(config model.GardenerConfig) apperrors.AppError {
	ret := _m.Called(config)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(model.GardenerConfig) apperrors.AppError); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateKubeconfig provides a mock function with given fields: runtimeID, kubeconfig
func (_m *WriteSession) UpdateKubeconfig(runtimeID string, kubeconfig string) apperrors.AppError {
	ret := _m.Called(runtimeID, kubeconfig)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(runtimeID, kubeconfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateKubernetesVersion provides a mock function with given fields: runtimeID, version
func (_m *WriteSession) UpdateKubernetesVersion(runtimeID string, version string) apperrors.AppError {
	ret := _m.Called(runtimeID, version)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(runtimeID, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateOperationLastError provides a mock function with given fields: operationID, msg, reason, component
func (_m *WriteSession) UpdateOperationLastError(operationID string, msg string, reason string, component string) apperrors.AppError {
	ret := _m.Called(operationID, msg, reason, component)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, string, string) apperrors.AppError); ok {
		r0 = rf(operationID, msg, reason, component)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateOperationState provides a mock function with given fields: operationID, message, state, endTime
func (_m *WriteSession) UpdateOperationState(operationID string, message string, state model.OperationState, endTime time.Time) apperrors.AppError {
	ret := _m.Called(operationID, message, state, endTime)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, model.OperationState, time.Time) apperrors.AppError); ok {
		r0 = rf(operationID, message, state, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateShootNetworkingFilterDisabled provides a mock function with given fields: runtimeID, shootNetworkingFilterDisabled
func (_m *WriteSession) UpdateShootNetworkingFilterDisabled(runtimeID string, shootNetworkingFilterDisabled *bool) apperrors.AppError {
	ret := _m.Called(runtimeID, shootNetworkingFilterDisabled)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, *bool) apperrors.AppError); ok {
		r0 = rf(runtimeID, shootNetworkingFilterDisabled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpdateTenant provides a mock function with given fields: runtimeID, tenant
func (_m *WriteSession) UpdateTenant(runtimeID string, tenant string) apperrors.AppError {
	ret := _m.Called(runtimeID, tenant)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(runtimeID, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// NewWriteSession creates a new instance of WriteSession. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWriteSession(t interface {
	mock.TestingT
	Cleanup(func())
}) *WriteSession {
	mock := &WriteSession{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
