// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v1.0.0
package mocks

import component "github.com/ksonnet/ksonnet/component"
import mock "github.com/stretchr/testify/mock"

// Namespace is an autogenerated mock type for the Namespace type
type Namespace struct {
	mock.Mock
}

// Components provides a mock function with given fields:
func (_m *Namespace) Components() ([]component.Component, error) {
	ret := _m.Called()

	var r0 []component.Component
	if rf, ok := ret.Get(0).(func() []component.Component); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.Component)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dir provides a mock function with given fields:
func (_m *Namespace) Dir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Namespace) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Params provides a mock function with given fields: envName
func (_m *Namespace) Params(envName string) ([]component.NamespaceParameter, error) {
	ret := _m.Called(envName)

	var r0 []component.NamespaceParameter
	if rf, ok := ret.Get(0).(func(string) []component.NamespaceParameter); ok {
		r0 = rf(envName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.NamespaceParameter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(envName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParamsPath provides a mock function with given fields:
func (_m *Namespace) ParamsPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResolvedParams provides a mock function with given fields:
func (_m *Namespace) ResolvedParams() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetParam provides a mock function with given fields: path, value
func (_m *Namespace) SetParam(path []string, value interface{}) error {
	ret := _m.Called(path, value)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, interface{}) error); ok {
		r0 = rf(path, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
