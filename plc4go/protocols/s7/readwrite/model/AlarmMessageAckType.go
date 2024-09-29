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

// AlarmMessageAckType is the corresponding interface of AlarmMessageAckType
type AlarmMessageAckType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectAckType
	// IsAlarmMessageAckType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarmMessageAckType()
	// CreateBuilder creates a AlarmMessageAckTypeBuilder
	CreateAlarmMessageAckTypeBuilder() AlarmMessageAckTypeBuilder
}

// _AlarmMessageAckType is the data-structure of this message
type _AlarmMessageAckType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []AlarmMessageObjectAckType
}

var _ AlarmMessageAckType = (*_AlarmMessageAckType)(nil)

// NewAlarmMessageAckType factory function for _AlarmMessageAckType
func NewAlarmMessageAckType(functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectAckType) *_AlarmMessageAckType {
	return &_AlarmMessageAckType{FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AlarmMessageAckTypeBuilder is a builder for AlarmMessageAckType
type AlarmMessageAckTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectAckType) AlarmMessageAckTypeBuilder
	// WithFunctionId adds FunctionId (property field)
	WithFunctionId(uint8) AlarmMessageAckTypeBuilder
	// WithNumberOfObjects adds NumberOfObjects (property field)
	WithNumberOfObjects(uint8) AlarmMessageAckTypeBuilder
	// WithMessageObjects adds MessageObjects (property field)
	WithMessageObjects(...AlarmMessageObjectAckType) AlarmMessageAckTypeBuilder
	// Build builds the AlarmMessageAckType or returns an error if something is wrong
	Build() (AlarmMessageAckType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AlarmMessageAckType
}

// NewAlarmMessageAckTypeBuilder() creates a AlarmMessageAckTypeBuilder
func NewAlarmMessageAckTypeBuilder() AlarmMessageAckTypeBuilder {
	return &_AlarmMessageAckTypeBuilder{_AlarmMessageAckType: new(_AlarmMessageAckType)}
}

type _AlarmMessageAckTypeBuilder struct {
	*_AlarmMessageAckType

	err *utils.MultiError
}

var _ (AlarmMessageAckTypeBuilder) = (*_AlarmMessageAckTypeBuilder)(nil)

func (b *_AlarmMessageAckTypeBuilder) WithMandatoryFields(functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectAckType) AlarmMessageAckTypeBuilder {
	return b.WithFunctionId(functionId).WithNumberOfObjects(numberOfObjects).WithMessageObjects(messageObjects...)
}

func (b *_AlarmMessageAckTypeBuilder) WithFunctionId(functionId uint8) AlarmMessageAckTypeBuilder {
	b.FunctionId = functionId
	return b
}

func (b *_AlarmMessageAckTypeBuilder) WithNumberOfObjects(numberOfObjects uint8) AlarmMessageAckTypeBuilder {
	b.NumberOfObjects = numberOfObjects
	return b
}

func (b *_AlarmMessageAckTypeBuilder) WithMessageObjects(messageObjects ...AlarmMessageObjectAckType) AlarmMessageAckTypeBuilder {
	b.MessageObjects = messageObjects
	return b
}

func (b *_AlarmMessageAckTypeBuilder) Build() (AlarmMessageAckType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AlarmMessageAckType.deepCopy(), nil
}

func (b *_AlarmMessageAckTypeBuilder) MustBuild() AlarmMessageAckType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AlarmMessageAckTypeBuilder) DeepCopy() any {
	_copy := b.CreateAlarmMessageAckTypeBuilder().(*_AlarmMessageAckTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAlarmMessageAckTypeBuilder creates a AlarmMessageAckTypeBuilder
func (b *_AlarmMessageAckType) CreateAlarmMessageAckTypeBuilder() AlarmMessageAckTypeBuilder {
	if b == nil {
		return NewAlarmMessageAckTypeBuilder()
	}
	return &_AlarmMessageAckTypeBuilder{_AlarmMessageAckType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageAckType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessageAckType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessageAckType) GetMessageObjects() []AlarmMessageObjectAckType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAlarmMessageAckType(structType any) AlarmMessageAckType {
	if casted, ok := structType.(AlarmMessageAckType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageAckType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageAckType) GetTypeName() string {
	return "AlarmMessageAckType"
}

func (m *_AlarmMessageAckType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessageAckType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageAckTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageAckType, error) {
	return AlarmMessageAckTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageAckTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckType, error) {
		return AlarmMessageAckTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageAckTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckType, error) {
	v, err := (&_AlarmMessageAckType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AlarmMessageAckType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessageAckType AlarmMessageAckType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageAckType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadSimpleField(ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	m.FunctionId = functionId

	numberOfObjects, err := ReadSimpleField(ctx, "numberOfObjects", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfObjects' field"))
	}
	m.NumberOfObjects = numberOfObjects

	messageObjects, err := ReadCountArrayField[AlarmMessageObjectAckType](ctx, "messageObjects", ReadComplex[AlarmMessageObjectAckType](AlarmMessageObjectAckTypeParseWithBuffer, readBuffer), uint64(numberOfObjects))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageObjects' field"))
	}
	m.MessageObjects = messageObjects

	if closeErr := readBuffer.CloseContext("AlarmMessageAckType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckType")
	}

	return m, nil
}

func (m *_AlarmMessageAckType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageAckType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageAckType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckType")
	}

	if err := WriteSimpleField[uint8](ctx, "functionId", m.GetFunctionId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "numberOfObjects", m.GetNumberOfObjects(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfObjects' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "messageObjects", m.GetMessageObjects(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'messageObjects' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckType")
	}
	return nil
}

func (m *_AlarmMessageAckType) IsAlarmMessageAckType() {}

func (m *_AlarmMessageAckType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AlarmMessageAckType) deepCopy() *_AlarmMessageAckType {
	if m == nil {
		return nil
	}
	_AlarmMessageAckTypeCopy := &_AlarmMessageAckType{
		m.FunctionId,
		m.NumberOfObjects,
		utils.DeepCopySlice[AlarmMessageObjectAckType, AlarmMessageObjectAckType](m.MessageObjects),
	}
	return _AlarmMessageAckTypeCopy
}

func (m *_AlarmMessageAckType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
