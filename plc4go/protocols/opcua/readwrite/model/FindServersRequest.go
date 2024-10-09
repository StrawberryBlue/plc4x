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

// FindServersRequest is the corresponding interface of FindServersRequest
type FindServersRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetNoOfLocaleIds returns NoOfLocaleIds (property field)
	GetNoOfLocaleIds() int32
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetNoOfServerUris returns NoOfServerUris (property field)
	GetNoOfServerUris() int32
	// GetServerUris returns ServerUris (property field)
	GetServerUris() []PascalString
	// IsFindServersRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersRequest()
	// CreateBuilder creates a FindServersRequestBuilder
	CreateFindServersRequestBuilder() FindServersRequestBuilder
}

// _FindServersRequest is the data-structure of this message
type _FindServersRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader  ExtensionObjectDefinition
	EndpointUrl    PascalString
	NoOfLocaleIds  int32
	LocaleIds      []PascalString
	NoOfServerUris int32
	ServerUris     []PascalString
}

var _ FindServersRequest = (*_FindServersRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersRequest)(nil)

// NewFindServersRequest factory function for _FindServersRequest
func NewFindServersRequest(requestHeader ExtensionObjectDefinition, endpointUrl PascalString, noOfLocaleIds int32, localeIds []PascalString, noOfServerUris int32, serverUris []PascalString) *_FindServersRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for FindServersRequest must not be nil")
	}
	if endpointUrl == nil {
		panic("endpointUrl of type PascalString for FindServersRequest must not be nil")
	}
	_result := &_FindServersRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		EndpointUrl:                       endpointUrl,
		NoOfLocaleIds:                     noOfLocaleIds,
		LocaleIds:                         localeIds,
		NoOfServerUris:                    noOfServerUris,
		ServerUris:                        serverUris,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FindServersRequestBuilder is a builder for FindServersRequest
type FindServersRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, endpointUrl PascalString, noOfLocaleIds int32, localeIds []PascalString, noOfServerUris int32, serverUris []PascalString) FindServersRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) FindServersRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) FindServersRequestBuilder
	// WithEndpointUrl adds EndpointUrl (property field)
	WithEndpointUrl(PascalString) FindServersRequestBuilder
	// WithEndpointUrlBuilder adds EndpointUrl (property field) which is build by the builder
	WithEndpointUrlBuilder(func(PascalStringBuilder) PascalStringBuilder) FindServersRequestBuilder
	// WithNoOfLocaleIds adds NoOfLocaleIds (property field)
	WithNoOfLocaleIds(int32) FindServersRequestBuilder
	// WithLocaleIds adds LocaleIds (property field)
	WithLocaleIds(...PascalString) FindServersRequestBuilder
	// WithNoOfServerUris adds NoOfServerUris (property field)
	WithNoOfServerUris(int32) FindServersRequestBuilder
	// WithServerUris adds ServerUris (property field)
	WithServerUris(...PascalString) FindServersRequestBuilder
	// Build builds the FindServersRequest or returns an error if something is wrong
	Build() (FindServersRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FindServersRequest
}

// NewFindServersRequestBuilder() creates a FindServersRequestBuilder
func NewFindServersRequestBuilder() FindServersRequestBuilder {
	return &_FindServersRequestBuilder{_FindServersRequest: new(_FindServersRequest)}
}

type _FindServersRequestBuilder struct {
	*_FindServersRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FindServersRequestBuilder) = (*_FindServersRequestBuilder)(nil)

func (b *_FindServersRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_FindServersRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, endpointUrl PascalString, noOfLocaleIds int32, localeIds []PascalString, noOfServerUris int32, serverUris []PascalString) FindServersRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithEndpointUrl(endpointUrl).WithNoOfLocaleIds(noOfLocaleIds).WithLocaleIds(localeIds...).WithNoOfServerUris(noOfServerUris).WithServerUris(serverUris...)
}

func (b *_FindServersRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) FindServersRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_FindServersRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) FindServersRequestBuilder {
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

func (b *_FindServersRequestBuilder) WithEndpointUrl(endpointUrl PascalString) FindServersRequestBuilder {
	b.EndpointUrl = endpointUrl
	return b
}

func (b *_FindServersRequestBuilder) WithEndpointUrlBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) FindServersRequestBuilder {
	builder := builderSupplier(b.EndpointUrl.CreatePascalStringBuilder())
	var err error
	b.EndpointUrl, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_FindServersRequestBuilder) WithNoOfLocaleIds(noOfLocaleIds int32) FindServersRequestBuilder {
	b.NoOfLocaleIds = noOfLocaleIds
	return b
}

