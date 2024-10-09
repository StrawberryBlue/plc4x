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

// SALDataIrrigationControl is the corresponding interface of SALDataIrrigationControl
type SALDataIrrigationControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetIrrigationControlData returns IrrigationControlData (property field)
	GetIrrigationControlData() LightingData
	// IsSALDataIrrigationControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataIrrigationControl()
	// CreateBuilder creates a SALDataIrrigationControlBuilder
	CreateSALDataIrrigationControlBuilder() SALDataIrrigationControlBuilder
}

// _SALDataIrrigationControl is the data-structure of this message
type _SALDataIrrigationControl struct {
	SALDataContract
	IrrigationControlData LightingData
}

var _ SALDataIrrigationControl = (*_SALDataIrrigationControl)(nil)
var _ SALDataRequirements = (*_SALDataIrrigationControl)(nil)

// NewSALDataIrrigationControl factory function for _SALDataIrrigationControl
func NewSALDataIrrigationControl(salData SALData, irrigationControlData LightingData) *_SALDataIrrigationControl {
	if irrigationControlData == nil {
		panic("irrigationControlData of type LightingData for SALDataIrrigationControl must not be nil")
	}
	_result := &_SALDataIrrigationControl{
		SALDataContract:       NewSALData(salData),
		IrrigationControlData: irrigationControlData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataIrrigationControlBuilder is a builder for SALDataIrrigationControl
type SALDataIrrigationControlBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(irrigationControlData LightingData) SALDataIrrigationControlBuilder
	// WithIrrigationControlData adds IrrigationControlData (property field)
	WithIrrigationControlData(LightingData) SALDataIrrigationControlBuilder
	// WithIrrigationControlDataBuilder adds IrrigationControlData (property field) which is build by the builder
	WithIrrigationControlDataBuilder(func(LightingDataBuilder) LightingDataBuilder) SALDataIrrigationControlBuilder
	// Build builds the SALDataIrrigationControl or returns an error if something is wrong
	Build() (SALDataIrrigationControl, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataIrrigationControl
}

// NewSALDataIrrigationControlBuilder() creates a SALDataIrrigationControlBuilder
func NewSALDataIrrigationControlBuilder() SALDataIrrigationControlBuilder {
	return &_SALDataIrrigationControlBuilder{_SALDataIrrigationControl: new(_SALDataIrrigationControl)}
}

type _SALDataIrrigationControlBuilder struct {
	*_SALDataIrrigationControl

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataIrrigationControlBuilder) = (*_SALDataIrrigationControlBuilder)(nil)

func (b *_SALDataIrrigationControlBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
}

func (b *_SALDataIrrigationControlBuilder) WithMandatoryFields(irrigationControlData LightingData) SALDataIrrigationControlBuilder {
	return b.WithIrrigationControlData(irrigationControlData)
}

func (b *_SALDataIrrigationControlBuilder) WithIrrigationControlData(irrigationControlData LightingData) SALDataIrrigationControlBuilder {
	b.IrrigationControlData = irrigationControlData
	return b
}

func (b *_SALDataIrrigationControlBuilder) WithIrrigationControlDataBuilder(builderSupplier func(LightingDataBuilder) LightingDataBuilder) SALDataIrrigationControlBuilder {
	builder := builderSupplier(b.IrrigationControlData.CreateLightingDataBuilder())
	var err error
	b.IrrigationControlData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LightingDataBuilder failed"))
	}
	return b
}

func (b *_SALDataIrrigationControlBuilder) Build() (SALDataIrrigationControl, error) {
	if b.IrrigationControlData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'irrigationControlData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataIrrigationControl.deepCopy(), nil
}

func (b *_SALDataIrrigationControlBuilder) MustBuild() SALDataIrrigationControl {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SALDataIrrigationControlBuilder) Done() SALDataBuilder {
	return b.parentBuilder
}

func (b *_SALDataIrrigationControlBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataIrrigationControlBuilder) DeepCopy() any {
	_copy := b.CreateSALDataIrrigationControlBuilder().(*_SALDataIrrigationControlBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataIrrigationControlBuilder creates a SALDataIrrigationControlBuilder
func (b *_SALDataIrrigationControl) CreateSALDataIrrigationControlBuilder() SALDataIrrigationControlBuilder {
	if b == nil {
		return NewSALDataIrrigationControlBuilder()
	}
	return &_SALDataIrrigationControlBuilder{_SALDataIrrigationControl: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataIrrigationControl) GetApplicationId() ApplicationId {
	return ApplicationId_IRRIGATION_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataIrrigationControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataIrrigationControl) GetIrrigationControlData() LightingData {
	return m.IrrigationControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataIrrigationControl(structType any) SALDataIrrigationControl {
	if casted, ok := structType.(SALDataIrrigationControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataIrrigationControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataIrrigationControl) GetTypeName() string {
	return "SALDataIrrigationControl"
}

func (m *_SALDataIrrigationControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (irrigationControlData)
	lengthInBits += m.IrrigationControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataIrrigationControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataIrrigationControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataIrrigationControl SALDataIrrigationControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataIrrigationControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataIrrigationControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	irrigationControlData, err := ReadSimpleField[LightingData](ctx, "irrigationControlData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'irrigationControlData' field"))
	}
	m.IrrigationControlData = irrigationControlData

	if closeErr := readBuffer.CloseContext("SALDataIrrigationControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataIrrigationControl")
	}

	return m, nil
}

func (m *_SALDataIrrigationControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataIrrigationControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataIrrigationControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataIrrigationControl")
		}

		if err := WriteSimpleField[LightingData](ctx, "irrigationControlData", m.GetIrrigationControlData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'irrigationControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataIrrigationControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataIrrigationControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataIrrigationControl) IsSALDataIrrigationControl() {}

func (m *_SALDataIrrigationControl) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataIrrigationControl) deepCopy() *_SALDataIrrigationControl {
	if m == nil {
		return nil
	}
	_SALDataIrrigationControlCopy := &_SALDataIrrigationControl{
		m.SALDataContract.(*_SALData).deepCopy(),
		m.IrrigationControlData.DeepCopy().(LightingData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataIrrigationControlCopy
}

func (m *_SALDataIrrigationControl) String() string {
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
