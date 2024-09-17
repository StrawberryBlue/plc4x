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

// BACnetLightingInProgressTagged is the corresponding interface of BACnetLightingInProgressTagged
type BACnetLightingInProgressTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLightingInProgress
	// IsBACnetLightingInProgressTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLightingInProgressTagged()
}

// _BACnetLightingInProgressTagged is the data-structure of this message
type _BACnetLightingInProgressTagged struct {
	Header BACnetTagHeader
	Value  BACnetLightingInProgress

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLightingInProgressTagged = (*_BACnetLightingInProgressTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingInProgressTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLightingInProgressTagged) GetValue() BACnetLightingInProgress {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLightingInProgressTagged factory function for _BACnetLightingInProgressTagged
func NewBACnetLightingInProgressTagged(header BACnetTagHeader, value BACnetLightingInProgress, tagNumber uint8, tagClass TagClass) *_BACnetLightingInProgressTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLightingInProgressTagged must not be nil")
	}
	return &_BACnetLightingInProgressTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLightingInProgressTagged(structType any) BACnetLightingInProgressTagged {
	if casted, ok := structType.(BACnetLightingInProgressTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingInProgressTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingInProgressTagged) GetTypeName() string {
	return "BACnetLightingInProgressTagged"
}

func (m *_BACnetLightingInProgressTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetLightingInProgressTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingInProgressTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLightingInProgressTagged, error) {
	return BACnetLightingInProgressTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLightingInProgressTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingInProgressTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingInProgressTagged, error) {
		return BACnetLightingInProgressTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLightingInProgressTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLightingInProgressTagged, error) {
	v, err := (&_BACnetLightingInProgressTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLightingInProgressTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLightingInProgressTagged BACnetLightingInProgressTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingInProgressTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingInProgressTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetLightingInProgress](ctx, "value", readBuffer, EnsureType[BACnetLightingInProgress](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetLightingInProgress_IDLE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetLightingInProgressTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingInProgressTagged")
	}

	return m, nil
}

func (m *_BACnetLightingInProgressTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingInProgressTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLightingInProgressTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingInProgressTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLightingInProgress](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingInProgressTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingInProgressTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLightingInProgressTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLightingInProgressTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLightingInProgressTagged) IsBACnetLightingInProgressTagged() {}

func (m *_BACnetLightingInProgressTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
