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

// ConnectionRequestInformation is the corresponding interface of ConnectionRequestInformation
type ConnectionRequestInformation interface {
	ConnectionRequestInformationContract
	ConnectionRequestInformationRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsConnectionRequestInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionRequestInformation()
	// CreateBuilder creates a ConnectionRequestInformationBuilder
	CreateConnectionRequestInformationBuilder() ConnectionRequestInformationBuilder
}

// ConnectionRequestInformationContract provides a set of functions which can be overwritten by a sub struct
type ConnectionRequestInformationContract interface {
	// IsConnectionRequestInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionRequestInformation()
	// CreateBuilder creates a ConnectionRequestInformationBuilder
	CreateConnectionRequestInformationBuilder() ConnectionRequestInformationBuilder
}

// ConnectionRequestInformationRequirements provides a set of functions which need to be implemented by a sub struct
type ConnectionRequestInformationRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetConnectionType returns ConnectionType (discriminator field)
	GetConnectionType() uint8
}

// _ConnectionRequestInformation is the data-structure of this message
type _ConnectionRequestInformation struct {
	_SubType ConnectionRequestInformation
}

var _ ConnectionRequestInformationContract = (*_ConnectionRequestInformation)(nil)

// NewConnectionRequestInformation factory function for _ConnectionRequestInformation
func NewConnectionRequestInformation() *_ConnectionRequestInformation {
	return &_ConnectionRequestInformation{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionRequestInformationBuilder is a builder for ConnectionRequestInformation
type ConnectionRequestInformationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ConnectionRequestInformationBuilder
	// AsConnectionRequestInformationDeviceManagement converts this build to a subType of ConnectionRequestInformation. It is always possible to return to current builder using Done()
	AsConnectionRequestInformationDeviceManagement() interface {
		ConnectionRequestInformationDeviceManagementBuilder
		Done() ConnectionRequestInformationBuilder
	}
	// AsConnectionRequestInformationTunnelConnection converts this build to a subType of ConnectionRequestInformation. It is always possible to return to current builder using Done()
	AsConnectionRequestInformationTunnelConnection() interface {
		ConnectionRequestInformationTunnelConnectionBuilder
		Done() ConnectionRequestInformationBuilder
	}
	// Build builds the ConnectionRequestInformation or returns an error if something is wrong
	PartialBuild() (ConnectionRequestInformationContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ConnectionRequestInformationContract
	// Build builds the ConnectionRequestInformation or returns an error if something is wrong
	Build() (ConnectionRequestInformation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionRequestInformation
}

// NewConnectionRequestInformationBuilder() creates a ConnectionRequestInformationBuilder
func NewConnectionRequestInformationBuilder() ConnectionRequestInformationBuilder {
	return &_ConnectionRequestInformationBuilder{_ConnectionRequestInformation: new(_ConnectionRequestInformation)}
}

type _ConnectionRequestInformationChildBuilder interface {
	utils.Copyable
	setParent(ConnectionRequestInformationContract)
	buildForConnectionRequestInformation() (ConnectionRequestInformation, error)
}

type _ConnectionRequestInformationBuilder struct {
	*_ConnectionRequestInformation

	childBuilder _ConnectionRequestInformationChildBuilder

	err *utils.MultiError
}

var _ (ConnectionRequestInformationBuilder) = (*_ConnectionRequestInformationBuilder)(nil)

func (b *_ConnectionRequestInformationBuilder) WithMandatoryFields() ConnectionRequestInformationBuilder {
	return b
}

func (b *_ConnectionRequestInformationBuilder) PartialBuild() (ConnectionRequestInformationContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectionRequestInformation.deepCopy(), nil
}

func (b *_ConnectionRequestInformationBuilder) PartialMustBuild() ConnectionRequestInformationContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionRequestInformationBuilder) AsConnectionRequestInformationDeviceManagement() interface {
	ConnectionRequestInformationDeviceManagementBuilder
	Done() ConnectionRequestInformationBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ConnectionRequestInformationDeviceManagementBuilder
		Done() ConnectionRequestInformationBuilder
	}); ok {
		return cb
	}
	cb := NewConnectionRequestInformationDeviceManagementBuilder().(*_ConnectionRequestInformationDeviceManagementBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ConnectionRequestInformationBuilder) AsConnectionRequestInformationTunnelConnection() interface {
	ConnectionRequestInformationTunnelConnectionBuilder
	Done() ConnectionRequestInformationBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ConnectionRequestInformationTunnelConnectionBuilder
		Done() ConnectionRequestInformationBuilder
	}); ok {
		return cb
	}
	cb := NewConnectionRequestInformationTunnelConnectionBuilder().(*_ConnectionRequestInformationTunnelConnectionBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ConnectionRequestInformationBuilder) Build() (ConnectionRequestInformation, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForConnectionRequestInformation()
}

func (b *_ConnectionRequestInformationBuilder) MustBuild() ConnectionRequestInformation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionRequestInformationBuilder) DeepCopy() any {
	_copy := b.CreateConnectionRequestInformationBuilder().(*_ConnectionRequestInformationBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ConnectionRequestInformationChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectionRequestInformationBuilder creates a ConnectionRequestInformationBuilder
func (b *_ConnectionRequestInformation) CreateConnectionRequestInformationBuilder() ConnectionRequestInformationBuilder {
	if b == nil {
		return NewConnectionRequestInformationBuilder()
	}
	return &_ConnectionRequestInformationBuilder{_ConnectionRequestInformation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformation(structType any) ConnectionRequestInformation {
	if casted, ok := structType.(ConnectionRequestInformation); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformation); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformation) GetTypeName() string {
	return "ConnectionRequestInformation"
}

func (m *_ConnectionRequestInformation) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8
	// Discriminator Field (connectionType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionRequestInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ConnectionRequestInformationParse[T ConnectionRequestInformation](ctx context.Context, theBytes []byte) (T, error) {
	return ConnectionRequestInformationParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConnectionRequestInformationParseWithBufferProducer[T ConnectionRequestInformation]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ConnectionRequestInformationParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ConnectionRequestInformationParseWithBuffer[T ConnectionRequestInformation](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ConnectionRequestInformation{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ConnectionRequestInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__connectionRequestInformation ConnectionRequestInformation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	connectionType, err := ReadDiscriminatorField[uint8](ctx, "connectionType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ConnectionRequestInformation
	switch {
	case connectionType == 0x03: // ConnectionRequestInformationDeviceManagement
		if _child, err = new(_ConnectionRequestInformationDeviceManagement).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionRequestInformationDeviceManagement for type-switch of ConnectionRequestInformation")
		}
	case connectionType == 0x04: // ConnectionRequestInformationTunnelConnection
		if _child, err = new(_ConnectionRequestInformationTunnelConnection).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionRequestInformationTunnelConnection for type-switch of ConnectionRequestInformation")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [connectionType=%v]", connectionType)
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformation")
	}

	return _child, nil
}

func (pm *_ConnectionRequestInformation) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ConnectionRequestInformation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ConnectionRequestInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformation")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteDiscriminatorField(ctx, "connectionType", m.GetConnectionType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'connectionType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ConnectionRequestInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ConnectionRequestInformation")
	}
	return nil
}

func (m *_ConnectionRequestInformation) IsConnectionRequestInformation() {}

func (m *_ConnectionRequestInformation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionRequestInformation) deepCopy() *_ConnectionRequestInformation {
	if m == nil {
		return nil
	}
	_ConnectionRequestInformationCopy := &_ConnectionRequestInformation{
		nil, // will be set by child
	}
	return _ConnectionRequestInformationCopy
}
