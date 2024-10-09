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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TrimmedString is the corresponding interface of TrimmedString
type TrimmedString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsTrimmedString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTrimmedString()
	// CreateBuilder creates a TrimmedStringBuilder
	CreateTrimmedStringBuilder() TrimmedStringBuilder
}

// _TrimmedString is the data-structure of this message
type _TrimmedString struct {
}

var _ TrimmedString = (*_TrimmedString)(nil)

// NewTrimmedString factory function for _TrimmedString
func NewTrimmedString() *_TrimmedString {
	return &_TrimmedString{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TrimmedStringBuilder is a builder for TrimmedString
type TrimmedStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() TrimmedStringBuilder
	// Build builds the TrimmedString or returns an error if something is wrong
	Build() (TrimmedString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TrimmedString
}

// NewTrimmedStringBuilder() creates a TrimmedStringBuilder
func NewTrimmedStringBuilder() TrimmedStringBuilder {
	return &_TrimmedStringBuilder{_TrimmedString: new(_TrimmedString)}
}

type _TrimmedStringBuilder struct {
	*_TrimmedString

	err *utils.MultiError
}

var _ (TrimmedStringBuilder) = (*_TrimmedStringBuilder)(nil)

func (b *_TrimmedStringBuilder) WithMandatoryFields() TrimmedStringBuilder {
	return b
}

func (b *_TrimmedStringBuilder) Build() (TrimmedString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TrimmedString.deepCopy(), nil
}

func (b *_TrimmedStringBuilder) MustBuild() TrimmedString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TrimmedStringBuilder) DeepCopy() any {
	_copy := b.CreateTrimmedStringBuilder().(*_TrimmedStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTrimmedStringBuilder creates a TrimmedStringBuilder
func (b *_TrimmedString) CreateTrimmedStringBuilder() TrimmedStringBuilder {
	if b == nil {
		return NewTrimmedStringBuilder()
	}
	return &_TrimmedStringBuilder{_TrimmedString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTrimmedString(structType any) TrimmedString {
	if casted, ok := structType.(TrimmedString); ok {
		return casted
	}
	if casted, ok := structType.(*TrimmedString); ok {
		return *casted
	}
	return nil
}

func (m *_TrimmedString) GetTypeName() string {
	return "TrimmedString"
}

func (m *_TrimmedString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_TrimmedString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TrimmedStringParse(ctx context.Context, theBytes []byte) (TrimmedString, error) {
	return TrimmedStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TrimmedStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TrimmedString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TrimmedString, error) {
		return TrimmedStringParseWithBuffer(ctx, readBuffer)
	}
}

func TrimmedStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TrimmedString, error) {
	v, err := (&_TrimmedString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_TrimmedString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__trimmedString TrimmedString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TrimmedString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TrimmedString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TrimmedString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TrimmedString")
	}

	return m, nil
}

func (m *_TrimmedString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TrimmedString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TrimmedString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TrimmedString")
	}

	if popErr := writeBuffer.PopContext("TrimmedString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TrimmedString")
	}
	return nil
}

func (m *_TrimmedString) IsTrimmedString() {}

func (m *_TrimmedString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TrimmedString) deepCopy() *_TrimmedString {
	if m == nil {
		return nil
	}
	_TrimmedStringCopy := &_TrimmedString{}
	return _TrimmedStringCopy
}

func (m *_TrimmedString) String() string {
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
