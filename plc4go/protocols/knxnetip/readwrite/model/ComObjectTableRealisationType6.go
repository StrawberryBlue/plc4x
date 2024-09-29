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

// ComObjectTableRealisationType6 is the corresponding interface of ComObjectTableRealisationType6
type ComObjectTableRealisationType6 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ComObjectTable
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() GroupObjectDescriptorRealisationType6
	// IsComObjectTableRealisationType6 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsComObjectTableRealisationType6()
	// CreateBuilder creates a ComObjectTableRealisationType6Builder
	CreateComObjectTableRealisationType6Builder() ComObjectTableRealisationType6Builder
}

// _ComObjectTableRealisationType6 is the data-structure of this message
type _ComObjectTableRealisationType6 struct {
	ComObjectTableContract
	ComObjectDescriptors GroupObjectDescriptorRealisationType6
}

var _ ComObjectTableRealisationType6 = (*_ComObjectTableRealisationType6)(nil)
var _ ComObjectTableRequirements = (*_ComObjectTableRealisationType6)(nil)

// NewComObjectTableRealisationType6 factory function for _ComObjectTableRealisationType6
func NewComObjectTableRealisationType6(comObjectDescriptors GroupObjectDescriptorRealisationType6) *_ComObjectTableRealisationType6 {
	if comObjectDescriptors == nil {
		panic("comObjectDescriptors of type GroupObjectDescriptorRealisationType6 for ComObjectTableRealisationType6 must not be nil")
	}
	_result := &_ComObjectTableRealisationType6{
		ComObjectTableContract: NewComObjectTable(),
		ComObjectDescriptors:   comObjectDescriptors,
	}
	_result.ComObjectTableContract.(*_ComObjectTable)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ComObjectTableRealisationType6Builder is a builder for ComObjectTableRealisationType6
type ComObjectTableRealisationType6Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(comObjectDescriptors GroupObjectDescriptorRealisationType6) ComObjectTableRealisationType6Builder
	// WithComObjectDescriptors adds ComObjectDescriptors (property field)
	WithComObjectDescriptors(GroupObjectDescriptorRealisationType6) ComObjectTableRealisationType6Builder
	// WithComObjectDescriptorsBuilder adds ComObjectDescriptors (property field) which is build by the builder
	WithComObjectDescriptorsBuilder(func(GroupObjectDescriptorRealisationType6Builder) GroupObjectDescriptorRealisationType6Builder) ComObjectTableRealisationType6Builder
	// Build builds the ComObjectTableRealisationType6 or returns an error if something is wrong
	Build() (ComObjectTableRealisationType6, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ComObjectTableRealisationType6
}

// NewComObjectTableRealisationType6Builder() creates a ComObjectTableRealisationType6Builder
func NewComObjectTableRealisationType6Builder() ComObjectTableRealisationType6Builder {
	return &_ComObjectTableRealisationType6Builder{_ComObjectTableRealisationType6: new(_ComObjectTableRealisationType6)}
}

type _ComObjectTableRealisationType6Builder struct {
	*_ComObjectTableRealisationType6

	parentBuilder *_ComObjectTableBuilder

	err *utils.MultiError
}

var _ (ComObjectTableRealisationType6Builder) = (*_ComObjectTableRealisationType6Builder)(nil)

func (b *_ComObjectTableRealisationType6Builder) setParent(contract ComObjectTableContract) {
	b.ComObjectTableContract = contract
}

func (b *_ComObjectTableRealisationType6Builder) WithMandatoryFields(comObjectDescriptors GroupObjectDescriptorRealisationType6) ComObjectTableRealisationType6Builder {
	return b.WithComObjectDescriptors(comObjectDescriptors)
}

func (b *_ComObjectTableRealisationType6Builder) WithComObjectDescriptors(comObjectDescriptors GroupObjectDescriptorRealisationType6) ComObjectTableRealisationType6Builder {
	b.ComObjectDescriptors = comObjectDescriptors
	return b
}

func (b *_ComObjectTableRealisationType6Builder) WithComObjectDescriptorsBuilder(builderSupplier func(GroupObjectDescriptorRealisationType6Builder) GroupObjectDescriptorRealisationType6Builder) ComObjectTableRealisationType6Builder {
	builder := builderSupplier(b.ComObjectDescriptors.CreateGroupObjectDescriptorRealisationType6Builder())
	var err error
	b.ComObjectDescriptors, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "GroupObjectDescriptorRealisationType6Builder failed"))
	}
	return b
}

