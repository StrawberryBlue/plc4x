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

// ApduDataExtWriteRoutingTableRequest is the corresponding interface of ApduDataExtWriteRoutingTableRequest
type ApduDataExtWriteRoutingTableRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtWriteRoutingTableRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtWriteRoutingTableRequest()
	// CreateBuilder creates a ApduDataExtWriteRoutingTableRequestBuilder
	CreateApduDataExtWriteRoutingTableRequestBuilder() ApduDataExtWriteRoutingTableRequestBuilder
}

// _ApduDataExtWriteRoutingTableRequest is the data-structure of this message
type _ApduDataExtWriteRoutingTableRequest struct {
	ApduDataExtContract
}

var _ ApduDataExtWriteRoutingTableRequest = (*_ApduDataExtWriteRoutingTableRequest)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtWriteRoutingTableRequest)(nil)

// NewApduDataExtWriteRoutingTableRequest factory function for _ApduDataExtWriteRoutingTableRequest
func NewApduDataExtWriteRoutingTableRequest(length uint8) *_ApduDataExtWriteRoutingTableRequest {
	_result := &_ApduDataExtWriteRoutingTableRequest{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtWriteRoutingTableRequestBuilder is a builder for ApduDataExtWriteRoutingTableRequest
type ApduDataExtWriteRoutingTableRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtWriteRoutingTableRequestBuilder
	// Build builds the ApduDataExtWriteRoutingTableRequest or returns an error if something is wrong
	Build() (ApduDataExtWriteRoutingTableRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtWriteRoutingTableRequest
}

// NewApduDataExtWriteRoutingTableRequestBuilder() creates a ApduDataExtWriteRoutingTableRequestBuilder
func NewApduDataExtWriteRoutingTableRequestBuilder() ApduDataExtWriteRoutingTableRequestBuilder {
	return &_ApduDataExtWriteRoutingTableRequestBuilder{_ApduDataExtWriteRoutingTableRequest: new(_ApduDataExtWriteRoutingTableRequest)}
}

type _ApduDataExtWriteRoutingTableRequestBuilder struct {
	*_ApduDataExtWriteRoutingTableRequest

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtWriteRoutingTableRequestBuilder) = (*_ApduDataExtWriteRoutingTableRequestBuilder)(nil)

func (b *_ApduDataExtWriteRoutingTableRequestBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtWriteRoutingTableRequestBuilder) WithMandatoryFields() ApduDataExtWriteRoutingTableRequestBuilder {
	return b
}

func (b *_ApduDataExtWriteRoutingTableRequestBuilder) Build() (ApduDataExtWriteRoutingTableRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtWriteRoutingTableRequest.deepCopy(), nil
}

func (b *_ApduDataExtWriteRoutingTableRequestBuilder) MustBuild() ApduDataExtWriteRoutingTableRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtWriteRoutingTableRequestBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtWriteRoutingTableRequestBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtWriteRoutingTableRequestBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtWriteRoutingTableRequestBuilder().(*_ApduDataExtWriteRoutingTableRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtWriteRoutingTableRequestBuilder creates a ApduDataExtWriteRoutingTableRequestBuilder
func (b *_ApduDataExtWriteRoutingTableRequest) CreateApduDataExtWriteRoutingTableRequestBuilder() ApduDataExtWriteRoutingTableRequestBuilder {
	if b == nil {
		return NewApduDataExtWriteRoutingTableRequestBuilder()
	}
	return &_ApduDataExtWriteRoutingTableRequestBuilder{_ApduDataExtWriteRoutingTableRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtWriteRoutingTableRequest) GetExtApciType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtWriteRoutingTableRequest) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtWriteRoutingTableRequest(structType any) ApduDataExtWriteRoutingTableRequest {
	if casted, ok := structType.(ApduDataExtWriteRoutingTableRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtWriteRoutingTableRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtWriteRoutingTableRequest) GetTypeName() string {
	return "ApduDataExtWriteRoutingTableRequest"
}

func (m *_ApduDataExtWriteRoutingTableRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtWriteRoutingTableRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtWriteRoutingTableRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtWriteRoutingTableRequest ApduDataExtWriteRoutingTableRequest, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtWriteRoutingTableRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtWriteRoutingTableRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtWriteRoutingTableRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtWriteRoutingTableRequest")
	}

	return m, nil
}

func (m *_ApduDataExtWriteRoutingTableRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtWriteRoutingTableRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtWriteRoutingTableRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtWriteRoutingTableRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtWriteRoutingTableRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtWriteRoutingTableRequest")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtWriteRoutingTableRequest) IsApduDataExtWriteRoutingTableRequest() {}

func (m *_ApduDataExtWriteRoutingTableRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtWriteRoutingTableRequest) deepCopy() *_ApduDataExtWriteRoutingTableRequest {
	if m == nil {
		return nil
	}
	_ApduDataExtWriteRoutingTableRequestCopy := &_ApduDataExtWriteRoutingTableRequest{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtWriteRoutingTableRequestCopy
}

func (m *_ApduDataExtWriteRoutingTableRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
