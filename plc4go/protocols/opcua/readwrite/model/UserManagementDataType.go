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

// UserManagementDataType is the corresponding interface of UserManagementDataType
type UserManagementDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetUserName returns UserName (property field)
	GetUserName() PascalString
	// GetUserConfiguration returns UserConfiguration (property field)
	GetUserConfiguration() UserConfigurationMask
	// GetDescription returns Description (property field)
	GetDescription() PascalString
	// IsUserManagementDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUserManagementDataType()
	// CreateBuilder creates a UserManagementDataTypeBuilder
	CreateUserManagementDataTypeBuilder() UserManagementDataTypeBuilder
}

// _UserManagementDataType is the data-structure of this message
type _UserManagementDataType struct {
	ExtensionObjectDefinitionContract
	UserName          PascalString
	UserConfiguration UserConfigurationMask
	Description       PascalString
}

var _ UserManagementDataType = (*_UserManagementDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UserManagementDataType)(nil)

// NewUserManagementDataType factory function for _UserManagementDataType
func NewUserManagementDataType(userName PascalString, userConfiguration UserConfigurationMask, description PascalString) *_UserManagementDataType {
	if userName == nil {
		panic("userName of type PascalString for UserManagementDataType must not be nil")
	}
	if description == nil {
		panic("description of type PascalString for UserManagementDataType must not be nil")
	}
	_result := &_UserManagementDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		UserName:                          userName,
		UserConfiguration:                 userConfiguration,
		Description:                       description,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UserManagementDataTypeBuilder is a builder for UserManagementDataType
type UserManagementDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(userName PascalString, userConfiguration UserConfigurationMask, description PascalString) UserManagementDataTypeBuilder
	// WithUserName adds UserName (property field)
	WithUserName(PascalString) UserManagementDataTypeBuilder
	// WithUserNameBuilder adds UserName (property field) which is build by the builder
	WithUserNameBuilder(func(PascalStringBuilder) PascalStringBuilder) UserManagementDataTypeBuilder
	// WithUserConfiguration adds UserConfiguration (property field)
	WithUserConfiguration(UserConfigurationMask) UserManagementDataTypeBuilder
	// WithDescription adds Description (property field)
	WithDescription(PascalString) UserManagementDataTypeBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(PascalStringBuilder) PascalStringBuilder) UserManagementDataTypeBuilder
	// Build builds the UserManagementDataType or returns an error if something is wrong
	Build() (UserManagementDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UserManagementDataType
}

// NewUserManagementDataTypeBuilder() creates a UserManagementDataTypeBuilder
func NewUserManagementDataTypeBuilder() UserManagementDataTypeBuilder {
	return &_UserManagementDataTypeBuilder{_UserManagementDataType: new(_UserManagementDataType)}
}

type _UserManagementDataTypeBuilder struct {
	*_UserManagementDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UserManagementDataTypeBuilder) = (*_UserManagementDataTypeBuilder)(nil)

func (b *_UserManagementDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_UserManagementDataTypeBuilder) WithMandatoryFields(userName PascalString, userConfiguration UserConfigurationMask, description PascalString) UserManagementDataTypeBuilder {
	return b.WithUserName(userName).WithUserConfiguration(userConfiguration).WithDescription(description)
}

func (b *_UserManagementDataTypeBuilder) WithUserName(userName PascalString) UserManagementDataTypeBuilder {
	b.UserName = userName
	return b
}

func (b *_UserManagementDataTypeBuilder) WithUserNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserManagementDataTypeBuilder {
	builder := builderSupplier(b.UserName.CreatePascalStringBuilder())
	var err error
	b.UserName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserManagementDataTypeBuilder) WithUserConfiguration(userConfiguration UserConfigurationMask) UserManagementDataTypeBuilder {
	b.UserConfiguration = userConfiguration
	return b
}

func (b *_UserManagementDataTypeBuilder) WithDescription(description PascalString) UserManagementDataTypeBuilder {
	b.Description = description
	return b
}

func (b *_UserManagementDataTypeBuilder) WithDescriptionBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserManagementDataTypeBuilder {
	builder := builderSupplier(b.Description.CreatePascalStringBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserManagementDataTypeBuilder) Build() (UserManagementDataType, error) {
	if b.UserName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UserManagementDataType.deepCopy(), nil
}

func (b *_UserManagementDataTypeBuilder) MustBuild() UserManagementDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_UserManagementDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_UserManagementDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UserManagementDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateUserManagementDataTypeBuilder().(*_UserManagementDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUserManagementDataTypeBuilder creates a UserManagementDataTypeBuilder
func (b *_UserManagementDataType) CreateUserManagementDataTypeBuilder() UserManagementDataTypeBuilder {
	if b == nil {
		return NewUserManagementDataTypeBuilder()
	}
	return &_UserManagementDataTypeBuilder{_UserManagementDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserManagementDataType) GetIdentifier() string {
	return "24283"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserManagementDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserManagementDataType) GetUserName() PascalString {
	return m.UserName
}

func (m *_UserManagementDataType) GetUserConfiguration() UserConfigurationMask {
	return m.UserConfiguration
}

func (m *_UserManagementDataType) GetDescription() PascalString {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUserManagementDataType(structType any) UserManagementDataType {
	if casted, ok := structType.(UserManagementDataType); ok {
		return casted
	}
	if casted, ok := structType.(*UserManagementDataType); ok {
		return *casted
	}
	return nil
}

func (m *_UserManagementDataType) GetTypeName() string {
	return "UserManagementDataType"
}

func (m *_UserManagementDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	// Simple field (userConfiguration)
	lengthInBits += 32

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UserManagementDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UserManagementDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__userManagementDataType UserManagementDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserManagementDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserManagementDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	userName, err := ReadSimpleField[PascalString](ctx, "userName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userName' field"))
	}
	m.UserName = userName

	userConfiguration, err := ReadEnumField[UserConfigurationMask](ctx, "userConfiguration", "UserConfigurationMask", ReadEnum(UserConfigurationMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userConfiguration' field"))
	}
	m.UserConfiguration = userConfiguration

	description, err := ReadSimpleField[PascalString](ctx, "description", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	if closeErr := readBuffer.CloseContext("UserManagementDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserManagementDataType")
	}

	return m, nil
}

func (m *_UserManagementDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UserManagementDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UserManagementDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UserManagementDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "userName", m.GetUserName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userName' field")
		}

		if err := WriteSimpleEnumField[UserConfigurationMask](ctx, "userConfiguration", "UserConfigurationMask", m.GetUserConfiguration(), WriteEnum[UserConfigurationMask, uint32](UserConfigurationMask.GetValue, UserConfigurationMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'userConfiguration' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "description", m.GetDescription(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("UserManagementDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UserManagementDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserManagementDataType) IsUserManagementDataType() {}

func (m *_UserManagementDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UserManagementDataType) deepCopy() *_UserManagementDataType {
	if m == nil {
		return nil
	}
	_UserManagementDataTypeCopy := &_UserManagementDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.UserName.DeepCopy().(PascalString),
		m.UserConfiguration,
		m.Description.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UserManagementDataTypeCopy
}

func (m *_UserManagementDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
