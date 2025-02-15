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

// BACnetEventParameterFloatingLimit is the corresponding interface of BACnetEventParameterFloatingLimit
type BACnetEventParameterFloatingLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetSetpointReference returns SetpointReference (property field)
	GetSetpointReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetLowDiffLimit returns LowDiffLimit (property field)
	GetLowDiffLimit() BACnetContextTagReal
	// GetHighDiffLimit returns HighDiffLimit (property field)
	GetHighDiffLimit() BACnetContextTagReal
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagReal
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterFloatingLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterFloatingLimit()
	// CreateBuilder creates a BACnetEventParameterFloatingLimitBuilder
	CreateBACnetEventParameterFloatingLimitBuilder() BACnetEventParameterFloatingLimitBuilder
}

// _BACnetEventParameterFloatingLimit is the data-structure of this message
type _BACnetEventParameterFloatingLimit struct {
	BACnetEventParameterContract
	OpeningTag        BACnetOpeningTag
	TimeDelay         BACnetContextTagUnsignedInteger
	SetpointReference BACnetDeviceObjectPropertyReferenceEnclosed
	LowDiffLimit      BACnetContextTagReal
	HighDiffLimit     BACnetContextTagReal
	Deadband          BACnetContextTagReal
	ClosingTag        BACnetClosingTag
}

var _ BACnetEventParameterFloatingLimit = (*_BACnetEventParameterFloatingLimit)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterFloatingLimit)(nil)

// NewBACnetEventParameterFloatingLimit factory function for _BACnetEventParameterFloatingLimit
func NewBACnetEventParameterFloatingLimit(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, setpointReference BACnetDeviceObjectPropertyReferenceEnclosed, lowDiffLimit BACnetContextTagReal, highDiffLimit BACnetContextTagReal, deadband BACnetContextTagReal, closingTag BACnetClosingTag) *_BACnetEventParameterFloatingLimit {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterFloatingLimit must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterFloatingLimit must not be nil")
	}
	if setpointReference == nil {
		panic("setpointReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetEventParameterFloatingLimit must not be nil")
	}
	if lowDiffLimit == nil {
		panic("lowDiffLimit of type BACnetContextTagReal for BACnetEventParameterFloatingLimit must not be nil")
	}
	if highDiffLimit == nil {
		panic("highDiffLimit of type BACnetContextTagReal for BACnetEventParameterFloatingLimit must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagReal for BACnetEventParameterFloatingLimit must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterFloatingLimit must not be nil")
	}
	_result := &_BACnetEventParameterFloatingLimit{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		SetpointReference:            setpointReference,
		LowDiffLimit:                 lowDiffLimit,
		HighDiffLimit:                highDiffLimit,
		Deadband:                     deadband,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterFloatingLimitBuilder is a builder for BACnetEventParameterFloatingLimit
type BACnetEventParameterFloatingLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, setpointReference BACnetDeviceObjectPropertyReferenceEnclosed, lowDiffLimit BACnetContextTagReal, highDiffLimit BACnetContextTagReal, deadband BACnetContextTagReal, closingTag BACnetClosingTag) BACnetEventParameterFloatingLimitBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterFloatingLimitBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterFloatingLimitBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterFloatingLimitBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterFloatingLimitBuilder
	// WithSetpointReference adds SetpointReference (property field)
	WithSetpointReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterFloatingLimitBuilder
	// WithSetpointReferenceBuilder adds SetpointReference (property field) which is build by the builder
	WithSetpointReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterFloatingLimitBuilder
	// WithLowDiffLimit adds LowDiffLimit (property field)
	WithLowDiffLimit(BACnetContextTagReal) BACnetEventParameterFloatingLimitBuilder
	// WithLowDiffLimitBuilder adds LowDiffLimit (property field) which is build by the builder
	WithLowDiffLimitBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterFloatingLimitBuilder
	// WithHighDiffLimit adds HighDiffLimit (property field)
	WithHighDiffLimit(BACnetContextTagReal) BACnetEventParameterFloatingLimitBuilder
	// WithHighDiffLimitBuilder adds HighDiffLimit (property field) which is build by the builder
	WithHighDiffLimitBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterFloatingLimitBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetContextTagReal) BACnetEventParameterFloatingLimitBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterFloatingLimitBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterFloatingLimitBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterFloatingLimitBuilder
	// Build builds the BACnetEventParameterFloatingLimit or returns an error if something is wrong
	Build() (BACnetEventParameterFloatingLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterFloatingLimit
}

// NewBACnetEventParameterFloatingLimitBuilder() creates a BACnetEventParameterFloatingLimitBuilder
func NewBACnetEventParameterFloatingLimitBuilder() BACnetEventParameterFloatingLimitBuilder {
	return &_BACnetEventParameterFloatingLimitBuilder{_BACnetEventParameterFloatingLimit: new(_BACnetEventParameterFloatingLimit)}
}

type _BACnetEventParameterFloatingLimitBuilder struct {
	*_BACnetEventParameterFloatingLimit

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterFloatingLimitBuilder) = (*_BACnetEventParameterFloatingLimitBuilder)(nil)

