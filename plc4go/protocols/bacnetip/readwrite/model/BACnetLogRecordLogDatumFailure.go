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

// BACnetLogRecordLogDatumFailure is the corresponding interface of BACnetLogRecordLogDatumFailure
type BACnetLogRecordLogDatumFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetFailure returns Failure (property field)
	GetFailure() ErrorEnclosed
	// IsBACnetLogRecordLogDatumFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumFailure()
	// CreateBuilder creates a BACnetLogRecordLogDatumFailureBuilder
	CreateBACnetLogRecordLogDatumFailureBuilder() BACnetLogRecordLogDatumFailureBuilder
}

// _BACnetLogRecordLogDatumFailure is the data-structure of this message
type _BACnetLogRecordLogDatumFailure struct {
	BACnetLogRecordLogDatumContract
	Failure ErrorEnclosed
}

var _ BACnetLogRecordLogDatumFailure = (*_BACnetLogRecordLogDatumFailure)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumFailure)(nil)

// NewBACnetLogRecordLogDatumFailure factory function for _BACnetLogRecordLogDatumFailure
func NewBACnetLogRecordLogDatumFailure(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, failure ErrorEnclosed, tagNumber uint8) *_BACnetLogRecordLogDatumFailure {
	if failure == nil {
		panic("failure of type ErrorEnclosed for BACnetLogRecordLogDatumFailure must not be nil")
	}
	_result := &_BACnetLogRecordLogDatumFailure{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		Failure:                         failure,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogRecordLogDatumFailureBuilder is a builder for BACnetLogRecordLogDatumFailure
type BACnetLogRecordLogDatumFailureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(failure ErrorEnclosed) BACnetLogRecordLogDatumFailureBuilder
	// WithFailure adds Failure (property field)
	WithFailure(ErrorEnclosed) BACnetLogRecordLogDatumFailureBuilder
	// WithFailureBuilder adds Failure (property field) which is build by the builder
	WithFailureBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) BACnetLogRecordLogDatumFailureBuilder
	// Build builds the BACnetLogRecordLogDatumFailure or returns an error if something is wrong
	Build() (BACnetLogRecordLogDatumFailure, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogRecordLogDatumFailure
}

// NewBACnetLogRecordLogDatumFailureBuilder() creates a BACnetLogRecordLogDatumFailureBuilder
func NewBACnetLogRecordLogDatumFailureBuilder() BACnetLogRecordLogDatumFailureBuilder {
	return &_BACnetLogRecordLogDatumFailureBuilder{_BACnetLogRecordLogDatumFailure: new(_BACnetLogRecordLogDatumFailure)}
}

type _BACnetLogRecordLogDatumFailureBuilder struct {
	*_BACnetLogRecordLogDatumFailure

	parentBuilder *_BACnetLogRecordLogDatumBuilder

	err *utils.MultiError
}

var _ (BACnetLogRecordLogDatumFailureBuilder) = (*_BACnetLogRecordLogDatumFailureBuilder)(nil)

func (b *_BACnetLogRecordLogDatumFailureBuilder) setParent(contract BACnetLogRecordLogDatumContract) {
	b.BACnetLogRecordLogDatumContract = contract
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) WithMandatoryFields(failure ErrorEnclosed) BACnetLogRecordLogDatumFailureBuilder {
	return b.WithFailure(failure)
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) WithFailure(failure ErrorEnclosed) BACnetLogRecordLogDatumFailureBuilder {
	b.Failure = failure
	return b
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) WithFailureBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) BACnetLogRecordLogDatumFailureBuilder {
	builder := builderSupplier(b.Failure.CreateErrorEnclosedBuilder())
	var err error
	b.Failure, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) Build() (BACnetLogRecordLogDatumFailure, error) {
	if b.Failure == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'failure' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogRecordLogDatumFailure.deepCopy(), nil
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) MustBuild() BACnetLogRecordLogDatumFailure {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetLogRecordLogDatumFailureBuilder) Done() BACnetLogRecordLogDatumBuilder {
	return b.parentBuilder
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) buildForBACnetLogRecordLogDatum() (BACnetLogRecordLogDatum, error) {
	return b.Build()
}

func (b *_BACnetLogRecordLogDatumFailureBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogRecordLogDatumFailureBuilder().(*_BACnetLogRecordLogDatumFailureBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogRecordLogDatumFailureBuilder creates a BACnetLogRecordLogDatumFailureBuilder
func (b *_BACnetLogRecordLogDatumFailure) CreateBACnetLogRecordLogDatumFailureBuilder() BACnetLogRecordLogDatumFailureBuilder {
	if b == nil {
		return NewBACnetLogRecordLogDatumFailureBuilder()
	}
	return &_BACnetLogRecordLogDatumFailureBuilder{_BACnetLogRecordLogDatumFailure: b.deepCopy()}
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

func (m *_BACnetLogRecordLogDatumFailure) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumFailure) GetFailure() ErrorEnclosed {
	return m.Failure
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumFailure(structType any) BACnetLogRecordLogDatumFailure {
	if casted, ok := structType.(BACnetLogRecordLogDatumFailure); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumFailure); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumFailure) GetTypeName() string {
	return "BACnetLogRecordLogDatumFailure"
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (failure)
	lengthInBits += m.Failure.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumFailure BACnetLogRecordLogDatumFailure, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	failure, err := ReadSimpleField[ErrorEnclosed](ctx, "failure", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(8))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'failure' field"))
	}
	m.Failure = failure

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumFailure")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumFailure")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "failure", m.GetFailure(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'failure' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumFailure")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumFailure) IsBACnetLogRecordLogDatumFailure() {}

func (m *_BACnetLogRecordLogDatumFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumFailure) deepCopy() *_BACnetLogRecordLogDatumFailure {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumFailureCopy := &_BACnetLogRecordLogDatumFailure{
		m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).deepCopy(),
		m.Failure.DeepCopy().(ErrorEnclosed),
	}
	m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumFailureCopy
}

func (m *_BACnetLogRecordLogDatumFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
