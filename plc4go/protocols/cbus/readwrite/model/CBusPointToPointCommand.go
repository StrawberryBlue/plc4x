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

// CBusPointToPointCommand is the corresponding interface of CBusPointToPointCommand
type CBusPointToPointCommand interface {
	CBusPointToPointCommandContract
	CBusPointToPointCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCBusPointToPointCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToPointCommand()
	// CreateBuilder creates a CBusPointToPointCommandBuilder
	CreateCBusPointToPointCommandBuilder() CBusPointToPointCommandBuilder
}

// CBusPointToPointCommandContract provides a set of functions which can be overwritten by a sub struct
type CBusPointToPointCommandContract interface {
	// GetBridgeAddressCountPeek returns BridgeAddressCountPeek (property field)
	GetBridgeAddressCountPeek() uint16
	// GetCalData returns CalData (property field)
	GetCalData() CALData
	// GetIsDirect returns IsDirect (virtual field)
	GetIsDirect() bool
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// IsCBusPointToPointCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToPointCommand()
	// CreateBuilder creates a CBusPointToPointCommandBuilder
	CreateCBusPointToPointCommandBuilder() CBusPointToPointCommandBuilder
}

// CBusPointToPointCommandRequirements provides a set of functions which need to be implemented by a sub struct
type CBusPointToPointCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsDirect returns IsDirect (discriminator field)
	GetIsDirect() bool
}

// _CBusPointToPointCommand is the data-structure of this message
type _CBusPointToPointCommand struct {
	_SubType               CBusPointToPointCommand
	BridgeAddressCountPeek uint16
	CalData                CALData

	// Arguments.
	CBusOptions CBusOptions
}

var _ CBusPointToPointCommandContract = (*_CBusPointToPointCommand)(nil)

// NewCBusPointToPointCommand factory function for _CBusPointToPointCommand
func NewCBusPointToPointCommand(bridgeAddressCountPeek uint16, calData CALData, cBusOptions CBusOptions) *_CBusPointToPointCommand {
	if calData == nil {
		panic("calData of type CALData for CBusPointToPointCommand must not be nil")
	}
	return &_CBusPointToPointCommand{BridgeAddressCountPeek: bridgeAddressCountPeek, CalData: calData, CBusOptions: cBusOptions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusPointToPointCommandBuilder is a builder for CBusPointToPointCommand
type CBusPointToPointCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bridgeAddressCountPeek uint16, calData CALData) CBusPointToPointCommandBuilder
	// WithBridgeAddressCountPeek adds BridgeAddressCountPeek (property field)
	WithBridgeAddressCountPeek(uint16) CBusPointToPointCommandBuilder
	// WithCalData adds CalData (property field)
	WithCalData(CALData) CBusPointToPointCommandBuilder
	// WithCalDataBuilder adds CalData (property field) which is build by the builder
	WithCalDataBuilder(func(CALDataBuilder) CALDataBuilder) CBusPointToPointCommandBuilder
	// AsCBusPointToPointCommandDirect converts this build to a subType of CBusPointToPointCommand. It is always possible to return to current builder using Done()
	AsCBusPointToPointCommandDirect() interface {
		CBusPointToPointCommandDirectBuilder
		Done() CBusPointToPointCommandBuilder
	}
	// AsCBusPointToPointCommandIndirect converts this build to a subType of CBusPointToPointCommand. It is always possible to return to current builder using Done()
	AsCBusPointToPointCommandIndirect() interface {
		CBusPointToPointCommandIndirectBuilder
		Done() CBusPointToPointCommandBuilder
	}
	// Build builds the CBusPointToPointCommand or returns an error if something is wrong
	PartialBuild() (CBusPointToPointCommandContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() CBusPointToPointCommandContract
	// Build builds the CBusPointToPointCommand or returns an error if something is wrong
	Build() (CBusPointToPointCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusPointToPointCommand
}

// NewCBusPointToPointCommandBuilder() creates a CBusPointToPointCommandBuilder
func NewCBusPointToPointCommandBuilder() CBusPointToPointCommandBuilder {
	return &_CBusPointToPointCommandBuilder{_CBusPointToPointCommand: new(_CBusPointToPointCommand)}
}

type _CBusPointToPointCommandChildBuilder interface {
	utils.Copyable
	setParent(CBusPointToPointCommandContract)
	buildForCBusPointToPointCommand() (CBusPointToPointCommand, error)
}

type _CBusPointToPointCommandBuilder struct {
	*_CBusPointToPointCommand

	childBuilder _CBusPointToPointCommandChildBuilder

	err *utils.MultiError
}

var _ (CBusPointToPointCommandBuilder) = (*_CBusPointToPointCommandBuilder)(nil)

func (b *_CBusPointToPointCommandBuilder) WithMandatoryFields(bridgeAddressCountPeek uint16, calData CALData) CBusPointToPointCommandBuilder {
	return b.WithBridgeAddressCountPeek(bridgeAddressCountPeek).WithCalData(calData)
}

func (b *_CBusPointToPointCommandBuilder) WithBridgeAddressCountPeek(bridgeAddressCountPeek uint16) CBusPointToPointCommandBuilder {
	b.BridgeAddressCountPeek = bridgeAddressCountPeek
	return b
}

func (b *_CBusPointToPointCommandBuilder) WithCalData(calData CALData) CBusPointToPointCommandBuilder {
	b.CalData = calData
	return b
}

func (b *_CBusPointToPointCommandBuilder) WithCalDataBuilder(builderSupplier func(CALDataBuilder) CALDataBuilder) CBusPointToPointCommandBuilder {
	builder := builderSupplier(b.CalData.CreateCALDataBuilder())
	var err error
	b.CalData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CALDataBuilder failed"))
	}
	return b
}

func (b *_CBusPointToPointCommandBuilder) PartialBuild() (CBusPointToPointCommandContract, error) {
	if b.CalData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'calData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusPointToPointCommand.deepCopy(), nil
}

func (b *_CBusPointToPointCommandBuilder) PartialMustBuild() CBusPointToPointCommandContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusPointToPointCommandBuilder) AsCBusPointToPointCommandDirect() interface {
	CBusPointToPointCommandDirectBuilder
	Done() CBusPointToPointCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CBusPointToPointCommandDirectBuilder
		Done() CBusPointToPointCommandBuilder
	}); ok {
		return cb
	}
	cb := NewCBusPointToPointCommandDirectBuilder().(*_CBusPointToPointCommandDirectBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusPointToPointCommandBuilder) AsCBusPointToPointCommandIndirect() interface {
	CBusPointToPointCommandIndirectBuilder
	Done() CBusPointToPointCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CBusPointToPointCommandIndirectBuilder
		Done() CBusPointToPointCommandBuilder
	}); ok {
		return cb
	}
	cb := NewCBusPointToPointCommandIndirectBuilder().(*_CBusPointToPointCommandIndirectBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusPointToPointCommandBuilder) Build() (CBusPointToPointCommand, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCBusPointToPointCommand()
}

