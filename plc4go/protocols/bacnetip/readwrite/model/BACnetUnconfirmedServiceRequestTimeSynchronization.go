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

// BACnetUnconfirmedServiceRequestTimeSynchronization is the corresponding interface of BACnetUnconfirmedServiceRequestTimeSynchronization
type BACnetUnconfirmedServiceRequestTimeSynchronization interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetSynchronizedDate returns SynchronizedDate (property field)
	GetSynchronizedDate() BACnetApplicationTagDate
	// GetSynchronizedTime returns SynchronizedTime (property field)
	GetSynchronizedTime() BACnetApplicationTagTime
	// IsBACnetUnconfirmedServiceRequestTimeSynchronization is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestTimeSynchronization()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
	CreateBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder() BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
}

// _BACnetUnconfirmedServiceRequestTimeSynchronization is the data-structure of this message
type _BACnetUnconfirmedServiceRequestTimeSynchronization struct {
	BACnetUnconfirmedServiceRequestContract
	SynchronizedDate BACnetApplicationTagDate
	SynchronizedTime BACnetApplicationTagTime
}

var _ BACnetUnconfirmedServiceRequestTimeSynchronization = (*_BACnetUnconfirmedServiceRequestTimeSynchronization)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestTimeSynchronization)(nil)

// NewBACnetUnconfirmedServiceRequestTimeSynchronization factory function for _BACnetUnconfirmedServiceRequestTimeSynchronization
func NewBACnetUnconfirmedServiceRequestTimeSynchronization(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestTimeSynchronization {
	if synchronizedDate == nil {
		panic("synchronizedDate of type BACnetApplicationTagDate for BACnetUnconfirmedServiceRequestTimeSynchronization must not be nil")
	}
	if synchronizedTime == nil {
		panic("synchronizedTime of type BACnetApplicationTagTime for BACnetUnconfirmedServiceRequestTimeSynchronization must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestTimeSynchronization{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		SynchronizedDate:                        synchronizedDate,
		SynchronizedTime:                        synchronizedTime,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder is a builder for BACnetUnconfirmedServiceRequestTimeSynchronization
type BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
	// WithSynchronizedDate adds SynchronizedDate (property field)
	WithSynchronizedDate(BACnetApplicationTagDate) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
	// WithSynchronizedDateBuilder adds SynchronizedDate (property field) which is build by the builder
	WithSynchronizedDateBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
	// WithSynchronizedTime adds SynchronizedTime (property field)
	WithSynchronizedTime(BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
	// WithSynchronizedTimeBuilder adds SynchronizedTime (property field) which is build by the builder
	WithSynchronizedTimeBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
	// Build builds the BACnetUnconfirmedServiceRequestTimeSynchronization or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestTimeSynchronization, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestTimeSynchronization
}

// NewBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder() creates a BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
func NewBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder() BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	return &_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder{_BACnetUnconfirmedServiceRequestTimeSynchronization: new(_BACnetUnconfirmedServiceRequestTimeSynchronization)}
}

type _BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder struct {
	*_BACnetUnconfirmedServiceRequestTimeSynchronization

	parentBuilder *_BACnetUnconfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) = (*_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder)(nil)

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) setParent(contract BACnetUnconfirmedServiceRequestContract) {
	b.BACnetUnconfirmedServiceRequestContract = contract
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) WithMandatoryFields(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	return b.WithSynchronizedDate(synchronizedDate).WithSynchronizedTime(synchronizedTime)
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) WithSynchronizedDate(synchronizedDate BACnetApplicationTagDate) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	b.SynchronizedDate = synchronizedDate
	return b
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) WithSynchronizedDateBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	builder := builderSupplier(b.SynchronizedDate.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.SynchronizedDate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) WithSynchronizedTime(synchronizedTime BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	b.SynchronizedTime = synchronizedTime
	return b
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) WithSynchronizedTimeBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	builder := builderSupplier(b.SynchronizedTime.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.SynchronizedTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) Build() (BACnetUnconfirmedServiceRequestTimeSynchronization, error) {
	if b.SynchronizedDate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'synchronizedDate' not set"))
	}
	if b.SynchronizedTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'synchronizedTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetUnconfirmedServiceRequestTimeSynchronization.deepCopy(), nil
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) MustBuild() BACnetUnconfirmedServiceRequestTimeSynchronization {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) Done() BACnetUnconfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) buildForBACnetUnconfirmedServiceRequest() (BACnetUnconfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder().(*_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder creates a BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder
func (b *_BACnetUnconfirmedServiceRequestTimeSynchronization) CreateBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder() BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder {
	if b == nil {
		return NewBACnetUnconfirmedServiceRequestTimeSynchronizationBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestTimeSynchronizationBuilder{_BACnetUnconfirmedServiceRequestTimeSynchronization: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetSynchronizedDate() BACnetApplicationTagDate {
	return m.SynchronizedDate
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetSynchronizedTime() BACnetApplicationTagTime {
	return m.SynchronizedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestTimeSynchronization(structType any) BACnetUnconfirmedServiceRequestTimeSynchronization {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestTimeSynchronization"
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.GetLengthInBits(ctx)

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestTimeSynchronization BACnetUnconfirmedServiceRequestTimeSynchronization, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestTimeSynchronization")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	synchronizedDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "synchronizedDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'synchronizedDate' field"))
	}
	m.SynchronizedDate = synchronizedDate

	synchronizedTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "synchronizedTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'synchronizedTime' field"))
	}
	m.SynchronizedTime = synchronizedTime

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestTimeSynchronization")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestTimeSynchronization")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "synchronizedDate", m.GetSynchronizedDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'synchronizedDate' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "synchronizedTime", m.GetSynchronizedTime(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestTimeSynchronization")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) IsBACnetUnconfirmedServiceRequestTimeSynchronization() {
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) deepCopy() *_BACnetUnconfirmedServiceRequestTimeSynchronization {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestTimeSynchronizationCopy := &_BACnetUnconfirmedServiceRequestTimeSynchronization{
		m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).deepCopy(),
		m.SynchronizedDate.DeepCopy().(BACnetApplicationTagDate),
		m.SynchronizedTime.DeepCopy().(BACnetApplicationTagTime),
	}
	m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestTimeSynchronizationCopy
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) String() string {
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