func (b *_ComObjectTableRealisationType6Builder) Build() (ComObjectTableRealisationType6, error) {
	if b.ComObjectDescriptors == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'comObjectDescriptors' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ComObjectTableRealisationType6.deepCopy(), nil
}

func (b *_ComObjectTableRealisationType6Builder) MustBuild() ComObjectTableRealisationType6 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ComObjectTableRealisationType6Builder) Done() ComObjectTableBuilder {
	return b.parentBuilder
}

func (b *_ComObjectTableRealisationType6Builder) buildForComObjectTable() (ComObjectTable, error) {
	return b.Build()
}

func (b *_ComObjectTableRealisationType6Builder) DeepCopy() any {
	_copy := b.CreateComObjectTableRealisationType6Builder().(*_ComObjectTableRealisationType6Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateComObjectTableRealisationType6Builder creates a ComObjectTableRealisationType6Builder
func (b *_ComObjectTableRealisationType6) CreateComObjectTableRealisationType6Builder() ComObjectTableRealisationType6Builder {
	if b == nil {
		return NewComObjectTableRealisationType6Builder()
	}
	return &_ComObjectTableRealisationType6Builder{_ComObjectTableRealisationType6: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType6) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType6) GetParent() ComObjectTableContract {
	return m.ComObjectTableContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType6) GetComObjectDescriptors() GroupObjectDescriptorRealisationType6 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType6(structType any) ComObjectTableRealisationType6 {
	if casted, ok := structType.(ComObjectTableRealisationType6); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType6); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType6) GetTypeName() string {
	return "ComObjectTableRealisationType6"
}

func (m *_ComObjectTableRealisationType6) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ComObjectTableContract.(*_ComObjectTable).getLengthInBits(ctx))

	// Simple field (comObjectDescriptors)
	lengthInBits += m.ComObjectDescriptors.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ComObjectTableRealisationType6) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ComObjectTableRealisationType6) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ComObjectTable, firmwareType FirmwareType) (__comObjectTableRealisationType6 ComObjectTableRealisationType6, err error) {
	m.ComObjectTableContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType6"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType6")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	comObjectDescriptors, err := ReadSimpleField[GroupObjectDescriptorRealisationType6](ctx, "comObjectDescriptors", ReadComplex[GroupObjectDescriptorRealisationType6](GroupObjectDescriptorRealisationType6ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comObjectDescriptors' field"))
	}
	m.ComObjectDescriptors = comObjectDescriptors

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType6"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType6")
	}

	return m, nil
}

func (m *_ComObjectTableRealisationType6) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComObjectTableRealisationType6) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType6"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType6")
		}

		if err := WriteSimpleField[GroupObjectDescriptorRealisationType6](ctx, "comObjectDescriptors", m.GetComObjectDescriptors(), WriteComplex[GroupObjectDescriptorRealisationType6](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType6"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType6")
		}
		return nil
	}
	return m.ComObjectTableContract.(*_ComObjectTable).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType6) IsComObjectTableRealisationType6() {}

func (m *_ComObjectTableRealisationType6) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ComObjectTableRealisationType6) deepCopy() *_ComObjectTableRealisationType6 {
	if m == nil {
		return nil
	}
	_ComObjectTableRealisationType6Copy := &_ComObjectTableRealisationType6{
		m.ComObjectTableContract.(*_ComObjectTable).deepCopy(),
		m.ComObjectDescriptors.DeepCopy().(GroupObjectDescriptorRealisationType6),
	}
	m.ComObjectTableContract.(*_ComObjectTable)._SubType = m
	return _ComObjectTableRealisationType6Copy
}

func (m *_ComObjectTableRealisationType6) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
