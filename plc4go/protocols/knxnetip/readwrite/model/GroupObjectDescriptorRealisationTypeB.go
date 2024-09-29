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

// GroupObjectDescriptorRealisationTypeB is the corresponding interface of GroupObjectDescriptorRealisationTypeB
type GroupObjectDescriptorRealisationTypeB interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetUpdateEnable returns UpdateEnable (property field)
	GetUpdateEnable() bool
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
	// IsGroupObjectDescriptorRealisationTypeB is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGroupObjectDescriptorRealisationTypeB()
	// CreateBuilder creates a GroupObjectDescriptorRealisationTypeBBuilder
	CreateGroupObjectDescriptorRealisationTypeBBuilder() GroupObjectDescriptorRealisationTypeBBuilder
}

// _GroupObjectDescriptorRealisationTypeB is the data-structure of this message
type _GroupObjectDescriptorRealisationTypeB struct {
	UpdateEnable          bool
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

var _ GroupObjectDescriptorRealisationTypeB = (*_GroupObjectDescriptorRealisationTypeB)(nil)

// NewGroupObjectDescriptorRealisationTypeB factory function for _GroupObjectDescriptorRealisationTypeB
func NewGroupObjectDescriptorRealisationTypeB(updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *_GroupObjectDescriptorRealisationTypeB {
	return &_GroupObjectDescriptorRealisationTypeB{UpdateEnable: updateEnable, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GroupObjectDescriptorRealisationTypeBBuilder is a builder for GroupObjectDescriptorRealisationTypeB
type GroupObjectDescriptorRealisationTypeBBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) GroupObjectDescriptorRealisationTypeBBuilder
	// WithUpdateEnable adds UpdateEnable (property field)
	WithUpdateEnable(bool) GroupObjectDescriptorRealisationTypeBBuilder
	// WithTransmitEnable adds TransmitEnable (property field)
	WithTransmitEnable(bool) GroupObjectDescriptorRealisationTypeBBuilder
	// WithSegmentSelectorEnable adds SegmentSelectorEnable (property field)
	WithSegmentSelectorEnable(bool) GroupObjectDescriptorRealisationTypeBBuilder
	// WithWriteEnable adds WriteEnable (property field)
	WithWriteEnable(bool) GroupObjectDescriptorRealisationTypeBBuilder
	// WithReadEnable adds ReadEnable (property field)
	WithReadEnable(bool) GroupObjectDescriptorRealisationTypeBBuilder
	// WithCommunicationEnable adds CommunicationEnable (property field)
	WithCommunicationEnable(bool) GroupObjectDescriptorRealisationTypeBBuilder
	// WithPriority adds Priority (property field)
	WithPriority(CEMIPriority) GroupObjectDescriptorRealisationTypeBBuilder
	// WithValueType adds ValueType (property field)
	WithValueType(ComObjectValueType) GroupObjectDescriptorRealisationTypeBBuilder
	// Build builds the GroupObjectDescriptorRealisationTypeB or returns an error if something is wrong
	Build() (GroupObjectDescriptorRealisationTypeB, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GroupObjectDescriptorRealisationTypeB
}

// NewGroupObjectDescriptorRealisationTypeBBuilder() creates a GroupObjectDescriptorRealisationTypeBBuilder
func NewGroupObjectDescriptorRealisationTypeBBuilder() GroupObjectDescriptorRealisationTypeBBuilder {
	return &_GroupObjectDescriptorRealisationTypeBBuilder{_GroupObjectDescriptorRealisationTypeB: new(_GroupObjectDescriptorRealisationTypeB)}
}

type _GroupObjectDescriptorRealisationTypeBBuilder struct {
	*_GroupObjectDescriptorRealisationTypeB

	err *utils.MultiError
}

var _ (GroupObjectDescriptorRealisationTypeBBuilder) = (*_GroupObjectDescriptorRealisationTypeBBuilder)(nil)

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithMandatoryFields(updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) GroupObjectDescriptorRealisationTypeBBuilder {
	return b.WithUpdateEnable(updateEnable).WithTransmitEnable(transmitEnable).WithSegmentSelectorEnable(segmentSelectorEnable).WithWriteEnable(writeEnable).WithReadEnable(readEnable).WithCommunicationEnable(communicationEnable).WithPriority(priority).WithValueType(valueType)
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithUpdateEnable(updateEnable bool) GroupObjectDescriptorRealisationTypeBBuilder {
	b.UpdateEnable = updateEnable
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithTransmitEnable(transmitEnable bool) GroupObjectDescriptorRealisationTypeBBuilder {
	b.TransmitEnable = transmitEnable
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithSegmentSelectorEnable(segmentSelectorEnable bool) GroupObjectDescriptorRealisationTypeBBuilder {
	b.SegmentSelectorEnable = segmentSelectorEnable
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithWriteEnable(writeEnable bool) GroupObjectDescriptorRealisationTypeBBuilder {
	b.WriteEnable = writeEnable
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithReadEnable(readEnable bool) GroupObjectDescriptorRealisationTypeBBuilder {
	b.ReadEnable = readEnable
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithCommunicationEnable(communicationEnable bool) GroupObjectDescriptorRealisationTypeBBuilder {
	b.CommunicationEnable = communicationEnable
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithPriority(priority CEMIPriority) GroupObjectDescriptorRealisationTypeBBuilder {
	b.Priority = priority
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) WithValueType(valueType ComObjectValueType) GroupObjectDescriptorRealisationTypeBBuilder {
	b.ValueType = valueType
	return b
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) Build() (GroupObjectDescriptorRealisationTypeB, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._GroupObjectDescriptorRealisationTypeB.deepCopy(), nil
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) MustBuild() GroupObjectDescriptorRealisationTypeB {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_GroupObjectDescriptorRealisationTypeBBuilder) DeepCopy() any {
	_copy := b.CreateGroupObjectDescriptorRealisationTypeBBuilder().(*_GroupObjectDescriptorRealisationTypeBBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGroupObjectDescriptorRealisationTypeBBuilder creates a GroupObjectDescriptorRealisationTypeBBuilder
func (b *_GroupObjectDescriptorRealisationTypeB) CreateGroupObjectDescriptorRealisationTypeBBuilder() GroupObjectDescriptorRealisationTypeBBuilder {
	if b == nil {
		return NewGroupObjectDescriptorRealisationTypeBBuilder()
	}
	return &_GroupObjectDescriptorRealisationTypeBBuilder{_GroupObjectDescriptorRealisationTypeB: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GroupObjectDescriptorRealisationTypeB) GetUpdateEnable() bool {
	return m.UpdateEnable
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationTypeB(structType any) GroupObjectDescriptorRealisationTypeB {
	if casted, ok := structType.(GroupObjectDescriptorRealisationTypeB); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationTypeB); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetTypeName() string {
	return "GroupObjectDescriptorRealisationTypeB"
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (updateEnable)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationTypeB) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GroupObjectDescriptorRealisationTypeBParse(ctx context.Context, theBytes []byte) (GroupObjectDescriptorRealisationTypeB, error) {
	return GroupObjectDescriptorRealisationTypeBParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GroupObjectDescriptorRealisationTypeBParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationTypeB, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationTypeB, error) {
		return GroupObjectDescriptorRealisationTypeBParseWithBuffer(ctx, readBuffer)
	}
}

func GroupObjectDescriptorRealisationTypeBParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationTypeB, error) {
	v, err := (&_GroupObjectDescriptorRealisationTypeB{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_GroupObjectDescriptorRealisationTypeB) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__groupObjectDescriptorRealisationTypeB GroupObjectDescriptorRealisationTypeB, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationTypeB"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationTypeB")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	updateEnable, err := ReadSimpleField(ctx, "updateEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateEnable' field"))
	}
	m.UpdateEnable = updateEnable

	transmitEnable, err := ReadSimpleField(ctx, "transmitEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transmitEnable' field"))
	}
	m.TransmitEnable = transmitEnable

	segmentSelectorEnable, err := ReadSimpleField(ctx, "segmentSelectorEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentSelectorEnable' field"))
	}
	m.SegmentSelectorEnable = segmentSelectorEnable

	writeEnable, err := ReadSimpleField(ctx, "writeEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeEnable' field"))
	}
	m.WriteEnable = writeEnable

	readEnable, err := ReadSimpleField(ctx, "readEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readEnable' field"))
	}
	m.ReadEnable = readEnable

	communicationEnable, err := ReadSimpleField(ctx, "communicationEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationEnable' field"))
	}
	m.CommunicationEnable = communicationEnable

	priority, err := ReadEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", ReadEnum(CEMIPriorityByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	valueType, err := ReadEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", ReadEnum(ComObjectValueTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueType' field"))
	}
	m.ValueType = valueType

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationTypeB"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationTypeB")
	}

	return m, nil
}

