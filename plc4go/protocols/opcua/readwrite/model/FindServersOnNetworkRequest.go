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

// FindServersOnNetworkRequest is the corresponding interface of FindServersOnNetworkRequest
type FindServersOnNetworkRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetStartingRecordId returns StartingRecordId (property field)
	GetStartingRecordId() uint32
	// GetMaxRecordsToReturn returns MaxRecordsToReturn (property field)
	GetMaxRecordsToReturn() uint32
	// GetNoOfServerCapabilityFilter returns NoOfServerCapabilityFilter (property field)
	GetNoOfServerCapabilityFilter() int32
	// GetServerCapabilityFilter returns ServerCapabilityFilter (property field)
	GetServerCapabilityFilter() []PascalString
	// IsFindServersOnNetworkRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersOnNetworkRequest()
	// CreateBuilder creates a FindServersOnNetworkRequestBuilder
	CreateFindServersOnNetworkRequestBuilder() FindServersOnNetworkRequestBuilder
}

// _FindServersOnNetworkRequest is the data-structure of this message
type _FindServersOnNetworkRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader              ExtensionObjectDefinition
	StartingRecordId           uint32
	MaxRecordsToReturn         uint32
	NoOfServerCapabilityFilter int32
	ServerCapabilityFilter     []PascalString
}

var _ FindServersOnNetworkRequest = (*_FindServersOnNetworkRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersOnNetworkRequest)(nil)

// NewFindServersOnNetworkRequest factory function for _FindServersOnNetworkRequest
func NewFindServersOnNetworkRequest(requestHeader ExtensionObjectDefinition, startingRecordId uint32, maxRecordsToReturn uint32, noOfServerCapabilityFilter int32, serverCapabilityFilter []PascalString) *_FindServersOnNetworkRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for FindServersOnNetworkRequest must not be nil")
	}
	_result := &_FindServersOnNetworkRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		StartingRecordId:                  startingRecordId,
		MaxRecordsToReturn:                maxRecordsToReturn,
		NoOfServerCapabilityFilter:        noOfServerCapabilityFilter,
		ServerCapabilityFilter:            serverCapabilityFilter,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FindServersOnNetworkRequestBuilder is a builder for FindServersOnNetworkRequest
type FindServersOnNetworkRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, startingRecordId uint32, maxRecordsToReturn uint32, noOfServerCapabilityFilter int32, serverCapabilityFilter []PascalString) FindServersOnNetworkRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) FindServersOnNetworkRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) FindServersOnNetworkRequestBuilder
	// WithStartingRecordId adds StartingRecordId (property field)
	WithStartingRecordId(uint32) FindServersOnNetworkRequestBuilder
	// WithMaxRecordsToReturn adds MaxRecordsToReturn (property field)
	WithMaxRecordsToReturn(uint32) FindServersOnNetworkRequestBuilder
	// WithNoOfServerCapabilityFilter adds NoOfServerCapabilityFilter (property field)
	WithNoOfServerCapabilityFilter(int32) FindServersOnNetworkRequestBuilder
	// WithServerCapabilityFilter adds ServerCapabilityFilter (property field)
	WithServerCapabilityFilter(...PascalString) FindServersOnNetworkRequestBuilder
	// Build builds the FindServersOnNetworkRequest or returns an error if something is wrong
	Build() (FindServersOnNetworkRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FindServersOnNetworkRequest
}

// NewFindServersOnNetworkRequestBuilder() creates a FindServersOnNetworkRequestBuilder
func NewFindServersOnNetworkRequestBuilder() FindServersOnNetworkRequestBuilder {
	return &_FindServersOnNetworkRequestBuilder{_FindServersOnNetworkRequest: new(_FindServersOnNetworkRequest)}
}

type _FindServersOnNetworkRequestBuilder struct {
	*_FindServersOnNetworkRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FindServersOnNetworkRequestBuilder) = (*_FindServersOnNetworkRequestBuilder)(nil)

func (b *_FindServersOnNetworkRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_FindServersOnNetworkRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, startingRecordId uint32, maxRecordsToReturn uint32, noOfServerCapabilityFilter int32, serverCapabilityFilter []PascalString) FindServersOnNetworkRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithStartingRecordId(startingRecordId).WithMaxRecordsToReturn(maxRecordsToReturn).WithNoOfServerCapabilityFilter(noOfServerCapabilityFilter).WithServerCapabilityFilter(serverCapabilityFilter...)
}

func (b *_FindServersOnNetworkRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) FindServersOnNetworkRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_FindServersOnNetworkRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) FindServersOnNetworkRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_FindServersOnNetworkRequestBuilder) WithStartingRecordId(startingRecordId uint32) FindServersOnNetworkRequestBuilder {
	b.StartingRecordId = startingRecordId
	return b
}

func (b *_FindServersOnNetworkRequestBuilder) WithMaxRecordsToReturn(maxRecordsToReturn uint32) FindServersOnNetworkRequestBuilder {
	b.MaxRecordsToReturn = maxRecordsToReturn
	return b
}

