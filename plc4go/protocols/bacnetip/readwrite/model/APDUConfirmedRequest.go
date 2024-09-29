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

// APDUConfirmedRequest is the corresponding interface of APDUConfirmedRequest
type APDUConfirmedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetSegmentedMessage returns SegmentedMessage (property field)
	GetSegmentedMessage() bool
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() bool
	// GetSegmentedResponseAccepted returns SegmentedResponseAccepted (property field)
	GetSegmentedResponseAccepted() bool
	// GetMaxSegmentsAccepted returns MaxSegmentsAccepted (property field)
	GetMaxSegmentsAccepted() MaxSegmentsAccepted
	// GetMaxApduLengthAccepted returns MaxApduLengthAccepted (property field)
	GetMaxApduLengthAccepted() MaxApduLengthAccepted
	// GetInvokeId returns InvokeId (property field)
	GetInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() *uint8
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() BACnetConfirmedServiceRequest
	// GetSegmentServiceChoice returns SegmentServiceChoice (property field)
	GetSegmentServiceChoice() *BACnetConfirmedServiceChoice
	// GetSegment returns Segment (property field)
	GetSegment() []byte
	// GetApduHeaderReduction returns ApduHeaderReduction (virtual field)
	GetApduHeaderReduction() uint16
	// GetSegmentReduction returns SegmentReduction (virtual field)
	GetSegmentReduction() uint16
	// IsAPDUConfirmedRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUConfirmedRequest()
	// CreateBuilder creates a APDUConfirmedRequestBuilder
	CreateAPDUConfirmedRequestBuilder() APDUConfirmedRequestBuilder
}