func (m *_GroupObjectDescriptorRealisationTypeB) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GroupObjectDescriptorRealisationTypeB) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationTypeB"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationTypeB")
	}

	if err := WriteSimpleField[bool](ctx, "updateEnable", m.GetUpdateEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'updateEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "transmitEnable", m.GetTransmitEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'transmitEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "segmentSelectorEnable", m.GetSegmentSelectorEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'segmentSelectorEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "writeEnable", m.GetWriteEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'writeEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "readEnable", m.GetReadEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'readEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "communicationEnable", m.GetCommunicationEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationEnable' field")
	}

	if err := WriteSimpleEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", m.GetPriority(), WriteEnum[CEMIPriority, uint8](CEMIPriority.GetValue, CEMIPriority.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if err := WriteSimpleEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", m.GetValueType(), WriteEnum[ComObjectValueType, uint8](ComObjectValueType.GetValue, ComObjectValueType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationTypeB"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationTypeB")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationTypeB) IsGroupObjectDescriptorRealisationTypeB() {}

func (m *_GroupObjectDescriptorRealisationTypeB) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GroupObjectDescriptorRealisationTypeB) deepCopy() *_GroupObjectDescriptorRealisationTypeB {
	if m == nil {
		return nil
	}
	_GroupObjectDescriptorRealisationTypeBCopy := &_GroupObjectDescriptorRealisationTypeB{
		m.UpdateEnable,
		m.TransmitEnable,
		m.SegmentSelectorEnable,
		m.WriteEnable,
		m.ReadEnable,
		m.CommunicationEnable,
		m.Priority,
		m.ValueType,
	}
	return _GroupObjectDescriptorRealisationTypeBCopy
}

func (m *_GroupObjectDescriptorRealisationTypeB) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
