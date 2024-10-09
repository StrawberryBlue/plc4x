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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataGasAlarmRaised is the corresponding interface of SecurityDataGasAlarmRaised
type SecurityDataGasAlarmRaised interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataGasAlarmRaised is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataGasAlarmRaised()
	// CreateBuilder creates a SecurityDataGasAlarmRaisedBuilder
	CreateSecurityDataGasAlarmRaisedBuilder() SecurityDataGasAlarmRaisedBuilder
}

// _SecurityDataGasAlarmRaised is the data-structure of this message
type _SecurityDataGasAlarmRaised struct {
	SecurityDataContract
}

var _ SecurityDataGasAlarmRaised = (*_SecurityDataGasAlarmRaised)(nil)
var _ SecurityDataRequirements = (*_SecurityDataGasAlarmRaised)(nil)

// NewSecurityDataGasAlarmRaised factory function for _SecurityDataGasAlarmRaised
func NewSecurityDataGasAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataGasAlarmRaised {
	_result := &_SecurityDataGasAlarmRaised{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataGasAlarmRaisedBuilder is a builder for SecurityDataGasAlarmRaised
type SecurityDataGasAlarmRaisedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataGasAlarmRaisedBuilder
	// Build builds the SecurityDataGasAlarmRaised or returns an error if something is wrong
	Build() (SecurityDataGasAlarmRaised, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataGasAlarmRaised
}

// NewSecurityDataGasAlarmRaisedBuilder() creates a SecurityDataGasAlarmRaisedBuilder
func NewSecurityDataGasAlarmRaisedBuilder() SecurityDataGasAlarmRaisedBuilder {
	return &_SecurityDataGasAlarmRaisedBuilder{_SecurityDataGasAlarmRaised: new(_SecurityDataGasAlarmRaised)}
}

type _SecurityDataGasAlarmRaisedBuilder struct {
	*_SecurityDataGasAlarmRaised

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataGasAlarmRaisedBuilder) = (*_SecurityDataGasAlarmRaisedBuilder)(nil)

func (b *_SecurityDataGasAlarmRaisedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
}

func (b *_SecurityDataGasAlarmRaisedBuilder) WithMandatoryFields() SecurityDataGasAlarmRaisedBuilder {
	return b
}

func (b *_SecurityDataGasAlarmRaisedBuilder) Build() (SecurityDataGasAlarmRaised, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataGasAlarmRaised.deepCopy(), nil
}

func (b *_SecurityDataGasAlarmRaisedBuilder) MustBuild() SecurityDataGasAlarmRaised {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SecurityDataGasAlarmRaisedBuilder) Done() SecurityDataBuilder {
	return b.parentBuilder
}

func (b *_SecurityDataGasAlarmRaisedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataGasAlarmRaisedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataGasAlarmRaisedBuilder().(*_SecurityDataGasAlarmRaisedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataGasAlarmRaisedBuilder creates a SecurityDataGasAlarmRaisedBuilder
func (b *_SecurityDataGasAlarmRaised) CreateSecurityDataGasAlarmRaisedBuilder() SecurityDataGasAlarmRaisedBuilder {
	if b == nil {
		return NewSecurityDataGasAlarmRaisedBuilder()
	}
	return &_SecurityDataGasAlarmRaisedBuilder{_SecurityDataGasAlarmRaised: b.deepCopy()}
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

func (m *_SecurityDataGasAlarmRaised) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataGasAlarmRaised(structType any) SecurityDataGasAlarmRaised {
	if casted, ok := structType.(SecurityDataGasAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataGasAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataGasAlarmRaised) GetTypeName() string {
	return "SecurityDataGasAlarmRaised"
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataGasAlarmRaised) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataGasAlarmRaised SecurityDataGasAlarmRaised, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataGasAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataGasAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataGasAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataGasAlarmRaised")
	}

	return m, nil
}

func (m *_SecurityDataGasAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataGasAlarmRaised) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataGasAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataGasAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataGasAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataGasAlarmRaised")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataGasAlarmRaised) IsSecurityDataGasAlarmRaised() {}

func (m *_SecurityDataGasAlarmRaised) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataGasAlarmRaised) deepCopy() *_SecurityDataGasAlarmRaised {
	if m == nil {
		return nil
	}
	_SecurityDataGasAlarmRaisedCopy := &_SecurityDataGasAlarmRaised{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataGasAlarmRaisedCopy
}

func (m *_SecurityDataGasAlarmRaised) String() string {
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
