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

// ListServicesRequest is the corresponding interface of ListServicesRequest
type ListServicesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// IsListServicesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsListServicesRequest()
	// CreateBuilder creates a ListServicesRequestBuilder
	CreateListServicesRequestBuilder() ListServicesRequestBuilder
}

// _ListServicesRequest is the data-structure of this message
type _ListServicesRequest struct {
	EipPacketContract
}

var _ ListServicesRequest = (*_ListServicesRequest)(nil)
var _ EipPacketRequirements = (*_ListServicesRequest)(nil)

// NewListServicesRequest factory function for _ListServicesRequest
func NewListServicesRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_ListServicesRequest {
	_result := &_ListServicesRequest{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ListServicesRequestBuilder is a builder for ListServicesRequest
type ListServicesRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ListServicesRequestBuilder
	// Build builds the ListServicesRequest or returns an error if something is wrong
	Build() (ListServicesRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ListServicesRequest
}

// NewListServicesRequestBuilder() creates a ListServicesRequestBuilder
func NewListServicesRequestBuilder() ListServicesRequestBuilder {
	return &_ListServicesRequestBuilder{_ListServicesRequest: new(_ListServicesRequest)}
}

type _ListServicesRequestBuilder struct {
	*_ListServicesRequest

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (ListServicesRequestBuilder) = (*_ListServicesRequestBuilder)(nil)

func (b *_ListServicesRequestBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
}

func (b *_ListServicesRequestBuilder) WithMandatoryFields() ListServicesRequestBuilder {
	return b
}

func (b *_ListServicesRequestBuilder) Build() (ListServicesRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ListServicesRequest.deepCopy(), nil
}

func (b *_ListServicesRequestBuilder) MustBuild() ListServicesRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ListServicesRequestBuilder) Done() EipPacketBuilder {
	return b.parentBuilder
}

func (b *_ListServicesRequestBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_ListServicesRequestBuilder) DeepCopy() any {
	_copy := b.CreateListServicesRequestBuilder().(*_ListServicesRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateListServicesRequestBuilder creates a ListServicesRequestBuilder
func (b *_ListServicesRequest) CreateListServicesRequestBuilder() ListServicesRequestBuilder {
	if b == nil {
		return NewListServicesRequestBuilder()
	}
	return &_ListServicesRequestBuilder{_ListServicesRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ListServicesRequest) GetCommand() uint16 {
	return 0x0004
}

func (m *_ListServicesRequest) GetResponse() bool {
	return bool(false)
}

func (m *_ListServicesRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ListServicesRequest) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// Deprecated: use the interface for direct cast
func CastListServicesRequest(structType any) ListServicesRequest {
	if casted, ok := structType.(ListServicesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ListServicesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ListServicesRequest) GetTypeName() string {
	return "ListServicesRequest"
}

func (m *_ListServicesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ListServicesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ListServicesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__listServicesRequest ListServicesRequest, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ListServicesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListServicesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ListServicesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListServicesRequest")
	}

	return m, nil
}

func (m *_ListServicesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListServicesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ListServicesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ListServicesRequest")
		}

		if popErr := writeBuffer.PopContext("ListServicesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ListServicesRequest")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ListServicesRequest) IsListServicesRequest() {}

func (m *_ListServicesRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ListServicesRequest) deepCopy() *_ListServicesRequest {
	if m == nil {
		return nil
	}
	_ListServicesRequestCopy := &_ListServicesRequest{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
	}
	m.EipPacketContract.(*_EipPacket)._SubType = m
	return _ListServicesRequestCopy
}

func (m *_ListServicesRequest) String() string {
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
