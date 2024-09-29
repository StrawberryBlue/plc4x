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

// S7PayloadUserDataItemCpuFunctionReadSzlResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionReadSzlResponse
type S7PayloadUserDataItemCpuFunctionReadSzlResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetItems returns Items (property field)
	GetItems() []byte
	// IsS7PayloadUserDataItemCpuFunctionReadSzlResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionReadSzlResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder
	CreateS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder() S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder
}

// _S7PayloadUserDataItemCpuFunctionReadSzlResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionReadSzlResponse struct {
	S7PayloadUserDataItemContract
	Items []byte
}

var _ S7PayloadUserDataItemCpuFunctionReadSzlResponse = (*_S7PayloadUserDataItemCpuFunctionReadSzlResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionReadSzlResponse)(nil)

// NewS7PayloadUserDataItemCpuFunctionReadSzlResponse factory function for _S7PayloadUserDataItemCpuFunctionReadSzlResponse
func NewS7PayloadUserDataItemCpuFunctionReadSzlResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, items []byte) *_S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Items:                         items,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder is a builder for S7PayloadUserDataItemCpuFunctionReadSzlResponse
type S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []byte) S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder
	// WithItems adds Items (property field)
	WithItems(...byte) S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionReadSzlResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionReadSzlResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionReadSzlResponse
}

// NewS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder() creates a S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder
func NewS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder() S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder {
	return &_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder{_S7PayloadUserDataItemCpuFunctionReadSzlResponse: new(_S7PayloadUserDataItemCpuFunctionReadSzlResponse)}
}

type _S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionReadSzlResponse

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) = (*_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder)(nil)

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) WithMandatoryFields(items []byte) S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder {
	return b.WithItems(items...)
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) WithItems(items ...byte) S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder {
	b.Items = items
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) Build() (S7PayloadUserDataItemCpuFunctionReadSzlResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCpuFunctionReadSzlResponse.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) Done() S7PayloadUserDataItemBuilder {
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder().(*_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder creates a S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder
func (b *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) CreateS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder() S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionReadSzlResponseBuilder{_S7PayloadUserDataItemCpuFunctionReadSzlResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetItems() []byte {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(structType any) S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		lengthInBits += 8 * uint16(len(m.Items))
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionReadSzlResponse S7PayloadUserDataItemCpuFunctionReadSzlResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := readBuffer.ReadByteArray("items", int(dataLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
		}

		if err := WriteByteArrayField(ctx, "items", m.GetItems(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) IsS7PayloadUserDataItemCpuFunctionReadSzlResponse() {
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) deepCopy() *_S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionReadSzlResponseCopy := &_S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Items),
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionReadSzlResponseCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
