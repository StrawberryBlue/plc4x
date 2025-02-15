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

// BACnetPropertyStatesLightningOperation is the corresponding interface of BACnetPropertyStatesLightningOperation
type BACnetPropertyStatesLightningOperation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLightningOperation returns LightningOperation (property field)
	GetLightningOperation() BACnetLightingOperationTagged
	// IsBACnetPropertyStatesLightningOperation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLightningOperation()
	// CreateBuilder creates a BACnetPropertyStatesLightningOperationBuilder
	CreateBACnetPropertyStatesLightningOperationBuilder() BACnetPropertyStatesLightningOperationBuilder
}

// _BACnetPropertyStatesLightningOperation is the data-structure of this message
type _BACnetPropertyStatesLightningOperation struct {
	BACnetPropertyStatesContract
	LightningOperation BACnetLightingOperationTagged
}

var _ BACnetPropertyStatesLightningOperation = (*_BACnetPropertyStatesLightningOperation)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLightningOperation)(nil)

// NewBACnetPropertyStatesLightningOperation factory function for _BACnetPropertyStatesLightningOperation
func NewBACnetPropertyStatesLightningOperation(peekedTagHeader BACnetTagHeader, lightningOperation BACnetLightingOperationTagged) *_BACnetPropertyStatesLightningOperation {
	if lightningOperation == nil {
		panic("lightningOperation of type BACnetLightingOperationTagged for BACnetPropertyStatesLightningOperation must not be nil")
	}
	_result := &_BACnetPropertyStatesLightningOperation{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LightningOperation:           lightningOperation,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLightningOperationBuilder is a builder for BACnetPropertyStatesLightningOperation
type BACnetPropertyStatesLightningOperationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lightningOperation BACnetLightingOperationTagged) BACnetPropertyStatesLightningOperationBuilder
	// WithLightningOperation adds LightningOperation (property field)
	WithLightningOperation(BACnetLightingOperationTagged) BACnetPropertyStatesLightningOperationBuilder
	// WithLightningOperationBuilder adds LightningOperation (property field) which is build by the builder
	WithLightningOperationBuilder(func(BACnetLightingOperationTaggedBuilder) BACnetLightingOperationTaggedBuilder) BACnetPropertyStatesLightningOperationBuilder
	// Build builds the BACnetPropertyStatesLightningOperation or returns an error if something is wrong
	Build() (BACnetPropertyStatesLightningOperation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLightningOperation
}

// NewBACnetPropertyStatesLightningOperationBuilder() creates a BACnetPropertyStatesLightningOperationBuilder
func NewBACnetPropertyStatesLightningOperationBuilder() BACnetPropertyStatesLightningOperationBuilder {
	return &_BACnetPropertyStatesLightningOperationBuilder{_BACnetPropertyStatesLightningOperation: new(_BACnetPropertyStatesLightningOperation)}
}

type _BACnetPropertyStatesLightningOperationBuilder struct {
	*_BACnetPropertyStatesLightningOperation

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLightningOperationBuilder) = (*_BACnetPropertyStatesLightningOperationBuilder)(nil)

func (b *_BACnetPropertyStatesLightningOperationBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) WithMandatoryFields(lightningOperation BACnetLightingOperationTagged) BACnetPropertyStatesLightningOperationBuilder {
	return b.WithLightningOperation(lightningOperation)
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) WithLightningOperation(lightningOperation BACnetLightingOperationTagged) BACnetPropertyStatesLightningOperationBuilder {
	b.LightningOperation = lightningOperation
	return b
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) WithLightningOperationBuilder(builderSupplier func(BACnetLightingOperationTaggedBuilder) BACnetLightingOperationTaggedBuilder) BACnetPropertyStatesLightningOperationBuilder {
	builder := builderSupplier(b.LightningOperation.CreateBACnetLightingOperationTaggedBuilder())
	var err error
	b.LightningOperation, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingOperationTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) Build() (BACnetPropertyStatesLightningOperation, error) {
	if b.LightningOperation == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lightningOperation' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesLightningOperation.deepCopy(), nil
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) MustBuild() BACnetPropertyStatesLightningOperation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesLightningOperationBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesLightningOperationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesLightningOperationBuilder().(*_BACnetPropertyStatesLightningOperationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesLightningOperationBuilder creates a BACnetPropertyStatesLightningOperationBuilder
func (b *_BACnetPropertyStatesLightningOperation) CreateBACnetPropertyStatesLightningOperationBuilder() BACnetPropertyStatesLightningOperationBuilder {
	if b == nil {
		return NewBACnetPropertyStatesLightningOperationBuilder()
	}
	return &_BACnetPropertyStatesLightningOperationBuilder{_BACnetPropertyStatesLightningOperation: b.deepCopy()}
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

func (m *_BACnetPropertyStatesLightningOperation) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningOperation) GetLightningOperation() BACnetLightingOperationTagged {
	return m.LightningOperation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningOperation(structType any) BACnetPropertyStatesLightningOperation {
	if casted, ok := structType.(BACnetPropertyStatesLightningOperation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningOperation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningOperation) GetTypeName() string {
	return "BACnetPropertyStatesLightningOperation"
}

func (m *_BACnetPropertyStatesLightningOperation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lightningOperation)
	lengthInBits += m.LightningOperation.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningOperation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLightningOperation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLightningOperation BACnetPropertyStatesLightningOperation, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningOperation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningOperation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningOperation, err := ReadSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", ReadComplex[BACnetLightingOperationTagged](BACnetLightingOperationTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningOperation' field"))
	}
	m.LightningOperation = lightningOperation

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningOperation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningOperation")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLightningOperation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningOperation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningOperation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningOperation")
		}

		if err := WriteSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", m.GetLightningOperation(), WriteComplex[BACnetLightingOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightningOperation' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningOperation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningOperation")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningOperation) IsBACnetPropertyStatesLightningOperation() {}

func (m *_BACnetPropertyStatesLightningOperation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLightningOperation) deepCopy() *_BACnetPropertyStatesLightningOperation {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLightningOperationCopy := &_BACnetPropertyStatesLightningOperation{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.LightningOperation.DeepCopy().(BACnetLightingOperationTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLightningOperationCopy
}

func (m *_BACnetPropertyStatesLightningOperation) String() string {
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
