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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AmsTCPPacket is the corresponding interface of AmsTCPPacket
type AmsTCPPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetUserdata returns Userdata (property field)
	GetUserdata() AmsPacket
	// IsAmsTCPPacket is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAmsTCPPacket()
	// CreateBuilder creates a AmsTCPPacketBuilder
	CreateAmsTCPPacketBuilder() AmsTCPPacketBuilder
}

// _AmsTCPPacket is the data-structure of this message
type _AmsTCPPacket struct {
	Userdata AmsPacket
	// Reserved Fields
	reservedField0 *uint16
}

var _ AmsTCPPacket = (*_AmsTCPPacket)(nil)

// NewAmsTCPPacket factory function for _AmsTCPPacket
func NewAmsTCPPacket(userdata AmsPacket) *_AmsTCPPacket {
	if userdata == nil {
		panic("userdata of type AmsPacket for AmsTCPPacket must not be nil")
	}
	return &_AmsTCPPacket{Userdata: userdata}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AmsTCPPacketBuilder is a builder for AmsTCPPacket
type AmsTCPPacketBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(userdata AmsPacket) AmsTCPPacketBuilder
	// WithUserdata adds Userdata (property field)
	WithUserdata(AmsPacket) AmsTCPPacketBuilder
	// WithUserdataBuilder adds Userdata (property field) which is build by the builder
	WithUserdataBuilder(func(AmsPacketBuilder) AmsPacketBuilder) AmsTCPPacketBuilder
	// Build builds the AmsTCPPacket or returns an error if something is wrong
	Build() (AmsTCPPacket, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AmsTCPPacket
}

// NewAmsTCPPacketBuilder() creates a AmsTCPPacketBuilder
func NewAmsTCPPacketBuilder() AmsTCPPacketBuilder {
	return &_AmsTCPPacketBuilder{_AmsTCPPacket: new(_AmsTCPPacket)}
}

type _AmsTCPPacketBuilder struct {
	*_AmsTCPPacket

	err *utils.MultiError
}

var _ (AmsTCPPacketBuilder) = (*_AmsTCPPacketBuilder)(nil)

func (b *_AmsTCPPacketBuilder) WithMandatoryFields(userdata AmsPacket) AmsTCPPacketBuilder {
	return b.WithUserdata(userdata)
}

func (b *_AmsTCPPacketBuilder) WithUserdata(userdata AmsPacket) AmsTCPPacketBuilder {
	b.Userdata = userdata
	return b
}

func (b *_AmsTCPPacketBuilder) WithUserdataBuilder(builderSupplier func(AmsPacketBuilder) AmsPacketBuilder) AmsTCPPacketBuilder {
	builder := builderSupplier(b.Userdata.CreateAmsPacketBuilder())
	var err error
	b.Userdata, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AmsPacketBuilder failed"))
	}
	return b
}

func (b *_AmsTCPPacketBuilder) Build() (AmsTCPPacket, error) {
	if b.Userdata == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userdata' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AmsTCPPacket.deepCopy(), nil
}

func (b *_AmsTCPPacketBuilder) MustBuild() AmsTCPPacket {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AmsTCPPacketBuilder) DeepCopy() any {
	_copy := b.CreateAmsTCPPacketBuilder().(*_AmsTCPPacketBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAmsTCPPacketBuilder creates a AmsTCPPacketBuilder
func (b *_AmsTCPPacket) CreateAmsTCPPacketBuilder() AmsTCPPacketBuilder {
	if b == nil {
		return NewAmsTCPPacketBuilder()
	}
	return &_AmsTCPPacketBuilder{_AmsTCPPacket: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsTCPPacket) GetUserdata() AmsPacket {
	return m.Userdata
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAmsTCPPacket(structType any) AmsTCPPacket {
	if casted, ok := structType.(AmsTCPPacket); ok {
		return casted
	}
	if casted, ok := structType.(*AmsTCPPacket); ok {
		return *casted
	}
	return nil
}

func (m *_AmsTCPPacket) GetTypeName() string {
	return "AmsTCPPacket"
}

func (m *_AmsTCPPacket) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (userdata)
	lengthInBits += m.Userdata.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AmsTCPPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsTCPPacketParse(ctx context.Context, theBytes []byte) (AmsTCPPacket, error) {
	return AmsTCPPacketParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AmsTCPPacketParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsTCPPacket, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsTCPPacket, error) {
		return AmsTCPPacketParseWithBuffer(ctx, readBuffer)
	}
}

func AmsTCPPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsTCPPacket, error) {
	v, err := (&_AmsTCPPacket{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AmsTCPPacket) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__amsTCPPacket AmsTCPPacket, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsTCPPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsTCPPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	length, err := ReadImplicitField[uint32](ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	userdata, err := ReadSimpleField[AmsPacket](ctx, "userdata", ReadComplex[AmsPacket](AmsPacketParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userdata' field"))
	}
	m.Userdata = userdata

	if closeErr := readBuffer.CloseContext("AmsTCPPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsTCPPacket")
	}

	return m, nil
}

func (m *_AmsTCPPacket) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsTCPPacket) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AmsTCPPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsTCPPacket")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}
	length := uint32(m.GetUserdata().GetLengthInBytes(ctx))
	if err := WriteImplicitField(ctx, "length", length, WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'length' field")
	}

	if err := WriteSimpleField[AmsPacket](ctx, "userdata", m.GetUserdata(), WriteComplex[AmsPacket](writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'userdata' field")
	}

	if popErr := writeBuffer.PopContext("AmsTCPPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsTCPPacket")
	}
	return nil
}

func (m *_AmsTCPPacket) IsAmsTCPPacket() {}

func (m *_AmsTCPPacket) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AmsTCPPacket) deepCopy() *_AmsTCPPacket {
	if m == nil {
		return nil
	}
	_AmsTCPPacketCopy := &_AmsTCPPacket{
		m.Userdata.DeepCopy().(AmsPacket),
		m.reservedField0,
	}
	return _AmsTCPPacketCopy
}

func (m *_AmsTCPPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
