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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CipReadRequest is the corresponding interface of CipReadRequest
type CipReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetTag returns Tag (property field)
	GetTag() []byte
	// GetElementNb returns ElementNb (property field)
	GetElementNb() uint16
	// IsCipReadRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipReadRequest()
	// CreateBuilder creates a CipReadRequestBuilder
	CreateCipReadRequestBuilder() CipReadRequestBuilder
}

// _CipReadRequest is the data-structure of this message
type _CipReadRequest struct {
	CipServiceContract
	Tag       []byte
	ElementNb uint16
}

var _ CipReadRequest = (*_CipReadRequest)(nil)
var _ CipServiceRequirements = (*_CipReadRequest)(nil)

// NewCipReadRequest factory function for _CipReadRequest
func NewCipReadRequest(tag []byte, elementNb uint16, serviceLen uint16) *_CipReadRequest {
	_result := &_CipReadRequest{
		CipServiceContract: NewCipService(serviceLen),
		Tag:                tag,
		ElementNb:          elementNb,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipReadRequestBuilder is a builder for CipReadRequest
type CipReadRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tag []byte, elementNb uint16) CipReadRequestBuilder
	// WithTag adds Tag (property field)
	WithTag(...byte) CipReadRequestBuilder
	// WithElementNb adds ElementNb (property field)
	WithElementNb(uint16) CipReadRequestBuilder
	// Build builds the CipReadRequest or returns an error if something is wrong
	Build() (CipReadRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipReadRequest
}

// NewCipReadRequestBuilder() creates a CipReadRequestBuilder
func NewCipReadRequestBuilder() CipReadRequestBuilder {
	return &_CipReadRequestBuilder{_CipReadRequest: new(_CipReadRequest)}
}

type _CipReadRequestBuilder struct {
	*_CipReadRequest

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (CipReadRequestBuilder) = (*_CipReadRequestBuilder)(nil)

func (b *_CipReadRequestBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
}

func (b *_CipReadRequestBuilder) WithMandatoryFields(tag []byte, elementNb uint16) CipReadRequestBuilder {
	return b.WithTag(tag...).WithElementNb(elementNb)
}

func (b *_CipReadRequestBuilder) WithTag(tag ...byte) CipReadRequestBuilder {
	b.Tag = tag
	return b
}

func (b *_CipReadRequestBuilder) WithElementNb(elementNb uint16) CipReadRequestBuilder {
	b.ElementNb = elementNb
	return b
}

func (b *_CipReadRequestBuilder) Build() (CipReadRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipReadRequest.deepCopy(), nil
}

func (b *_CipReadRequestBuilder) MustBuild() CipReadRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CipReadRequestBuilder) Done() CipServiceBuilder {
	return b.parentBuilder
}

func (b *_CipReadRequestBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_CipReadRequestBuilder) DeepCopy() any {
	_copy := b.CreateCipReadRequestBuilder().(*_CipReadRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipReadRequestBuilder creates a CipReadRequestBuilder
func (b *_CipReadRequest) CreateCipReadRequestBuilder() CipReadRequestBuilder {
	if b == nil {
		return NewCipReadRequestBuilder()
	}
	return &_CipReadRequestBuilder{_CipReadRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipReadRequest) GetService() uint8 {
	return 0x4C
}

func (m *_CipReadRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipReadRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipReadRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipReadRequest) GetTag() []byte {
	return m.Tag
}

func (m *_CipReadRequest) GetElementNb() uint16 {
	return m.ElementNb
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipReadRequest(structType any) CipReadRequest {
	if casted, ok := structType.(CipReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipReadRequest) GetTypeName() string {
	return "CipReadRequest"
}

func (m *_CipReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.Tag) > 0 {
		lengthInBits += 8 * uint16(len(m.Tag))
	}

	// Simple field (elementNb)
	lengthInBits += 16

	return lengthInBits
}

func (m *_CipReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipReadRequest CipReadRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	tag, err := readBuffer.ReadByteArray("tag", int((int32(requestPathSize) * int32(int32(2)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tag' field"))
	}
	m.Tag = tag

	elementNb, err := ReadSimpleField(ctx, "elementNb", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementNb' field"))
	}
	m.ElementNb = elementNb

	if closeErr := readBuffer.CloseContext("CipReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipReadRequest")
	}

	return m, nil
}

func (m *_CipReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipReadRequest")
		}
		requestPathSize := uint8(uint8(uint8(len(m.GetTag()))) / uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "requestPathSize", requestPathSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteByteArrayField(ctx, "tag", m.GetTag(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'tag' field")
		}

		if err := WriteSimpleField[uint16](ctx, "elementNb", m.GetElementNb(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'elementNb' field")
		}

		if popErr := writeBuffer.PopContext("CipReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipReadRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipReadRequest) IsCipReadRequest() {}

func (m *_CipReadRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipReadRequest) deepCopy() *_CipReadRequest {
	if m == nil {
		return nil
	}
	_CipReadRequestCopy := &_CipReadRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Tag),
		m.ElementNb,
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _CipReadRequestCopy
}

func (m *_CipReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
