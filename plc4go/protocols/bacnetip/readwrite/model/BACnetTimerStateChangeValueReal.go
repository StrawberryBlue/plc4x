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

// BACnetTimerStateChangeValueReal is the corresponding interface of BACnetTimerStateChangeValueReal
type BACnetTimerStateChangeValueReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// IsBACnetTimerStateChangeValueReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueReal()
	// CreateBuilder creates a BACnetTimerStateChangeValueRealBuilder
	CreateBACnetTimerStateChangeValueRealBuilder() BACnetTimerStateChangeValueRealBuilder
}

// _BACnetTimerStateChangeValueReal is the data-structure of this message
type _BACnetTimerStateChangeValueReal struct {
	BACnetTimerStateChangeValueContract
	RealValue BACnetApplicationTagReal
}

var _ BACnetTimerStateChangeValueReal = (*_BACnetTimerStateChangeValueReal)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueReal)(nil)

// NewBACnetTimerStateChangeValueReal factory function for _BACnetTimerStateChangeValueReal
func NewBACnetTimerStateChangeValueReal(peekedTagHeader BACnetTagHeader, realValue BACnetApplicationTagReal, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueReal {
	if realValue == nil {
		panic("realValue of type BACnetApplicationTagReal for BACnetTimerStateChangeValueReal must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueReal{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		RealValue:                           realValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueRealBuilder is a builder for BACnetTimerStateChangeValueReal
type BACnetTimerStateChangeValueRealBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetTimerStateChangeValueRealBuilder
	// WithRealValue adds RealValue (property field)
	WithRealValue(BACnetApplicationTagReal) BACnetTimerStateChangeValueRealBuilder
	// WithRealValueBuilder adds RealValue (property field) which is build by the builder
	WithRealValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetTimerStateChangeValueRealBuilder
	// Build builds the BACnetTimerStateChangeValueReal or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueReal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueReal
}

// NewBACnetTimerStateChangeValueRealBuilder() creates a BACnetTimerStateChangeValueRealBuilder
func NewBACnetTimerStateChangeValueRealBuilder() BACnetTimerStateChangeValueRealBuilder {
	return &_BACnetTimerStateChangeValueRealBuilder{_BACnetTimerStateChangeValueReal: new(_BACnetTimerStateChangeValueReal)}
}

type _BACnetTimerStateChangeValueRealBuilder struct {
	*_BACnetTimerStateChangeValueReal

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueRealBuilder) = (*_BACnetTimerStateChangeValueRealBuilder)(nil)

func (b *_BACnetTimerStateChangeValueRealBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
}

func (b *_BACnetTimerStateChangeValueRealBuilder) WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetTimerStateChangeValueRealBuilder {
	return b.WithRealValue(realValue)
}

func (b *_BACnetTimerStateChangeValueRealBuilder) WithRealValue(realValue BACnetApplicationTagReal) BACnetTimerStateChangeValueRealBuilder {
	b.RealValue = realValue
	return b
}

func (b *_BACnetTimerStateChangeValueRealBuilder) WithRealValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetTimerStateChangeValueRealBuilder {
	builder := builderSupplier(b.RealValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.RealValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueRealBuilder) Build() (BACnetTimerStateChangeValueReal, error) {
	if b.RealValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'realValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueReal.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueRealBuilder) MustBuild() BACnetTimerStateChangeValueReal {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetTimerStateChangeValueRealBuilder) Done() BACnetTimerStateChangeValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueRealBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueRealBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueRealBuilder().(*_BACnetTimerStateChangeValueRealBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueRealBuilder creates a BACnetTimerStateChangeValueRealBuilder
func (b *_BACnetTimerStateChangeValueReal) CreateBACnetTimerStateChangeValueRealBuilder() BACnetTimerStateChangeValueRealBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueRealBuilder()
	}
	return &_BACnetTimerStateChangeValueRealBuilder{_BACnetTimerStateChangeValueReal: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueReal) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueReal(structType any) BACnetTimerStateChangeValueReal {
	if casted, ok := structType.(BACnetTimerStateChangeValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueReal) GetTypeName() string {
	return "BACnetTimerStateChangeValueReal"
}

func (m *_BACnetTimerStateChangeValueReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueReal BACnetTimerStateChangeValueReal, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	m.RealValue = realValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueReal")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueReal")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueReal")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueReal) IsBACnetTimerStateChangeValueReal() {}

func (m *_BACnetTimerStateChangeValueReal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueReal) deepCopy() *_BACnetTimerStateChangeValueReal {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueRealCopy := &_BACnetTimerStateChangeValueReal{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		m.RealValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueRealCopy
}

func (m *_BACnetTimerStateChangeValueReal) String() string {
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