func (b *_BACnetEventParameterFloatingLimitBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, setpointReference BACnetDeviceObjectPropertyReferenceEnclosed, lowDiffLimit BACnetContextTagReal, highDiffLimit BACnetContextTagReal, deadband BACnetContextTagReal, closingTag BACnetClosingTag) BACnetEventParameterFloatingLimitBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithSetpointReference(setpointReference).WithLowDiffLimit(lowDiffLimit).WithHighDiffLimit(highDiffLimit).WithDeadband(deadband).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterFloatingLimitBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterFloatingLimitBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithSetpointReference(setpointReference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterFloatingLimitBuilder {
	b.SetpointReference = setpointReference
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithSetpointReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.SetpointReference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.SetpointReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithLowDiffLimit(lowDiffLimit BACnetContextTagReal) BACnetEventParameterFloatingLimitBuilder {
	b.LowDiffLimit = lowDiffLimit
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithLowDiffLimitBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.LowDiffLimit.CreateBACnetContextTagRealBuilder())
	var err error
	b.LowDiffLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithHighDiffLimit(highDiffLimit BACnetContextTagReal) BACnetEventParameterFloatingLimitBuilder {
	b.HighDiffLimit = highDiffLimit
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithHighDiffLimitBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.HighDiffLimit.CreateBACnetContextTagRealBuilder())
	var err error
	b.HighDiffLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithDeadband(deadband BACnetContextTagReal) BACnetEventParameterFloatingLimitBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithDeadbandBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetContextTagRealBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterFloatingLimitBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterFloatingLimitBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterFloatingLimitBuilder) Build() (BACnetEventParameterFloatingLimit, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.SetpointReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'setpointReference' not set"))
	}
	if b.LowDiffLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowDiffLimit' not set"))
	}
	if b.HighDiffLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highDiffLimit' not set"))
	}
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventParameterFloatingLimit.deepCopy(), nil
}

func (b *_BACnetEventParameterFloatingLimitBuilder) MustBuild() BACnetEventParameterFloatingLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetEventParameterFloatingLimitBuilder) Done() BACnetEventParameterBuilder {
	return b.parentBuilder
}

func (b *_BACnetEventParameterFloatingLimitBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterFloatingLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterFloatingLimitBuilder().(*_BACnetEventParameterFloatingLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterFloatingLimitBuilder creates a BACnetEventParameterFloatingLimitBuilder
func (b *_BACnetEventParameterFloatingLimit) CreateBACnetEventParameterFloatingLimitBuilder() BACnetEventParameterFloatingLimitBuilder {
	if b == nil {
		return NewBACnetEventParameterFloatingLimitBuilder()
	}
	return &_BACnetEventParameterFloatingLimitBuilder{_BACnetEventParameterFloatingLimit: b.deepCopy()}
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

func (m *_BACnetEventParameterFloatingLimit) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterFloatingLimit) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterFloatingLimit) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterFloatingLimit) GetSetpointReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.SetpointReference
}

func (m *_BACnetEventParameterFloatingLimit) GetLowDiffLimit() BACnetContextTagReal {
	return m.LowDiffLimit
}

func (m *_BACnetEventParameterFloatingLimit) GetHighDiffLimit() BACnetContextTagReal {
	return m.HighDiffLimit
}

func (m *_BACnetEventParameterFloatingLimit) GetDeadband() BACnetContextTagReal {
	return m.Deadband
}

func (m *_BACnetEventParameterFloatingLimit) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterFloatingLimit(structType any) BACnetEventParameterFloatingLimit {
	if casted, ok := structType.(BACnetEventParameterFloatingLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterFloatingLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterFloatingLimit) GetTypeName() string {
	return "BACnetEventParameterFloatingLimit"
}

func (m *_BACnetEventParameterFloatingLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (setpointReference)
	lengthInBits += m.SetpointReference.GetLengthInBits(ctx)

	// Simple field (lowDiffLimit)
	lengthInBits += m.LowDiffLimit.GetLengthInBits(ctx)

	// Simple field (highDiffLimit)
	lengthInBits += m.HighDiffLimit.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterFloatingLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterFloatingLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterFloatingLimit BACnetEventParameterFloatingLimit, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterFloatingLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterFloatingLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	setpointReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "setpointReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'setpointReference' field"))
	}
	m.SetpointReference = setpointReference

	lowDiffLimit, err := ReadSimpleField[BACnetContextTagReal](ctx, "lowDiffLimit", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowDiffLimit' field"))
	}
	m.LowDiffLimit = lowDiffLimit

	highDiffLimit, err := ReadSimpleField[BACnetContextTagReal](ctx, "highDiffLimit", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highDiffLimit' field"))
	}
	m.HighDiffLimit = highDiffLimit

	deadband, err := ReadSimpleField[BACnetContextTagReal](ctx, "deadband", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterFloatingLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterFloatingLimit")
	}

	return m, nil
}

func (m *_BACnetEventParameterFloatingLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterFloatingLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterFloatingLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterFloatingLimit")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "setpointReference", m.GetSetpointReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'setpointReference' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "lowDiffLimit", m.GetLowDiffLimit(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowDiffLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "highDiffLimit", m.GetHighDiffLimit(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highDiffLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterFloatingLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterFloatingLimit")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterFloatingLimit) IsBACnetEventParameterFloatingLimit() {}

func (m *_BACnetEventParameterFloatingLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterFloatingLimit) deepCopy() *_BACnetEventParameterFloatingLimit {
	if m == nil {
		return nil
	}
	_BACnetEventParameterFloatingLimitCopy := &_BACnetEventParameterFloatingLimit{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.TimeDelay.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.SetpointReference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.LowDiffLimit.DeepCopy().(BACnetContextTagReal),
		m.HighDiffLimit.DeepCopy().(BACnetContextTagReal),
		m.Deadband.DeepCopy().(BACnetContextTagReal),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterFloatingLimitCopy
}

func (m *_BACnetEventParameterFloatingLimit) String() string {
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
