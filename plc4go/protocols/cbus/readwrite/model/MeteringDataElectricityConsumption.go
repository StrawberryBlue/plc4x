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

// MeteringDataElectricityConsumption is the corresponding interface of MeteringDataElectricityConsumption
type MeteringDataElectricityConsumption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MeteringData
	// GetKWhr returns KWhr (property field)
	GetKWhr() uint32
	// IsMeteringDataElectricityConsumption is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataElectricityConsumption()
	// CreateBuilder creates a MeteringDataElectricityConsumptionBuilder
	CreateMeteringDataElectricityConsumptionBuilder() MeteringDataElectricityConsumptionBuilder
}

// _MeteringDataElectricityConsumption is the data-structure of this message
type _MeteringDataElectricityConsumption struct {
	MeteringDataContract
	KWhr uint32
}

var _ MeteringDataElectricityConsumption = (*_MeteringDataElectricityConsumption)(nil)
var _ MeteringDataRequirements = (*_MeteringDataElectricityConsumption)(nil)

// NewMeteringDataElectricityConsumption factory function for _MeteringDataElectricityConsumption
func NewMeteringDataElectricityConsumption(commandTypeContainer MeteringCommandTypeContainer, argument byte, kWhr uint32) *_MeteringDataElectricityConsumption {
	_result := &_MeteringDataElectricityConsumption{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
		KWhr:                 kWhr,
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MeteringDataElectricityConsumptionBuilder is a builder for MeteringDataElectricityConsumption
type MeteringDataElectricityConsumptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(kWhr uint32) MeteringDataElectricityConsumptionBuilder
	// WithKWhr adds KWhr (property field)
	WithKWhr(uint32) MeteringDataElectricityConsumptionBuilder
	// Build builds the MeteringDataElectricityConsumption or returns an error if something is wrong
	Build() (MeteringDataElectricityConsumption, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MeteringDataElectricityConsumption
}

// NewMeteringDataElectricityConsumptionBuilder() creates a MeteringDataElectricityConsumptionBuilder
func NewMeteringDataElectricityConsumptionBuilder() MeteringDataElectricityConsumptionBuilder {
	return &_MeteringDataElectricityConsumptionBuilder{_MeteringDataElectricityConsumption: new(_MeteringDataElectricityConsumption)}
}

type _MeteringDataElectricityConsumptionBuilder struct {
	*_MeteringDataElectricityConsumption

	parentBuilder *_MeteringDataBuilder

	err *utils.MultiError
}

var _ (MeteringDataElectricityConsumptionBuilder) = (*_MeteringDataElectricityConsumptionBuilder)(nil)

func (b *_MeteringDataElectricityConsumptionBuilder) setParent(contract MeteringDataContract) {
	b.MeteringDataContract = contract
}

func (b *_MeteringDataElectricityConsumptionBuilder) WithMandatoryFields(kWhr uint32) MeteringDataElectricityConsumptionBuilder {
	return b.WithKWhr(kWhr)
}

func (b *_MeteringDataElectricityConsumptionBuilder) WithKWhr(kWhr uint32) MeteringDataElectricityConsumptionBuilder {
	b.KWhr = kWhr
	return b
}

func (b *_MeteringDataElectricityConsumptionBuilder) Build() (MeteringDataElectricityConsumption, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MeteringDataElectricityConsumption.deepCopy(), nil
}

func (b *_MeteringDataElectricityConsumptionBuilder) MustBuild() MeteringDataElectricityConsumption {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MeteringDataElectricityConsumptionBuilder) Done() MeteringDataBuilder {
	return b.parentBuilder
}

func (b *_MeteringDataElectricityConsumptionBuilder) buildForMeteringData() (MeteringData, error) {
	return b.Build()
}

func (b *_MeteringDataElectricityConsumptionBuilder) DeepCopy() any {
	_copy := b.CreateMeteringDataElectricityConsumptionBuilder().(*_MeteringDataElectricityConsumptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMeteringDataElectricityConsumptionBuilder creates a MeteringDataElectricityConsumptionBuilder
func (b *_MeteringDataElectricityConsumption) CreateMeteringDataElectricityConsumptionBuilder() MeteringDataElectricityConsumptionBuilder {
	if b == nil {
		return NewMeteringDataElectricityConsumptionBuilder()
	}
	return &_MeteringDataElectricityConsumptionBuilder{_MeteringDataElectricityConsumption: b.deepCopy()}
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

func (m *_MeteringDataElectricityConsumption) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataElectricityConsumption) GetKWhr() uint32 {
	return m.KWhr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMeteringDataElectricityConsumption(structType any) MeteringDataElectricityConsumption {
	if casted, ok := structType.(MeteringDataElectricityConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataElectricityConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataElectricityConsumption) GetTypeName() string {
	return "MeteringDataElectricityConsumption"
}

func (m *_MeteringDataElectricityConsumption) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	// Simple field (kWhr)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MeteringDataElectricityConsumption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataElectricityConsumption) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataElectricityConsumption MeteringDataElectricityConsumption, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataElectricityConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataElectricityConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	kWhr, err := ReadSimpleField(ctx, "kWhr", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'kWhr' field"))
	}
	m.KWhr = kWhr

	if closeErr := readBuffer.CloseContext("MeteringDataElectricityConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataElectricityConsumption")
	}

	return m, nil
}

func (m *_MeteringDataElectricityConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataElectricityConsumption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataElectricityConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataElectricityConsumption")
		}

		if err := WriteSimpleField[uint32](ctx, "kWhr", m.GetKWhr(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'kWhr' field")
		}

		if popErr := writeBuffer.PopContext("MeteringDataElectricityConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataElectricityConsumption")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataElectricityConsumption) IsMeteringDataElectricityConsumption() {}

func (m *_MeteringDataElectricityConsumption) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataElectricityConsumption) deepCopy() *_MeteringDataElectricityConsumption {
	if m == nil {
		return nil
	}
	_MeteringDataElectricityConsumptionCopy := &_MeteringDataElectricityConsumption{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
		m.KWhr,
	}
	m.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataElectricityConsumptionCopy
}

func (m *_MeteringDataElectricityConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