// _APDUConfirmedRequest is the data-structure of this message
type _APDUConfirmedRequest struct {
	APDUContract
	SegmentedMessage          bool
	MoreFollows               bool
	SegmentedResponseAccepted bool
	MaxSegmentsAccepted       MaxSegmentsAccepted
	MaxApduLengthAccepted     MaxApduLengthAccepted
	InvokeId                  uint8
	SequenceNumber            *uint8
	ProposedWindowSize        *uint8
	ServiceRequest            BACnetConfirmedServiceRequest
	SegmentServiceChoice      *BACnetConfirmedServiceChoice
	Segment                   []byte
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUConfirmedRequest = (*_APDUConfirmedRequest)(nil)
var _ APDURequirements = (*_APDUConfirmedRequest)(nil)

// NewAPDUConfirmedRequest factory function for _APDUConfirmedRequest
func NewAPDUConfirmedRequest(segmentedMessage bool, moreFollows bool, segmentedResponseAccepted bool, maxSegmentsAccepted MaxSegmentsAccepted, maxApduLengthAccepted MaxApduLengthAccepted, invokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceRequest BACnetConfirmedServiceRequest, segmentServiceChoice *BACnetConfirmedServiceChoice, segment []byte, apduLength uint16) *_APDUConfirmedRequest {
	_result := &_APDUConfirmedRequest{
		APDUContract:              NewAPDU(apduLength),
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            serviceRequest,
		SegmentServiceChoice:      segmentServiceChoice,
		Segment:                   segment,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// APDUConfirmedRequestBuilder is a builder for APDUConfirmedRequest
type APDUConfirmedRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(segmentedMessage bool, moreFollows bool, segmentedResponseAccepted bool, maxSegmentsAccepted MaxSegmentsAccepted, maxApduLengthAccepted MaxApduLengthAccepted, invokeId uint8, segment []byte) APDUConfirmedRequestBuilder
	// WithSegmentedMessage adds SegmentedMessage (property field)
	WithSegmentedMessage(bool) APDUConfirmedRequestBuilder
	// WithMoreFollows adds MoreFollows (property field)
	WithMoreFollows(bool) APDUConfirmedRequestBuilder
	// WithSegmentedResponseAccepted adds SegmentedResponseAccepted (property field)
	WithSegmentedResponseAccepted(bool) APDUConfirmedRequestBuilder
	// WithMaxSegmentsAccepted adds MaxSegmentsAccepted (property field)
	WithMaxSegmentsAccepted(MaxSegmentsAccepted) APDUConfirmedRequestBuilder
	// WithMaxApduLengthAccepted adds MaxApduLengthAccepted (property field)
	WithMaxApduLengthAccepted(MaxApduLengthAccepted) APDUConfirmedRequestBuilder
	// WithInvokeId adds InvokeId (property field)
	WithInvokeId(uint8) APDUConfirmedRequestBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithOptionalSequenceNumber(uint8) APDUConfirmedRequestBuilder
	// WithProposedWindowSize adds ProposedWindowSize (property field)
	WithOptionalProposedWindowSize(uint8) APDUConfirmedRequestBuilder
	// WithServiceRequest adds ServiceRequest (property field)
	WithOptionalServiceRequest(BACnetConfirmedServiceRequest) APDUConfirmedRequestBuilder
	// WithOptionalServiceRequestBuilder adds ServiceRequest (property field) which is build by the builder
	WithOptionalServiceRequestBuilder(func(BACnetConfirmedServiceRequestBuilder) BACnetConfirmedServiceRequestBuilder) APDUConfirmedRequestBuilder
	// WithSegmentServiceChoice adds SegmentServiceChoice (property field)
	WithOptionalSegmentServiceChoice(BACnetConfirmedServiceChoice) APDUConfirmedRequestBuilder
	// WithSegment adds Segment (property field)
	WithSegment(...byte) APDUConfirmedRequestBuilder
	// Build builds the APDUConfirmedRequest or returns an error if something is wrong
	Build() (APDUConfirmedRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUConfirmedRequest
}

// NewAPDUConfirmedRequestBuilder() creates a APDUConfirmedRequestBuilder
func NewAPDUConfirmedRequestBuilder() APDUConfirmedRequestBuilder {
	return &_APDUConfirmedRequestBuilder{_APDUConfirmedRequest: new(_APDUConfirmedRequest)}
}

type _APDUConfirmedRequestBuilder struct {
	*_APDUConfirmedRequest

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDUConfirmedRequestBuilder) = (*_APDUConfirmedRequestBuilder)(nil)

func (b *_APDUConfirmedRequestBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
}

func (b *_APDUConfirmedRequestBuilder) WithMandatoryFields(segmentedMessage bool, moreFollows bool, segmentedResponseAccepted bool, maxSegmentsAccepted MaxSegmentsAccepted, maxApduLengthAccepted MaxApduLengthAccepted, invokeId uint8, segment []byte) APDUConfirmedRequestBuilder {
	return b.WithSegmentedMessage(segmentedMessage).WithMoreFollows(moreFollows).WithSegmentedResponseAccepted(segmentedResponseAccepted).WithMaxSegmentsAccepted(maxSegmentsAccepted).WithMaxApduLengthAccepted(maxApduLengthAccepted).WithInvokeId(invokeId).WithSegment(segment...)
}

func (b *_APDUConfirmedRequestBuilder) WithSegmentedMessage(segmentedMessage bool) APDUConfirmedRequestBuilder {
	b.SegmentedMessage = segmentedMessage
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithMoreFollows(moreFollows bool) APDUConfirmedRequestBuilder {
	b.MoreFollows = moreFollows
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithSegmentedResponseAccepted(segmentedResponseAccepted bool) APDUConfirmedRequestBuilder {
	b.SegmentedResponseAccepted = segmentedResponseAccepted
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithMaxSegmentsAccepted(maxSegmentsAccepted MaxSegmentsAccepted) APDUConfirmedRequestBuilder {
	b.MaxSegmentsAccepted = maxSegmentsAccepted
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithMaxApduLengthAccepted(maxApduLengthAccepted MaxApduLengthAccepted) APDUConfirmedRequestBuilder {
	b.MaxApduLengthAccepted = maxApduLengthAccepted
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithInvokeId(invokeId uint8) APDUConfirmedRequestBuilder {
	b.InvokeId = invokeId
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithOptionalSequenceNumber(sequenceNumber uint8) APDUConfirmedRequestBuilder {
	b.SequenceNumber = &sequenceNumber
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithOptionalProposedWindowSize(proposedWindowSize uint8) APDUConfirmedRequestBuilder {
	b.ProposedWindowSize = &proposedWindowSize
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithOptionalServiceRequest(serviceRequest BACnetConfirmedServiceRequest) APDUConfirmedRequestBuilder {
	b.ServiceRequest = serviceRequest
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithOptionalServiceRequestBuilder(builderSupplier func(BACnetConfirmedServiceRequestBuilder) BACnetConfirmedServiceRequestBuilder) APDUConfirmedRequestBuilder {
	builder := builderSupplier(b.ServiceRequest.CreateBACnetConfirmedServiceRequestBuilder())
	var err error
	b.ServiceRequest, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestBuilder failed"))
	}
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithOptionalSegmentServiceChoice(segmentServiceChoice BACnetConfirmedServiceChoice) APDUConfirmedRequestBuilder {
	b.SegmentServiceChoice = &segmentServiceChoice
	return b
}

func (b *_APDUConfirmedRequestBuilder) WithSegment(segment ...byte) APDUConfirmedRequestBuilder {
	b.Segment = segment
	return b
}

func (b *_APDUConfirmedRequestBuilder) Build() (APDUConfirmedRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUConfirmedRequest.deepCopy(), nil
}

func (b *_APDUConfirmedRequestBuilder) MustBuild() APDUConfirmedRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_APDUConfirmedRequestBuilder) Done() APDUBuilder {
	return b.parentBuilder
}

func (b *_APDUConfirmedRequestBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDUConfirmedRequestBuilder) DeepCopy() any {
	_copy := b.CreateAPDUConfirmedRequestBuilder().(*_APDUConfirmedRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDUConfirmedRequestBuilder creates a APDUConfirmedRequestBuilder
func (b *_APDUConfirmedRequest) CreateAPDUConfirmedRequestBuilder() APDUConfirmedRequestBuilder {
	if b == nil {
		return NewAPDUConfirmedRequestBuilder()
	}
	return &_APDUConfirmedRequestBuilder{_APDUConfirmedRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUConfirmedRequest) GetApduType() ApduType {
	return ApduType_CONFIRMED_REQUEST_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUConfirmedRequest) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUConfirmedRequest) GetSegmentedMessage() bool {
	return m.SegmentedMessage
}

func (m *_APDUConfirmedRequest) GetMoreFollows() bool {
	return m.MoreFollows
}

func (m *_APDUConfirmedRequest) GetSegmentedResponseAccepted() bool {
	return m.SegmentedResponseAccepted
}

func (m *_APDUConfirmedRequest) GetMaxSegmentsAccepted() MaxSegmentsAccepted {
	return m.MaxSegmentsAccepted
}

func (m *_APDUConfirmedRequest) GetMaxApduLengthAccepted() MaxApduLengthAccepted {
	return m.MaxApduLengthAccepted
}

func (m *_APDUConfirmedRequest) GetInvokeId() uint8 {
	return m.InvokeId
}

func (m *_APDUConfirmedRequest) GetSequenceNumber() *uint8 {
	return m.SequenceNumber
}

func (m *_APDUConfirmedRequest) GetProposedWindowSize() *uint8 {
	return m.ProposedWindowSize
}

func (m *_APDUConfirmedRequest) GetServiceRequest() BACnetConfirmedServiceRequest {
	return m.ServiceRequest
}

func (m *_APDUConfirmedRequest) GetSegmentServiceChoice() *BACnetConfirmedServiceChoice {
	return m.SegmentServiceChoice
}

func (m *_APDUConfirmedRequest) GetSegment() []byte {
	return m.Segment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_APDUConfirmedRequest) GetApduHeaderReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.GetSequenceNumber()
	_ = sequenceNumber
	proposedWindowSize := m.GetProposedWindowSize()
	_ = proposedWindowSize
	serviceRequest := m.GetServiceRequest()
	_ = serviceRequest
	segmentServiceChoice := m.GetSegmentServiceChoice()
	_ = segmentServiceChoice
	return uint16(uint16(uint16(3)) + uint16((utils.InlineIf(m.GetSegmentedMessage(), func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
}

func (m *_APDUConfirmedRequest) GetSegmentReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.GetSequenceNumber()
	_ = sequenceNumber
	proposedWindowSize := m.GetProposedWindowSize()
	_ = proposedWindowSize
	serviceRequest := m.GetServiceRequest()
	_ = serviceRequest
	segmentServiceChoice := m.GetSegmentServiceChoice()
	_ = segmentServiceChoice
	return uint16(utils.InlineIf((bool((m.GetSegmentServiceChoice()) != (nil))), func() any { return uint16((uint16(m.GetApduHeaderReduction()) + uint16(uint16(1)))) }, func() any { return uint16(m.GetApduHeaderReduction()) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUConfirmedRequest(structType any) APDUConfirmedRequest {
	if casted, ok := structType.(APDUConfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*APDUConfirmedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_APDUConfirmedRequest) GetTypeName() string {
	return "APDUConfirmedRequest"
}

func (m *_APDUConfirmedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Simple field (segmentedResponseAccepted)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (maxSegmentsAccepted)
	lengthInBits += 3

	// Simple field (maxApduLengthAccepted)
	lengthInBits += 4

	// Simple field (invokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (serviceRequest)
	if m.ServiceRequest != nil {
		lengthInBits += m.ServiceRequest.GetLengthInBits(ctx)
	}

	// Optional Field (segmentServiceChoice)
	if m.SegmentServiceChoice != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Segment) > 0 {
		lengthInBits += 8 * uint16(len(m.Segment))
	}

	return lengthInBits
}

func (m *_APDUConfirmedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUConfirmedRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUConfirmedRequest APDUConfirmedRequest, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUConfirmedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUConfirmedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentedMessage, err := ReadSimpleField(ctx, "segmentedMessage", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentedMessage' field"))
	}
	m.SegmentedMessage = segmentedMessage

	moreFollows, err := ReadSimpleField(ctx, "moreFollows", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreFollows' field"))
	}
	m.MoreFollows = moreFollows

	segmentedResponseAccepted, err := ReadSimpleField(ctx, "segmentedResponseAccepted", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentedResponseAccepted' field"))
	}
	m.SegmentedResponseAccepted = segmentedResponseAccepted

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(2)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	maxSegmentsAccepted, err := ReadEnumField[MaxSegmentsAccepted](ctx, "maxSegmentsAccepted", "MaxSegmentsAccepted", ReadEnum(MaxSegmentsAcceptedByValue, ReadUnsignedByte(readBuffer, uint8(3))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxSegmentsAccepted' field"))
	}
	m.MaxSegmentsAccepted = maxSegmentsAccepted

	maxApduLengthAccepted, err := ReadEnumField[MaxApduLengthAccepted](ctx, "maxApduLengthAccepted", "MaxApduLengthAccepted", ReadEnum(MaxApduLengthAcceptedByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxApduLengthAccepted' field"))
	}
	m.MaxApduLengthAccepted = maxApduLengthAccepted

	invokeId, err := ReadSimpleField(ctx, "invokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'invokeId' field"))
	}
	m.InvokeId = invokeId

	var sequenceNumber *uint8
	sequenceNumber, err = ReadOptionalField[uint8](ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	var proposedWindowSize *uint8
	proposedWindowSize, err = ReadOptionalField[uint8](ctx, "proposedWindowSize", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proposedWindowSize' field"))
	}
	m.ProposedWindowSize = proposedWindowSize

	apduHeaderReduction, err := ReadVirtualField[uint16](ctx, "apduHeaderReduction", (*uint16)(nil), uint16(uint16(3))+uint16((utils.InlineIf(segmentedMessage, func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduHeaderReduction' field"))
	}
	_ = apduHeaderReduction

	var serviceRequest BACnetConfirmedServiceRequest
	_serviceRequest, err := ReadOptionalField[BACnetConfirmedServiceRequest](ctx, "serviceRequest", ReadComplex[BACnetConfirmedServiceRequest](BACnetConfirmedServiceRequestParseWithBufferProducer[BACnetConfirmedServiceRequest]((uint32)(uint32(apduLength)-uint32(apduHeaderReduction))), readBuffer), !(segmentedMessage))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceRequest' field"))
	}
	if _serviceRequest != nil {
		serviceRequest = *_serviceRequest
		m.ServiceRequest = serviceRequest
	}

	// Validation
	if !(bool((bool(!(segmentedMessage)) && bool(bool((serviceRequest) != (nil))))) || bool(segmentedMessage)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "service request should be set"})
	}

	var segmentServiceChoice *BACnetConfirmedServiceChoice
	segmentServiceChoice, err = ReadOptionalField[BACnetConfirmedServiceChoice](ctx, "segmentServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool(segmentedMessage) && bool(bool((*sequenceNumber) != (0))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentServiceChoice' field"))
	}
	m.SegmentServiceChoice = segmentServiceChoice

	segmentReduction, err := ReadVirtualField[uint16](ctx, "segmentReduction", (*uint16)(nil), utils.InlineIf((bool((segmentServiceChoice) != (nil))), func() any { return uint16((uint16(apduHeaderReduction) + uint16(uint16(1)))) }, func() any { return uint16(apduHeaderReduction) }).(uint16))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentReduction' field"))
	}
	_ = segmentReduction

	segment, err := readBuffer.ReadByteArray("segment", int(utils.InlineIf(segmentedMessage, func() any {
		return int32((utils.InlineIf((bool((apduLength) > (0))), func() any { return int32((int32(apduLength) - int32(segmentReduction))) }, func() any { return int32(int32(0)) }).(int32)))
	}, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segment' field"))
	}
	m.Segment = segment

	if closeErr := readBuffer.CloseContext("APDUConfirmedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUConfirmedRequest")
	}

	return m, nil
}

func (m *_APDUConfirmedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUConfirmedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUConfirmedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUConfirmedRequest")
		}

		if err := WriteSimpleField[bool](ctx, "segmentedMessage", m.GetSegmentedMessage(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentedMessage' field")
		}

		if err := WriteSimpleField[bool](ctx, "moreFollows", m.GetMoreFollows(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'moreFollows' field")
		}

		if err := WriteSimpleField[bool](ctx, "segmentedResponseAccepted", m.GetSegmentedResponseAccepted(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentedResponseAccepted' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 2)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleEnumField[MaxSegmentsAccepted](ctx, "maxSegmentsAccepted", "MaxSegmentsAccepted", m.GetMaxSegmentsAccepted(), WriteEnum[MaxSegmentsAccepted, uint8](MaxSegmentsAccepted.GetValue, MaxSegmentsAccepted.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 3))); err != nil {
			return errors.Wrap(err, "Error serializing 'maxSegmentsAccepted' field")
		}

		if err := WriteSimpleEnumField[MaxApduLengthAccepted](ctx, "maxApduLengthAccepted", "MaxApduLengthAccepted", m.GetMaxApduLengthAccepted(), WriteEnum[MaxApduLengthAccepted, uint8](MaxApduLengthAccepted.GetValue, MaxApduLengthAccepted.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'maxApduLengthAccepted' field")
		}

		if err := WriteSimpleField[uint8](ctx, "invokeId", m.GetInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'invokeId' field")
		}

		if err := WriteOptionalField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteOptionalField[uint8](ctx, "proposedWindowSize", m.GetProposedWindowSize(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'proposedWindowSize' field")
		}
		// Virtual field
		apduHeaderReduction := m.GetApduHeaderReduction()
		_ = apduHeaderReduction
		if _apduHeaderReductionErr := writeBuffer.WriteVirtual(ctx, "apduHeaderReduction", m.GetApduHeaderReduction()); _apduHeaderReductionErr != nil {
			return errors.Wrap(_apduHeaderReductionErr, "Error serializing 'apduHeaderReduction' field")
		}

		if err := WriteOptionalField[BACnetConfirmedServiceRequest](ctx, "serviceRequest", GetRef(m.GetServiceRequest()), WriteComplex[BACnetConfirmedServiceRequest](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceRequest' field")
		}

		if err := WriteOptionalEnumField[BACnetConfirmedServiceChoice](ctx, "segmentServiceChoice", "BACnetConfirmedServiceChoice", m.GetSegmentServiceChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool(m.GetSegmentedMessage()) && bool(bool((*m.GetSequenceNumber()) != (0)))); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentServiceChoice' field")
		}
		// Virtual field
		segmentReduction := m.GetSegmentReduction()
		_ = segmentReduction
		if _segmentReductionErr := writeBuffer.WriteVirtual(ctx, "segmentReduction", m.GetSegmentReduction()); _segmentReductionErr != nil {
			return errors.Wrap(_segmentReductionErr, "Error serializing 'segmentReduction' field")
		}

		if err := WriteByteArrayField(ctx, "segment", m.GetSegment(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'segment' field")
		}

		if popErr := writeBuffer.PopContext("APDUConfirmedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUConfirmedRequest")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUConfirmedRequest) IsAPDUConfirmedRequest() {}

func (m *_APDUConfirmedRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUConfirmedRequest) deepCopy() *_APDUConfirmedRequest {
	if m == nil {
		return nil
	}
	_APDUConfirmedRequestCopy := &_APDUConfirmedRequest{
		m.APDUContract.(*_APDU).deepCopy(),
		m.SegmentedMessage,
		m.MoreFollows,
		m.SegmentedResponseAccepted,
		m.MaxSegmentsAccepted,
		m.MaxApduLengthAccepted,
		m.InvokeId,
		utils.CopyPtr[uint8](m.SequenceNumber),
		utils.CopyPtr[uint8](m.ProposedWindowSize),
		m.ServiceRequest.DeepCopy().(BACnetConfirmedServiceRequest),
		utils.CopyPtr[BACnetConfirmedServiceChoice](m.SegmentServiceChoice),
		utils.DeepCopySlice[byte, byte](m.Segment),
		m.reservedField0,
	}
	m.APDUContract.(*_APDU)._SubType = m
	return _APDUConfirmedRequestCopy
}

func (m *_APDUConfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
