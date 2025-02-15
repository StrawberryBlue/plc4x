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

// ApduData is the corresponding interface of ApduData
type ApduData interface {
	ApduDataContract
	ApduDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsApduData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduData()
	// CreateBuilder creates a ApduDataBuilder
	CreateApduDataBuilder() ApduDataBuilder
}

// ApduDataContract provides a set of functions which can be overwritten by a sub struct
type ApduDataContract interface {
	// GetDataLength() returns a parser argument
	GetDataLength() uint8
	// IsApduData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduData()
	// CreateBuilder creates a ApduDataBuilder
	CreateApduDataBuilder() ApduDataBuilder
}

// ApduDataRequirements provides a set of functions which need to be implemented by a sub struct
type ApduDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetApciType returns ApciType (discriminator field)
	GetApciType() uint8
}

// _ApduData is the data-structure of this message
type _ApduData struct {
	_SubType ApduData

	// Arguments.
	DataLength uint8
}

var _ ApduDataContract = (*_ApduData)(nil)

// NewApduData factory function for _ApduData
func NewApduData(dataLength uint8) *_ApduData {
	return &_ApduData{DataLength: dataLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataBuilder is a builder for ApduData
type ApduDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataBuilder
	// AsApduDataGroupValueRead converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataGroupValueRead() interface {
		ApduDataGroupValueReadBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataGroupValueResponse converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataGroupValueResponse() interface {
		ApduDataGroupValueResponseBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataGroupValueWrite converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataGroupValueWrite() interface {
		ApduDataGroupValueWriteBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataIndividualAddressWrite converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataIndividualAddressWrite() interface {
		ApduDataIndividualAddressWriteBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataIndividualAddressRead converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataIndividualAddressRead() interface {
		ApduDataIndividualAddressReadBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataIndividualAddressResponse converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataIndividualAddressResponse() interface {
		ApduDataIndividualAddressResponseBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataAdcRead converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataAdcRead() interface {
		ApduDataAdcReadBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataAdcResponse converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataAdcResponse() interface {
		ApduDataAdcResponseBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataMemoryRead converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataMemoryRead() interface {
		ApduDataMemoryReadBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataMemoryResponse converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataMemoryResponse() interface {
		ApduDataMemoryResponseBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataMemoryWrite converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataMemoryWrite() interface {
		ApduDataMemoryWriteBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataUserMessage converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataUserMessage() interface {
		ApduDataUserMessageBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataDeviceDescriptorRead converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataDeviceDescriptorRead() interface {
		ApduDataDeviceDescriptorReadBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataDeviceDescriptorResponse converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataDeviceDescriptorResponse() interface {
		ApduDataDeviceDescriptorResponseBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataRestart converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataRestart() interface {
		ApduDataRestartBuilder
		Done() ApduDataBuilder
	}
	// AsApduDataOther converts this build to a subType of ApduData. It is always possible to return to current builder using Done()
	AsApduDataOther() interface {
		ApduDataOtherBuilder
		Done() ApduDataBuilder
	}
	// Build builds the ApduData or returns an error if something is wrong
	PartialBuild() (ApduDataContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ApduDataContract
	// Build builds the ApduData or returns an error if something is wrong
	Build() (ApduData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduData
}

// NewApduDataBuilder() creates a ApduDataBuilder
func NewApduDataBuilder() ApduDataBuilder {
	return &_ApduDataBuilder{_ApduData: new(_ApduData)}
}

type _ApduDataChildBuilder interface {
	utils.Copyable
	setParent(ApduDataContract)
	buildForApduData() (ApduData, error)
}

type _ApduDataBuilder struct {
	*_ApduData

	childBuilder _ApduDataChildBuilder

	err *utils.MultiError
}

var _ (ApduDataBuilder) = (*_ApduDataBuilder)(nil)

func (b *_ApduDataBuilder) WithMandatoryFields() ApduDataBuilder {
	return b
}

func (b *_ApduDataBuilder) PartialBuild() (ApduDataContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduData.deepCopy(), nil
}

func (b *_ApduDataBuilder) PartialMustBuild() ApduDataContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataBuilder) AsApduDataGroupValueRead() interface {
	ApduDataGroupValueReadBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataGroupValueReadBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataGroupValueReadBuilder().(*_ApduDataGroupValueReadBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataGroupValueResponse() interface {
	ApduDataGroupValueResponseBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataGroupValueResponseBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataGroupValueResponseBuilder().(*_ApduDataGroupValueResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataGroupValueWrite() interface {
	ApduDataGroupValueWriteBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataGroupValueWriteBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataGroupValueWriteBuilder().(*_ApduDataGroupValueWriteBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataIndividualAddressWrite() interface {
	ApduDataIndividualAddressWriteBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataIndividualAddressWriteBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataIndividualAddressWriteBuilder().(*_ApduDataIndividualAddressWriteBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataIndividualAddressRead() interface {
	ApduDataIndividualAddressReadBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataIndividualAddressReadBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataIndividualAddressReadBuilder().(*_ApduDataIndividualAddressReadBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataIndividualAddressResponse() interface {
	ApduDataIndividualAddressResponseBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataIndividualAddressResponseBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataIndividualAddressResponseBuilder().(*_ApduDataIndividualAddressResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataAdcRead() interface {
	ApduDataAdcReadBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataAdcReadBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataAdcReadBuilder().(*_ApduDataAdcReadBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataAdcResponse() interface {
	ApduDataAdcResponseBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataAdcResponseBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataAdcResponseBuilder().(*_ApduDataAdcResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataMemoryRead() interface {
	ApduDataMemoryReadBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataMemoryReadBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataMemoryReadBuilder().(*_ApduDataMemoryReadBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataMemoryResponse() interface {
	ApduDataMemoryResponseBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataMemoryResponseBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataMemoryResponseBuilder().(*_ApduDataMemoryResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataMemoryWrite() interface {
	ApduDataMemoryWriteBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataMemoryWriteBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataMemoryWriteBuilder().(*_ApduDataMemoryWriteBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataUserMessage() interface {
	ApduDataUserMessageBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataUserMessageBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataUserMessageBuilder().(*_ApduDataUserMessageBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataDeviceDescriptorRead() interface {
	ApduDataDeviceDescriptorReadBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataDeviceDescriptorReadBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataDeviceDescriptorReadBuilder().(*_ApduDataDeviceDescriptorReadBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataDeviceDescriptorResponse() interface {
	ApduDataDeviceDescriptorResponseBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataDeviceDescriptorResponseBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataDeviceDescriptorResponseBuilder().(*_ApduDataDeviceDescriptorResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataRestart() interface {
	ApduDataRestartBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataRestartBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataRestartBuilder().(*_ApduDataRestartBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) AsApduDataOther() interface {
	ApduDataOtherBuilder
	Done() ApduDataBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ApduDataOtherBuilder
		Done() ApduDataBuilder
	}); ok {
		return cb
	}
	cb := NewApduDataOtherBuilder().(*_ApduDataOtherBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ApduDataBuilder) Build() (ApduData, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForApduData()
}

func (b *_ApduDataBuilder) MustBuild() ApduData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataBuilder) DeepCopy() any {
	_copy := b.CreateApduDataBuilder().(*_ApduDataBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ApduDataChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataBuilder creates a ApduDataBuilder
func (b *_ApduData) CreateApduDataBuilder() ApduDataBuilder {
	if b == nil {
		return NewApduDataBuilder()
	}
	return &_ApduDataBuilder{_ApduData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduData(structType any) ApduData {
	if casted, ok := structType.(ApduData); ok {
		return casted
	}
	if casted, ok := structType.(*ApduData); ok {
		return *casted
	}
	return nil
}

func (m *_ApduData) GetTypeName() string {
	return "ApduData"
}

func (m *_ApduData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apciType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ApduData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ApduDataParse[T ApduData](ctx context.Context, theBytes []byte, dataLength uint8) (T, error) {
	return ApduDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataParseWithBufferProducer[T ApduData](dataLength uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ApduDataParseWithBuffer[T](ctx, readBuffer, dataLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ApduDataParseWithBuffer[T ApduData](ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (T, error) {
	v, err := (&_ApduData{DataLength: dataLength}).parse(ctx, readBuffer, dataLength)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ApduData) parse(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (__apduData ApduData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	apciType, err := ReadDiscriminatorField[uint8](ctx, "apciType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apciType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ApduData
	switch {
	case apciType == 0x0: // ApduDataGroupValueRead
		if _child, err = new(_ApduDataGroupValueRead).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataGroupValueRead for type-switch of ApduData")
		}
	case apciType == 0x1: // ApduDataGroupValueResponse
		if _child, err = new(_ApduDataGroupValueResponse).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataGroupValueResponse for type-switch of ApduData")
		}
	case apciType == 0x2: // ApduDataGroupValueWrite
		if _child, err = new(_ApduDataGroupValueWrite).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataGroupValueWrite for type-switch of ApduData")
		}
	case apciType == 0x3: // ApduDataIndividualAddressWrite
		if _child, err = new(_ApduDataIndividualAddressWrite).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataIndividualAddressWrite for type-switch of ApduData")
		}
	case apciType == 0x4: // ApduDataIndividualAddressRead
		if _child, err = new(_ApduDataIndividualAddressRead).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataIndividualAddressRead for type-switch of ApduData")
		}
	case apciType == 0x5: // ApduDataIndividualAddressResponse
		if _child, err = new(_ApduDataIndividualAddressResponse).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataIndividualAddressResponse for type-switch of ApduData")
		}
	case apciType == 0x6: // ApduDataAdcRead
		if _child, err = new(_ApduDataAdcRead).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataAdcRead for type-switch of ApduData")
		}
	case apciType == 0x7: // ApduDataAdcResponse
		if _child, err = new(_ApduDataAdcResponse).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataAdcResponse for type-switch of ApduData")
		}
	case apciType == 0x8: // ApduDataMemoryRead
		if _child, err = new(_ApduDataMemoryRead).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataMemoryRead for type-switch of ApduData")
		}
	case apciType == 0x9: // ApduDataMemoryResponse
		if _child, err = new(_ApduDataMemoryResponse).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataMemoryResponse for type-switch of ApduData")
		}
	case apciType == 0xA: // ApduDataMemoryWrite
		if _child, err = new(_ApduDataMemoryWrite).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataMemoryWrite for type-switch of ApduData")
		}
	case apciType == 0xB: // ApduDataUserMessage
		if _child, err = new(_ApduDataUserMessage).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataUserMessage for type-switch of ApduData")
		}
	case apciType == 0xC: // ApduDataDeviceDescriptorRead
		if _child, err = new(_ApduDataDeviceDescriptorRead).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataDeviceDescriptorRead for type-switch of ApduData")
		}
	case apciType == 0xD: // ApduDataDeviceDescriptorResponse
		if _child, err = new(_ApduDataDeviceDescriptorResponse).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataDeviceDescriptorResponse for type-switch of ApduData")
		}
	case apciType == 0xE: // ApduDataRestart
		if _child, err = new(_ApduDataRestart).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataRestart for type-switch of ApduData")
		}
	case apciType == 0xF: // ApduDataOther
		if _child, err = new(_ApduDataOther).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataOther for type-switch of ApduData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [apciType=%v]", apciType)
	}

	if closeErr := readBuffer.CloseContext("ApduData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduData")
	}

	return _child, nil
}

func (pm *_ApduData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApduData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduData")
	}

	if err := WriteDiscriminatorField(ctx, "apciType", m.GetApciType(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
		return errors.Wrap(err, "Error serializing 'apciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduData")
	}
	return nil
}

////
// Arguments Getter

func (m *_ApduData) GetDataLength() uint8 {
	return m.DataLength
}

//
////

func (m *_ApduData) IsApduData() {}

func (m *_ApduData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduData) deepCopy() *_ApduData {
	if m == nil {
		return nil
	}
	_ApduDataCopy := &_ApduData{
		nil, // will be set by child
		m.DataLength,
	}
	return _ApduDataCopy
}
