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

// BACnetAssignedLandingCalls is the corresponding interface of BACnetAssignedLandingCalls
type BACnetAssignedLandingCalls interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLandingCalls returns LandingCalls (property field)
	GetLandingCalls() BACnetAssignedLandingCallsLandingCallsList
	// IsBACnetAssignedLandingCalls is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAssignedLandingCalls()
}

// _BACnetAssignedLandingCalls is the data-structure of this message
type _BACnetAssignedLandingCalls struct {
	LandingCalls BACnetAssignedLandingCallsLandingCallsList
}

var _ BACnetAssignedLandingCalls = (*_BACnetAssignedLandingCalls)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAssignedLandingCalls) GetLandingCalls() BACnetAssignedLandingCallsLandingCallsList {
	return m.LandingCalls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedLandingCalls factory function for _BACnetAssignedLandingCalls
func NewBACnetAssignedLandingCalls(landingCalls BACnetAssignedLandingCallsLandingCallsList) *_BACnetAssignedLandingCalls {
	if landingCalls == nil {
		panic("landingCalls of type BACnetAssignedLandingCallsLandingCallsList for BACnetAssignedLandingCalls must not be nil")
	}
	return &_BACnetAssignedLandingCalls{LandingCalls: landingCalls}
}

// Deprecated: use the interface for direct cast
func CastBACnetAssignedLandingCalls(structType any) BACnetAssignedLandingCalls {
	if casted, ok := structType.(BACnetAssignedLandingCalls); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAssignedLandingCalls); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAssignedLandingCalls) GetTypeName() string {
	return "BACnetAssignedLandingCalls"
}

func (m *_BACnetAssignedLandingCalls) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (landingCalls)
	lengthInBits += m.LandingCalls.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAssignedLandingCalls) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAssignedLandingCallsParse(ctx context.Context, theBytes []byte) (BACnetAssignedLandingCalls, error) {
	return BACnetAssignedLandingCallsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAssignedLandingCallsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCalls, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCalls, error) {
		return BACnetAssignedLandingCallsParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAssignedLandingCallsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCalls, error) {
	v, err := (&_BACnetAssignedLandingCalls{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAssignedLandingCalls) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAssignedLandingCalls BACnetAssignedLandingCalls, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedLandingCalls"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedLandingCalls")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	landingCalls, err := ReadSimpleField[BACnetAssignedLandingCallsLandingCallsList](ctx, "landingCalls", ReadComplex[BACnetAssignedLandingCallsLandingCallsList](BACnetAssignedLandingCallsLandingCallsListParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'landingCalls' field"))
	}
	m.LandingCalls = landingCalls

	if closeErr := readBuffer.CloseContext("BACnetAssignedLandingCalls"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedLandingCalls")
	}

	return m, nil
}

func (m *_BACnetAssignedLandingCalls) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAssignedLandingCalls) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAssignedLandingCalls"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedLandingCalls")
	}

	if err := WriteSimpleField[BACnetAssignedLandingCallsLandingCallsList](ctx, "landingCalls", m.GetLandingCalls(), WriteComplex[BACnetAssignedLandingCallsLandingCallsList](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'landingCalls' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedLandingCalls"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedLandingCalls")
	}
	return nil
}

func (m *_BACnetAssignedLandingCalls) IsBACnetAssignedLandingCalls() {}

func (m *_BACnetAssignedLandingCalls) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
