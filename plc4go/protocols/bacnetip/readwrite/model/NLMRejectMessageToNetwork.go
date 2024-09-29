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

// NLMRejectMessageToNetwork is the corresponding interface of NLMRejectMessageToNetwork
type NLMRejectMessageToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() NLMRejectMessageToNetworkRejectReason
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// IsNLMRejectMessageToNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMRejectMessageToNetwork()
	// CreateBuilder creates a NLMRejectMessageToNetworkBuilder
	CreateNLMRejectMessageToNetworkBuilder() NLMRejectMessageToNetworkBuilder
}

// _NLMRejectMessageToNetwork is the data-structure of this message
type _NLMRejectMessageToNetwork struct {
	NLMContract
	RejectReason              NLMRejectMessageToNetworkRejectReason
	DestinationNetworkAddress uint16
}

var _ NLMRejectMessageToNetwork = (*_NLMRejectMessageToNetwork)(nil)
var _ NLMRequirements = (*_NLMRejectMessageToNetwork)(nil)

// NewNLMRejectMessageToNetwork factory function for _NLMRejectMessageToNetwork
func NewNLMRejectMessageToNetwork(rejectReason NLMRejectMessageToNetworkRejectReason, destinationNetworkAddress uint16, apduLength uint16) *_NLMRejectMessageToNetwork {
	_result := &_NLMRejectMessageToNetwork{
		NLMContract:               NewNLM(apduLength),
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMRejectMessageToNetworkBuilder is a builder for NLMRejectMessageToNetwork
type NLMRejectMessageToNetworkBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(rejectReason NLMRejectMessageToNetworkRejectReason, destinationNetworkAddress uint16) NLMRejectMessageToNetworkBuilder
	// WithRejectReason adds RejectReason (property field)
	WithRejectReason(NLMRejectMessageToNetworkRejectReason) NLMRejectMessageToNetworkBuilder
	// WithDestinationNetworkAddress adds DestinationNetworkAddress (property field)
	WithDestinationNetworkAddress(uint16) NLMRejectMessageToNetworkBuilder
	// Build builds the NLMRejectMessageToNetwork or returns an error if something is wrong
	Build() (NLMRejectMessageToNetwork, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMRejectMessageToNetwork
}

// NewNLMRejectMessageToNetworkBuilder() creates a NLMRejectMessageToNetworkBuilder
func NewNLMRejectMessageToNetworkBuilder() NLMRejectMessageToNetworkBuilder {
	return &_NLMRejectMessageToNetworkBuilder{_NLMRejectMessageToNetwork: new(_NLMRejectMessageToNetwork)}
}

type _NLMRejectMessageToNetworkBuilder struct {
	*_NLMRejectMessageToNetwork

	parentBuilder *_NLMBuilder

	err *utils.MultiError
}

var _ (NLMRejectMessageToNetworkBuilder) = (*_NLMRejectMessageToNetworkBuilder)(nil)

func (b *_NLMRejectMessageToNetworkBuilder) setParent(contract NLMContract) {
	b.NLMContract = contract
}

func (b *_NLMRejectMessageToNetworkBuilder) WithMandatoryFields(rejectReason NLMRejectMessageToNetworkRejectReason, destinationNetworkAddress uint16) NLMRejectMessageToNetworkBuilder {
	return b.WithRejectReason(rejectReason).WithDestinationNetworkAddress(destinationNetworkAddress)
}

func (b *_NLMRejectMessageToNetworkBuilder) WithRejectReason(rejectReason NLMRejectMessageToNetworkRejectReason) NLMRejectMessageToNetworkBuilder {
	b.RejectReason = rejectReason
	return b
}

func (b *_NLMRejectMessageToNetworkBuilder) WithDestinationNetworkAddress(destinationNetworkAddress uint16) NLMRejectMessageToNetworkBuilder {
	b.DestinationNetworkAddress = destinationNetworkAddress
	return b
}

func (b *_NLMRejectMessageToNetworkBuilder) Build() (NLMRejectMessageToNetwork, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLMRejectMessageToNetwork.deepCopy(), nil
}

func (b *_NLMRejectMessageToNetworkBuilder) MustBuild() NLMRejectMessageToNetwork {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NLMRejectMessageToNetworkBuilder) Done() NLMBuilder {
	return b.parentBuilder
}

func (b *_NLMRejectMessageToNetworkBuilder) buildForNLM() (NLM, error) {
	return b.Build()
}

func (b *_NLMRejectMessageToNetworkBuilder) DeepCopy() any {
	_copy := b.CreateNLMRejectMessageToNetworkBuilder().(*_NLMRejectMessageToNetworkBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMRejectMessageToNetworkBuilder creates a NLMRejectMessageToNetworkBuilder
func (b *_NLMRejectMessageToNetwork) CreateNLMRejectMessageToNetworkBuilder() NLMRejectMessageToNetworkBuilder {
	if b == nil {
		return NewNLMRejectMessageToNetworkBuilder()
	}
	return &_NLMRejectMessageToNetworkBuilder{_NLMRejectMessageToNetwork: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRejectMessageToNetwork) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRejectMessageToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRejectMessageToNetwork) GetRejectReason() NLMRejectMessageToNetworkRejectReason {
	return m.RejectReason
}

func (m *_NLMRejectMessageToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMRejectMessageToNetwork(structType any) NLMRejectMessageToNetwork {
	if casted, ok := structType.(NLMRejectMessageToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRejectMessageToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRejectMessageToNetwork) GetTypeName() string {
	return "NLMRejectMessageToNetwork"
}

func (m *_NLMRejectMessageToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (rejectReason)
	lengthInBits += 8

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_NLMRejectMessageToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMRejectMessageToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMRejectMessageToNetwork NLMRejectMessageToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRejectMessageToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRejectMessageToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	rejectReason, err := ReadEnumField[NLMRejectMessageToNetworkRejectReason](ctx, "rejectReason", "NLMRejectMessageToNetworkRejectReason", ReadEnum(NLMRejectMessageToNetworkRejectReasonByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectReason' field"))
	}
	m.RejectReason = rejectReason

	destinationNetworkAddress, err := ReadSimpleField(ctx, "destinationNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddress' field"))
	}
	m.DestinationNetworkAddress = destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMRejectMessageToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRejectMessageToNetwork")
	}

	return m, nil
}

func (m *_NLMRejectMessageToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRejectMessageToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRejectMessageToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRejectMessageToNetwork")
		}

		if err := WriteSimpleEnumField[NLMRejectMessageToNetworkRejectReason](ctx, "rejectReason", "NLMRejectMessageToNetworkRejectReason", m.GetRejectReason(), WriteEnum[NLMRejectMessageToNetworkRejectReason, uint8](NLMRejectMessageToNetworkRejectReason.GetValue, NLMRejectMessageToNetworkRejectReason.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectReason' field")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationNetworkAddress", m.GetDestinationNetworkAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMRejectMessageToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRejectMessageToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMRejectMessageToNetwork) IsNLMRejectMessageToNetwork() {}

func (m *_NLMRejectMessageToNetwork) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMRejectMessageToNetwork) deepCopy() *_NLMRejectMessageToNetwork {
	if m == nil {
		return nil
	}
	_NLMRejectMessageToNetworkCopy := &_NLMRejectMessageToNetwork{
		m.NLMContract.(*_NLM).deepCopy(),
		m.RejectReason,
		m.DestinationNetworkAddress,
	}
	m.NLMContract.(*_NLM)._SubType = m
	return _NLMRejectMessageToNetworkCopy
}

func (m *_NLMRejectMessageToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
