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

// StatusByte is the corresponding interface of StatusByte
type StatusByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetGav3 returns Gav3 (property field)
	GetGav3() GAVState
	// GetGav2 returns Gav2 (property field)
	GetGav2() GAVState
	// GetGav1 returns Gav1 (property field)
	GetGav1() GAVState
	// GetGav0 returns Gav0 (property field)
	GetGav0() GAVState
	// IsStatusByte is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusByte()
}

// _StatusByte is the data-structure of this message
type _StatusByte struct {
	Gav3 GAVState
	Gav2 GAVState
	Gav1 GAVState
	Gav0 GAVState
}

var _ StatusByte = (*_StatusByte)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusByte) GetGav3() GAVState {
	return m.Gav3
}

func (m *_StatusByte) GetGav2() GAVState {
	return m.Gav2
}

func (m *_StatusByte) GetGav1() GAVState {
	return m.Gav1
}

func (m *_StatusByte) GetGav0() GAVState {
	return m.Gav0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusByte factory function for _StatusByte
func NewStatusByte(gav3 GAVState, gav2 GAVState, gav1 GAVState, gav0 GAVState) *_StatusByte {
	return &_StatusByte{Gav3: gav3, Gav2: gav2, Gav1: gav1, Gav0: gav0}
}

// Deprecated: use the interface for direct cast
func CastStatusByte(structType any) StatusByte {
	if casted, ok := structType.(StatusByte); ok {
		return casted
	}
	if casted, ok := structType.(*StatusByte); ok {
		return *casted
	}
	return nil
}

func (m *_StatusByte) GetTypeName() string {
	return "StatusByte"
}

func (m *_StatusByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (gav3)
	lengthInBits += 2

	// Simple field (gav2)
	lengthInBits += 2

	// Simple field (gav1)
	lengthInBits += 2

	// Simple field (gav0)
	lengthInBits += 2

	return lengthInBits
}

func (m *_StatusByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusByteParse(ctx context.Context, theBytes []byte) (StatusByte, error) {
	return StatusByteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StatusByteParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (StatusByte, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (StatusByte, error) {
		return StatusByteParseWithBuffer(ctx, readBuffer)
	}
}

func StatusByteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StatusByte, error) {
	v, err := (&_StatusByte{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_StatusByte) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__statusByte StatusByte, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	gav3, err := ReadEnumField[GAVState](ctx, "gav3", "GAVState", ReadEnum(GAVStateByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gav3' field"))
	}
	m.Gav3 = gav3

	gav2, err := ReadEnumField[GAVState](ctx, "gav2", "GAVState", ReadEnum(GAVStateByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gav2' field"))
	}
	m.Gav2 = gav2

	gav1, err := ReadEnumField[GAVState](ctx, "gav1", "GAVState", ReadEnum(GAVStateByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gav1' field"))
	}
	m.Gav1 = gav1

	gav0, err := ReadEnumField[GAVState](ctx, "gav0", "GAVState", ReadEnum(GAVStateByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gav0' field"))
	}
	m.Gav0 = gav0

	if closeErr := readBuffer.CloseContext("StatusByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusByte")
	}

	return m, nil
}

func (m *_StatusByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("StatusByte"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StatusByte")
	}

	if err := WriteSimpleEnumField[GAVState](ctx, "gav3", "GAVState", m.GetGav3(), WriteEnum[GAVState, uint8](GAVState.GetValue, GAVState.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'gav3' field")
	}

	if err := WriteSimpleEnumField[GAVState](ctx, "gav2", "GAVState", m.GetGav2(), WriteEnum[GAVState, uint8](GAVState.GetValue, GAVState.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'gav2' field")
	}

	if err := WriteSimpleEnumField[GAVState](ctx, "gav1", "GAVState", m.GetGav1(), WriteEnum[GAVState, uint8](GAVState.GetValue, GAVState.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'gav1' field")
	}

	if err := WriteSimpleEnumField[GAVState](ctx, "gav0", "GAVState", m.GetGav0(), WriteEnum[GAVState, uint8](GAVState.GetValue, GAVState.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'gav0' field")
	}

	if popErr := writeBuffer.PopContext("StatusByte"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StatusByte")
	}
	return nil
}

func (m *_StatusByte) IsStatusByte() {}

func (m *_StatusByte) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
