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

// SALDataTelephonyStatusAndControl is the corresponding interface of SALDataTelephonyStatusAndControl
type SALDataTelephonyStatusAndControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetTelephonyData returns TelephonyData (property field)
	GetTelephonyData() TelephonyData
	// IsSALDataTelephonyStatusAndControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataTelephonyStatusAndControl()
	// CreateBuilder creates a SALDataTelephonyStatusAndControlBuilder
	CreateSALDataTelephonyStatusAndControlBuilder() SALDataTelephonyStatusAndControlBuilder
}

// _SALDataTelephonyStatusAndControl is the data-structure of this message
type _SALDataTelephonyStatusAndControl struct {
	SALDataContract
	TelephonyData TelephonyData
}

var _ SALDataTelephonyStatusAndControl = (*_SALDataTelephonyStatusAndControl)(nil)
var _ SALDataRequirements = (*_SALDataTelephonyStatusAndControl)(nil)

// NewSALDataTelephonyStatusAndControl factory function for _SALDataTelephonyStatusAndControl
func NewSALDataTelephonyStatusAndControl(salData SALData, telephonyData TelephonyData) *_SALDataTelephonyStatusAndControl {
	if telephonyData == nil {
		panic("telephonyData of type TelephonyData for SALDataTelephonyStatusAndControl must not be nil")
	}
	_result := &_SALDataTelephonyStatusAndControl{
		SALDataContract: NewSALData(salData),
		TelephonyData:   telephonyData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataTelephonyStatusAndControlBuilder is a builder for SALDataTelephonyStatusAndControl
type SALDataTelephonyStatusAndControlBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(telephonyData TelephonyData) SALDataTelephonyStatusAndControlBuilder
	// WithTelephonyData adds TelephonyData (property field)
	WithTelephonyData(TelephonyData) SALDataTelephonyStatusAndControlBuilder
	// WithTelephonyDataBuilder adds TelephonyData (property field) which is build by the builder
	WithTelephonyDataBuilder(func(TelephonyDataBuilder) TelephonyDataBuilder) SALDataTelephonyStatusAndControlBuilder
	// Build builds the SALDataTelephonyStatusAndControl or returns an error if something is wrong
	Build() (SALDataTelephonyStatusAndControl, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataTelephonyStatusAndControl
}

// NewSALDataTelephonyStatusAndControlBuilder() creates a SALDataTelephonyStatusAndControlBuilder
func NewSALDataTelephonyStatusAndControlBuilder() SALDataTelephonyStatusAndControlBuilder {
	return &_SALDataTelephonyStatusAndControlBuilder{_SALDataTelephonyStatusAndControl: new(_SALDataTelephonyStatusAndControl)}
}

type _SALDataTelephonyStatusAndControlBuilder struct {
	*_SALDataTelephonyStatusAndControl

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataTelephonyStatusAndControlBuilder) = (*_SALDataTelephonyStatusAndControlBuilder)(nil)

func (b *_SALDataTelephonyStatusAndControlBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
}

func (b *_SALDataTelephonyStatusAndControlBuilder) WithMandatoryFields(telephonyData TelephonyData) SALDataTelephonyStatusAndControlBuilder {
	return b.WithTelephonyData(telephonyData)
}

func (b *_SALDataTelephonyStatusAndControlBuilder) WithTelephonyData(telephonyData TelephonyData) SALDataTelephonyStatusAndControlBuilder {
	b.TelephonyData = telephonyData
	return b
}

func (b *_SALDataTelephonyStatusAndControlBuilder) WithTelephonyDataBuilder(builderSupplier func(TelephonyDataBuilder) TelephonyDataBuilder) SALDataTelephonyStatusAndControlBuilder {
	builder := builderSupplier(b.TelephonyData.CreateTelephonyDataBuilder())
	var err error
	b.TelephonyData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "TelephonyDataBuilder failed"))
	}
	return b
}

func (b *_SALDataTelephonyStatusAndControlBuilder) Build() (SALDataTelephonyStatusAndControl, error) {
	if b.TelephonyData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'telephonyData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataTelephonyStatusAndControl.deepCopy(), nil
}

func (b *_SALDataTelephonyStatusAndControlBuilder) MustBuild() SALDataTelephonyStatusAndControl {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SALDataTelephonyStatusAndControlBuilder) Done() SALDataBuilder {
	return b.parentBuilder
}

func (b *_SALDataTelephonyStatusAndControlBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataTelephonyStatusAndControlBuilder) DeepCopy() any {
	_copy := b.CreateSALDataTelephonyStatusAndControlBuilder().(*_SALDataTelephonyStatusAndControlBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataTelephonyStatusAndControlBuilder creates a SALDataTelephonyStatusAndControlBuilder
func (b *_SALDataTelephonyStatusAndControl) CreateSALDataTelephonyStatusAndControlBuilder() SALDataTelephonyStatusAndControlBuilder {
	if b == nil {
		return NewSALDataTelephonyStatusAndControlBuilder()
	}
	return &_SALDataTelephonyStatusAndControlBuilder{_SALDataTelephonyStatusAndControl: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetApplicationId() ApplicationId {
	return ApplicationId_TELEPHONY_STATUS_AND_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetTelephonyData() TelephonyData {
	return m.TelephonyData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataTelephonyStatusAndControl(structType any) SALDataTelephonyStatusAndControl {
	if casted, ok := structType.(SALDataTelephonyStatusAndControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTelephonyStatusAndControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTelephonyStatusAndControl) GetTypeName() string {
	return "SALDataTelephonyStatusAndControl"
}

func (m *_SALDataTelephonyStatusAndControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (telephonyData)
	lengthInBits += m.TelephonyData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTelephonyStatusAndControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataTelephonyStatusAndControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataTelephonyStatusAndControl SALDataTelephonyStatusAndControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTelephonyStatusAndControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTelephonyStatusAndControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	telephonyData, err := ReadSimpleField[TelephonyData](ctx, "telephonyData", ReadComplex[TelephonyData](TelephonyDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'telephonyData' field"))
	}
	m.TelephonyData = telephonyData

	if closeErr := readBuffer.CloseContext("SALDataTelephonyStatusAndControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTelephonyStatusAndControl")
	}

	return m, nil
}

func (m *_SALDataTelephonyStatusAndControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTelephonyStatusAndControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTelephonyStatusAndControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTelephonyStatusAndControl")
		}

		if err := WriteSimpleField[TelephonyData](ctx, "telephonyData", m.GetTelephonyData(), WriteComplex[TelephonyData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'telephonyData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTelephonyStatusAndControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTelephonyStatusAndControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTelephonyStatusAndControl) IsSALDataTelephonyStatusAndControl() {}

func (m *_SALDataTelephonyStatusAndControl) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataTelephonyStatusAndControl) deepCopy() *_SALDataTelephonyStatusAndControl {
	if m == nil {
		return nil
	}
	_SALDataTelephonyStatusAndControlCopy := &_SALDataTelephonyStatusAndControl{
		m.SALDataContract.(*_SALData).deepCopy(),
		m.TelephonyData.DeepCopy().(TelephonyData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataTelephonyStatusAndControlCopy
}

func (m *_SALDataTelephonyStatusAndControl) String() string {
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
