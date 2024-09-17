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

// BACnetEventSummariesList is the corresponding interface of BACnetEventSummariesList
type BACnetEventSummariesList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfEventSummaries returns ListOfEventSummaries (property field)
	GetListOfEventSummaries() []BACnetEventSummary
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventSummariesList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventSummariesList()
}

// _BACnetEventSummariesList is the data-structure of this message
type _BACnetEventSummariesList struct {
	OpeningTag           BACnetOpeningTag
	ListOfEventSummaries []BACnetEventSummary
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventSummariesList = (*_BACnetEventSummariesList)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventSummariesList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventSummariesList) GetListOfEventSummaries() []BACnetEventSummary {
	return m.ListOfEventSummaries
}

func (m *_BACnetEventSummariesList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventSummariesList factory function for _BACnetEventSummariesList
func NewBACnetEventSummariesList(openingTag BACnetOpeningTag, listOfEventSummaries []BACnetEventSummary, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventSummariesList {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventSummariesList must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventSummariesList must not be nil")
	}
	return &_BACnetEventSummariesList{OpeningTag: openingTag, ListOfEventSummaries: listOfEventSummaries, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventSummariesList(structType any) BACnetEventSummariesList {
	if casted, ok := structType.(BACnetEventSummariesList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventSummariesList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventSummariesList) GetTypeName() string {
	return "BACnetEventSummariesList"
}

func (m *_BACnetEventSummariesList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfEventSummaries) > 0 {
		for _, element := range m.ListOfEventSummaries {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventSummariesList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventSummariesListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventSummariesList, error) {
	return BACnetEventSummariesListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventSummariesListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventSummariesList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventSummariesList, error) {
		return BACnetEventSummariesListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventSummariesListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventSummariesList, error) {
	v, err := (&_BACnetEventSummariesList{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventSummariesList) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventSummariesList BACnetEventSummariesList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventSummariesList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventSummariesList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfEventSummaries, err := ReadTerminatedArrayField[BACnetEventSummary](ctx, "listOfEventSummaries", ReadComplex[BACnetEventSummary](BACnetEventSummaryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfEventSummaries' field"))
	}
	m.ListOfEventSummaries = listOfEventSummaries

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventSummariesList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventSummariesList")
	}

	return m, nil
}

func (m *_BACnetEventSummariesList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventSummariesList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventSummariesList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventSummariesList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfEventSummaries", m.GetListOfEventSummaries(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfEventSummaries' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventSummariesList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventSummariesList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventSummariesList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventSummariesList) IsBACnetEventSummariesList() {}

func (m *_BACnetEventSummariesList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
