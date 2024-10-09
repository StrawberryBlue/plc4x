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

// BACnetTagPayloadDate is the corresponding interface of BACnetTagPayloadDate
type BACnetTagPayloadDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetYearMinus1900 returns YearMinus1900 (property field)
	GetYearMinus1900() uint8
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetDayOfMonth returns DayOfMonth (property field)
	GetDayOfMonth() uint8
	// GetDayOfWeek returns DayOfWeek (property field)
	GetDayOfWeek() uint8
	// GetWildcard returns Wildcard (virtual field)
	GetWildcard() uint8
	// GetYearIsWildcard returns YearIsWildcard (virtual field)
	GetYearIsWildcard() bool
	// GetYear returns Year (virtual field)
	GetYear() uint16
	// GetMonthIsWildcard returns MonthIsWildcard (virtual field)
	GetMonthIsWildcard() bool
	// GetOddMonthWildcard returns OddMonthWildcard (virtual field)
	GetOddMonthWildcard() bool
	// GetEvenMonthWildcard returns EvenMonthWildcard (virtual field)
	GetEvenMonthWildcard() bool
	// GetDayOfMonthIsWildcard returns DayOfMonthIsWildcard (virtual field)
	GetDayOfMonthIsWildcard() bool
	// GetLastDayOfMonthWildcard returns LastDayOfMonthWildcard (virtual field)
	GetLastDayOfMonthWildcard() bool
	// GetOddDayOfMonthWildcard returns OddDayOfMonthWildcard (virtual field)
	GetOddDayOfMonthWildcard() bool
	// GetEvenDayOfMonthWildcard returns EvenDayOfMonthWildcard (virtual field)
	GetEvenDayOfMonthWildcard() bool
	// GetDayOfWeekIsWildcard returns DayOfWeekIsWildcard (virtual field)
	GetDayOfWeekIsWildcard() bool
	// IsBACnetTagPayloadDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadDate()
	// CreateBuilder creates a BACnetTagPayloadDateBuilder
	CreateBACnetTagPayloadDateBuilder() BACnetTagPayloadDateBuilder
}

// _BACnetTagPayloadDate is the data-structure of this message
type _BACnetTagPayloadDate struct {
	YearMinus1900 uint8
	Month         uint8
	DayOfMonth    uint8
	DayOfWeek     uint8
}

var _ BACnetTagPayloadDate = (*_BACnetTagPayloadDate)(nil)

