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

// BACnetPropertyReference is the corresponding interface of BACnetPropertyReference
type BACnetPropertyReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// IsBACnetPropertyReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyReference()
}

// _BACnetPropertyReference is the data-structure of this message
type _BACnetPropertyReference struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
}

var _ BACnetPropertyReference = (*_BACnetPropertyReference)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyReference) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetPropertyReference) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyReference factory function for _BACnetPropertyReference
func NewBACnetPropertyReference(propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger) *_BACnetPropertyReference {
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetPropertyReference must not be nil")
	}
	return &_BACnetPropertyReference{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyReference(structType any) BACnetPropertyReference {
	if casted, ok := structType.(BACnetPropertyReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyReference) GetTypeName() string {
	return "BACnetPropertyReference"
}

func (m *_BACnetPropertyReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetPropertyReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyReferenceParse(ctx context.Context, theBytes []byte) (BACnetPropertyReference, error) {
	return BACnetPropertyReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPropertyReferenceParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReference, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReference, error) {
		return BACnetPropertyReferenceParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetPropertyReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReference, error) {
	v, err := (&_BACnetPropertyReference{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPropertyReference) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetPropertyReference BACnetPropertyReference, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var arrayIndex BACnetContextTagUnsignedInteger
	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
		m.ArrayIndex = arrayIndex
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyReference")
	}

	return m, nil
}

func (m *_BACnetPropertyReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyReference")
	}

	if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayIndex' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyReference")
	}
	return nil
}

func (m *_BACnetPropertyReference) IsBACnetPropertyReference() {}

func (m *_BACnetPropertyReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
