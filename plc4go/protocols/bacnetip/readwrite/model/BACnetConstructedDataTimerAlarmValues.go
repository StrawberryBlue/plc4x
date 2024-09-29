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

// BACnetConstructedDataTimerAlarmValues is the corresponding interface of BACnetConstructedDataTimerAlarmValues
type BACnetConstructedDataTimerAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetTimerStateTagged
	// IsBACnetConstructedDataTimerAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimerAlarmValues()
	// CreateBuilder creates a BACnetConstructedDataTimerAlarmValuesBuilder
	CreateBACnetConstructedDataTimerAlarmValuesBuilder() BACnetConstructedDataTimerAlarmValuesBuilder
}

// _BACnetConstructedDataTimerAlarmValues is the data-structure of this message
type _BACnetConstructedDataTimerAlarmValues struct {
	BACnetConstructedDataContract
	AlarmValues []BACnetTimerStateTagged
}

var _ BACnetConstructedDataTimerAlarmValues = (*_BACnetConstructedDataTimerAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimerAlarmValues)(nil)

// NewBACnetConstructedDataTimerAlarmValues factory function for _BACnetConstructedDataTimerAlarmValues
func NewBACnetConstructedDataTimerAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, alarmValues []BACnetTimerStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerAlarmValues {
	_result := &_BACnetConstructedDataTimerAlarmValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AlarmValues:                   alarmValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimerAlarmValuesBuilder is a builder for BACnetConstructedDataTimerAlarmValues
type BACnetConstructedDataTimerAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alarmValues []BACnetTimerStateTagged) BACnetConstructedDataTimerAlarmValuesBuilder
	// WithAlarmValues adds AlarmValues (property field)
	WithAlarmValues(...BACnetTimerStateTagged) BACnetConstructedDataTimerAlarmValuesBuilder
	// Build builds the BACnetConstructedDataTimerAlarmValues or returns an error if something is wrong
	Build() (BACnetConstructedDataTimerAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimerAlarmValues
}

// NewBACnetConstructedDataTimerAlarmValuesBuilder() creates a BACnetConstructedDataTimerAlarmValuesBuilder
func NewBACnetConstructedDataTimerAlarmValuesBuilder() BACnetConstructedDataTimerAlarmValuesBuilder {
	return &_BACnetConstructedDataTimerAlarmValuesBuilder{_BACnetConstructedDataTimerAlarmValues: new(_BACnetConstructedDataTimerAlarmValues)}
}

type _BACnetConstructedDataTimerAlarmValuesBuilder struct {
	*_BACnetConstructedDataTimerAlarmValues

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimerAlarmValuesBuilder) = (*_BACnetConstructedDataTimerAlarmValuesBuilder)(nil)

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) WithMandatoryFields(alarmValues []BACnetTimerStateTagged) BACnetConstructedDataTimerAlarmValuesBuilder {
	return b.WithAlarmValues(alarmValues...)
}

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) WithAlarmValues(alarmValues ...BACnetTimerStateTagged) BACnetConstructedDataTimerAlarmValuesBuilder {
	b.AlarmValues = alarmValues
	return b
}

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) Build() (BACnetConstructedDataTimerAlarmValues, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimerAlarmValues.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) MustBuild() BACnetConstructedDataTimerAlarmValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimerAlarmValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimerAlarmValuesBuilder().(*_BACnetConstructedDataTimerAlarmValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimerAlarmValuesBuilder creates a BACnetConstructedDataTimerAlarmValuesBuilder
func (b *_BACnetConstructedDataTimerAlarmValues) CreateBACnetConstructedDataTimerAlarmValuesBuilder() BACnetConstructedDataTimerAlarmValuesBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimerAlarmValuesBuilder()
	}
	return &_BACnetConstructedDataTimerAlarmValuesBuilder{_BACnetConstructedDataTimerAlarmValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerAlarmValues) GetAlarmValues() []BACnetTimerStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerAlarmValues(structType any) BACnetConstructedDataTimerAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataTimerAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataTimerAlarmValues"
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimerAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimerAlarmValues BACnetConstructedDataTimerAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmValues, err := ReadTerminatedArrayField[BACnetTimerStateTagged](ctx, "alarmValues", ReadComplex[BACnetTimerStateTagged](BACnetTimerStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimerAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimerAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerAlarmValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerAlarmValues) IsBACnetConstructedDataTimerAlarmValues() {}

func (m *_BACnetConstructedDataTimerAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimerAlarmValues) deepCopy() *_BACnetConstructedDataTimerAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimerAlarmValuesCopy := &_BACnetConstructedDataTimerAlarmValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetTimerStateTagged, BACnetTimerStateTagged](m.AlarmValues),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimerAlarmValuesCopy
}

func (m *_BACnetConstructedDataTimerAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
