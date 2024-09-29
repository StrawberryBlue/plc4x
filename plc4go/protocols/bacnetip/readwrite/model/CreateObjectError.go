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

// CreateObjectError is the corresponding interface of CreateObjectError
type CreateObjectError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger
	// IsCreateObjectError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateObjectError()
	// CreateBuilder creates a CreateObjectErrorBuilder
	CreateCreateObjectErrorBuilder() CreateObjectErrorBuilder
}

// _CreateObjectError is the data-structure of this message
type _CreateObjectError struct {
	BACnetErrorContract
	ErrorType                ErrorEnclosed
	FirstFailedElementNumber BACnetContextTagUnsignedInteger
}

var _ CreateObjectError = (*_CreateObjectError)(nil)
var _ BACnetErrorRequirements = (*_CreateObjectError)(nil)

// NewCreateObjectError factory function for _CreateObjectError
func NewCreateObjectError(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) *_CreateObjectError {
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for CreateObjectError must not be nil")
	}
	if firstFailedElementNumber == nil {
		panic("firstFailedElementNumber of type BACnetContextTagUnsignedInteger for CreateObjectError must not be nil")
	}
	_result := &_CreateObjectError{
		BACnetErrorContract:      NewBACnetError(),
		ErrorType:                errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CreateObjectErrorBuilder is a builder for CreateObjectError
type CreateObjectErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) CreateObjectErrorBuilder
	// WithErrorType adds ErrorType (property field)
	WithErrorType(ErrorEnclosed) CreateObjectErrorBuilder
	// WithErrorTypeBuilder adds ErrorType (property field) which is build by the builder
	WithErrorTypeBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) CreateObjectErrorBuilder
	// WithFirstFailedElementNumber adds FirstFailedElementNumber (property field)
	WithFirstFailedElementNumber(BACnetContextTagUnsignedInteger) CreateObjectErrorBuilder
	// WithFirstFailedElementNumberBuilder adds FirstFailedElementNumber (property field) which is build by the builder
	WithFirstFailedElementNumberBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) CreateObjectErrorBuilder
	// Build builds the CreateObjectError or returns an error if something is wrong
	Build() (CreateObjectError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CreateObjectError
}

// NewCreateObjectErrorBuilder() creates a CreateObjectErrorBuilder
func NewCreateObjectErrorBuilder() CreateObjectErrorBuilder {
	return &_CreateObjectErrorBuilder{_CreateObjectError: new(_CreateObjectError)}
}

type _CreateObjectErrorBuilder struct {
	*_CreateObjectError

	parentBuilder *_BACnetErrorBuilder

	err *utils.MultiError
}

var _ (CreateObjectErrorBuilder) = (*_CreateObjectErrorBuilder)(nil)

func (b *_CreateObjectErrorBuilder) setParent(contract BACnetErrorContract) {
	b.BACnetErrorContract = contract
}

func (b *_CreateObjectErrorBuilder) WithMandatoryFields(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) CreateObjectErrorBuilder {
	return b.WithErrorType(errorType).WithFirstFailedElementNumber(firstFailedElementNumber)
}

func (b *_CreateObjectErrorBuilder) WithErrorType(errorType ErrorEnclosed) CreateObjectErrorBuilder {
	b.ErrorType = errorType
	return b
}

func (b *_CreateObjectErrorBuilder) WithErrorTypeBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) CreateObjectErrorBuilder {
	builder := builderSupplier(b.ErrorType.CreateErrorEnclosedBuilder())
	var err error
	b.ErrorType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_CreateObjectErrorBuilder) WithFirstFailedElementNumber(firstFailedElementNumber BACnetContextTagUnsignedInteger) CreateObjectErrorBuilder {
	b.FirstFailedElementNumber = firstFailedElementNumber
	return b
}

func (b *_CreateObjectErrorBuilder) WithFirstFailedElementNumberBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) CreateObjectErrorBuilder {
	builder := builderSupplier(b.FirstFailedElementNumber.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.FirstFailedElementNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_CreateObjectErrorBuilder) Build() (CreateObjectError, error) {
	if b.ErrorType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorType' not set"))
	}
	if b.FirstFailedElementNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'firstFailedElementNumber' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CreateObjectError.deepCopy(), nil
}

func (b *_CreateObjectErrorBuilder) MustBuild() CreateObjectError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CreateObjectErrorBuilder) Done() BACnetErrorBuilder {
	return b.parentBuilder
}

func (b *_CreateObjectErrorBuilder) buildForBACnetError() (BACnetError, error) {
	return b.Build()
}

func (b *_CreateObjectErrorBuilder) DeepCopy() any {
	_copy := b.CreateCreateObjectErrorBuilder().(*_CreateObjectErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCreateObjectErrorBuilder creates a CreateObjectErrorBuilder
func (b *_CreateObjectError) CreateCreateObjectErrorBuilder() CreateObjectErrorBuilder {
	if b == nil {
		return NewCreateObjectErrorBuilder()
	}
	return &_CreateObjectErrorBuilder{_CreateObjectError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateObjectError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateObjectError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateObjectError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_CreateObjectError) GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCreateObjectError(structType any) CreateObjectError {
	if casted, ok := structType.(CreateObjectError); ok {
		return casted
	}
	if casted, ok := structType.(*CreateObjectError); ok {
		return *casted
	}
	return nil
}

func (m *_CreateObjectError) GetTypeName() string {
	return "CreateObjectError"
}

func (m *_CreateObjectError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CreateObjectError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateObjectError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__createObjectError CreateObjectError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateObjectError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateObjectError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	firstFailedElementNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstFailedElementNumber' field"))
	}
	m.FirstFailedElementNumber = firstFailedElementNumber

	if closeErr := readBuffer.CloseContext("CreateObjectError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateObjectError")
	}

	return m, nil
}

func (m *_CreateObjectError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateObjectError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateObjectError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateObjectError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", m.GetFirstFailedElementNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'firstFailedElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("CreateObjectError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateObjectError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateObjectError) IsCreateObjectError() {}

func (m *_CreateObjectError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CreateObjectError) deepCopy() *_CreateObjectError {
	if m == nil {
		return nil
	}
	_CreateObjectErrorCopy := &_CreateObjectError{
		m.BACnetErrorContract.(*_BACnetError).deepCopy(),
		m.ErrorType.DeepCopy().(ErrorEnclosed),
		m.FirstFailedElementNumber.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _CreateObjectErrorCopy
}

func (m *_CreateObjectError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
