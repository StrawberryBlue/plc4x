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

// BACnetTagPayloadReal is the corresponding interface of BACnetTagPayloadReal
type BACnetTagPayloadReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() float32
	// IsBACnetTagPayloadReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadReal()
}

// _BACnetTagPayloadReal is the data-structure of this message
type _BACnetTagPayloadReal struct {
	Value float32
}

var _ BACnetTagPayloadReal = (*_BACnetTagPayloadReal)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadReal) GetValue() float32 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadReal factory function for _BACnetTagPayloadReal
func NewBACnetTagPayloadReal(value float32) *_BACnetTagPayloadReal {
	return &_BACnetTagPayloadReal{Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadReal(structType any) BACnetTagPayloadReal {
	if casted, ok := structType.(BACnetTagPayloadReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadReal) GetTypeName() string {
	return "BACnetTagPayloadReal"
}

func (m *_BACnetTagPayloadReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m *_BACnetTagPayloadReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadRealParse(ctx context.Context, theBytes []byte) (BACnetTagPayloadReal, error) {
	return BACnetTagPayloadRealParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadRealParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadReal, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadReal, error) {
		return BACnetTagPayloadRealParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTagPayloadRealParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadReal, error) {
	v, err := (&_BACnetTagPayloadReal{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadReal) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTagPayloadReal BACnetTagPayloadReal, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField(ctx, "value", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadReal")
	}

	return m, nil
}

func (m *_BACnetTagPayloadReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadReal"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadReal")
	}

	if err := WriteSimpleField[float32](ctx, "value", m.GetValue(), WriteFloat(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadReal"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadReal")
	}
	return nil
}

func (m *_BACnetTagPayloadReal) IsBACnetTagPayloadReal() {}

func (m *_BACnetTagPayloadReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
