/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadExceptionStatusRequest is the corresponding interface of ModbusPDUReadExceptionStatusRequest
type ModbusPDUReadExceptionStatusRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// IsModbusPDUReadExceptionStatusRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadExceptionStatusRequest()
	// CreateBuilder creates a ModbusPDUReadExceptionStatusRequestBuilder
	CreateModbusPDUReadExceptionStatusRequestBuilder() ModbusPDUReadExceptionStatusRequestBuilder
}

// _ModbusPDUReadExceptionStatusRequest is the data-structure of this message
type _ModbusPDUReadExceptionStatusRequest struct {
	ModbusPDUContract
}

var _ ModbusPDUReadExceptionStatusRequest = (*_ModbusPDUReadExceptionStatusRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadExceptionStatusRequest)(nil)

// NewModbusPDUReadExceptionStatusRequest factory function for _ModbusPDUReadExceptionStatusRequest
func NewModbusPDUReadExceptionStatusRequest() *_ModbusPDUReadExceptionStatusRequest {
	_result := &_ModbusPDUReadExceptionStatusRequest{
		ModbusPDUContract: NewModbusPDU(),
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadExceptionStatusRequestBuilder is a builder for ModbusPDUReadExceptionStatusRequest
type ModbusPDUReadExceptionStatusRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ModbusPDUReadExceptionStatusRequestBuilder
	// Build builds the ModbusPDUReadExceptionStatusRequest or returns an error if something is wrong
	Build() (ModbusPDUReadExceptionStatusRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadExceptionStatusRequest
}

// NewModbusPDUReadExceptionStatusRequestBuilder() creates a ModbusPDUReadExceptionStatusRequestBuilder
func NewModbusPDUReadExceptionStatusRequestBuilder() ModbusPDUReadExceptionStatusRequestBuilder {
	return &_ModbusPDUReadExceptionStatusRequestBuilder{_ModbusPDUReadExceptionStatusRequest: new(_ModbusPDUReadExceptionStatusRequest)}
}

type _ModbusPDUReadExceptionStatusRequestBuilder struct {
	*_ModbusPDUReadExceptionStatusRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadExceptionStatusRequestBuilder) = (*_ModbusPDUReadExceptionStatusRequestBuilder)(nil)

func (b *_ModbusPDUReadExceptionStatusRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
}

func (b *_ModbusPDUReadExceptionStatusRequestBuilder) WithMandatoryFields() ModbusPDUReadExceptionStatusRequestBuilder {
	return b
}

func (b *_ModbusPDUReadExceptionStatusRequestBuilder) Build() (ModbusPDUReadExceptionStatusRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadExceptionStatusRequest.deepCopy(), nil
}

func (b *_ModbusPDUReadExceptionStatusRequestBuilder) MustBuild() ModbusPDUReadExceptionStatusRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ModbusPDUReadExceptionStatusRequestBuilder) Done() ModbusPDUBuilder {
	return b.parentBuilder
}

func (b *_ModbusPDUReadExceptionStatusRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadExceptionStatusRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadExceptionStatusRequestBuilder().(*_ModbusPDUReadExceptionStatusRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadExceptionStatusRequestBuilder creates a ModbusPDUReadExceptionStatusRequestBuilder
func (b *_ModbusPDUReadExceptionStatusRequest) CreateModbusPDUReadExceptionStatusRequestBuilder() ModbusPDUReadExceptionStatusRequestBuilder {
	if b == nil {
		return NewModbusPDUReadExceptionStatusRequestBuilder()
	}
	return &_ModbusPDUReadExceptionStatusRequestBuilder{_ModbusPDUReadExceptionStatusRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadExceptionStatusRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadExceptionStatusRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadExceptionStatusRequest(structType any) ModbusPDUReadExceptionStatusRequest {
	if casted, ok := structType.(ModbusPDUReadExceptionStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadExceptionStatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusRequest"
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadExceptionStatusRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadExceptionStatusRequest ModbusPDUReadExceptionStatusRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadExceptionStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadExceptionStatusRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadExceptionStatusRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadExceptionStatusRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadExceptionStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadExceptionStatusRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadExceptionStatusRequest) IsModbusPDUReadExceptionStatusRequest() {}

func (m *_ModbusPDUReadExceptionStatusRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadExceptionStatusRequest) deepCopy() *_ModbusPDUReadExceptionStatusRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReadExceptionStatusRequestCopy := &_ModbusPDUReadExceptionStatusRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadExceptionStatusRequestCopy
}

func (m *_ModbusPDUReadExceptionStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
