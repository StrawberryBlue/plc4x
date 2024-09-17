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

// BACnetConfirmedServiceRequestReadRangeRange is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRange
type BACnetConfirmedServiceRequestReadRangeRange interface {
	BACnetConfirmedServiceRequestReadRangeRangeContract
	BACnetConfirmedServiceRequestReadRangeRangeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetConfirmedServiceRequestReadRangeRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadRangeRange()
}

// BACnetConfirmedServiceRequestReadRangeRangeContract provides a set of functions which can be overwritten by a sub struct
type BACnetConfirmedServiceRequestReadRangeRangeContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetConfirmedServiceRequestReadRangeRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadRangeRange()
}

// BACnetConfirmedServiceRequestReadRangeRangeRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetConfirmedServiceRequestReadRangeRangeRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetConfirmedServiceRequestReadRangeRange is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRange struct {
	_SubType        BACnetConfirmedServiceRequestReadRangeRange
	PeekedTagHeader BACnetTagHeader
	OpeningTag      BACnetOpeningTag
	ClosingTag      BACnetClosingTag
}

var _ BACnetConfirmedServiceRequestReadRangeRangeContract = (*_BACnetConfirmedServiceRequestReadRangeRange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetConfirmedServiceRequestReadRangeRange) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRange factory function for _BACnetConfirmedServiceRequestReadRangeRange
func NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestReadRangeRange {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetConfirmedServiceRequestReadRangeRange must not be nil")
	}
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetConfirmedServiceRequestReadRangeRange must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetConfirmedServiceRequestReadRangeRange must not be nil")
	}
	return &_BACnetConfirmedServiceRequestReadRangeRange{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRange(structType any) BACnetConfirmedServiceRequestReadRangeRange {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRange"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeParse[T BACnetConfirmedServiceRequestReadRangeRange](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceRequestReadRangeRangeParseWithBufferProducer[T BACnetConfirmedServiceRequestReadRangeRange]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer[T BACnetConfirmedServiceRequestReadRangeRange](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetConfirmedServiceRequestReadRangeRange{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetConfirmedServiceRequestReadRangeRange BACnetConfirmedServiceRequestReadRangeRange, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetConfirmedServiceRequestReadRangeRange
	switch {
	case peekedTagNumber == 0x3: // BACnetConfirmedServiceRequestReadRangeRangeByPosition
		if _child, err = (&_BACnetConfirmedServiceRequestReadRangeRangeByPosition{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadRangeRangeByPosition for type-switch of BACnetConfirmedServiceRequestReadRangeRange")
		}
	case peekedTagNumber == 0x6: // BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
		if _child, err = (&_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber for type-switch of BACnetConfirmedServiceRequestReadRangeRange")
		}
	case peekedTagNumber == 0x7: // BACnetConfirmedServiceRequestReadRangeRangeByTime
		if _child, err = (&_BACnetConfirmedServiceRequestReadRangeRangeByTime{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadRangeRangeByTime for type-switch of BACnetConfirmedServiceRequestReadRangeRange")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRange")
	}

	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequestReadRangeRange) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequestReadRangeRange, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRange")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRange")
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) IsBACnetConfirmedServiceRequestReadRangeRange() {
}
