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

// AssociatedValueType is the corresponding interface of AssociatedValueType
type AssociatedValueType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetValueLength returns ValueLength (property field)
	GetValueLength() uint16
	// GetData returns Data (property field)
	GetData() []uint8
	// IsAssociatedValueType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAssociatedValueType()
}

// _AssociatedValueType is the data-structure of this message
type _AssociatedValueType struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	ValueLength   uint16
	Data          []uint8
}

var _ AssociatedValueType = (*_AssociatedValueType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AssociatedValueType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_AssociatedValueType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_AssociatedValueType) GetValueLength() uint16 {
	return m.ValueLength
}

func (m *_AssociatedValueType) GetData() []uint8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAssociatedValueType factory function for _AssociatedValueType
func NewAssociatedValueType(returnCode DataTransportErrorCode, transportSize DataTransportSize, valueLength uint16, data []uint8) *_AssociatedValueType {
	return &_AssociatedValueType{ReturnCode: returnCode, TransportSize: transportSize, ValueLength: valueLength, Data: data}
}

// Deprecated: use the interface for direct cast
func CastAssociatedValueType(structType any) AssociatedValueType {
	if casted, ok := structType.(AssociatedValueType); ok {
		return casted
	}
	if casted, ok := structType.(*AssociatedValueType); ok {
		return *casted
	}
	return nil
}

func (m *_AssociatedValueType) GetTypeName() string {
	return "AssociatedValueType"
}

func (m *_AssociatedValueType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Manual Field (valueLength)
	lengthInBits += uint16(int32(2))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AssociatedValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AssociatedValueTypeParse(ctx context.Context, theBytes []byte) (AssociatedValueType, error) {
	return AssociatedValueTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AssociatedValueTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedValueType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedValueType, error) {
		return AssociatedValueTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AssociatedValueTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedValueType, error) {
	v, err := (&_AssociatedValueType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AssociatedValueType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__associatedValueType AssociatedValueType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AssociatedValueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AssociatedValueType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	returnCode, err := ReadEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", ReadEnum(DataTransportErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnCode' field"))
	}
	m.ReturnCode = returnCode

	transportSize, err := ReadEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", ReadEnum(DataTransportSizeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}
	m.TransportSize = transportSize

	valueLength, err := ReadManualField[uint16](ctx, "valueLength", readBuffer, EnsureType[uint16](RightShift3(ctx, readBuffer, transportSize)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueLength' field"))
	}
	m.ValueLength = valueLength

	data, err := ReadCountArrayField[uint8](ctx, "data", ReadUnsignedByte(readBuffer, uint8(8)), uint64(EventItemLength(ctx, readBuffer, valueLength)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AssociatedValueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AssociatedValueType")
	}

	return m, nil
}

func (m *_AssociatedValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AssociatedValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AssociatedValueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AssociatedValueType")
	}

	if err := WriteSimpleEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", m.GetReturnCode(), WriteEnum[DataTransportErrorCode, uint8](DataTransportErrorCode.GetValue, DataTransportErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'returnCode' field")
	}

	if err := WriteSimpleEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", m.GetTransportSize(), WriteEnum[DataTransportSize, uint8](DataTransportSize.GetValue, DataTransportSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'transportSize' field")
	}

	if err := WriteManualField[uint16](ctx, "valueLength", func(ctx context.Context) error { return LeftShift3(ctx, writeBuffer, m.GetValueLength()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'valueLength' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("AssociatedValueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AssociatedValueType")
	}
	return nil
}

func (m *_AssociatedValueType) IsAssociatedValueType() {}

func (m *_AssociatedValueType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
