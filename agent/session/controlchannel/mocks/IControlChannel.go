// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by mockery v1.0.0
package mocks

import (
	context "github.com/aws/amazon-ssm-agent/agent/context"
	processor "github.com/aws/amazon-ssm-agent/agent/framework/processor"
	log "github.com/aws/amazon-ssm-agent/agent/log"
	service "github.com/aws/amazon-ssm-agent/agent/session/service"
	mock "github.com/stretchr/testify/mock"
)

// IControlChannel is an autogenerated mock type for the IControlChannel type
type IControlChannel struct {
	mock.Mock
}

// Close provides a mock function with given fields: _a0
func (_m *IControlChannel) Close(_a0 log.T) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: _a0, mgsService, _a2, instanceId
func (_m *IControlChannel) Initialize(_a0 context.T, mgsService service.Service, _a2 processor.Processor, instanceId string) {
	_m.Called(_a0, mgsService, _a2, instanceId)
}

// Open provides a mock function with given fields: _a0
func (_m *IControlChannel) Open(_a0 log.T) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reconnect provides a mock function with given fields: _a0
func (_m *IControlChannel) Reconnect(_a0 log.T) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessage provides a mock function with given fields: _a0, input, inputType
func (_m *IControlChannel) SendMessage(_a0 log.T, input []byte, inputType int) error {
	ret := _m.Called(_a0, input, inputType)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T, []byte, int) error); ok {
		r0 = rf(_a0, input, inputType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWebSocket provides a mock function with given fields: _a0, mgsService, _a2, instanceId
func (_m *IControlChannel) SetWebSocket(_a0 context.T, mgsService service.Service, _a2 processor.Processor, instanceId string) error {
	ret := _m.Called(_a0, mgsService, _a2, instanceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.T, service.Service, processor.Processor, string) error); ok {
		r0 = rf(_a0, mgsService, _a2, instanceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
