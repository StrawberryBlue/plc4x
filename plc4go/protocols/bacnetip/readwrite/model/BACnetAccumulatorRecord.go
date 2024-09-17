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

// BACnetAccumulatorRecord is the corresponding interface of BACnetAccumulatorRecord
type BACnetAccumulatorRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetContextTagSignedInteger
	// GetAccumulatedValue returns AccumulatedValue (property field)
	GetAccumulatedValue() BACnetContextTagSignedInteger
	// GetAccumulatorStatus returns AccumulatorStatus (property field)
	GetAccumulatorStatus() BACnetAccumulatorRecordAccumulatorStatusTagged
	// IsBACnetAccumulatorRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccumulatorRecord()
}

// _BACnetAccumulatorRecord is the data-structure of this message
type _BACnetAccumulatorRecord struct {
	Timestamp         BACnetDateTimeEnclosed
	PresentValue      BACnetContextTagSignedInteger
	AccumulatedValue  BACnetContextTagSignedInteger
	AccumulatorStatus BACnetAccumulatorRecordAccumulatorStatusTagged
}

var _ BACnetAccumulatorRecord = (*_BACnetAccumulatorRecord)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccumulatorRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetAccumulatorRecord) GetPresentValue() BACnetContextTagSignedInteger {
	return m.PresentValue
}

func (m *_BACnetAccumulatorRecord) GetAccumulatedValue() BACnetContextTagSignedInteger {
	return m.AccumulatedValue
}

func (m *_BACnetAccumulatorRecord) GetAccumulatorStatus() BACnetAccumulatorRecordAccumulatorStatusTagged {
	return m.AccumulatorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccumulatorRecord factory function for _BACnetAccumulatorRecord
func NewBACnetAccumulatorRecord(timestamp BACnetDateTimeEnclosed, presentValue BACnetContextTagSignedInteger, accumulatedValue BACnetContextTagSignedInteger, accumulatorStatus BACnetAccumulatorRecordAccumulatorStatusTagged) *_BACnetAccumulatorRecord {
	if timestamp == nil {
		panic("timestamp of type BACnetDateTimeEnclosed for BACnetAccumulatorRecord must not be nil")
	}
	if presentValue == nil {
		panic("presentValue of type BACnetContextTagSignedInteger for BACnetAccumulatorRecord must not be nil")
	}
	if accumulatedValue == nil {
		panic("accumulatedValue of type BACnetContextTagSignedInteger for BACnetAccumulatorRecord must not be nil")
	}
	if accumulatorStatus == nil {
		panic("accumulatorStatus of type BACnetAccumulatorRecordAccumulatorStatusTagged for BACnetAccumulatorRecord must not be nil")
	}
	return &_BACnetAccumulatorRecord{Timestamp: timestamp, PresentValue: presentValue, AccumulatedValue: accumulatedValue, AccumulatorStatus: accumulatorStatus}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccumulatorRecord(structType any) BACnetAccumulatorRecord {
	if casted, ok := structType.(BACnetAccumulatorRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccumulatorRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccumulatorRecord) GetTypeName() string {
	return "BACnetAccumulatorRecord"
}

func (m *_BACnetAccumulatorRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// Simple field (accumulatedValue)
	lengthInBits += m.AccumulatedValue.GetLengthInBits(ctx)

	// Simple field (accumulatorStatus)
	lengthInBits += m.AccumulatorStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAccumulatorRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccumulatorRecordParse(ctx context.Context, theBytes []byte) (BACnetAccumulatorRecord, error) {
	return BACnetAccumulatorRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccumulatorRecordParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
		return BACnetAccumulatorRecordParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAccumulatorRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecord, error) {
	v, err := (&_BACnetAccumulatorRecord{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccumulatorRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAccumulatorRecord BACnetAccumulatorRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccumulatorRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccumulatorRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	presentValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "presentValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	accumulatedValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "accumulatedValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accumulatedValue' field"))
	}
	m.AccumulatedValue = accumulatedValue

	accumulatorStatus, err := ReadSimpleField[BACnetAccumulatorRecordAccumulatorStatusTagged](ctx, "accumulatorStatus", ReadComplex[BACnetAccumulatorRecordAccumulatorStatusTagged](BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accumulatorStatus' field"))
	}
	m.AccumulatorStatus = accumulatorStatus

	if closeErr := readBuffer.CloseContext("BACnetAccumulatorRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccumulatorRecord")
	}

	return m, nil
}

func (m *_BACnetAccumulatorRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccumulatorRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccumulatorRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccumulatorRecord")
	}

	if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'presentValue' field")
	}

	if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "accumulatedValue", m.GetAccumulatedValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'accumulatedValue' field")
	}

	if err := WriteSimpleField[BACnetAccumulatorRecordAccumulatorStatusTagged](ctx, "accumulatorStatus", m.GetAccumulatorStatus(), WriteComplex[BACnetAccumulatorRecordAccumulatorStatusTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'accumulatorStatus' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccumulatorRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccumulatorRecord")
	}
	return nil
}

func (m *_BACnetAccumulatorRecord) IsBACnetAccumulatorRecord() {}

func (m *_BACnetAccumulatorRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
