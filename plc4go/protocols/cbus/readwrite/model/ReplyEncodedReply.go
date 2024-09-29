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
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ReplyEncodedReply is the corresponding interface of ReplyEncodedReply
type ReplyEncodedReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Reply
	// GetEncodedReply returns EncodedReply (property field)
	GetEncodedReply() EncodedReply
	// GetChksum returns Chksum (property field)
	GetChksum() Checksum
	// GetEncodedReplyDecoded returns EncodedReplyDecoded (virtual field)
	GetEncodedReplyDecoded() EncodedReply
	// GetChksumDecoded returns ChksumDecoded (virtual field)
	GetChksumDecoded() Checksum
	// IsReplyEncodedReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReplyEncodedReply()
	// CreateBuilder creates a ReplyEncodedReplyBuilder
	CreateReplyEncodedReplyBuilder() ReplyEncodedReplyBuilder
}

// _ReplyEncodedReply is the data-structure of this message
type _ReplyEncodedReply struct {
	ReplyContract
	EncodedReply EncodedReply
	Chksum       Checksum
}

var _ ReplyEncodedReply = (*_ReplyEncodedReply)(nil)
var _ ReplyRequirements = (*_ReplyEncodedReply)(nil)

// NewReplyEncodedReply factory function for _ReplyEncodedReply
func NewReplyEncodedReply(peekedByte byte, encodedReply EncodedReply, chksum Checksum, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyEncodedReply {
	_result := &_ReplyEncodedReply{
		ReplyContract: NewReply(peekedByte, cBusOptions, requestContext),
		EncodedReply:  encodedReply,
		Chksum:        chksum,
	}
	_result.ReplyContract.(*_Reply)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReplyEncodedReplyBuilder is a builder for ReplyEncodedReply
type ReplyEncodedReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(encodedReply EncodedReply, chksum Checksum) ReplyEncodedReplyBuilder
	// WithEncodedReply adds EncodedReply (property field)
	WithEncodedReply(EncodedReply) ReplyEncodedReplyBuilder
	// WithEncodedReplyBuilder adds EncodedReply (property field) which is build by the builder
	WithEncodedReplyBuilder(func(EncodedReplyBuilder) EncodedReplyBuilder) ReplyEncodedReplyBuilder
	// WithChksum adds Chksum (property field)
	WithChksum(Checksum) ReplyEncodedReplyBuilder
	// WithChksumBuilder adds Chksum (property field) which is build by the builder
	WithChksumBuilder(func(ChecksumBuilder) ChecksumBuilder) ReplyEncodedReplyBuilder
	// Build builds the ReplyEncodedReply or returns an error if something is wrong
	Build() (ReplyEncodedReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReplyEncodedReply
}

// NewReplyEncodedReplyBuilder() creates a ReplyEncodedReplyBuilder
func NewReplyEncodedReplyBuilder() ReplyEncodedReplyBuilder {
	return &_ReplyEncodedReplyBuilder{_ReplyEncodedReply: new(_ReplyEncodedReply)}
}

type _ReplyEncodedReplyBuilder struct {
	*_ReplyEncodedReply

	parentBuilder *_ReplyBuilder

	err *utils.MultiError
}

var _ (ReplyEncodedReplyBuilder) = (*_ReplyEncodedReplyBuilder)(nil)

func (b *_ReplyEncodedReplyBuilder) setParent(contract ReplyContract) {
	b.ReplyContract = contract
}

func (b *_ReplyEncodedReplyBuilder) WithMandatoryFields(encodedReply EncodedReply, chksum Checksum) ReplyEncodedReplyBuilder {
	return b.WithEncodedReply(encodedReply).WithChksum(chksum)
}

func (b *_ReplyEncodedReplyBuilder) WithEncodedReply(encodedReply EncodedReply) ReplyEncodedReplyBuilder {
	b.EncodedReply = encodedReply
	return b
}

func (b *_ReplyEncodedReplyBuilder) WithEncodedReplyBuilder(builderSupplier func(EncodedReplyBuilder) EncodedReplyBuilder) ReplyEncodedReplyBuilder {
	builder := builderSupplier(b.EncodedReply.CreateEncodedReplyBuilder())
	var err error
	b.EncodedReply, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "EncodedReplyBuilder failed"))
	}
	return b
}

func (b *_ReplyEncodedReplyBuilder) WithChksum(chksum Checksum) ReplyEncodedReplyBuilder {
	b.Chksum = chksum
	return b
}

func (b *_ReplyEncodedReplyBuilder) WithChksumBuilder(builderSupplier func(ChecksumBuilder) ChecksumBuilder) ReplyEncodedReplyBuilder {
	builder := builderSupplier(b.Chksum.CreateChecksumBuilder())
	var err error
	b.Chksum, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ChecksumBuilder failed"))
	}
	return b
}

func (b *_ReplyEncodedReplyBuilder) Build() (ReplyEncodedReply, error) {
	if b.EncodedReply == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'encodedReply' not set"))
	}
	if b.Chksum == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'chksum' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReplyEncodedReply.deepCopy(), nil
}

func (b *_ReplyEncodedReplyBuilder) MustBuild() ReplyEncodedReply {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ReplyEncodedReplyBuilder) Done() ReplyBuilder {
	return b.parentBuilder
}

func (b *_ReplyEncodedReplyBuilder) buildForReply() (Reply, error) {
	return b.Build()
}

