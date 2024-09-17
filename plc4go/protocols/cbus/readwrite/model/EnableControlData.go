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

// EnableControlData is the corresponding interface of EnableControlData
type EnableControlData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() EnableControlCommandTypeContainer
	// GetEnableNetworkVariable returns EnableNetworkVariable (property field)
	GetEnableNetworkVariable() byte
	// GetValue returns Value (property field)
	GetValue() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() EnableControlCommandType
	// IsEnableControlData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEnableControlData()
}

// _EnableControlData is the data-structure of this message
type _EnableControlData struct {
	CommandTypeContainer  EnableControlCommandTypeContainer
	EnableNetworkVariable byte
	Value                 byte
}

var _ EnableControlData = (*_EnableControlData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EnableControlData) GetCommandTypeContainer() EnableControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_EnableControlData) GetEnableNetworkVariable() byte {
	return m.EnableNetworkVariable
}

func (m *_EnableControlData) GetValue() byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_EnableControlData) GetCommandType() EnableControlCommandType {
	ctx := context.Background()
	_ = ctx
	return CastEnableControlCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEnableControlData factory function for _EnableControlData
func NewEnableControlData(commandTypeContainer EnableControlCommandTypeContainer, enableNetworkVariable byte, value byte) *_EnableControlData {
	return &_EnableControlData{CommandTypeContainer: commandTypeContainer, EnableNetworkVariable: enableNetworkVariable, Value: value}
}

// Deprecated: use the interface for direct cast
func CastEnableControlData(structType any) EnableControlData {
	if casted, ok := structType.(EnableControlData); ok {
		return casted
	}
	if casted, ok := structType.(*EnableControlData); ok {
		return *casted
	}
	return nil
}

func (m *_EnableControlData) GetTypeName() string {
	return "EnableControlData"
}

func (m *_EnableControlData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (enableNetworkVariable)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *_EnableControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EnableControlDataParse(ctx context.Context, theBytes []byte) (EnableControlData, error) {
	return EnableControlDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func EnableControlDataParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (EnableControlData, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EnableControlData, error) {
		return EnableControlDataParseWithBuffer(ctx, readBuffer)
	}
}

func EnableControlDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EnableControlData, error) {
	v, err := (&_EnableControlData{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_EnableControlData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__enableControlData EnableControlData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EnableControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EnableControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsEnableControlCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[EnableControlCommandTypeContainer](ctx, "commandTypeContainer", "EnableControlCommandTypeContainer", ReadEnum(EnableControlCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[EnableControlCommandType](ctx, "commandType", (*EnableControlCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	enableNetworkVariable, err := ReadSimpleField(ctx, "enableNetworkVariable", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enableNetworkVariable' field"))
	}
	m.EnableNetworkVariable = enableNetworkVariable

	value, err := ReadSimpleField(ctx, "value", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("EnableControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EnableControlData")
	}

	return m, nil
}

func (m *_EnableControlData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EnableControlData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("EnableControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EnableControlData")
	}

	if err := WriteSimpleEnumField[EnableControlCommandTypeContainer](ctx, "commandTypeContainer", "EnableControlCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[EnableControlCommandTypeContainer, uint8](EnableControlCommandTypeContainer.GetValue, EnableControlCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "enableNetworkVariable", m.GetEnableNetworkVariable(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'enableNetworkVariable' field")
	}

	if err := WriteSimpleField[byte](ctx, "value", m.GetValue(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("EnableControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EnableControlData")
	}
	return nil
}

func (m *_EnableControlData) IsEnableControlData() {}

func (m *_EnableControlData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
