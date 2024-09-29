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

// ErrorResponse is the corresponding interface of ErrorResponse
type ErrorResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// IsErrorResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorResponse()
	// CreateBuilder creates a ErrorResponseBuilder
	CreateErrorResponseBuilder() ErrorResponseBuilder
}

// _ErrorResponse is the data-structure of this message
type _ErrorResponse struct {
	AmsPacketContract
}

var _ ErrorResponse = (*_ErrorResponse)(nil)
var _ AmsPacketRequirements = (*_ErrorResponse)(nil)

// NewErrorResponse factory function for _ErrorResponse
func NewErrorResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_ErrorResponse {
	_result := &_ErrorResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorResponseBuilder is a builder for ErrorResponse
type ErrorResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ErrorResponseBuilder
	// Build builds the ErrorResponse or returns an error if something is wrong
	Build() (ErrorResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorResponse
}

// NewErrorResponseBuilder() creates a ErrorResponseBuilder
func NewErrorResponseBuilder() ErrorResponseBuilder {
	return &_ErrorResponseBuilder{_ErrorResponse: new(_ErrorResponse)}
}

type _ErrorResponseBuilder struct {
	*_ErrorResponse

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (ErrorResponseBuilder) = (*_ErrorResponseBuilder)(nil)

func (b *_ErrorResponseBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
}

func (b *_ErrorResponseBuilder) WithMandatoryFields() ErrorResponseBuilder {
	return b
}

func (b *_ErrorResponseBuilder) Build() (ErrorResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorResponse.deepCopy(), nil
}

func (b *_ErrorResponseBuilder) MustBuild() ErrorResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ErrorResponseBuilder) Done() AmsPacketBuilder {
	return b.parentBuilder
}

func (b *_ErrorResponseBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_ErrorResponseBuilder) DeepCopy() any {
	_copy := b.CreateErrorResponseBuilder().(*_ErrorResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorResponseBuilder creates a ErrorResponseBuilder
func (b *_ErrorResponse) CreateErrorResponseBuilder() ErrorResponseBuilder {
	if b == nil {
		return NewErrorResponseBuilder()
	}
	return &_ErrorResponseBuilder{_ErrorResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorResponse) GetCommandId() CommandId {
	return 0
}

func (m *_ErrorResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

// Deprecated: use the interface for direct cast
func CastErrorResponse(structType any) ErrorResponse {
	if casted, ok := structType.(ErrorResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorResponse) GetTypeName() string {
	return "ErrorResponse"
}

func (m *_ErrorResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ErrorResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__errorResponse ErrorResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ErrorResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorResponse")
	}

	return m, nil
}

func (m *_ErrorResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorResponse")
		}

		if popErr := writeBuffer.PopContext("ErrorResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorResponse) IsErrorResponse() {}

func (m *_ErrorResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorResponse) deepCopy() *_ErrorResponse {
	if m == nil {
		return nil
	}
	_ErrorResponseCopy := &_ErrorResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _ErrorResponseCopy
}

func (m *_ErrorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
