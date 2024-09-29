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

// SecurityDataStatus2Request is the corresponding interface of SecurityDataStatus2Request
type SecurityDataStatus2Request interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataStatus2Request is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataStatus2Request()
	// CreateBuilder creates a SecurityDataStatus2RequestBuilder
	CreateSecurityDataStatus2RequestBuilder() SecurityDataStatus2RequestBuilder
}

// _SecurityDataStatus2Request is the data-structure of this message
type _SecurityDataStatus2Request struct {
	SecurityDataContract
}

var _ SecurityDataStatus2Request = (*_SecurityDataStatus2Request)(nil)
var _ SecurityDataRequirements = (*_SecurityDataStatus2Request)(nil)

// NewSecurityDataStatus2Request factory function for _SecurityDataStatus2Request
func NewSecurityDataStatus2Request(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatus2Request {
	_result := &_SecurityDataStatus2Request{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataStatus2RequestBuilder is a builder for SecurityDataStatus2Request
type SecurityDataStatus2RequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataStatus2RequestBuilder
	// Build builds the SecurityDataStatus2Request or returns an error if something is wrong
	Build() (SecurityDataStatus2Request, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataStatus2Request
}

// NewSecurityDataStatus2RequestBuilder() creates a SecurityDataStatus2RequestBuilder
func NewSecurityDataStatus2RequestBuilder() SecurityDataStatus2RequestBuilder {
	return &_SecurityDataStatus2RequestBuilder{_SecurityDataStatus2Request: new(_SecurityDataStatus2Request)}
}

type _SecurityDataStatus2RequestBuilder struct {
	*_SecurityDataStatus2Request

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataStatus2RequestBuilder) = (*_SecurityDataStatus2RequestBuilder)(nil)

func (b *_SecurityDataStatus2RequestBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
}

func (b *_SecurityDataStatus2RequestBuilder) WithMandatoryFields() SecurityDataStatus2RequestBuilder {
	return b
}

func (b *_SecurityDataStatus2RequestBuilder) Build() (SecurityDataStatus2Request, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataStatus2Request.deepCopy(), nil
}

func (b *_SecurityDataStatus2RequestBuilder) MustBuild() SecurityDataStatus2Request {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SecurityDataStatus2RequestBuilder) Done() SecurityDataBuilder {
	return b.parentBuilder
}

func (b *_SecurityDataStatus2RequestBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataStatus2RequestBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataStatus2RequestBuilder().(*_SecurityDataStatus2RequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataStatus2RequestBuilder creates a SecurityDataStatus2RequestBuilder
func (b *_SecurityDataStatus2Request) CreateSecurityDataStatus2RequestBuilder() SecurityDataStatus2RequestBuilder {
	if b == nil {
		return NewSecurityDataStatus2RequestBuilder()
	}
	return &_SecurityDataStatus2RequestBuilder{_SecurityDataStatus2Request: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataStatus2Request) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatus2Request(structType any) SecurityDataStatus2Request {
	if casted, ok := structType.(SecurityDataStatus2Request); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatus2Request); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatus2Request) GetTypeName() string {
	return "SecurityDataStatus2Request"
}

func (m *_SecurityDataStatus2Request) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataStatus2Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataStatus2Request) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataStatus2Request SecurityDataStatus2Request, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataStatus2Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatus2Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataStatus2Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatus2Request")
	}

	return m, nil
}

func (m *_SecurityDataStatus2Request) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatus2Request) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatus2Request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatus2Request")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatus2Request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatus2Request")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatus2Request) IsSecurityDataStatus2Request() {}

func (m *_SecurityDataStatus2Request) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataStatus2Request) deepCopy() *_SecurityDataStatus2Request {
	if m == nil {
		return nil
	}
	_SecurityDataStatus2RequestCopy := &_SecurityDataStatus2Request{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataStatus2RequestCopy
}

func (m *_SecurityDataStatus2Request) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
