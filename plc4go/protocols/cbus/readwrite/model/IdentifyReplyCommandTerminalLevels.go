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

// IdentifyReplyCommandTerminalLevels is the corresponding interface of IdentifyReplyCommandTerminalLevels
type IdentifyReplyCommandTerminalLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetTerminalLevels returns TerminalLevels (property field)
	GetTerminalLevels() []byte
	// IsIdentifyReplyCommandTerminalLevels is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandTerminalLevels()
	// CreateBuilder creates a IdentifyReplyCommandTerminalLevelsBuilder
	CreateIdentifyReplyCommandTerminalLevelsBuilder() IdentifyReplyCommandTerminalLevelsBuilder
}

// _IdentifyReplyCommandTerminalLevels is the data-structure of this message
type _IdentifyReplyCommandTerminalLevels struct {
	IdentifyReplyCommandContract
	TerminalLevels []byte
}

var _ IdentifyReplyCommandTerminalLevels = (*_IdentifyReplyCommandTerminalLevels)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandTerminalLevels)(nil)

// NewIdentifyReplyCommandTerminalLevels factory function for _IdentifyReplyCommandTerminalLevels
func NewIdentifyReplyCommandTerminalLevels(terminalLevels []byte, numBytes uint8) *_IdentifyReplyCommandTerminalLevels {
	_result := &_IdentifyReplyCommandTerminalLevels{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		TerminalLevels:               terminalLevels,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandTerminalLevelsBuilder is a builder for IdentifyReplyCommandTerminalLevels
type IdentifyReplyCommandTerminalLevelsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(terminalLevels []byte) IdentifyReplyCommandTerminalLevelsBuilder
	// WithTerminalLevels adds TerminalLevels (property field)
	WithTerminalLevels(...byte) IdentifyReplyCommandTerminalLevelsBuilder
	// Build builds the IdentifyReplyCommandTerminalLevels or returns an error if something is wrong
	Build() (IdentifyReplyCommandTerminalLevels, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandTerminalLevels
}

// NewIdentifyReplyCommandTerminalLevelsBuilder() creates a IdentifyReplyCommandTerminalLevelsBuilder
func NewIdentifyReplyCommandTerminalLevelsBuilder() IdentifyReplyCommandTerminalLevelsBuilder {
	return &_IdentifyReplyCommandTerminalLevelsBuilder{_IdentifyReplyCommandTerminalLevels: new(_IdentifyReplyCommandTerminalLevels)}
}

type _IdentifyReplyCommandTerminalLevelsBuilder struct {
	*_IdentifyReplyCommandTerminalLevels

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandTerminalLevelsBuilder) = (*_IdentifyReplyCommandTerminalLevelsBuilder)(nil)

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
}

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) WithMandatoryFields(terminalLevels []byte) IdentifyReplyCommandTerminalLevelsBuilder {
	return b.WithTerminalLevels(terminalLevels...)
}

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) WithTerminalLevels(terminalLevels ...byte) IdentifyReplyCommandTerminalLevelsBuilder {
	b.TerminalLevels = terminalLevels
	return b
}

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) Build() (IdentifyReplyCommandTerminalLevels, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandTerminalLevels.deepCopy(), nil
}

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) MustBuild() IdentifyReplyCommandTerminalLevels {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_IdentifyReplyCommandTerminalLevelsBuilder) Done() IdentifyReplyCommandBuilder {
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandTerminalLevelsBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandTerminalLevelsBuilder().(*_IdentifyReplyCommandTerminalLevelsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandTerminalLevelsBuilder creates a IdentifyReplyCommandTerminalLevelsBuilder
func (b *_IdentifyReplyCommandTerminalLevels) CreateIdentifyReplyCommandTerminalLevelsBuilder() IdentifyReplyCommandTerminalLevelsBuilder {
	if b == nil {
		return NewIdentifyReplyCommandTerminalLevelsBuilder()
	}
	return &_IdentifyReplyCommandTerminalLevelsBuilder{_IdentifyReplyCommandTerminalLevels: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandTerminalLevels) GetAttribute() Attribute {
	return Attribute_TerminalLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandTerminalLevels) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandTerminalLevels) GetTerminalLevels() []byte {
	return m.TerminalLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandTerminalLevels(structType any) IdentifyReplyCommandTerminalLevels {
	if casted, ok := structType.(IdentifyReplyCommandTerminalLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandTerminalLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandTerminalLevels) GetTypeName() string {
	return "IdentifyReplyCommandTerminalLevels"
}

func (m *_IdentifyReplyCommandTerminalLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.TerminalLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.TerminalLevels))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandTerminalLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandTerminalLevels) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandTerminalLevels IdentifyReplyCommandTerminalLevels, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandTerminalLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandTerminalLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	terminalLevels, err := readBuffer.ReadByteArray("terminalLevels", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'terminalLevels' field"))
	}
	m.TerminalLevels = terminalLevels

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandTerminalLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandTerminalLevels")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandTerminalLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandTerminalLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandTerminalLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandTerminalLevels")
		}

		if err := WriteByteArrayField(ctx, "terminalLevels", m.GetTerminalLevels(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'terminalLevels' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandTerminalLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandTerminalLevels")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandTerminalLevels) IsIdentifyReplyCommandTerminalLevels() {}

func (m *_IdentifyReplyCommandTerminalLevels) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandTerminalLevels) deepCopy() *_IdentifyReplyCommandTerminalLevels {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandTerminalLevelsCopy := &_IdentifyReplyCommandTerminalLevels{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.TerminalLevels),
	}
	m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandTerminalLevelsCopy
}

func (m *_IdentifyReplyCommandTerminalLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
