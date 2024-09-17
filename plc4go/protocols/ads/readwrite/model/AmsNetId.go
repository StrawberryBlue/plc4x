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

// AmsNetId is the corresponding interface of AmsNetId
type AmsNetId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOctet1 returns Octet1 (property field)
	GetOctet1() uint8
	// GetOctet2 returns Octet2 (property field)
	GetOctet2() uint8
	// GetOctet3 returns Octet3 (property field)
	GetOctet3() uint8
	// GetOctet4 returns Octet4 (property field)
	GetOctet4() uint8
	// GetOctet5 returns Octet5 (property field)
	GetOctet5() uint8
	// GetOctet6 returns Octet6 (property field)
	GetOctet6() uint8
	// IsAmsNetId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAmsNetId()
}

// _AmsNetId is the data-structure of this message
type _AmsNetId struct {
	Octet1 uint8
	Octet2 uint8
	Octet3 uint8
	Octet4 uint8
	Octet5 uint8
	Octet6 uint8
}

var _ AmsNetId = (*_AmsNetId)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsNetId) GetOctet1() uint8 {
	return m.Octet1
}

func (m *_AmsNetId) GetOctet2() uint8 {
	return m.Octet2
}

func (m *_AmsNetId) GetOctet3() uint8 {
	return m.Octet3
}

func (m *_AmsNetId) GetOctet4() uint8 {
	return m.Octet4
}

func (m *_AmsNetId) GetOctet5() uint8 {
	return m.Octet5
}

func (m *_AmsNetId) GetOctet6() uint8 {
	return m.Octet6
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAmsNetId factory function for _AmsNetId
func NewAmsNetId(octet1 uint8, octet2 uint8, octet3 uint8, octet4 uint8, octet5 uint8, octet6 uint8) *_AmsNetId {
	return &_AmsNetId{Octet1: octet1, Octet2: octet2, Octet3: octet3, Octet4: octet4, Octet5: octet5, Octet6: octet6}
}

// Deprecated: use the interface for direct cast
func CastAmsNetId(structType any) AmsNetId {
	if casted, ok := structType.(AmsNetId); ok {
		return casted
	}
	if casted, ok := structType.(*AmsNetId); ok {
		return *casted
	}
	return nil
}

func (m *_AmsNetId) GetTypeName() string {
	return "AmsNetId"
}

func (m *_AmsNetId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (octet1)
	lengthInBits += 8

	// Simple field (octet2)
	lengthInBits += 8

	// Simple field (octet3)
	lengthInBits += 8

	// Simple field (octet4)
	lengthInBits += 8

	// Simple field (octet5)
	lengthInBits += 8

	// Simple field (octet6)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AmsNetId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsNetIdParse(ctx context.Context, theBytes []byte) (AmsNetId, error) {
	return AmsNetIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AmsNetIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsNetId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsNetId, error) {
		return AmsNetIdParseWithBuffer(ctx, readBuffer)
	}
}

func AmsNetIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsNetId, error) {
	v, err := (&_AmsNetId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AmsNetId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__amsNetId AmsNetId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsNetId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octet1, err := ReadSimpleField(ctx, "octet1", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet1' field"))
	}
	m.Octet1 = octet1

	octet2, err := ReadSimpleField(ctx, "octet2", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet2' field"))
	}
	m.Octet2 = octet2

	octet3, err := ReadSimpleField(ctx, "octet3", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet3' field"))
	}
	m.Octet3 = octet3

	octet4, err := ReadSimpleField(ctx, "octet4", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet4' field"))
	}
	m.Octet4 = octet4

	octet5, err := ReadSimpleField(ctx, "octet5", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet5' field"))
	}
	m.Octet5 = octet5

	octet6, err := ReadSimpleField(ctx, "octet6", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet6' field"))
	}
	m.Octet6 = octet6

	if closeErr := readBuffer.CloseContext("AmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsNetId")
	}

	return m, nil
}

func (m *_AmsNetId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsNetId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AmsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsNetId")
	}

	if err := WriteSimpleField[uint8](ctx, "octet1", m.GetOctet1(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet1' field")
	}

	if err := WriteSimpleField[uint8](ctx, "octet2", m.GetOctet2(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet2' field")
	}

	if err := WriteSimpleField[uint8](ctx, "octet3", m.GetOctet3(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet3' field")
	}

	if err := WriteSimpleField[uint8](ctx, "octet4", m.GetOctet4(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet4' field")
	}

	if err := WriteSimpleField[uint8](ctx, "octet5", m.GetOctet5(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet5' field")
	}

	if err := WriteSimpleField[uint8](ctx, "octet6", m.GetOctet6(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet6' field")
	}

	if popErr := writeBuffer.PopContext("AmsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsNetId")
	}
	return nil
}

func (m *_AmsNetId) IsAmsNetId() {}

func (m *_AmsNetId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
