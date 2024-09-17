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

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetTagPayloadCharacterString is the corresponding interface of BACnetTagPayloadCharacterString
type BACnetTagPayloadCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetEncoding returns Encoding (property field)
	GetEncoding() BACnetCharacterEncoding
	// GetValue returns Value (property field)
	GetValue() string
	// GetActualLengthInBit returns ActualLengthInBit (virtual field)
	GetActualLengthInBit() uint16
	// IsBACnetTagPayloadCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadCharacterString()
}

// _BACnetTagPayloadCharacterString is the data-structure of this message
type _BACnetTagPayloadCharacterString struct {
	Encoding BACnetCharacterEncoding
	Value    string

	// Arguments.
	ActualLength uint32
}

var _ BACnetTagPayloadCharacterString = (*_BACnetTagPayloadCharacterString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadCharacterString) GetEncoding() BACnetCharacterEncoding {
	return m.Encoding
}

func (m *_BACnetTagPayloadCharacterString) GetValue() string {
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

func (m *_BACnetTagPayloadCharacterString) GetActualLengthInBit() uint16 {
	ctx := context.Background()
	_ = ctx
	return uint16(uint16(uint16(m.GetActualLength())*uint16(uint16(8))) - uint16(uint16(8)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadCharacterString factory function for _BACnetTagPayloadCharacterString
func NewBACnetTagPayloadCharacterString(encoding BACnetCharacterEncoding, value string, actualLength uint32) *_BACnetTagPayloadCharacterString {
	return &_BACnetTagPayloadCharacterString{Encoding: encoding, Value: value, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadCharacterString(structType any) BACnetTagPayloadCharacterString {
	if casted, ok := structType.(BACnetTagPayloadCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadCharacterString) GetTypeName() string {
	return "BACnetTagPayloadCharacterString"
}

func (m *_BACnetTagPayloadCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (encoding)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.GetActualLengthInBit())

	return lengthInBits
}

func (m *_BACnetTagPayloadCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadCharacterStringParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadCharacterString, error) {
	return BACnetTagPayloadCharacterStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadCharacterStringParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadCharacterString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadCharacterString, error) {
		return BACnetTagPayloadCharacterStringParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetTagPayloadCharacterStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadCharacterString, error) {
	v, err := (&_BACnetTagPayloadCharacterString{ActualLength: actualLength}).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetTagPayloadCharacterString BACnetTagPayloadCharacterString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	encoding, err := ReadEnumField[BACnetCharacterEncoding](ctx, "encoding", "BACnetCharacterEncoding", ReadEnum(BACnetCharacterEncodingByValue, ReadByte(readBuffer, 8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encoding' field"))
	}
	m.Encoding = encoding

	actualLengthInBit, err := ReadVirtualField[uint16](ctx, "actualLengthInBit", (*uint16)(nil), uint16(uint16(actualLength)*uint16(uint16(8)))-uint16(uint16(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualLengthInBit' field"))
	}
	_ = actualLengthInBit

	value, err := ReadSimpleField(ctx, "value", ReadString(readBuffer, uint32(actualLengthInBit)), codegen.WithEncoding("UTF-8"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadCharacterString")
	}

	return m, nil
}

func (m *_BACnetTagPayloadCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadCharacterString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadCharacterString")
	}

	if err := WriteSimpleEnumField[BACnetCharacterEncoding](ctx, "encoding", "BACnetCharacterEncoding", m.GetEncoding(), WriteEnum[BACnetCharacterEncoding, byte](BACnetCharacterEncoding.GetValue, BACnetCharacterEncoding.PLC4XEnumName, WriteByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'encoding' field")
	}
	// Virtual field
	actualLengthInBit := m.GetActualLengthInBit()
	_ = actualLengthInBit
	if _actualLengthInBitErr := writeBuffer.WriteVirtual(ctx, "actualLengthInBit", m.GetActualLengthInBit()); _actualLengthInBitErr != nil {
		return errors.Wrap(_actualLengthInBitErr, "Error serializing 'actualLengthInBit' field")
	}

	if err := WriteSimpleField[string](ctx, "value", m.GetValue(), WriteString(writeBuffer, int32(m.GetActualLengthInBit())), codegen.WithEncoding("UTF-8")); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadCharacterString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadCharacterString")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadCharacterString) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadCharacterString) IsBACnetTagPayloadCharacterString() {}

func (m *_BACnetTagPayloadCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
