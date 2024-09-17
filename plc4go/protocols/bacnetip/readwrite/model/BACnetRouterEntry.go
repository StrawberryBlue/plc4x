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

// BACnetRouterEntry is the corresponding interface of BACnetRouterEntry
type BACnetRouterEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() BACnetContextTagUnsignedInteger
	// GetMacAddress returns MacAddress (property field)
	GetMacAddress() BACnetContextTagOctetString
	// GetStatus returns Status (property field)
	GetStatus() BACnetRouterEntryStatusTagged
	// GetPerformanceIndex returns PerformanceIndex (property field)
	GetPerformanceIndex() BACnetContextTagOctetString
	// IsBACnetRouterEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRouterEntry()
}

// _BACnetRouterEntry is the data-structure of this message
type _BACnetRouterEntry struct {
	NetworkNumber    BACnetContextTagUnsignedInteger
	MacAddress       BACnetContextTagOctetString
	Status           BACnetRouterEntryStatusTagged
	PerformanceIndex BACnetContextTagOctetString
}

var _ BACnetRouterEntry = (*_BACnetRouterEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRouterEntry) GetNetworkNumber() BACnetContextTagUnsignedInteger {
	return m.NetworkNumber
}

func (m *_BACnetRouterEntry) GetMacAddress() BACnetContextTagOctetString {
	return m.MacAddress
}

func (m *_BACnetRouterEntry) GetStatus() BACnetRouterEntryStatusTagged {
	return m.Status
}

func (m *_BACnetRouterEntry) GetPerformanceIndex() BACnetContextTagOctetString {
	return m.PerformanceIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRouterEntry factory function for _BACnetRouterEntry
func NewBACnetRouterEntry(networkNumber BACnetContextTagUnsignedInteger, macAddress BACnetContextTagOctetString, status BACnetRouterEntryStatusTagged, performanceIndex BACnetContextTagOctetString) *_BACnetRouterEntry {
	if networkNumber == nil {
		panic("networkNumber of type BACnetContextTagUnsignedInteger for BACnetRouterEntry must not be nil")
	}
	if macAddress == nil {
		panic("macAddress of type BACnetContextTagOctetString for BACnetRouterEntry must not be nil")
	}
	if status == nil {
		panic("status of type BACnetRouterEntryStatusTagged for BACnetRouterEntry must not be nil")
	}
	return &_BACnetRouterEntry{NetworkNumber: networkNumber, MacAddress: macAddress, Status: status, PerformanceIndex: performanceIndex}
}

// Deprecated: use the interface for direct cast
func CastBACnetRouterEntry(structType any) BACnetRouterEntry {
	if casted, ok := structType.(BACnetRouterEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRouterEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRouterEntry) GetTypeName() string {
	return "BACnetRouterEntry"
}

func (m *_BACnetRouterEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkNumber)
	lengthInBits += m.NetworkNumber.GetLengthInBits(ctx)

	// Simple field (macAddress)
	lengthInBits += m.MacAddress.GetLengthInBits(ctx)

	// Simple field (status)
	lengthInBits += m.Status.GetLengthInBits(ctx)

	// Optional Field (performanceIndex)
	if m.PerformanceIndex != nil {
		lengthInBits += m.PerformanceIndex.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetRouterEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRouterEntryParse(ctx context.Context, theBytes []byte) (BACnetRouterEntry, error) {
	return BACnetRouterEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRouterEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRouterEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRouterEntry, error) {
		return BACnetRouterEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetRouterEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRouterEntry, error) {
	v, err := (&_BACnetRouterEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetRouterEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetRouterEntry BACnetRouterEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRouterEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRouterEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "networkNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumber' field"))
	}
	m.NetworkNumber = networkNumber

	macAddress, err := ReadSimpleField[BACnetContextTagOctetString](ctx, "macAddress", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'macAddress' field"))
	}
	m.MacAddress = macAddress

	status, err := ReadSimpleField[BACnetRouterEntryStatusTagged](ctx, "status", ReadComplex[BACnetRouterEntryStatusTagged](BACnetRouterEntryStatusTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	var performanceIndex BACnetContextTagOctetString
	_performanceIndex, err := ReadOptionalField[BACnetContextTagOctetString](ctx, "performanceIndex", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'performanceIndex' field"))
	}
	if _performanceIndex != nil {
		performanceIndex = *_performanceIndex
		m.PerformanceIndex = performanceIndex
	}

	if closeErr := readBuffer.CloseContext("BACnetRouterEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRouterEntry")
	}

	return m, nil
}

func (m *_BACnetRouterEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRouterEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRouterEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRouterEntry")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "networkNumber", m.GetNetworkNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'networkNumber' field")
	}

	if err := WriteSimpleField[BACnetContextTagOctetString](ctx, "macAddress", m.GetMacAddress(), WriteComplex[BACnetContextTagOctetString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'macAddress' field")
	}

	if err := WriteSimpleField[BACnetRouterEntryStatusTagged](ctx, "status", m.GetStatus(), WriteComplex[BACnetRouterEntryStatusTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if err := WriteOptionalField[BACnetContextTagOctetString](ctx, "performanceIndex", GetRef(m.GetPerformanceIndex()), WriteComplex[BACnetContextTagOctetString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'performanceIndex' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRouterEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRouterEntry")
	}
	return nil
}

func (m *_BACnetRouterEntry) IsBACnetRouterEntry() {}

func (m *_BACnetRouterEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