func (b *_CBusPointToPointCommandBuilder) MustBuild() CBusPointToPointCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusPointToPointCommandBuilder) DeepCopy() any {
	_copy := b.CreateCBusPointToPointCommandBuilder().(*_CBusPointToPointCommandBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_CBusPointToPointCommandChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusPointToPointCommandBuilder creates a CBusPointToPointCommandBuilder
func (b *_CBusPointToPointCommand) CreateCBusPointToPointCommandBuilder() CBusPointToPointCommandBuilder {
	if b == nil {
		return NewCBusPointToPointCommandBuilder()
	}
	return &_CBusPointToPointCommandBuilder{_CBusPointToPointCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommand) GetBridgeAddressCountPeek() uint16 {
	return m.BridgeAddressCountPeek
}

func (m *_CBusPointToPointCommand) GetCalData() CALData {
	return m.CalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_CBusPointToPointCommand) GetIsDirect() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetBridgeAddressCountPeek() & 0x00FF) == (0x0000)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommand(structType any) CBusPointToPointCommand {
	if casted, ok := structType.(CBusPointToPointCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommand) GetTypeName() string {
	return "CBusPointToPointCommand"
}

func (m *_CBusPointToPointCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CBusPointToPointCommandParse[T CBusPointToPointCommand](ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (T, error) {
	return CBusPointToPointCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToPointCommandParseWithBufferProducer[T CBusPointToPointCommand](cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CBusPointToPointCommandParseWithBuffer[T](ctx, readBuffer, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CBusPointToPointCommandParseWithBuffer[T CBusPointToPointCommand](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (T, error) {
	v, err := (&_CBusPointToPointCommand{CBusOptions: cBusOptions}).parse(ctx, readBuffer, cBusOptions)
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

func (m *_CBusPointToPointCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (__cBusPointToPointCommand CBusPointToPointCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bridgeAddressCountPeek, err := ReadPeekField[uint16](ctx, "bridgeAddressCountPeek", ReadUnsignedShort(readBuffer, uint8(16)), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddressCountPeek' field"))
	}
	m.BridgeAddressCountPeek = bridgeAddressCountPeek

	isDirect, err := ReadVirtualField[bool](ctx, "isDirect", (*bool)(nil), bool((bridgeAddressCountPeek&0x00FF) == (0x0000)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDirect' field"))
	}
	_ = isDirect

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CBusPointToPointCommand
	switch {
	case isDirect == bool(true): // CBusPointToPointCommandDirect
		if _child, err = new(_CBusPointToPointCommandDirect).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusPointToPointCommandDirect for type-switch of CBusPointToPointCommand")
		}
	case isDirect == bool(false): // CBusPointToPointCommandIndirect
		if _child, err = new(_CBusPointToPointCommandIndirect).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusPointToPointCommandIndirect for type-switch of CBusPointToPointCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [isDirect=%v]", isDirect)
	}

	calData, err := ReadSimpleField[CALData](ctx, "calData", ReadComplex[CALData](CALDataParseWithBufferProducer[CALData]((RequestContext)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calData' field"))
	}
	m.CalData = calData

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommand")
	}

	return _child, nil
}

func (pm *_CBusPointToPointCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusPointToPointCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusPointToPointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommand")
	}
	// Virtual field
	isDirect := m.GetIsDirect()
	_ = isDirect
	if _isDirectErr := writeBuffer.WriteVirtual(ctx, "isDirect", m.GetIsDirect()); _isDirectErr != nil {
		return errors.Wrap(_isDirectErr, "Error serializing 'isDirect' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[CALData](ctx, "calData", m.GetCalData(), WriteComplex[CALData](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'calData' field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToPointCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusPointToPointCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusPointToPointCommand) IsCBusPointToPointCommand() {}

func (m *_CBusPointToPointCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusPointToPointCommand) deepCopy() *_CBusPointToPointCommand {
	if m == nil {
		return nil
	}
	_CBusPointToPointCommandCopy := &_CBusPointToPointCommand{
		nil, // will be set by child
		m.BridgeAddressCountPeek,
		m.CalData.DeepCopy().(CALData),
		m.CBusOptions,
	}
	return _CBusPointToPointCommandCopy
}
