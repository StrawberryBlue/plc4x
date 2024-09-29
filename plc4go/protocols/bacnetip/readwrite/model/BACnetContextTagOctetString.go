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

// BACnetContextTagOctetString is the corresponding interface of BACnetContextTagOctetString
type BACnetContextTagOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadOctetString
	// IsBACnetContextTagOctetString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagOctetString()
	// CreateBuilder creates a BACnetContextTagOctetStringBuilder
	CreateBACnetContextTagOctetStringBuilder() BACnetContextTagOctetStringBuilder
}

// _BACnetContextTagOctetString is the data-structure of this message
type _BACnetContextTagOctetString struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadOctetString
}

var _ BACnetContextTagOctetString = (*_BACnetContextTagOctetString)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagOctetString)(nil)

// NewBACnetContextTagOctetString factory function for _BACnetContextTagOctetString
func NewBACnetContextTagOctetString(header BACnetTagHeader, payload BACnetTagPayloadOctetString, tagNumberArgument uint8) *_BACnetContextTagOctetString {
	if payload == nil {
		panic("payload of type BACnetTagPayloadOctetString for BACnetContextTagOctetString must not be nil")
	}
	_result := &_BACnetContextTagOctetString{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagOctetStringBuilder is a builder for BACnetContextTagOctetString
type BACnetContextTagOctetStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadOctetString) BACnetContextTagOctetStringBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadOctetString) BACnetContextTagOctetStringBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadOctetStringBuilder) BACnetTagPayloadOctetStringBuilder) BACnetContextTagOctetStringBuilder
	// Build builds the BACnetContextTagOctetString or returns an error if something is wrong
	Build() (BACnetContextTagOctetString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagOctetString
}

// NewBACnetContextTagOctetStringBuilder() creates a BACnetContextTagOctetStringBuilder
func NewBACnetContextTagOctetStringBuilder() BACnetContextTagOctetStringBuilder {
	return &_BACnetContextTagOctetStringBuilder{_BACnetContextTagOctetString: new(_BACnetContextTagOctetString)}
}

type _BACnetContextTagOctetStringBuilder struct {
	*_BACnetContextTagOctetString

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagOctetStringBuilder) = (*_BACnetContextTagOctetStringBuilder)(nil)

func (b *_BACnetContextTagOctetStringBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
}

func (b *_BACnetContextTagOctetStringBuilder) WithMandatoryFields(payload BACnetTagPayloadOctetString) BACnetContextTagOctetStringBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetContextTagOctetStringBuilder) WithPayload(payload BACnetTagPayloadOctetString) BACnetContextTagOctetStringBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetContextTagOctetStringBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadOctetStringBuilder) BACnetTagPayloadOctetStringBuilder) BACnetContextTagOctetStringBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadOctetStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetContextTagOctetStringBuilder) Build() (BACnetContextTagOctetString, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagOctetString.deepCopy(), nil
}

func (b *_BACnetContextTagOctetStringBuilder) MustBuild() BACnetContextTagOctetString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetContextTagOctetStringBuilder) Done() BACnetContextTagBuilder {
	return b.parentBuilder
}

func (b *_BACnetContextTagOctetStringBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagOctetStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagOctetStringBuilder().(*_BACnetContextTagOctetStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagOctetStringBuilder creates a BACnetContextTagOctetStringBuilder
func (b *_BACnetContextTagOctetString) CreateBACnetContextTagOctetStringBuilder() BACnetContextTagOctetStringBuilder {
	if b == nil {
		return NewBACnetContextTagOctetStringBuilder()
	}
	return &_BACnetContextTagOctetStringBuilder{_BACnetContextTagOctetString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagOctetString) GetDataType() BACnetDataType {
	return BACnetDataType_OCTET_STRING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagOctetString) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagOctetString) GetPayload() BACnetTagPayloadOctetString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagOctetString(structType any) BACnetContextTagOctetString {
	if casted, ok := structType.(BACnetContextTagOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagOctetString) GetTypeName() string {
	return "BACnetContextTagOctetString"
}

func (m *_BACnetContextTagOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetContextTagOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagOctetString BACnetContextTagOctetString, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadOctetString](ctx, "payload", ReadComplex[BACnetTagPayloadOctetString](BACnetTagPayloadOctetStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetContextTagOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagOctetString")
	}

	return m, nil
}

func (m *_BACnetContextTagOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagOctetString")
		}

		if err := WriteSimpleField[BACnetTagPayloadOctetString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagOctetString")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagOctetString) IsBACnetContextTagOctetString() {}

func (m *_BACnetContextTagOctetString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagOctetString) deepCopy() *_BACnetContextTagOctetString {
	if m == nil {
		return nil
	}
	_BACnetContextTagOctetStringCopy := &_BACnetContextTagOctetString{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadOctetString),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagOctetStringCopy
}

func (m *_BACnetContextTagOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