// NewBACnetTagPayloadDate factory function for _BACnetTagPayloadDate
func NewBACnetTagPayloadDate(yearMinus1900 uint8, month uint8, dayOfMonth uint8, dayOfWeek uint8) *_BACnetTagPayloadDate {
	return &_BACnetTagPayloadDate{YearMinus1900: yearMinus1900, Month: month, DayOfMonth: dayOfMonth, DayOfWeek: dayOfWeek}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagPayloadDateBuilder is a builder for BACnetTagPayloadDate
type BACnetTagPayloadDateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(yearMinus1900 uint8, month uint8, dayOfMonth uint8, dayOfWeek uint8) BACnetTagPayloadDateBuilder
	// WithYearMinus1900 adds YearMinus1900 (property field)
	WithYearMinus1900(uint8) BACnetTagPayloadDateBuilder
	// WithMonth adds Month (property field)
	WithMonth(uint8) BACnetTagPayloadDateBuilder
	// WithDayOfMonth adds DayOfMonth (property field)
	WithDayOfMonth(uint8) BACnetTagPayloadDateBuilder
	// WithDayOfWeek adds DayOfWeek (property field)
	WithDayOfWeek(uint8) BACnetTagPayloadDateBuilder
	// Build builds the BACnetTagPayloadDate or returns an error if something is wrong
	Build() (BACnetTagPayloadDate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagPayloadDate
}

// NewBACnetTagPayloadDateBuilder() creates a BACnetTagPayloadDateBuilder
func NewBACnetTagPayloadDateBuilder() BACnetTagPayloadDateBuilder {
	return &_BACnetTagPayloadDateBuilder{_BACnetTagPayloadDate: new(_BACnetTagPayloadDate)}
}

type _BACnetTagPayloadDateBuilder struct {
	*_BACnetTagPayloadDate

	err *utils.MultiError
}

var _ (BACnetTagPayloadDateBuilder) = (*_BACnetTagPayloadDateBuilder)(nil)

func (b *_BACnetTagPayloadDateBuilder) WithMandatoryFields(yearMinus1900 uint8, month uint8, dayOfMonth uint8, dayOfWeek uint8) BACnetTagPayloadDateBuilder {
	return b.WithYearMinus1900(yearMinus1900).WithMonth(month).WithDayOfMonth(dayOfMonth).WithDayOfWeek(dayOfWeek)
}

func (b *_BACnetTagPayloadDateBuilder) WithYearMinus1900(yearMinus1900 uint8) BACnetTagPayloadDateBuilder {
	b.YearMinus1900 = yearMinus1900
	return b
}

func (b *_BACnetTagPayloadDateBuilder) WithMonth(month uint8) BACnetTagPayloadDateBuilder {
	b.Month = month
	return b
}

func (b *_BACnetTagPayloadDateBuilder) WithDayOfMonth(dayOfMonth uint8) BACnetTagPayloadDateBuilder {
	b.DayOfMonth = dayOfMonth
	return b
}

func (b *_BACnetTagPayloadDateBuilder) WithDayOfWeek(dayOfWeek uint8) BACnetTagPayloadDateBuilder {
	b.DayOfWeek = dayOfWeek
	return b
}

func (b *_BACnetTagPayloadDateBuilder) Build() (BACnetTagPayloadDate, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTagPayloadDate.deepCopy(), nil
}

func (b *_BACnetTagPayloadDateBuilder) MustBuild() BACnetTagPayloadDate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTagPayloadDateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTagPayloadDateBuilder().(*_BACnetTagPayloadDateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTagPayloadDateBuilder creates a BACnetTagPayloadDateBuilder
func (b *_BACnetTagPayloadDate) CreateBACnetTagPayloadDateBuilder() BACnetTagPayloadDateBuilder {
	if b == nil {
		return NewBACnetTagPayloadDateBuilder()
	}
	return &_BACnetTagPayloadDateBuilder{_BACnetTagPayloadDate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadDate) GetYearMinus1900() uint8 {
	return m.YearMinus1900
}

func (m *_BACnetTagPayloadDate) GetMonth() uint8 {
	return m.Month
}

func (m *_BACnetTagPayloadDate) GetDayOfMonth() uint8 {
	return m.DayOfMonth
}

func (m *_BACnetTagPayloadDate) GetDayOfWeek() uint8 {
	return m.DayOfWeek
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadDate) GetWildcard() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(0xFF)
}

func (m *_BACnetTagPayloadDate) GetYearIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetYearMinus1900()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadDate) GetYear() uint16 {
	ctx := context.Background()
	_ = ctx
	return uint16(uint16(m.GetYearMinus1900()) + uint16(uint16(1900)))
}

func (m *_BACnetTagPayloadDate) GetMonthIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadDate) GetOddMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (13)))
}

func (m *_BACnetTagPayloadDate) GetEvenMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (14)))
}

func (m *_BACnetTagPayloadDate) GetDayOfMonthIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadDate) GetLastDayOfMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (32)))
}

func (m *_BACnetTagPayloadDate) GetOddDayOfMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (33)))
}

func (m *_BACnetTagPayloadDate) GetEvenDayOfMonthWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfMonth()) == (34)))
}

