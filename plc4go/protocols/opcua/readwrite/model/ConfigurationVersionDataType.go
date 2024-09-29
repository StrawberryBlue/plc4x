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

// ConfigurationVersionDataType is the corresponding interface of ConfigurationVersionDataType
type ConfigurationVersionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint32
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint32
	// IsConfigurationVersionDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConfigurationVersionDataType()
	// CreateBuilder creates a ConfigurationVersionDataTypeBuilder
	CreateConfigurationVersionDataTypeBuilder() ConfigurationVersionDataTypeBuilder
}

// _ConfigurationVersionDataType is the data-structure of this message
type _ConfigurationVersionDataType struct {
	ExtensionObjectDefinitionContract
	MajorVersion uint32
	MinorVersion uint32
}

var _ ConfigurationVersionDataType = (*_ConfigurationVersionDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ConfigurationVersionDataType)(nil)

// NewConfigurationVersionDataType factory function for _ConfigurationVersionDataType
func NewConfigurationVersionDataType(majorVersion uint32, minorVersion uint32) *_ConfigurationVersionDataType {
	_result := &_ConfigurationVersionDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		MajorVersion:                      majorVersion,
		MinorVersion:                      minorVersion,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConfigurationVersionDataTypeBuilder is a builder for ConfigurationVersionDataType
type ConfigurationVersionDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(majorVersion uint32, minorVersion uint32) ConfigurationVersionDataTypeBuilder
	// WithMajorVersion adds MajorVersion (property field)
	WithMajorVersion(uint32) ConfigurationVersionDataTypeBuilder
	// WithMinorVersion adds MinorVersion (property field)
	WithMinorVersion(uint32) ConfigurationVersionDataTypeBuilder
	// Build builds the ConfigurationVersionDataType or returns an error if something is wrong
	Build() (ConfigurationVersionDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConfigurationVersionDataType
}

// NewConfigurationVersionDataTypeBuilder() creates a ConfigurationVersionDataTypeBuilder
func NewConfigurationVersionDataTypeBuilder() ConfigurationVersionDataTypeBuilder {
	return &_ConfigurationVersionDataTypeBuilder{_ConfigurationVersionDataType: new(_ConfigurationVersionDataType)}
}

type _ConfigurationVersionDataTypeBuilder struct {
	*_ConfigurationVersionDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ConfigurationVersionDataTypeBuilder) = (*_ConfigurationVersionDataTypeBuilder)(nil)

func (b *_ConfigurationVersionDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ConfigurationVersionDataTypeBuilder) WithMandatoryFields(majorVersion uint32, minorVersion uint32) ConfigurationVersionDataTypeBuilder {
	return b.WithMajorVersion(majorVersion).WithMinorVersion(minorVersion)
}

func (b *_ConfigurationVersionDataTypeBuilder) WithMajorVersion(majorVersion uint32) ConfigurationVersionDataTypeBuilder {
	b.MajorVersion = majorVersion
	return b
}

func (b *_ConfigurationVersionDataTypeBuilder) WithMinorVersion(minorVersion uint32) ConfigurationVersionDataTypeBuilder {
	b.MinorVersion = minorVersion
	return b
}

func (b *_ConfigurationVersionDataTypeBuilder) Build() (ConfigurationVersionDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConfigurationVersionDataType.deepCopy(), nil
}

func (b *_ConfigurationVersionDataTypeBuilder) MustBuild() ConfigurationVersionDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ConfigurationVersionDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ConfigurationVersionDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ConfigurationVersionDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateConfigurationVersionDataTypeBuilder().(*_ConfigurationVersionDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConfigurationVersionDataTypeBuilder creates a ConfigurationVersionDataTypeBuilder
func (b *_ConfigurationVersionDataType) CreateConfigurationVersionDataTypeBuilder() ConfigurationVersionDataTypeBuilder {
	if b == nil {
		return NewConfigurationVersionDataTypeBuilder()
	}
	return &_ConfigurationVersionDataTypeBuilder{_ConfigurationVersionDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConfigurationVersionDataType) GetIdentifier() string {
	return "14595"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConfigurationVersionDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConfigurationVersionDataType) GetMajorVersion() uint32 {
	return m.MajorVersion
}

func (m *_ConfigurationVersionDataType) GetMinorVersion() uint32 {
	return m.MinorVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConfigurationVersionDataType(structType any) ConfigurationVersionDataType {
	if casted, ok := structType.(ConfigurationVersionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ConfigurationVersionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ConfigurationVersionDataType) GetTypeName() string {
	return "ConfigurationVersionDataType"
}

func (m *_ConfigurationVersionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (majorVersion)
	lengthInBits += 32

	// Simple field (minorVersion)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ConfigurationVersionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConfigurationVersionDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__configurationVersionDataType ConfigurationVersionDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConfigurationVersionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConfigurationVersionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	majorVersion, err := ReadSimpleField(ctx, "majorVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'majorVersion' field"))
	}
	m.MajorVersion = majorVersion

	minorVersion, err := ReadSimpleField(ctx, "minorVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minorVersion' field"))
	}
	m.MinorVersion = minorVersion

	if closeErr := readBuffer.CloseContext("ConfigurationVersionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConfigurationVersionDataType")
	}

	return m, nil
}

func (m *_ConfigurationVersionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConfigurationVersionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConfigurationVersionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConfigurationVersionDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "majorVersion", m.GetMajorVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'majorVersion' field")
		}

		if err := WriteSimpleField[uint32](ctx, "minorVersion", m.GetMinorVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'minorVersion' field")
		}

		if popErr := writeBuffer.PopContext("ConfigurationVersionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConfigurationVersionDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConfigurationVersionDataType) IsConfigurationVersionDataType() {}

func (m *_ConfigurationVersionDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConfigurationVersionDataType) deepCopy() *_ConfigurationVersionDataType {
	if m == nil {
		return nil
	}
	_ConfigurationVersionDataTypeCopy := &_ConfigurationVersionDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.MajorVersion,
		m.MinorVersion,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ConfigurationVersionDataTypeCopy
}

func (m *_ConfigurationVersionDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
