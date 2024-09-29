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

// Constant values.
const AdsDiscoveryBlockStatus_STATUSLENGTH uint16 = 0x0004

// AdsDiscoveryBlockStatus is the corresponding interface of AdsDiscoveryBlockStatus
type AdsDiscoveryBlockStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetStatus returns Status (property field)
	GetStatus() Status
	// IsAdsDiscoveryBlockStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockStatus()
	// CreateBuilder creates a AdsDiscoveryBlockStatusBuilder
	CreateAdsDiscoveryBlockStatusBuilder() AdsDiscoveryBlockStatusBuilder
}

// _AdsDiscoveryBlockStatus is the data-structure of this message
type _AdsDiscoveryBlockStatus struct {
	AdsDiscoveryBlockContract
	Status Status
}

var _ AdsDiscoveryBlockStatus = (*_AdsDiscoveryBlockStatus)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockStatus)(nil)

// NewAdsDiscoveryBlockStatus factory function for _AdsDiscoveryBlockStatus
func NewAdsDiscoveryBlockStatus(status Status) *_AdsDiscoveryBlockStatus {
	_result := &_AdsDiscoveryBlockStatus{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		Status:                    status,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDiscoveryBlockStatusBuilder is a builder for AdsDiscoveryBlockStatus
type AdsDiscoveryBlockStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(status Status) AdsDiscoveryBlockStatusBuilder
	// WithStatus adds Status (property field)
	WithStatus(Status) AdsDiscoveryBlockStatusBuilder
	// Build builds the AdsDiscoveryBlockStatus or returns an error if something is wrong
	Build() (AdsDiscoveryBlockStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDiscoveryBlockStatus
}

// NewAdsDiscoveryBlockStatusBuilder() creates a AdsDiscoveryBlockStatusBuilder
func NewAdsDiscoveryBlockStatusBuilder() AdsDiscoveryBlockStatusBuilder {
	return &_AdsDiscoveryBlockStatusBuilder{_AdsDiscoveryBlockStatus: new(_AdsDiscoveryBlockStatus)}
}

type _AdsDiscoveryBlockStatusBuilder struct {
	*_AdsDiscoveryBlockStatus

	parentBuilder *_AdsDiscoveryBlockBuilder

	err *utils.MultiError
}

var _ (AdsDiscoveryBlockStatusBuilder) = (*_AdsDiscoveryBlockStatusBuilder)(nil)

func (b *_AdsDiscoveryBlockStatusBuilder) setParent(contract AdsDiscoveryBlockContract) {
	b.AdsDiscoveryBlockContract = contract
}

func (b *_AdsDiscoveryBlockStatusBuilder) WithMandatoryFields(status Status) AdsDiscoveryBlockStatusBuilder {
	return b.WithStatus(status)
}

func (b *_AdsDiscoveryBlockStatusBuilder) WithStatus(status Status) AdsDiscoveryBlockStatusBuilder {
	b.Status = status
	return b
}

func (b *_AdsDiscoveryBlockStatusBuilder) Build() (AdsDiscoveryBlockStatus, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDiscoveryBlockStatus.deepCopy(), nil
}

func (b *_AdsDiscoveryBlockStatusBuilder) MustBuild() AdsDiscoveryBlockStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AdsDiscoveryBlockStatusBuilder) Done() AdsDiscoveryBlockBuilder {
	return b.parentBuilder
}

func (b *_AdsDiscoveryBlockStatusBuilder) buildForAdsDiscoveryBlock() (AdsDiscoveryBlock, error) {
	return b.Build()
}

func (b *_AdsDiscoveryBlockStatusBuilder) DeepCopy() any {
	_copy := b.CreateAdsDiscoveryBlockStatusBuilder().(*_AdsDiscoveryBlockStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDiscoveryBlockStatusBuilder creates a AdsDiscoveryBlockStatusBuilder
func (b *_AdsDiscoveryBlockStatus) CreateAdsDiscoveryBlockStatusBuilder() AdsDiscoveryBlockStatusBuilder {
	if b == nil {
		return NewAdsDiscoveryBlockStatusBuilder()
	}
	return &_AdsDiscoveryBlockStatusBuilder{_AdsDiscoveryBlockStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockStatus) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockStatus) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockStatus) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDiscoveryBlockStatus) GetStatusLength() uint16 {
	return AdsDiscoveryBlockStatus_STATUSLENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockStatus(structType any) AdsDiscoveryBlockStatus {
	if casted, ok := structType.(AdsDiscoveryBlockStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockStatus) GetTypeName() string {
	return "AdsDiscoveryBlockStatus"
}

func (m *_AdsDiscoveryBlockStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Const Field (statusLength)
	lengthInBits += 16

	// Simple field (status)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDiscoveryBlockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockStatus AdsDiscoveryBlockStatus, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusLength, err := ReadConstField[uint16](ctx, "statusLength", ReadUnsignedShort(readBuffer, uint8(16)), AdsDiscoveryBlockStatus_STATUSLENGTH)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusLength' field"))
	}
	_ = statusLength

	status, err := ReadEnumField[Status](ctx, "status", "Status", ReadEnum(StatusByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockStatus")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockStatus")
		}

		if err := WriteConstField(ctx, "statusLength", AdsDiscoveryBlockStatus_STATUSLENGTH, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusLength' field")
		}

		if err := WriteSimpleEnumField[Status](ctx, "status", "Status", m.GetStatus(), WriteEnum[Status, uint32](Status.GetValue, Status.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockStatus")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockStatus) IsAdsDiscoveryBlockStatus() {}

func (m *_AdsDiscoveryBlockStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockStatus) deepCopy() *_AdsDiscoveryBlockStatus {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockStatusCopy := &_AdsDiscoveryBlockStatus{
		m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).deepCopy(),
		m.Status,
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockStatusCopy
}

func (m *_AdsDiscoveryBlockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