func (m *_BACnetTagPayloadDate) GetDayOfWeekIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfWeek()) == (m.GetWildcard())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadDate(structType any) BACnetTagPayloadDate {
	if casted, ok := structType.(BACnetTagPayloadDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadDate) GetTypeName() string {
	return "BACnetTagPayloadDate"
}

func (m *_BACnetTagPayloadDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (yearMinus1900)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (month)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfMonth)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfWeek)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadDateParse(ctx context.Context, theBytes []byte) (BACnetTagPayloadDate, error) {
	return BACnetTagPayloadDateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadDateParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDate, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDate, error) {
		return BACnetTagPayloadDateParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTagPayloadDateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDate, error) {
	v, err := (&_BACnetTagPayloadDate{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadDate) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTagPayloadDate BACnetTagPayloadDate, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	wildcard, err := ReadVirtualField[uint8](ctx, "wildcard", (*uint8)(nil), 0xFF)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'wildcard' field"))
	}
	_ = wildcard

	yearMinus1900, err := ReadSimpleField(ctx, "yearMinus1900", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'yearMinus1900' field"))
	}
	m.YearMinus1900 = yearMinus1900

	yearIsWildcard, err := ReadVirtualField[bool](ctx, "yearIsWildcard", (*bool)(nil), bool((yearMinus1900) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'yearIsWildcard' field"))
	}
	_ = yearIsWildcard

	year, err := ReadVirtualField[uint16](ctx, "year", (*uint16)(nil), uint16(yearMinus1900)+uint16(uint16(1900)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'year' field"))
	}
	_ = year

	month, err := ReadSimpleField(ctx, "month", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'month' field"))
	}
	m.Month = month

	monthIsWildcard, err := ReadVirtualField[bool](ctx, "monthIsWildcard", (*bool)(nil), bool((month) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monthIsWildcard' field"))
	}
	_ = monthIsWildcard

	oddMonthWildcard, err := ReadVirtualField[bool](ctx, "oddMonthWildcard", (*bool)(nil), bool((month) == (13)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'oddMonthWildcard' field"))
	}
	_ = oddMonthWildcard

	evenMonthWildcard, err := ReadVirtualField[bool](ctx, "evenMonthWildcard", (*bool)(nil), bool((month) == (14)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'evenMonthWildcard' field"))
	}
	_ = evenMonthWildcard

	dayOfMonth, err := ReadSimpleField(ctx, "dayOfMonth", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfMonth' field"))
	}
	m.DayOfMonth = dayOfMonth

	dayOfMonthIsWildcard, err := ReadVirtualField[bool](ctx, "dayOfMonthIsWildcard", (*bool)(nil), bool((dayOfMonth) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfMonthIsWildcard' field"))
	}
	_ = dayOfMonthIsWildcard

	lastDayOfMonthWildcard, err := ReadVirtualField[bool](ctx, "lastDayOfMonthWildcard", (*bool)(nil), bool((dayOfMonth) == (32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastDayOfMonthWildcard' field"))
	}
	_ = lastDayOfMonthWildcard

	oddDayOfMonthWildcard, err := ReadVirtualField[bool](ctx, "oddDayOfMonthWildcard", (*bool)(nil), bool((dayOfMonth) == (33)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'oddDayOfMonthWildcard' field"))
	}
	_ = oddDayOfMonthWildcard

	evenDayOfMonthWildcard, err := ReadVirtualField[bool](ctx, "evenDayOfMonthWildcard", (*bool)(nil), bool((dayOfMonth) == (34)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'evenDayOfMonthWildcard' field"))
	}
	_ = evenDayOfMonthWildcard

	dayOfWeek, err := ReadSimpleField(ctx, "dayOfWeek", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfWeek' field"))
	}
	m.DayOfWeek = dayOfWeek

	dayOfWeekIsWildcard, err := ReadVirtualField[bool](ctx, "dayOfWeekIsWildcard", (*bool)(nil), bool((dayOfWeek) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfWeekIsWildcard' field"))
	}
	_ = dayOfWeekIsWildcard

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadDate")
	}

	return m, nil
}

func (m *_BACnetTagPayloadDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadDate")
	}
	// Virtual field
	wildcard := m.GetWildcard()
	_ = wildcard
	if _wildcardErr := writeBuffer.WriteVirtual(ctx, "wildcard", m.GetWildcard()); _wildcardErr != nil {
		return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "yearMinus1900", m.GetYearMinus1900(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'yearMinus1900' field")
	}
	// Virtual field
	yearIsWildcard := m.GetYearIsWildcard()
	_ = yearIsWildcard
	if _yearIsWildcardErr := writeBuffer.WriteVirtual(ctx, "yearIsWildcard", m.GetYearIsWildcard()); _yearIsWildcardErr != nil {
		return errors.Wrap(_yearIsWildcardErr, "Error serializing 'yearIsWildcard' field")
	}
	// Virtual field
	year := m.GetYear()
	_ = year
	if _yearErr := writeBuffer.WriteVirtual(ctx, "year", m.GetYear()); _yearErr != nil {
		return errors.Wrap(_yearErr, "Error serializing 'year' field")
	}

	if err := WriteSimpleField[uint8](ctx, "month", m.GetMonth(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'month' field")
	}
	// Virtual field
	monthIsWildcard := m.GetMonthIsWildcard()
	_ = monthIsWildcard
	if _monthIsWildcardErr := writeBuffer.WriteVirtual(ctx, "monthIsWildcard", m.GetMonthIsWildcard()); _monthIsWildcardErr != nil {
		return errors.Wrap(_monthIsWildcardErr, "Error serializing 'monthIsWildcard' field")
	}
	// Virtual field
	oddMonthWildcard := m.GetOddMonthWildcard()
	_ = oddMonthWildcard
	if _oddMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "oddMonthWildcard", m.GetOddMonthWildcard()); _oddMonthWildcardErr != nil {
		return errors.Wrap(_oddMonthWildcardErr, "Error serializing 'oddMonthWildcard' field")
	}
	// Virtual field
	evenMonthWildcard := m.GetEvenMonthWildcard()
	_ = evenMonthWildcard
	if _evenMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "evenMonthWildcard", m.GetEvenMonthWildcard()); _evenMonthWildcardErr != nil {
		return errors.Wrap(_evenMonthWildcardErr, "Error serializing 'evenMonthWildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "dayOfMonth", m.GetDayOfMonth(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'dayOfMonth' field")
	}
	// Virtual field
	dayOfMonthIsWildcard := m.GetDayOfMonthIsWildcard()
	_ = dayOfMonthIsWildcard
	if _dayOfMonthIsWildcardErr := writeBuffer.WriteVirtual(ctx, "dayOfMonthIsWildcard", m.GetDayOfMonthIsWildcard()); _dayOfMonthIsWildcardErr != nil {
		return errors.Wrap(_dayOfMonthIsWildcardErr, "Error serializing 'dayOfMonthIsWildcard' field")
	}
	// Virtual field
	lastDayOfMonthWildcard := m.GetLastDayOfMonthWildcard()
	_ = lastDayOfMonthWildcard
	if _lastDayOfMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "lastDayOfMonthWildcard", m.GetLastDayOfMonthWildcard()); _lastDayOfMonthWildcardErr != nil {
		return errors.Wrap(_lastDayOfMonthWildcardErr, "Error serializing 'lastDayOfMonthWildcard' field")
	}
	// Virtual field
	oddDayOfMonthWildcard := m.GetOddDayOfMonthWildcard()
	_ = oddDayOfMonthWildcard
	if _oddDayOfMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "oddDayOfMonthWildcard", m.GetOddDayOfMonthWildcard()); _oddDayOfMonthWildcardErr != nil {
		return errors.Wrap(_oddDayOfMonthWildcardErr, "Error serializing 'oddDayOfMonthWildcard' field")
	}
	// Virtual field
	evenDayOfMonthWildcard := m.GetEvenDayOfMonthWildcard()
	_ = evenDayOfMonthWildcard
	if _evenDayOfMonthWildcardErr := writeBuffer.WriteVirtual(ctx, "evenDayOfMonthWildcard", m.GetEvenDayOfMonthWildcard()); _evenDayOfMonthWildcardErr != nil {
		return errors.Wrap(_evenDayOfMonthWildcardErr, "Error serializing 'evenDayOfMonthWildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "dayOfWeek", m.GetDayOfWeek(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'dayOfWeek' field")
	}
	// Virtual field
	dayOfWeekIsWildcard := m.GetDayOfWeekIsWildcard()
	_ = dayOfWeekIsWildcard
	if _dayOfWeekIsWildcardErr := writeBuffer.WriteVirtual(ctx, "dayOfWeekIsWildcard", m.GetDayOfWeekIsWildcard()); _dayOfWeekIsWildcardErr != nil {
		return errors.Wrap(_dayOfWeekIsWildcardErr, "Error serializing 'dayOfWeekIsWildcard' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadDate")
	}
	return nil
}

func (m *_BACnetTagPayloadDate) IsBACnetTagPayloadDate() {}

func (m *_BACnetTagPayloadDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadDate) deepCopy() *_BACnetTagPayloadDate {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadDateCopy := &_BACnetTagPayloadDate{
		m.YearMinus1900,
		m.Month,
		m.DayOfMonth,
		m.DayOfWeek,
	}
	return _BACnetTagPayloadDateCopy
}

func (m *_BACnetTagPayloadDate) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
