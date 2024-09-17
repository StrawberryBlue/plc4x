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

// ErrorReportingSystemCategory is the corresponding interface of ErrorReportingSystemCategory
type ErrorReportingSystemCategory interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSystemCategoryClass returns SystemCategoryClass (property field)
	GetSystemCategoryClass() ErrorReportingSystemCategoryClass
	// GetSystemCategoryType returns SystemCategoryType (property field)
	GetSystemCategoryType() ErrorReportingSystemCategoryType
	// GetSystemCategoryVariant returns SystemCategoryVariant (property field)
	GetSystemCategoryVariant() ErrorReportingSystemCategoryVariant
	// IsErrorReportingSystemCategory is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategory()
}

// _ErrorReportingSystemCategory is the data-structure of this message
type _ErrorReportingSystemCategory struct {
	SystemCategoryClass   ErrorReportingSystemCategoryClass
	SystemCategoryType    ErrorReportingSystemCategoryType
	SystemCategoryVariant ErrorReportingSystemCategoryVariant
}

var _ ErrorReportingSystemCategory = (*_ErrorReportingSystemCategory)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategory) GetSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return m.SystemCategoryClass
}

func (m *_ErrorReportingSystemCategory) GetSystemCategoryType() ErrorReportingSystemCategoryType {
	return m.SystemCategoryType
}

func (m *_ErrorReportingSystemCategory) GetSystemCategoryVariant() ErrorReportingSystemCategoryVariant {
	return m.SystemCategoryVariant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingSystemCategory factory function for _ErrorReportingSystemCategory
func NewErrorReportingSystemCategory(systemCategoryClass ErrorReportingSystemCategoryClass, systemCategoryType ErrorReportingSystemCategoryType, systemCategoryVariant ErrorReportingSystemCategoryVariant) *_ErrorReportingSystemCategory {
	if systemCategoryType == nil {
		panic("systemCategoryType of type ErrorReportingSystemCategoryType for ErrorReportingSystemCategory must not be nil")
	}
	return &_ErrorReportingSystemCategory{SystemCategoryClass: systemCategoryClass, SystemCategoryType: systemCategoryType, SystemCategoryVariant: systemCategoryVariant}
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategory(structType any) ErrorReportingSystemCategory {
	if casted, ok := structType.(ErrorReportingSystemCategory); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategory); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategory) GetTypeName() string {
	return "ErrorReportingSystemCategory"
}

func (m *_ErrorReportingSystemCategory) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (systemCategoryClass)
	lengthInBits += 4

	// Simple field (systemCategoryType)
	lengthInBits += m.SystemCategoryType.GetLengthInBits(ctx)

	// Simple field (systemCategoryVariant)
	lengthInBits += 2

	return lengthInBits
}

func (m *_ErrorReportingSystemCategory) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategory, error) {
	return ErrorReportingSystemCategoryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategory, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategory, error) {
		return ErrorReportingSystemCategoryParseWithBuffer(ctx, readBuffer)
	}
}

func ErrorReportingSystemCategoryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategory, error) {
	v, err := (&_ErrorReportingSystemCategory{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ErrorReportingSystemCategory) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__errorReportingSystemCategory ErrorReportingSystemCategory, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategory"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategory")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	systemCategoryClass, err := ReadEnumField[ErrorReportingSystemCategoryClass](ctx, "systemCategoryClass", "ErrorReportingSystemCategoryClass", ReadEnum(ErrorReportingSystemCategoryClassByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemCategoryClass' field"))
	}
	m.SystemCategoryClass = systemCategoryClass

	systemCategoryType, err := ReadSimpleField[ErrorReportingSystemCategoryType](ctx, "systemCategoryType", ReadComplex[ErrorReportingSystemCategoryType](ErrorReportingSystemCategoryTypeParseWithBufferProducer[ErrorReportingSystemCategoryType]((ErrorReportingSystemCategoryClass)(systemCategoryClass)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemCategoryType' field"))
	}
	m.SystemCategoryType = systemCategoryType

	systemCategoryVariant, err := ReadEnumField[ErrorReportingSystemCategoryVariant](ctx, "systemCategoryVariant", "ErrorReportingSystemCategoryVariant", ReadEnum(ErrorReportingSystemCategoryVariantByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemCategoryVariant' field"))
	}
	m.SystemCategoryVariant = systemCategoryVariant

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategory"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategory")
	}

	return m, nil
}

func (m *_ErrorReportingSystemCategory) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategory) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategory"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategory")
	}

	if err := WriteSimpleEnumField[ErrorReportingSystemCategoryClass](ctx, "systemCategoryClass", "ErrorReportingSystemCategoryClass", m.GetSystemCategoryClass(), WriteEnum[ErrorReportingSystemCategoryClass, uint8](ErrorReportingSystemCategoryClass.GetValue, ErrorReportingSystemCategoryClass.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
		return errors.Wrap(err, "Error serializing 'systemCategoryClass' field")
	}

	if err := WriteSimpleField[ErrorReportingSystemCategoryType](ctx, "systemCategoryType", m.GetSystemCategoryType(), WriteComplex[ErrorReportingSystemCategoryType](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'systemCategoryType' field")
	}

	if err := WriteSimpleEnumField[ErrorReportingSystemCategoryVariant](ctx, "systemCategoryVariant", "ErrorReportingSystemCategoryVariant", m.GetSystemCategoryVariant(), WriteEnum[ErrorReportingSystemCategoryVariant, uint8](ErrorReportingSystemCategoryVariant.GetValue, ErrorReportingSystemCategoryVariant.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'systemCategoryVariant' field")
	}

	if popErr := writeBuffer.PopContext("ErrorReportingSystemCategory"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategory")
	}
	return nil
}

func (m *_ErrorReportingSystemCategory) IsErrorReportingSystemCategory() {}

func (m *_ErrorReportingSystemCategory) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
