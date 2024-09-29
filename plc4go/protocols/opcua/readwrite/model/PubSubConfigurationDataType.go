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

// PubSubConfigurationDataType is the corresponding interface of PubSubConfigurationDataType
type PubSubConfigurationDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNoOfPublishedDataSets returns NoOfPublishedDataSets (property field)
	GetNoOfPublishedDataSets() int32
	// GetPublishedDataSets returns PublishedDataSets (property field)
	GetPublishedDataSets() []ExtensionObjectDefinition
	// GetNoOfConnections returns NoOfConnections (property field)
	GetNoOfConnections() int32
	// GetConnections returns Connections (property field)
	GetConnections() []ExtensionObjectDefinition
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// IsPubSubConfigurationDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubConfigurationDataType()
	// CreateBuilder creates a PubSubConfigurationDataTypeBuilder
	CreatePubSubConfigurationDataTypeBuilder() PubSubConfigurationDataTypeBuilder
}

// _PubSubConfigurationDataType is the data-structure of this message
type _PubSubConfigurationDataType struct {
	ExtensionObjectDefinitionContract
	NoOfPublishedDataSets int32
	PublishedDataSets     []ExtensionObjectDefinition
	NoOfConnections       int32
	Connections           []ExtensionObjectDefinition
	Enabled               bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ PubSubConfigurationDataType = (*_PubSubConfigurationDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubConfigurationDataType)(nil)

// NewPubSubConfigurationDataType factory function for _PubSubConfigurationDataType
func NewPubSubConfigurationDataType(noOfPublishedDataSets int32, publishedDataSets []ExtensionObjectDefinition, noOfConnections int32, connections []ExtensionObjectDefinition, enabled bool) *_PubSubConfigurationDataType {
	_result := &_PubSubConfigurationDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfPublishedDataSets:             noOfPublishedDataSets,
		PublishedDataSets:                 publishedDataSets,
		NoOfConnections:                   noOfConnections,
		Connections:                       connections,
		Enabled:                           enabled,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PubSubConfigurationDataTypeBuilder is a builder for PubSubConfigurationDataType
type PubSubConfigurationDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(noOfPublishedDataSets int32, publishedDataSets []ExtensionObjectDefinition, noOfConnections int32, connections []ExtensionObjectDefinition, enabled bool) PubSubConfigurationDataTypeBuilder
	// WithNoOfPublishedDataSets adds NoOfPublishedDataSets (property field)
	WithNoOfPublishedDataSets(int32) PubSubConfigurationDataTypeBuilder
	// WithPublishedDataSets adds PublishedDataSets (property field)
	WithPublishedDataSets(...ExtensionObjectDefinition) PubSubConfigurationDataTypeBuilder
	// WithNoOfConnections adds NoOfConnections (property field)
	WithNoOfConnections(int32) PubSubConfigurationDataTypeBuilder
	// WithConnections adds Connections (property field)
	WithConnections(...ExtensionObjectDefinition) PubSubConfigurationDataTypeBuilder
	// WithEnabled adds Enabled (property field)
	WithEnabled(bool) PubSubConfigurationDataTypeBuilder
	// Build builds the PubSubConfigurationDataType or returns an error if something is wrong
	Build() (PubSubConfigurationDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PubSubConfigurationDataType
}

// NewPubSubConfigurationDataTypeBuilder() creates a PubSubConfigurationDataTypeBuilder
func NewPubSubConfigurationDataTypeBuilder() PubSubConfigurationDataTypeBuilder {
	return &_PubSubConfigurationDataTypeBuilder{_PubSubConfigurationDataType: new(_PubSubConfigurationDataType)}
}

type _PubSubConfigurationDataTypeBuilder struct {
	*_PubSubConfigurationDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PubSubConfigurationDataTypeBuilder) = (*_PubSubConfigurationDataTypeBuilder)(nil)

func (b *_PubSubConfigurationDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_PubSubConfigurationDataTypeBuilder) WithMandatoryFields(noOfPublishedDataSets int32, publishedDataSets []ExtensionObjectDefinition, noOfConnections int32, connections []ExtensionObjectDefinition, enabled bool) PubSubConfigurationDataTypeBuilder {
	return b.WithNoOfPublishedDataSets(noOfPublishedDataSets).WithPublishedDataSets(publishedDataSets...).WithNoOfConnections(noOfConnections).WithConnections(connections...).WithEnabled(enabled)
}

func (b *_PubSubConfigurationDataTypeBuilder) WithNoOfPublishedDataSets(noOfPublishedDataSets int32) PubSubConfigurationDataTypeBuilder {
	b.NoOfPublishedDataSets = noOfPublishedDataSets
	return b
}

func (b *_PubSubConfigurationDataTypeBuilder) WithPublishedDataSets(publishedDataSets ...ExtensionObjectDefinition) PubSubConfigurationDataTypeBuilder {
	b.PublishedDataSets = publishedDataSets
	return b
}

func (b *_PubSubConfigurationDataTypeBuilder) WithNoOfConnections(noOfConnections int32) PubSubConfigurationDataTypeBuilder {
	b.NoOfConnections = noOfConnections
	return b
}

func (b *_PubSubConfigurationDataTypeBuilder) WithConnections(connections ...ExtensionObjectDefinition) PubSubConfigurationDataTypeBuilder {
	b.Connections = connections
	return b
}

func (b *_PubSubConfigurationDataTypeBuilder) WithEnabled(enabled bool) PubSubConfigurationDataTypeBuilder {
	b.Enabled = enabled
	return b
}

func (b *_PubSubConfigurationDataTypeBuilder) Build() (PubSubConfigurationDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PubSubConfigurationDataType.deepCopy(), nil
}

func (b *_PubSubConfigurationDataTypeBuilder) MustBuild() PubSubConfigurationDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_PubSubConfigurationDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_PubSubConfigurationDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PubSubConfigurationDataTypeBuilder) DeepCopy() any {
	_copy := b.CreatePubSubConfigurationDataTypeBuilder().(*_PubSubConfigurationDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePubSubConfigurationDataTypeBuilder creates a PubSubConfigurationDataTypeBuilder
func (b *_PubSubConfigurationDataType) CreatePubSubConfigurationDataTypeBuilder() PubSubConfigurationDataTypeBuilder {
	if b == nil {
		return NewPubSubConfigurationDataTypeBuilder()
	}
	return &_PubSubConfigurationDataTypeBuilder{_PubSubConfigurationDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConfigurationDataType) GetIdentifier() string {
	return "15532"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConfigurationDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConfigurationDataType) GetNoOfPublishedDataSets() int32 {
	return m.NoOfPublishedDataSets
}

func (m *_PubSubConfigurationDataType) GetPublishedDataSets() []ExtensionObjectDefinition {
	return m.PublishedDataSets
}

func (m *_PubSubConfigurationDataType) GetNoOfConnections() int32 {
	return m.NoOfConnections
}

func (m *_PubSubConfigurationDataType) GetConnections() []ExtensionObjectDefinition {
	return m.Connections
}

func (m *_PubSubConfigurationDataType) GetEnabled() bool {
	return m.Enabled
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPubSubConfigurationDataType(structType any) PubSubConfigurationDataType {
	if casted, ok := structType.(PubSubConfigurationDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConfigurationDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConfigurationDataType) GetTypeName() string {
	return "PubSubConfigurationDataType"
}

func (m *_PubSubConfigurationDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfPublishedDataSets)
	lengthInBits += 32

	// Array field
	if len(m.PublishedDataSets) > 0 {
		for _curItem, element := range m.PublishedDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PublishedDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfConnections)
	lengthInBits += 32

	// Array field
	if len(m.Connections) > 0 {
		for _curItem, element := range m.Connections {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Connections), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	return lengthInBits
}

func (m *_PubSubConfigurationDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PubSubConfigurationDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__pubSubConfigurationDataType PubSubConfigurationDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PubSubConfigurationDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConfigurationDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfPublishedDataSets, err := ReadSimpleField(ctx, "noOfPublishedDataSets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfPublishedDataSets' field"))
	}
	m.NoOfPublishedDataSets = noOfPublishedDataSets

	publishedDataSets, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "publishedDataSets", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("15580")), readBuffer), uint64(noOfPublishedDataSets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishedDataSets' field"))
	}
	m.PublishedDataSets = publishedDataSets

	noOfConnections, err := ReadSimpleField(ctx, "noOfConnections", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfConnections' field"))
	}
	m.NoOfConnections = noOfConnections

	connections, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "connections", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("15619")), readBuffer), uint64(noOfConnections))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connections' field"))
	}
	m.Connections = connections

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	if closeErr := readBuffer.CloseContext("PubSubConfigurationDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConfigurationDataType")
	}

	return m, nil
}

func (m *_PubSubConfigurationDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConfigurationDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConfigurationDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConfigurationDataType")
		}

		if err := WriteSimpleField[int32](ctx, "noOfPublishedDataSets", m.GetNoOfPublishedDataSets(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfPublishedDataSets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "publishedDataSets", m.GetPublishedDataSets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'publishedDataSets' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfConnections", m.GetNoOfConnections(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfConnections' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "connections", m.GetConnections(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'connections' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConfigurationDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConfigurationDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConfigurationDataType) IsPubSubConfigurationDataType() {}

func (m *_PubSubConfigurationDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PubSubConfigurationDataType) deepCopy() *_PubSubConfigurationDataType {
	if m == nil {
		return nil
	}
	_PubSubConfigurationDataTypeCopy := &_PubSubConfigurationDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NoOfPublishedDataSets,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.PublishedDataSets),
		m.NoOfConnections,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.Connections),
		m.Enabled,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PubSubConfigurationDataTypeCopy
}

func (m *_PubSubConfigurationDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
