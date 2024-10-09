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

// IdentifyReplyCommandSummary is the corresponding interface of IdentifyReplyCommandSummary
type IdentifyReplyCommandSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetPartName returns PartName (property field)
	GetPartName() string
	// GetUnitServiceType returns UnitServiceType (property field)
	GetUnitServiceType() byte
	// GetVersion returns Version (property field)
	GetVersion() string
	// IsIdentifyReplyCommandSummary is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandSummary()
	// CreateBuilder creates a IdentifyReplyCommandSummaryBuilder
	CreateIdentifyReplyCommandSummaryBuilder() IdentifyReplyCommandSummaryBuilder
}

// _IdentifyReplyCommandSummary is the data-structure of this message
type _IdentifyReplyCommandSummary struct {
	IdentifyReplyCommandContract
	PartName        string
	UnitServiceType byte
	Version         string
}

var _ IdentifyReplyCommandSummary = (*_IdentifyReplyCommandSummary)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandSummary)(nil)

// NewIdentifyReplyCommandSummary factory function for _IdentifyReplyCommandSummary
func NewIdentifyReplyCommandSummary(partName string, unitServiceType byte, version string, numBytes uint8) *_IdentifyReplyCommandSummary {
	_result := &_IdentifyReplyCommandSummary{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		PartName:                     partName,
		UnitServiceType:              unitServiceType,
		Version:                      version,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandSummaryBuilder is a builder for IdentifyReplyCommandSummary
type IdentifyReplyCommandSummaryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(partName string, unitServiceType byte, version string) IdentifyReplyCommandSummaryBuilder
	// WithPartName adds PartName (property field)
	WithPartName(string) IdentifyReplyCommandSummaryBuilder
	// WithUnitServiceType adds UnitServiceType (property field)
	WithUnitServiceType(byte) IdentifyReplyCommandSummaryBuilder
	// WithVersion adds Version (property field)
	WithVersion(string) IdentifyReplyCommandSummaryBuilder
	// Build builds the IdentifyReplyCommandSummary or returns an error if something is wrong
	Build() (IdentifyReplyCommandSummary, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandSummary
}

// NewIdentifyReplyCommandSummaryBuilder() creates a IdentifyReplyCommandSummaryBuilder
func NewIdentifyReplyCommandSummaryBuilder() IdentifyReplyCommandSummaryBuilder {
	return &_IdentifyReplyCommandSummaryBuilder{_IdentifyReplyCommandSummary: new(_IdentifyReplyCommandSummary)}
}

type _IdentifyReplyCommandSummaryBuilder struct {
	*_IdentifyReplyCommandSummary

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandSummaryBuilder) = (*_IdentifyReplyCommandSummaryBuilder)(nil)

func (b *_IdentifyReplyCommandSummaryBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
}

func (b *_IdentifyReplyCommandSummaryBuilder) WithMandatoryFields(partName string, unitServiceType byte, version string) IdentifyReplyCommandSummaryBuilder {
	return b.WithPartName(partName).WithUnitServiceType(unitServiceType).WithVersion(version)
}

func (b *_IdentifyReplyCommandSummaryBuilder) WithPartName(partName string) IdentifyReplyCommandSummaryBuilder {
	b.PartName = partName
	return b
}

func (b *_IdentifyReplyCommandSummaryBuilder) WithUnitServiceType(unitServiceType byte) IdentifyReplyCommandSummaryBuilder {
	b.UnitServiceType = unitServiceType
	return b
}

func (b *_IdentifyReplyCommandSummaryBuilder) WithVersion(version string) IdentifyReplyCommandSummaryBuilder {
	b.Version = version
	return b
}

func (b *_IdentifyReplyCommandSummaryBuilder) Build() (IdentifyReplyCommandSummary, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandSummary.deepCopy(), nil
}

func (b *_IdentifyReplyCommandSummaryBuilder) MustBuild() IdentifyReplyCommandSummary {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_IdentifyReplyCommandSummaryBuilder) Done() IdentifyReplyCommandBuilder {
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandSummaryBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandSummaryBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandSummaryBuilder().(*_IdentifyReplyCommandSummaryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandSummaryBuilder creates a IdentifyReplyCommandSummaryBuilder
func (b *_IdentifyReplyCommandSummary) CreateIdentifyReplyCommandSummaryBuilder() IdentifyReplyCommandSummaryBuilder {
	if b == nil {
		return NewIdentifyReplyCommandSummaryBuilder()
	}
	return &_IdentifyReplyCommandSummaryBuilder{_IdentifyReplyCommandSummary: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandSummary) GetAttribute() Attribute {
	return Attribute_Summary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandSummary) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandSummary) GetPartName() string {
	return m.PartName
}

func (m *_IdentifyReplyCommandSummary) GetUnitServiceType() byte {
	return m.UnitServiceType
}

func (m *_IdentifyReplyCommandSummary) GetVersion() string {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandSummary(structType any) IdentifyReplyCommandSummary {
	if casted, ok := structType.(IdentifyReplyCommandSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandSummary) GetTypeName() string {
	return "IdentifyReplyCommandSummary"
}

func (m *_IdentifyReplyCommandSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Simple field (partName)
	lengthInBits += 48

	// Simple field (unitServiceType)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 32

	return lengthInBits
}

func (m *_IdentifyReplyCommandSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandSummary IdentifyReplyCommandSummary, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	partName, err := ReadSimpleField(ctx, "partName", ReadString(readBuffer, uint32(48)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'partName' field"))
	}
	m.PartName = partName

	unitServiceType, err := ReadSimpleField(ctx, "unitServiceType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitServiceType' field"))
	}
	m.UnitServiceType = unitServiceType

	version, err := ReadSimpleField(ctx, "version", ReadString(readBuffer, uint32(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandSummary")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandSummary")
		}

		if err := WriteSimpleField[string](ctx, "partName", m.GetPartName(), WriteString(writeBuffer, 48)); err != nil {
			return errors.Wrap(err, "Error serializing 'partName' field")
		}

		if err := WriteSimpleField[byte](ctx, "unitServiceType", m.GetUnitServiceType(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitServiceType' field")
		}

		if err := WriteSimpleField[string](ctx, "version", m.GetVersion(), WriteString(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandSummary")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandSummary) IsIdentifyReplyCommandSummary() {}

func (m *_IdentifyReplyCommandSummary) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandSummary) deepCopy() *_IdentifyReplyCommandSummary {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandSummaryCopy := &_IdentifyReplyCommandSummary{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		m.PartName,
		m.UnitServiceType,
		m.Version,
	}
	m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandSummaryCopy
}

func (m *_IdentifyReplyCommandSummary) String() string {
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
