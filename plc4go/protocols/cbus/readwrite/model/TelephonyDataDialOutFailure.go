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

// TelephonyDataDialOutFailure is the corresponding interface of TelephonyDataDialOutFailure
type TelephonyDataDialOutFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TelephonyData
	// GetReason returns Reason (property field)
	GetReason() DialOutFailureReason
	// IsTelephonyDataDialOutFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataDialOutFailure()
	// CreateBuilder creates a TelephonyDataDialOutFailureBuilder
	CreateTelephonyDataDialOutFailureBuilder() TelephonyDataDialOutFailureBuilder
}

// _TelephonyDataDialOutFailure is the data-structure of this message
type _TelephonyDataDialOutFailure struct {
	TelephonyDataContract
	Reason DialOutFailureReason
}

var _ TelephonyDataDialOutFailure = (*_TelephonyDataDialOutFailure)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataDialOutFailure)(nil)

// NewTelephonyDataDialOutFailure factory function for _TelephonyDataDialOutFailure
func NewTelephonyDataDialOutFailure(commandTypeContainer TelephonyCommandTypeContainer, argument byte, reason DialOutFailureReason) *_TelephonyDataDialOutFailure {
	_result := &_TelephonyDataDialOutFailure{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		Reason:                reason,
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TelephonyDataDialOutFailureBuilder is a builder for TelephonyDataDialOutFailure
type TelephonyDataDialOutFailureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reason DialOutFailureReason) TelephonyDataDialOutFailureBuilder
	// WithReason adds Reason (property field)
	WithReason(DialOutFailureReason) TelephonyDataDialOutFailureBuilder
	// Build builds the TelephonyDataDialOutFailure or returns an error if something is wrong
	Build() (TelephonyDataDialOutFailure, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TelephonyDataDialOutFailure
}

// NewTelephonyDataDialOutFailureBuilder() creates a TelephonyDataDialOutFailureBuilder
func NewTelephonyDataDialOutFailureBuilder() TelephonyDataDialOutFailureBuilder {
	return &_TelephonyDataDialOutFailureBuilder{_TelephonyDataDialOutFailure: new(_TelephonyDataDialOutFailure)}
}

type _TelephonyDataDialOutFailureBuilder struct {
	*_TelephonyDataDialOutFailure

	parentBuilder *_TelephonyDataBuilder

	err *utils.MultiError
}

var _ (TelephonyDataDialOutFailureBuilder) = (*_TelephonyDataDialOutFailureBuilder)(nil)

func (b *_TelephonyDataDialOutFailureBuilder) setParent(contract TelephonyDataContract) {
	b.TelephonyDataContract = contract
}

func (b *_TelephonyDataDialOutFailureBuilder) WithMandatoryFields(reason DialOutFailureReason) TelephonyDataDialOutFailureBuilder {
	return b.WithReason(reason)
}

func (b *_TelephonyDataDialOutFailureBuilder) WithReason(reason DialOutFailureReason) TelephonyDataDialOutFailureBuilder {
	b.Reason = reason
	return b
}

func (b *_TelephonyDataDialOutFailureBuilder) Build() (TelephonyDataDialOutFailure, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TelephonyDataDialOutFailure.deepCopy(), nil
}

func (b *_TelephonyDataDialOutFailureBuilder) MustBuild() TelephonyDataDialOutFailure {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TelephonyDataDialOutFailureBuilder) Done() TelephonyDataBuilder {
	return b.parentBuilder
}

func (b *_TelephonyDataDialOutFailureBuilder) buildForTelephonyData() (TelephonyData, error) {
	return b.Build()
}

func (b *_TelephonyDataDialOutFailureBuilder) DeepCopy() any {
	_copy := b.CreateTelephonyDataDialOutFailureBuilder().(*_TelephonyDataDialOutFailureBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTelephonyDataDialOutFailureBuilder creates a TelephonyDataDialOutFailureBuilder
func (b *_TelephonyDataDialOutFailure) CreateTelephonyDataDialOutFailureBuilder() TelephonyDataDialOutFailureBuilder {
	if b == nil {
		return NewTelephonyDataDialOutFailureBuilder()
	}
	return &_TelephonyDataDialOutFailureBuilder{_TelephonyDataDialOutFailure: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataDialOutFailure) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDialOutFailure) GetReason() DialOutFailureReason {
	return m.Reason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTelephonyDataDialOutFailure(structType any) TelephonyDataDialOutFailure {
	if casted, ok := structType.(TelephonyDataDialOutFailure); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDialOutFailure); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDialOutFailure) GetTypeName() string {
	return "TelephonyDataDialOutFailure"
}

func (m *_TelephonyDataDialOutFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (reason)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyDataDialOutFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataDialOutFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataDialOutFailure TelephonyDataDialOutFailure, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDialOutFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDialOutFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reason, err := ReadEnumField[DialOutFailureReason](ctx, "reason", "DialOutFailureReason", ReadEnum(DialOutFailureReasonByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reason' field"))
	}
	m.Reason = reason

	if closeErr := readBuffer.CloseContext("TelephonyDataDialOutFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDialOutFailure")
	}

	return m, nil
}

func (m *_TelephonyDataDialOutFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataDialOutFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDialOutFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDialOutFailure")
		}

		if err := WriteSimpleEnumField[DialOutFailureReason](ctx, "reason", "DialOutFailureReason", m.GetReason(), WriteEnum[DialOutFailureReason, uint8](DialOutFailureReason.GetValue, DialOutFailureReason.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'reason' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDialOutFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDialOutFailure")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataDialOutFailure) IsTelephonyDataDialOutFailure() {}

func (m *_TelephonyDataDialOutFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TelephonyDataDialOutFailure) deepCopy() *_TelephonyDataDialOutFailure {
	if m == nil {
		return nil
	}
	_TelephonyDataDialOutFailureCopy := &_TelephonyDataDialOutFailure{
		m.TelephonyDataContract.(*_TelephonyData).deepCopy(),
		m.Reason,
	}
	m.TelephonyDataContract.(*_TelephonyData)._SubType = m
	return _TelephonyDataDialOutFailureCopy
}

func (m *_TelephonyDataDialOutFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
