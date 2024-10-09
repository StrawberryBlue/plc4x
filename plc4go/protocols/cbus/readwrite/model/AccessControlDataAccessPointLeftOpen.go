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

// AccessControlDataAccessPointLeftOpen is the corresponding interface of AccessControlDataAccessPointLeftOpen
type AccessControlDataAccessPointLeftOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AccessControlData
	// IsAccessControlDataAccessPointLeftOpen is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataAccessPointLeftOpen()
	// CreateBuilder creates a AccessControlDataAccessPointLeftOpenBuilder
	CreateAccessControlDataAccessPointLeftOpenBuilder() AccessControlDataAccessPointLeftOpenBuilder
}

// _AccessControlDataAccessPointLeftOpen is the data-structure of this message
type _AccessControlDataAccessPointLeftOpen struct {
	AccessControlDataContract
}

var _ AccessControlDataAccessPointLeftOpen = (*_AccessControlDataAccessPointLeftOpen)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataAccessPointLeftOpen)(nil)

// NewAccessControlDataAccessPointLeftOpen factory function for _AccessControlDataAccessPointLeftOpen
func NewAccessControlDataAccessPointLeftOpen(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataAccessPointLeftOpen {
	_result := &_AccessControlDataAccessPointLeftOpen{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AccessControlDataAccessPointLeftOpenBuilder is a builder for AccessControlDataAccessPointLeftOpen
type AccessControlDataAccessPointLeftOpenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AccessControlDataAccessPointLeftOpenBuilder
	// Build builds the AccessControlDataAccessPointLeftOpen or returns an error if something is wrong
	Build() (AccessControlDataAccessPointLeftOpen, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AccessControlDataAccessPointLeftOpen
}

// NewAccessControlDataAccessPointLeftOpenBuilder() creates a AccessControlDataAccessPointLeftOpenBuilder
func NewAccessControlDataAccessPointLeftOpenBuilder() AccessControlDataAccessPointLeftOpenBuilder {
	return &_AccessControlDataAccessPointLeftOpenBuilder{_AccessControlDataAccessPointLeftOpen: new(_AccessControlDataAccessPointLeftOpen)}
}

type _AccessControlDataAccessPointLeftOpenBuilder struct {
	*_AccessControlDataAccessPointLeftOpen

	parentBuilder *_AccessControlDataBuilder

	err *utils.MultiError
}

var _ (AccessControlDataAccessPointLeftOpenBuilder) = (*_AccessControlDataAccessPointLeftOpenBuilder)(nil)

func (b *_AccessControlDataAccessPointLeftOpenBuilder) setParent(contract AccessControlDataContract) {
	b.AccessControlDataContract = contract
}

func (b *_AccessControlDataAccessPointLeftOpenBuilder) WithMandatoryFields() AccessControlDataAccessPointLeftOpenBuilder {
	return b
}

func (b *_AccessControlDataAccessPointLeftOpenBuilder) Build() (AccessControlDataAccessPointLeftOpen, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AccessControlDataAccessPointLeftOpen.deepCopy(), nil
}

func (b *_AccessControlDataAccessPointLeftOpenBuilder) MustBuild() AccessControlDataAccessPointLeftOpen {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AccessControlDataAccessPointLeftOpenBuilder) Done() AccessControlDataBuilder {
	return b.parentBuilder
}

func (b *_AccessControlDataAccessPointLeftOpenBuilder) buildForAccessControlData() (AccessControlData, error) {
	return b.Build()
}

func (b *_AccessControlDataAccessPointLeftOpenBuilder) DeepCopy() any {
	_copy := b.CreateAccessControlDataAccessPointLeftOpenBuilder().(*_AccessControlDataAccessPointLeftOpenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAccessControlDataAccessPointLeftOpenBuilder creates a AccessControlDataAccessPointLeftOpenBuilder
func (b *_AccessControlDataAccessPointLeftOpen) CreateAccessControlDataAccessPointLeftOpenBuilder() AccessControlDataAccessPointLeftOpenBuilder {
	if b == nil {
		return NewAccessControlDataAccessPointLeftOpenBuilder()
	}
	return &_AccessControlDataAccessPointLeftOpenBuilder{_AccessControlDataAccessPointLeftOpen: b.deepCopy()}
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

func (m *_AccessControlDataAccessPointLeftOpen) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataAccessPointLeftOpen(structType any) AccessControlDataAccessPointLeftOpen {
	if casted, ok := structType.(AccessControlDataAccessPointLeftOpen); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataAccessPointLeftOpen); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataAccessPointLeftOpen) GetTypeName() string {
	return "AccessControlDataAccessPointLeftOpen"
}

func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataAccessPointLeftOpen) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataAccessPointLeftOpen AccessControlDataAccessPointLeftOpen, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataAccessPointLeftOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataAccessPointLeftOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataAccessPointLeftOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataAccessPointLeftOpen")
	}

	return m, nil
}

func (m *_AccessControlDataAccessPointLeftOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataAccessPointLeftOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataAccessPointLeftOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataAccessPointLeftOpen")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataAccessPointLeftOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataAccessPointLeftOpen")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataAccessPointLeftOpen) IsAccessControlDataAccessPointLeftOpen() {}

func (m *_AccessControlDataAccessPointLeftOpen) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AccessControlDataAccessPointLeftOpen) deepCopy() *_AccessControlDataAccessPointLeftOpen {
	if m == nil {
		return nil
	}
	_AccessControlDataAccessPointLeftOpenCopy := &_AccessControlDataAccessPointLeftOpen{
		m.AccessControlDataContract.(*_AccessControlData).deepCopy(),
	}
	m.AccessControlDataContract.(*_AccessControlData)._SubType = m
	return _AccessControlDataAccessPointLeftOpenCopy
}

func (m *_AccessControlDataAccessPointLeftOpen) String() string {
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
