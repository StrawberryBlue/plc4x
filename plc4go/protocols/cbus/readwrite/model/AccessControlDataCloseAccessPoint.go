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

// AccessControlDataCloseAccessPoint is the corresponding interface of AccessControlDataCloseAccessPoint
type AccessControlDataCloseAccessPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AccessControlData
	// IsAccessControlDataCloseAccessPoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataCloseAccessPoint()
	// CreateBuilder creates a AccessControlDataCloseAccessPointBuilder
	CreateAccessControlDataCloseAccessPointBuilder() AccessControlDataCloseAccessPointBuilder
}

// _AccessControlDataCloseAccessPoint is the data-structure of this message
type _AccessControlDataCloseAccessPoint struct {
	AccessControlDataContract
}

var _ AccessControlDataCloseAccessPoint = (*_AccessControlDataCloseAccessPoint)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataCloseAccessPoint)(nil)

// NewAccessControlDataCloseAccessPoint factory function for _AccessControlDataCloseAccessPoint
func NewAccessControlDataCloseAccessPoint(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataCloseAccessPoint {
	_result := &_AccessControlDataCloseAccessPoint{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AccessControlDataCloseAccessPointBuilder is a builder for AccessControlDataCloseAccessPoint
type AccessControlDataCloseAccessPointBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AccessControlDataCloseAccessPointBuilder
	// Build builds the AccessControlDataCloseAccessPoint or returns an error if something is wrong
	Build() (AccessControlDataCloseAccessPoint, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AccessControlDataCloseAccessPoint
}

// NewAccessControlDataCloseAccessPointBuilder() creates a AccessControlDataCloseAccessPointBuilder
func NewAccessControlDataCloseAccessPointBuilder() AccessControlDataCloseAccessPointBuilder {
	return &_AccessControlDataCloseAccessPointBuilder{_AccessControlDataCloseAccessPoint: new(_AccessControlDataCloseAccessPoint)}
}

type _AccessControlDataCloseAccessPointBuilder struct {
	*_AccessControlDataCloseAccessPoint

	parentBuilder *_AccessControlDataBuilder

	err *utils.MultiError
}

var _ (AccessControlDataCloseAccessPointBuilder) = (*_AccessControlDataCloseAccessPointBuilder)(nil)

func (b *_AccessControlDataCloseAccessPointBuilder) setParent(contract AccessControlDataContract) {
	b.AccessControlDataContract = contract
}

func (b *_AccessControlDataCloseAccessPointBuilder) WithMandatoryFields() AccessControlDataCloseAccessPointBuilder {
	return b
}

func (b *_AccessControlDataCloseAccessPointBuilder) Build() (AccessControlDataCloseAccessPoint, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AccessControlDataCloseAccessPoint.deepCopy(), nil
}

func (b *_AccessControlDataCloseAccessPointBuilder) MustBuild() AccessControlDataCloseAccessPoint {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AccessControlDataCloseAccessPointBuilder) Done() AccessControlDataBuilder {
	return b.parentBuilder
}

func (b *_AccessControlDataCloseAccessPointBuilder) buildForAccessControlData() (AccessControlData, error) {
	return b.Build()
}

func (b *_AccessControlDataCloseAccessPointBuilder) DeepCopy() any {
	_copy := b.CreateAccessControlDataCloseAccessPointBuilder().(*_AccessControlDataCloseAccessPointBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAccessControlDataCloseAccessPointBuilder creates a AccessControlDataCloseAccessPointBuilder
func (b *_AccessControlDataCloseAccessPoint) CreateAccessControlDataCloseAccessPointBuilder() AccessControlDataCloseAccessPointBuilder {
	if b == nil {
		return NewAccessControlDataCloseAccessPointBuilder()
	}
	return &_AccessControlDataCloseAccessPointBuilder{_AccessControlDataCloseAccessPoint: b.deepCopy()}
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

func (m *_AccessControlDataCloseAccessPoint) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataCloseAccessPoint(structType any) AccessControlDataCloseAccessPoint {
	if casted, ok := structType.(AccessControlDataCloseAccessPoint); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataCloseAccessPoint); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataCloseAccessPoint) GetTypeName() string {
	return "AccessControlDataCloseAccessPoint"
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataCloseAccessPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataCloseAccessPoint AccessControlDataCloseAccessPoint, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataCloseAccessPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataCloseAccessPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataCloseAccessPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataCloseAccessPoint")
	}

	return m, nil
}

func (m *_AccessControlDataCloseAccessPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataCloseAccessPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataCloseAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataCloseAccessPoint")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataCloseAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataCloseAccessPoint")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataCloseAccessPoint) IsAccessControlDataCloseAccessPoint() {}

func (m *_AccessControlDataCloseAccessPoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AccessControlDataCloseAccessPoint) deepCopy() *_AccessControlDataCloseAccessPoint {
	if m == nil {
		return nil
	}
	_AccessControlDataCloseAccessPointCopy := &_AccessControlDataCloseAccessPoint{
		m.AccessControlDataContract.(*_AccessControlData).deepCopy(),
	}
	m.AccessControlDataContract.(*_AccessControlData)._SubType = m
	return _AccessControlDataCloseAccessPointCopy
}

func (m *_AccessControlDataCloseAccessPoint) String() string {
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
