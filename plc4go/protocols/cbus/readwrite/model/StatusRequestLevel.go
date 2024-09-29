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

// StatusRequestLevel is the corresponding interface of StatusRequestLevel
type StatusRequestLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	StatusRequest
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetStartingGroupAddressLabel returns StartingGroupAddressLabel (property field)
	GetStartingGroupAddressLabel() byte
	// IsStatusRequestLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusRequestLevel()
	// CreateBuilder creates a StatusRequestLevelBuilder
	CreateStatusRequestLevelBuilder() StatusRequestLevelBuilder
}

// _StatusRequestLevel is the data-structure of this message
type _StatusRequestLevel struct {
	StatusRequestContract
	Application               ApplicationIdContainer
	StartingGroupAddressLabel byte
	// Reserved Fields
	reservedField0 *byte
	reservedField1 *byte
}

var _ StatusRequestLevel = (*_StatusRequestLevel)(nil)
var _ StatusRequestRequirements = (*_StatusRequestLevel)(nil)

// NewStatusRequestLevel factory function for _StatusRequestLevel
func NewStatusRequestLevel(statusType byte, application ApplicationIdContainer, startingGroupAddressLabel byte) *_StatusRequestLevel {
	_result := &_StatusRequestLevel{
		StatusRequestContract:     NewStatusRequest(statusType),
		Application:               application,
		StartingGroupAddressLabel: startingGroupAddressLabel,
	}
	_result.StatusRequestContract.(*_StatusRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StatusRequestLevelBuilder is a builder for StatusRequestLevel
type StatusRequestLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(application ApplicationIdContainer, startingGroupAddressLabel byte) StatusRequestLevelBuilder
	// WithApplication adds Application (property field)
	WithApplication(ApplicationIdContainer) StatusRequestLevelBuilder
	// WithStartingGroupAddressLabel adds StartingGroupAddressLabel (property field)
	WithStartingGroupAddressLabel(byte) StatusRequestLevelBuilder
	// Build builds the StatusRequestLevel or returns an error if something is wrong
	Build() (StatusRequestLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StatusRequestLevel
}

// NewStatusRequestLevelBuilder() creates a StatusRequestLevelBuilder
func NewStatusRequestLevelBuilder() StatusRequestLevelBuilder {
	return &_StatusRequestLevelBuilder{_StatusRequestLevel: new(_StatusRequestLevel)}
}

type _StatusRequestLevelBuilder struct {
	*_StatusRequestLevel

	parentBuilder *_StatusRequestBuilder

	err *utils.MultiError
}

var _ (StatusRequestLevelBuilder) = (*_StatusRequestLevelBuilder)(nil)

func (b *_StatusRequestLevelBuilder) setParent(contract StatusRequestContract) {
	b.StatusRequestContract = contract
}

func (b *_StatusRequestLevelBuilder) WithMandatoryFields(application ApplicationIdContainer, startingGroupAddressLabel byte) StatusRequestLevelBuilder {
	return b.WithApplication(application).WithStartingGroupAddressLabel(startingGroupAddressLabel)
}

func (b *_StatusRequestLevelBuilder) WithApplication(application ApplicationIdContainer) StatusRequestLevelBuilder {
	b.Application = application
	return b
}

func (b *_StatusRequestLevelBuilder) WithStartingGroupAddressLabel(startingGroupAddressLabel byte) StatusRequestLevelBuilder {
	b.StartingGroupAddressLabel = startingGroupAddressLabel
	return b
}

func (b *_StatusRequestLevelBuilder) Build() (StatusRequestLevel, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StatusRequestLevel.deepCopy(), nil
}

func (b *_StatusRequestLevelBuilder) MustBuild() StatusRequestLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_StatusRequestLevelBuilder) Done() StatusRequestBuilder {
	return b.parentBuilder
}

func (b *_StatusRequestLevelBuilder) buildForStatusRequest() (StatusRequest, error) {
	return b.Build()
}

func (b *_StatusRequestLevelBuilder) DeepCopy() any {
	_copy := b.CreateStatusRequestLevelBuilder().(*_StatusRequestLevelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStatusRequestLevelBuilder creates a StatusRequestLevelBuilder
func (b *_StatusRequestLevel) CreateStatusRequestLevelBuilder() StatusRequestLevelBuilder {
	if b == nil {
		return NewStatusRequestLevelBuilder()
	}
	return &_StatusRequestLevelBuilder{_StatusRequestLevel: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusRequestLevel) GetParent() StatusRequestContract {
	return m.StatusRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequestLevel) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_StatusRequestLevel) GetStartingGroupAddressLabel() byte {
	return m.StartingGroupAddressLabel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStatusRequestLevel(structType any) StatusRequestLevel {
	if casted, ok := structType.(StatusRequestLevel); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequestLevel); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequestLevel) GetTypeName() string {
	return "StatusRequestLevel"
}

func (m *_StatusRequestLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.StatusRequestContract.(*_StatusRequest).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (startingGroupAddressLabel)
	lengthInBits += 8

	return lengthInBits
}

func (m *_StatusRequestLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_StatusRequestLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_StatusRequest) (__statusRequestLevel StatusRequestLevel, err error) {
	m.StatusRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequestLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequestLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x73))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x07))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}
	m.Application = application

	startingGroupAddressLabel, err := ReadSimpleField(ctx, "startingGroupAddressLabel", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingGroupAddressLabel' field"))
	}
	m.StartingGroupAddressLabel = startingGroupAddressLabel

	// Validation
	if !(bool(bool(bool(bool(bool(bool(bool(bool((startingGroupAddressLabel) == (0x00))) || bool(bool((startingGroupAddressLabel) == (0x20)))) || bool(bool((startingGroupAddressLabel) == (0x40)))) || bool(bool((startingGroupAddressLabel) == (0x60)))) || bool(bool((startingGroupAddressLabel) == (0x80)))) || bool(bool((startingGroupAddressLabel) == (0xA0)))) || bool(bool((startingGroupAddressLabel) == (0xC0)))) || bool(bool((startingGroupAddressLabel) == (0xE0)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "invalid label"})
	}

	if closeErr := readBuffer.CloseContext("StatusRequestLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequestLevel")
	}

	return m, nil
}

func (m *_StatusRequestLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusRequestLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusRequestLevel")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x73), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x07), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", m.GetApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'application' field")
		}

		if err := WriteSimpleField[byte](ctx, "startingGroupAddressLabel", m.GetStartingGroupAddressLabel(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingGroupAddressLabel' field")
		}

		if popErr := writeBuffer.PopContext("StatusRequestLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusRequestLevel")
		}
		return nil
	}
	return m.StatusRequestContract.(*_StatusRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusRequestLevel) IsStatusRequestLevel() {}

func (m *_StatusRequestLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StatusRequestLevel) deepCopy() *_StatusRequestLevel {
	if m == nil {
		return nil
	}
	_StatusRequestLevelCopy := &_StatusRequestLevel{
		m.StatusRequestContract.(*_StatusRequest).deepCopy(),
		m.Application,
		m.StartingGroupAddressLabel,
		m.reservedField0,
		m.reservedField1,
	}
	m.StatusRequestContract.(*_StatusRequest)._SubType = m
	return _StatusRequestLevelCopy
}

func (m *_StatusRequestLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
