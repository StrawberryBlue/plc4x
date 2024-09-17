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

// BACnetPriorityValue is the corresponding interface of BACnetPriorityValue
type BACnetPriorityValue interface {
	BACnetPriorityValueContract
	BACnetPriorityValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetPriorityValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValue()
}

// BACnetPriorityValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetPriorityValueContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetObjectTypeArgument() returns a parser argument
	GetObjectTypeArgument() BACnetObjectType
	// IsBACnetPriorityValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValue()
}

// BACnetPriorityValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetPriorityValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedIsContextTag returns PeekedIsContextTag (discriminator field)
	GetPeekedIsContextTag() bool
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetPriorityValue is the data-structure of this message
type _BACnetPriorityValue struct {
	_SubType        BACnetPriorityValue
	PeekedTagHeader BACnetTagHeader

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

var _ BACnetPriorityValueContract = (*_BACnetPriorityValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValue) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetPriorityValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (pm *_BACnetPriorityValue) GetPeekedIsContextTag() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValue factory function for _BACnetPriorityValue
func NewBACnetPriorityValue(peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValue {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetPriorityValue must not be nil")
	}
	return &_BACnetPriorityValue{PeekedTagHeader: peekedTagHeader, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValue(structType any) BACnetPriorityValue {
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValue) GetTypeName() string {
	return "BACnetPriorityValue"
}

func (m *_BACnetPriorityValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPriorityValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetPriorityValueParse[T BACnetPriorityValue](ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (T, error) {
	return BACnetPriorityValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPriorityValueParseWithBufferProducer[T BACnetPriorityValue](objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetPriorityValueParseWithBuffer[T](ctx, readBuffer, objectTypeArgument)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetPriorityValueParseWithBuffer[T BACnetPriorityValue](ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (T, error) {
	v, err := (&_BACnetPriorityValue{ObjectTypeArgument: objectTypeArgument}).parse(ctx, readBuffer, objectTypeArgument)
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

func (m *_BACnetPriorityValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (__bACnetPriorityValue BACnetPriorityValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValue")
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

	peekedIsContextTag, err := ReadVirtualField[bool](ctx, "peekedIsContextTag", (*bool)(nil), bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedIsContextTag' field"))
	}
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetPriorityValue
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetPriorityValueNull
		if _child, err = (&_BACnetPriorityValueNull{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueNull for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetPriorityValueReal
		if _child, err = (&_BACnetPriorityValueReal{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueReal for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetPriorityValueEnumerated
		if _child, err = (&_BACnetPriorityValueEnumerated{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueEnumerated for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetPriorityValueUnsigned
		if _child, err = (&_BACnetPriorityValueUnsigned{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueUnsigned for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetPriorityValueBoolean
		if _child, err = (&_BACnetPriorityValueBoolean{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueBoolean for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetPriorityValueInteger
		if _child, err = (&_BACnetPriorityValueInteger{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueInteger for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetPriorityValueDouble
		if _child, err = (&_BACnetPriorityValueDouble{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueDouble for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetPriorityValueTime
		if _child, err = (&_BACnetPriorityValueTime{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueTime for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetPriorityValueCharacterString
		if _child, err = (&_BACnetPriorityValueCharacterString{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueCharacterString for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetPriorityValueOctetString
		if _child, err = (&_BACnetPriorityValueOctetString{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueOctetString for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetPriorityValueBitString
		if _child, err = (&_BACnetPriorityValueBitString{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueBitString for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetPriorityValueDate
		if _child, err = (&_BACnetPriorityValueDate{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueDate for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetPriorityValueObjectidentifier
		if _child, err = (&_BACnetPriorityValueObjectidentifier{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueObjectidentifier for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetPriorityValueConstructedValue
		if _child, err = (&_BACnetPriorityValueConstructedValue{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueConstructedValue for type-switch of BACnetPriorityValue")
		}
	case peekedTagNumber == uint8(1) && peekedIsContextTag == bool(true): // BACnetPriorityValueDateTime
		if _child, err = (&_BACnetPriorityValueDateTime{}).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPriorityValueDateTime for type-switch of BACnetPriorityValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValue")
	}

	return _child, nil
}

func (pm *_BACnetPriorityValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPriorityValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPriorityValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValue")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	peekedIsContextTag := m.GetPeekedIsContextTag()
	_ = peekedIsContextTag
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPriorityValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPriorityValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPriorityValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetPriorityValue) IsBACnetPriorityValue() {}