func (b *_ReplyEncodedReplyBuilder) DeepCopy() any {
	_copy := b.CreateReplyEncodedReplyBuilder().(*_ReplyEncodedReplyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReplyEncodedReplyBuilder creates a ReplyEncodedReplyBuilder
func (b *_ReplyEncodedReply) CreateReplyEncodedReplyBuilder() ReplyEncodedReplyBuilder {
	if b == nil {
		return NewReplyEncodedReplyBuilder()
	}
	return &_ReplyEncodedReplyBuilder{_ReplyEncodedReply: b.deepCopy()}
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

func (m *_ReplyEncodedReply) GetParent() ReplyContract {
	return m.ReplyContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyEncodedReply) GetEncodedReply() EncodedReply {
	return m.EncodedReply
}

func (m *_ReplyEncodedReply) GetChksum() Checksum {
	return m.Chksum
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ReplyEncodedReply) GetEncodedReplyDecoded() EncodedReply {
	ctx := context.Background()
	_ = ctx
	return CastEncodedReply(m.GetEncodedReply())
}

func (m *_ReplyEncodedReply) GetChksumDecoded() Checksum {
	ctx := context.Background()
	_ = ctx
	return CastChecksum(m.GetChksum())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReplyEncodedReply(structType any) ReplyEncodedReply {
	if casted, ok := structType.(ReplyEncodedReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyEncodedReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyEncodedReply) GetTypeName() string {
	return "ReplyEncodedReply"
}

func (m *_ReplyEncodedReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyContract.(*_Reply).getLengthInBits(ctx))

	// Manual Field (encodedReply)
	lengthInBits += uint16(int32((int32(m.GetEncodedReply().GetLengthInBytes(ctx)) * int32(int32(2)))) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Manual Field (chksum)
	lengthInBits += uint16(utils.InlineIf((m.GetCBusOptions().GetSrchk()), func() any { return int32((int32(16))) }, func() any { return int32((int32(0))) }).(int32))

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ReplyEncodedReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReplyEncodedReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Reply, cBusOptions CBusOptions, requestContext RequestContext) (__replyEncodedReply ReplyEncodedReply, err error) {
	m.ReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyEncodedReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyEncodedReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	encodedReply, err := ReadManualField[EncodedReply](ctx, "encodedReply", readBuffer, EnsureType[EncodedReply](ReadEncodedReply(ctx, readBuffer, cBusOptions, requestContext, cBusOptions.GetSrchk())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encodedReply' field"))
	}
	m.EncodedReply = encodedReply

	encodedReplyDecoded, err := ReadVirtualField[EncodedReply](ctx, "encodedReplyDecoded", (*EncodedReply)(nil), encodedReply)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encodedReplyDecoded' field"))
	}
	_ = encodedReplyDecoded

	chksum, err := ReadManualField[Checksum](ctx, "chksum", readBuffer, EnsureType[Checksum](ReadAndValidateChecksum(ctx, readBuffer, encodedReply, cBusOptions.GetSrchk())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chksum' field"))
	}
	m.Chksum = chksum

	chksumDecoded, err := ReadVirtualField[Checksum](ctx, "chksumDecoded", (*Checksum)(nil), chksum)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chksumDecoded' field"))
	}
	_ = chksumDecoded

	if closeErr := readBuffer.CloseContext("ReplyEncodedReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyEncodedReply")
	}

	return m, nil
}

func (m *_ReplyEncodedReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyEncodedReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyEncodedReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyEncodedReply")
		}

		if err := WriteManualField[EncodedReply](ctx, "encodedReply", func(ctx context.Context) error { return WriteEncodedReply(ctx, writeBuffer, m.GetEncodedReply()) }, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'encodedReply' field")
		}
		// Virtual field
		encodedReplyDecoded := m.GetEncodedReplyDecoded()
		_ = encodedReplyDecoded
		if _encodedReplyDecodedErr := writeBuffer.WriteVirtual(ctx, "encodedReplyDecoded", m.GetEncodedReplyDecoded()); _encodedReplyDecodedErr != nil {
			return errors.Wrap(_encodedReplyDecodedErr, "Error serializing 'encodedReplyDecoded' field")
		}

		if err := WriteManualField[Checksum](ctx, "chksum", func(ctx context.Context) error {
			return CalculateChecksum(ctx, writeBuffer, m.GetEncodedReply(), m.GetCBusOptions().GetSrchk())
		}, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'chksum' field")
		}
		// Virtual field
		chksumDecoded := m.GetChksumDecoded()
		_ = chksumDecoded
		if _chksumDecodedErr := writeBuffer.WriteVirtual(ctx, "chksumDecoded", m.GetChksumDecoded()); _chksumDecodedErr != nil {
			return errors.Wrap(_chksumDecodedErr, "Error serializing 'chksumDecoded' field")
		}

		if popErr := writeBuffer.PopContext("ReplyEncodedReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyEncodedReply")
		}
		return nil
	}
	return m.ReplyContract.(*_Reply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReplyEncodedReply) IsReplyEncodedReply() {}

func (m *_ReplyEncodedReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReplyEncodedReply) deepCopy() *_ReplyEncodedReply {
	if m == nil {
		return nil
	}
	_ReplyEncodedReplyCopy := &_ReplyEncodedReply{
		m.ReplyContract.(*_Reply).deepCopy(),
		m.EncodedReply.DeepCopy().(EncodedReply),
		m.Chksum.DeepCopy().(Checksum),
	}
	m.ReplyContract.(*_Reply)._SubType = m
	return _ReplyEncodedReplyCopy
}

func (m *_ReplyEncodedReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
