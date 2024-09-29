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

// MPropInfoInd is the corresponding interface of MPropInfoInd
type MPropInfoInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMPropInfoInd is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMPropInfoInd()
	// CreateBuilder creates a MPropInfoIndBuilder
	CreateMPropInfoIndBuilder() MPropInfoIndBuilder
}

// _MPropInfoInd is the data-structure of this message
type _MPropInfoInd struct {
	CEMIContract
}

var _ MPropInfoInd = (*_MPropInfoInd)(nil)
var _ CEMIRequirements = (*_MPropInfoInd)(nil)

// NewMPropInfoInd factory function for _MPropInfoInd
func NewMPropInfoInd(size uint16) *_MPropInfoInd {
	_result := &_MPropInfoInd{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MPropInfoIndBuilder is a builder for MPropInfoInd
type MPropInfoIndBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MPropInfoIndBuilder
	// Build builds the MPropInfoInd or returns an error if something is wrong
	Build() (MPropInfoInd, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MPropInfoInd
}

// NewMPropInfoIndBuilder() creates a MPropInfoIndBuilder
func NewMPropInfoIndBuilder() MPropInfoIndBuilder {
	return &_MPropInfoIndBuilder{_MPropInfoInd: new(_MPropInfoInd)}
}

type _MPropInfoIndBuilder struct {
	*_MPropInfoInd

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MPropInfoIndBuilder) = (*_MPropInfoIndBuilder)(nil)

func (b *_MPropInfoIndBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_MPropInfoIndBuilder) WithMandatoryFields() MPropInfoIndBuilder {
	return b
}

func (b *_MPropInfoIndBuilder) Build() (MPropInfoInd, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MPropInfoInd.deepCopy(), nil
}

func (b *_MPropInfoIndBuilder) MustBuild() MPropInfoInd {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MPropInfoIndBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_MPropInfoIndBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MPropInfoIndBuilder) DeepCopy() any {
	_copy := b.CreateMPropInfoIndBuilder().(*_MPropInfoIndBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMPropInfoIndBuilder creates a MPropInfoIndBuilder
func (b *_MPropInfoInd) CreateMPropInfoIndBuilder() MPropInfoIndBuilder {
	if b == nil {
		return NewMPropInfoIndBuilder()
	}
	return &_MPropInfoIndBuilder{_MPropInfoInd: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropInfoInd) GetMessageCode() uint8 {
	return 0xF7
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropInfoInd) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMPropInfoInd(structType any) MPropInfoInd {
	if casted, ok := structType.(MPropInfoInd); ok {
		return casted
	}
	if casted, ok := structType.(*MPropInfoInd); ok {
		return *casted
	}
	return nil
}

func (m *_MPropInfoInd) GetTypeName() string {
	return "MPropInfoInd"
}

func (m *_MPropInfoInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MPropInfoInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropInfoInd) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mPropInfoInd MPropInfoInd, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropInfoInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropInfoInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropInfoInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropInfoInd")
	}

	return m, nil
}

func (m *_MPropInfoInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropInfoInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropInfoInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropInfoInd")
		}

		if popErr := writeBuffer.PopContext("MPropInfoInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropInfoInd")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropInfoInd) IsMPropInfoInd() {}

func (m *_MPropInfoInd) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MPropInfoInd) deepCopy() *_MPropInfoInd {
	if m == nil {
		return nil
	}
	_MPropInfoIndCopy := &_MPropInfoInd{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MPropInfoIndCopy
}

func (m *_MPropInfoInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