func (b *_FindServersOnNetworkRequestBuilder) WithNoOfServerCapabilityFilter(noOfServerCapabilityFilter int32) FindServersOnNetworkRequestBuilder {
	b.NoOfServerCapabilityFilter = noOfServerCapabilityFilter
	return b
}

func (b *_FindServersOnNetworkRequestBuilder) WithServerCapabilityFilter(serverCapabilityFilter ...PascalString) FindServersOnNetworkRequestBuilder {
	b.ServerCapabilityFilter = serverCapabilityFilter
	return b
}

func (b *_FindServersOnNetworkRequestBuilder) Build() (FindServersOnNetworkRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FindServersOnNetworkRequest.deepCopy(), nil
}

func (b *_FindServersOnNetworkRequestBuilder) MustBuild() FindServersOnNetworkRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_FindServersOnNetworkRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_FindServersOnNetworkRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FindServersOnNetworkRequestBuilder) DeepCopy() any {
	_copy := b.CreateFindServersOnNetworkRequestBuilder().(*_FindServersOnNetworkRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFindServersOnNetworkRequestBuilder creates a FindServersOnNetworkRequestBuilder
func (b *_FindServersOnNetworkRequest) CreateFindServersOnNetworkRequestBuilder() FindServersOnNetworkRequestBuilder {
	if b == nil {
		return NewFindServersOnNetworkRequestBuilder()
	}
	return &_FindServersOnNetworkRequestBuilder{_FindServersOnNetworkRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersOnNetworkRequest) GetIdentifier() string {
	return "12192"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersOnNetworkRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersOnNetworkRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_FindServersOnNetworkRequest) GetStartingRecordId() uint32 {
	return m.StartingRecordId
}

func (m *_FindServersOnNetworkRequest) GetMaxRecordsToReturn() uint32 {
	return m.MaxRecordsToReturn
}

func (m *_FindServersOnNetworkRequest) GetNoOfServerCapabilityFilter() int32 {
	return m.NoOfServerCapabilityFilter
}

func (m *_FindServersOnNetworkRequest) GetServerCapabilityFilter() []PascalString {
	return m.ServerCapabilityFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFindServersOnNetworkRequest(structType any) FindServersOnNetworkRequest {
	if casted, ok := structType.(FindServersOnNetworkRequest); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersOnNetworkRequest); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersOnNetworkRequest) GetTypeName() string {
	return "FindServersOnNetworkRequest"
}

func (m *_FindServersOnNetworkRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (startingRecordId)
	lengthInBits += 32

	// Simple field (maxRecordsToReturn)
	lengthInBits += 32

	// Simple field (noOfServerCapabilityFilter)
	lengthInBits += 32

	// Array field
	if len(m.ServerCapabilityFilter) > 0 {
		for _curItem, element := range m.ServerCapabilityFilter {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerCapabilityFilter), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersOnNetworkRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FindServersOnNetworkRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__findServersOnNetworkRequest FindServersOnNetworkRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersOnNetworkRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersOnNetworkRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	startingRecordId, err := ReadSimpleField(ctx, "startingRecordId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingRecordId' field"))
	}
	m.StartingRecordId = startingRecordId

	maxRecordsToReturn, err := ReadSimpleField(ctx, "maxRecordsToReturn", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxRecordsToReturn' field"))
	}
	m.MaxRecordsToReturn = maxRecordsToReturn

	noOfServerCapabilityFilter, err := ReadSimpleField(ctx, "noOfServerCapabilityFilter", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerCapabilityFilter' field"))
	}
	m.NoOfServerCapabilityFilter = noOfServerCapabilityFilter

	serverCapabilityFilter, err := ReadCountArrayField[PascalString](ctx, "serverCapabilityFilter", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfServerCapabilityFilter))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverCapabilityFilter' field"))
	}
	m.ServerCapabilityFilter = serverCapabilityFilter

	if closeErr := readBuffer.CloseContext("FindServersOnNetworkRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersOnNetworkRequest")
	}

	return m, nil
}

func (m *_FindServersOnNetworkRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersOnNetworkRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersOnNetworkRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersOnNetworkRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "startingRecordId", m.GetStartingRecordId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingRecordId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxRecordsToReturn", m.GetMaxRecordsToReturn(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxRecordsToReturn' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServerCapabilityFilter", m.GetNoOfServerCapabilityFilter(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerCapabilityFilter' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverCapabilityFilter", m.GetServerCapabilityFilter(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverCapabilityFilter' field")
		}

		if popErr := writeBuffer.PopContext("FindServersOnNetworkRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersOnNetworkRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersOnNetworkRequest) IsFindServersOnNetworkRequest() {}

func (m *_FindServersOnNetworkRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FindServersOnNetworkRequest) deepCopy() *_FindServersOnNetworkRequest {
	if m == nil {
		return nil
	}
	_FindServersOnNetworkRequestCopy := &_FindServersOnNetworkRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.StartingRecordId,
		m.MaxRecordsToReturn,
		m.NoOfServerCapabilityFilter,
		utils.DeepCopySlice[PascalString, PascalString](m.ServerCapabilityFilter),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FindServersOnNetworkRequestCopy
}

func (m *_FindServersOnNetworkRequest) String() string {
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
