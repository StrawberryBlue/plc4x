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

// BACnetCalendarEntry is the corresponding interface of BACnetCalendarEntry
type BACnetCalendarEntry interface {
	BACnetCalendarEntryContract
	BACnetCalendarEntryRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetCalendarEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCalendarEntry()
}

// BACnetCalendarEntryContract provides a set of functions which can be overwritten by a sub struct
type BACnetCalendarEntryContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetCalendarEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCalendarEntry()
}

// BACnetCalendarEntryRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetCalendarEntryRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetCalendarEntry is the data-structure of this message
type _BACnetCalendarEntry struct {
	_SubType        BACnetCalendarEntry
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetCalendarEntryContract = (*_BACnetCalendarEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntry) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetCalendarEntry) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntry factory function for _BACnetCalendarEntry
func NewBACnetCalendarEntry(peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntry {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetCalendarEntry must not be nil")
	}
	return &_BACnetCalendarEntry{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntry(structType any) BACnetCalendarEntry {
	if casted, ok := structType.(BACnetCalendarEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntry) GetTypeName() string {
	return "BACnetCalendarEntry"
}

func (m *_BACnetCalendarEntry) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetCalendarEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetCalendarEntryParse[T BACnetCalendarEntry](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetCalendarEntryParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCalendarEntryParseWithBufferProducer[T BACnetCalendarEntry]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetCalendarEntryParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetCalendarEntryParseWithBuffer[T BACnetCalendarEntry](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetCalendarEntry{}).parse(ctx, readBuffer)
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

func (m *_BACnetCalendarEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetCalendarEntry BACnetCalendarEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Validation
	if !(bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "Validation failed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetCalendarEntry
	switch {
	case peekedTagNumber == uint8(0): // BACnetCalendarEntryDate
		if _child, err = (&_BACnetCalendarEntryDate{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetCalendarEntryDate for type-switch of BACnetCalendarEntry")
		}
	case peekedTagNumber == uint8(1): // BACnetCalendarEntryDateRange
		if _child, err = (&_BACnetCalendarEntryDateRange{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetCalendarEntryDateRange for type-switch of BACnetCalendarEntry")
		}
	case peekedTagNumber == uint8(2): // BACnetCalendarEntryWeekNDay
		if _child, err = (&_BACnetCalendarEntryWeekNDay{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetCalendarEntryWeekNDay for type-switch of BACnetCalendarEntry")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntry")
	}

	return _child, nil
}

func (pm *_BACnetCalendarEntry) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetCalendarEntry, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCalendarEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntry")
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

	if popErr := writeBuffer.PopContext("BACnetCalendarEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCalendarEntry")
	}
	return nil
}

func (m *_BACnetCalendarEntry) IsBACnetCalendarEntry() {}