func (b *_FindServersRequestBuilder) WithLocaleIds(localeIds ...PascalString) FindServersRequestBuilder {
	b.LocaleIds = localeIds
	return b
}

func (b *_FindServersRequestBuilder) WithNoOfServerUris(noOfServerUris int32) FindServersRequestBuilder {
	b.NoOfServerUris = noOfServerUris
	return b
}

func (b *_FindServersRequestBuilder) WithServerUris(serverUris ...PascalString) FindServersRequestBuilder {
	b.ServerUris = serverUris
	return b
}

func (b *_FindServersRequestBuilder) Build() (FindServersRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.EndpointUrl == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'endpointUrl' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FindServersRequest.deepCopy(), nil
}

func (b *_FindServersRequestBuilder) MustBuild() FindServersRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_FindServersRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_FindServersRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FindServersRequestBuilder) DeepCopy() any {
	_copy := b.CreateFindServersRequestBuilder().(*_FindServersRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFindServersRequestBuilder creates a FindServersRequestBuilder
func (b *_FindServersRequest) CreateFindServersRequestBuilder() FindServersRequestBuilder {
	if b == nil {
		return NewFindServersRequestBuilder()
	}
	return &_FindServersRequestBuilder{_FindServersRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersRequest) GetIdentifier() string {
	return "422"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_FindServersRequest) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_FindServersRequest) GetNoOfLocaleIds() int32 {
	return m.NoOfLocaleIds
}

func (m *_FindServersRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_FindServersRequest) GetNoOfServerUris() int32 {
	return m.NoOfServerUris
}

func (m *_FindServersRequest) GetServerUris() []PascalString {
	return m.ServerUris
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFindServersRequest(structType any) FindServersRequest {
	if casted, ok := structType.(FindServersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersRequest) GetTypeName() string {
	return "FindServersRequest"
}

func (m *_FindServersRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfServerUris)
	lengthInBits += 32

	// Array field
	if len(m.ServerUris) > 0 {
		for _curItem, element := range m.ServerUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FindServersRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__findServersRequest FindServersRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	endpointUrl, err := ReadSimpleField[PascalString](ctx, "endpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpointUrl' field"))
	}
	m.EndpointUrl = endpointUrl

	noOfLocaleIds, err := ReadSimpleField(ctx, "noOfLocaleIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLocaleIds' field"))
	}
	m.NoOfLocaleIds = noOfLocaleIds

	localeIds, err := ReadCountArrayField[PascalString](ctx, "localeIds", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfLocaleIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeIds' field"))
	}
	m.LocaleIds = localeIds

	noOfServerUris, err := ReadSimpleField(ctx, "noOfServerUris", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerUris' field"))
	}
	m.NoOfServerUris = noOfServerUris

	serverUris, err := ReadCountArrayField[PascalString](ctx, "serverUris", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfServerUris))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverUris' field"))
	}
	m.ServerUris = serverUris

	if closeErr := readBuffer.CloseContext("FindServersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersRequest")
	}

	return m, nil
}

func (m *_FindServersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpointUrl", m.GetEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpointUrl' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLocaleIds", m.GetNoOfLocaleIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServerUris", m.GetNoOfServerUris(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerUris' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverUris", m.GetServerUris(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverUris' field")
		}

		if popErr := writeBuffer.PopContext("FindServersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersRequest) IsFindServersRequest() {}

func (m *_FindServersRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FindServersRequest) deepCopy() *_FindServersRequest {
	if m == nil {
		return nil
	}
	_FindServersRequestCopy := &_FindServersRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.EndpointUrl.DeepCopy().(PascalString),
		m.NoOfLocaleIds,
		utils.DeepCopySlice[PascalString, PascalString](m.LocaleIds),
		m.NoOfServerUris,
		utils.DeepCopySlice[PascalString, PascalString](m.ServerUris),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FindServersRequestCopy
}

func (m *_FindServersRequest) String() string {
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
