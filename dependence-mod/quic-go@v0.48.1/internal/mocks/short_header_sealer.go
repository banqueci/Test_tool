// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/internal/handshake (interfaces: ShortHeaderSealer)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mocks -destination short_header_sealer.go github.com/quic-go/quic-go/internal/handshake ShortHeaderSealer
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	protocol "github.com/quic-go/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockShortHeaderSealer is a mock of ShortHeaderSealer interface.
type MockShortHeaderSealer struct {
	ctrl     *gomock.Controller
	recorder *MockShortHeaderSealerMockRecorder
}

// MockShortHeaderSealerMockRecorder is the mock recorder for MockShortHeaderSealer.
type MockShortHeaderSealerMockRecorder struct {
	mock *MockShortHeaderSealer
}

// NewMockShortHeaderSealer creates a new mock instance.
func NewMockShortHeaderSealer(ctrl *gomock.Controller) *MockShortHeaderSealer {
	mock := &MockShortHeaderSealer{ctrl: ctrl}
	mock.recorder = &MockShortHeaderSealerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortHeaderSealer) EXPECT() *MockShortHeaderSealerMockRecorder {
	return m.recorder
}

// EncryptHeader mocks base method.
func (m *MockShortHeaderSealer) EncryptHeader(arg0 []byte, arg1 *byte, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EncryptHeader", arg0, arg1, arg2)
}

// EncryptHeader indicates an expected call of EncryptHeader.
func (mr *MockShortHeaderSealerMockRecorder) EncryptHeader(arg0, arg1, arg2 any) *MockShortHeaderSealerEncryptHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptHeader", reflect.TypeOf((*MockShortHeaderSealer)(nil).EncryptHeader), arg0, arg1, arg2)
	return &MockShortHeaderSealerEncryptHeaderCall{Call: call}
}

// MockShortHeaderSealerEncryptHeaderCall wrap *gomock.Call
type MockShortHeaderSealerEncryptHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockShortHeaderSealerEncryptHeaderCall) Return() *MockShortHeaderSealerEncryptHeaderCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockShortHeaderSealerEncryptHeaderCall) Do(f func([]byte, *byte, []byte)) *MockShortHeaderSealerEncryptHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockShortHeaderSealerEncryptHeaderCall) DoAndReturn(f func([]byte, *byte, []byte)) *MockShortHeaderSealerEncryptHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// KeyPhase mocks base method.
func (m *MockShortHeaderSealer) KeyPhase() protocol.KeyPhaseBit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyPhase")
	ret0, _ := ret[0].(protocol.KeyPhaseBit)
	return ret0
}

// KeyPhase indicates an expected call of KeyPhase.
func (mr *MockShortHeaderSealerMockRecorder) KeyPhase() *MockShortHeaderSealerKeyPhaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyPhase", reflect.TypeOf((*MockShortHeaderSealer)(nil).KeyPhase))
	return &MockShortHeaderSealerKeyPhaseCall{Call: call}
}

// MockShortHeaderSealerKeyPhaseCall wrap *gomock.Call
type MockShortHeaderSealerKeyPhaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockShortHeaderSealerKeyPhaseCall) Return(arg0 protocol.KeyPhaseBit) *MockShortHeaderSealerKeyPhaseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockShortHeaderSealerKeyPhaseCall) Do(f func() protocol.KeyPhaseBit) *MockShortHeaderSealerKeyPhaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockShortHeaderSealerKeyPhaseCall) DoAndReturn(f func() protocol.KeyPhaseBit) *MockShortHeaderSealerKeyPhaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Overhead mocks base method.
func (m *MockShortHeaderSealer) Overhead() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Overhead")
	ret0, _ := ret[0].(int)
	return ret0
}

// Overhead indicates an expected call of Overhead.
func (mr *MockShortHeaderSealerMockRecorder) Overhead() *MockShortHeaderSealerOverheadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overhead", reflect.TypeOf((*MockShortHeaderSealer)(nil).Overhead))
	return &MockShortHeaderSealerOverheadCall{Call: call}
}

// MockShortHeaderSealerOverheadCall wrap *gomock.Call
type MockShortHeaderSealerOverheadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockShortHeaderSealerOverheadCall) Return(arg0 int) *MockShortHeaderSealerOverheadCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockShortHeaderSealerOverheadCall) Do(f func() int) *MockShortHeaderSealerOverheadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockShortHeaderSealerOverheadCall) DoAndReturn(f func() int) *MockShortHeaderSealerOverheadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Seal mocks base method.
func (m *MockShortHeaderSealer) Seal(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Seal indicates an expected call of Seal.
func (mr *MockShortHeaderSealerMockRecorder) Seal(arg0, arg1, arg2, arg3 any) *MockShortHeaderSealerSealCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockShortHeaderSealer)(nil).Seal), arg0, arg1, arg2, arg3)
	return &MockShortHeaderSealerSealCall{Call: call}
}

// MockShortHeaderSealerSealCall wrap *gomock.Call
type MockShortHeaderSealerSealCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockShortHeaderSealerSealCall) Return(arg0 []byte) *MockShortHeaderSealerSealCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockShortHeaderSealerSealCall) Do(f func([]byte, []byte, protocol.PacketNumber, []byte) []byte) *MockShortHeaderSealerSealCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockShortHeaderSealerSealCall) DoAndReturn(f func([]byte, []byte, protocol.PacketNumber, []byte) []byte) *MockShortHeaderSealerSealCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}