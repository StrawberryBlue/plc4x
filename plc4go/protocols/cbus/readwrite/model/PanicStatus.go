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

// PanicStatus is the corresponding interface of PanicStatus
type PanicStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetIsNoPanic returns IsNoPanic (virtual field)
	GetIsNoPanic() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsPanicCurrentlyActive returns IsPanicCurrentlyActive (virtual field)
	GetIsPanicCurrentlyActive() bool
	// IsPanicStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPanicStatus()
}

// _PanicStatus is the data-structure of this message
type _PanicStatus struct {
	Status uint8
}

var _ PanicStatus = (*_PanicStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PanicStatus) GetStatus() uint8 {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_PanicStatus) GetIsNoPanic() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStatus()) == (0x00)))
}

func (m *_PanicStatus) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetStatus()) >= (0x01))) && bool(bool((m.GetStatus()) <= (0xFE))))
}

func (m *_PanicStatus) GetIsPanicCurrentlyActive() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStatus()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPanicStatus factory function for _PanicStatus
func NewPanicStatus(status uint8) *_PanicStatus {
	return &_PanicStatus{Status: status}
}

// Deprecated: use the interface for direct cast
func CastPanicStatus(structType any) PanicStatus {
	if casted, ok := structType.(PanicStatus); ok {
		return casted
	}
	if casted, ok := structType.(*PanicStatus); ok {
		return *casted
	}
	return nil
}

func (m *_PanicStatus) GetTypeName() string {
	return "PanicStatus"
}

func (m *_PanicStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (status)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_PanicStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PanicStatusParse(ctx context.Context, theBytes []byte) (PanicStatus, error) {
	return PanicStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PanicStatusParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (PanicStatus, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (PanicStatus, error) {
		return PanicStatusParseWithBuffer(ctx, readBuffer)
	}
}

func PanicStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PanicStatus, error) {
	v, err := (&_PanicStatus{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_PanicStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__panicStatus PanicStatus, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PanicStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PanicStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	isNoPanic, err := ReadVirtualField[bool](ctx, "isNoPanic", (*bool)(nil), bool((status) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNoPanic' field"))
	}
	_ = isNoPanic

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool((status) >= (0x01))) && bool(bool((status) <= (0xFE))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	isPanicCurrentlyActive, err := ReadVirtualField[bool](ctx, "isPanicCurrentlyActive", (*bool)(nil), bool((status) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPanicCurrentlyActive' field"))
	}
	_ = isPanicCurrentlyActive

	if closeErr := readBuffer.CloseContext("PanicStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PanicStatus")
	}

	return m, nil
}

func (m *_PanicStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PanicStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PanicStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PanicStatus")
	}

	if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}
	// Virtual field
	isNoPanic := m.GetIsNoPanic()
	_ = isNoPanic
	if _isNoPanicErr := writeBuffer.WriteVirtual(ctx, "isNoPanic", m.GetIsNoPanic()); _isNoPanicErr != nil {
		return errors.Wrap(_isNoPanicErr, "Error serializing 'isNoPanic' field")
	}
	// Virtual field
	isReserved := m.GetIsReserved()
	_ = isReserved
	if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
		return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
	}
	// Virtual field
	isPanicCurrentlyActive := m.GetIsPanicCurrentlyActive()
	_ = isPanicCurrentlyActive
	if _isPanicCurrentlyActiveErr := writeBuffer.WriteVirtual(ctx, "isPanicCurrentlyActive", m.GetIsPanicCurrentlyActive()); _isPanicCurrentlyActiveErr != nil {
		return errors.Wrap(_isPanicCurrentlyActiveErr, "Error serializing 'isPanicCurrentlyActive' field")
	}

	if popErr := writeBuffer.PopContext("PanicStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PanicStatus")
	}
	return nil
}

func (m *_PanicStatus) IsPanicStatus() {}

func (m *_PanicStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
