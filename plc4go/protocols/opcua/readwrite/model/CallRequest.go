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

// CallRequest is the corresponding interface of CallRequest
type CallRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfMethodsToCall returns NoOfMethodsToCall (property field)
	GetNoOfMethodsToCall() int32
	// GetMethodsToCall returns MethodsToCall (property field)
	GetMethodsToCall() []ExtensionObjectDefinition
	// IsCallRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCallRequest()
	// CreateBuilder creates a CallRequestBuilder
	CreateCallRequestBuilder() CallRequestBuilder
}

// _CallRequest is the data-structure of this message
type _CallRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader     ExtensionObjectDefinition
	NoOfMethodsToCall int32
	MethodsToCall     []ExtensionObjectDefinition
}

var _ CallRequest = (*_CallRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CallRequest)(nil)

// NewCallRequest factory function for _CallRequest
func NewCallRequest(requestHeader ExtensionObjectDefinition, noOfMethodsToCall int32, methodsToCall []ExtensionObjectDefinition) *_CallRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for CallRequest must not be nil")
	}
	_result := &_CallRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfMethodsToCall:                 noOfMethodsToCall,
		MethodsToCall:                     methodsToCall,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CallRequestBuilder is a builder for CallRequest
type CallRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfMethodsToCall int32, methodsToCall []ExtensionObjectDefinition) CallRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) CallRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CallRequestBuilder
	// WithNoOfMethodsToCall adds NoOfMethodsToCall (property field)
	WithNoOfMethodsToCall(int32) CallRequestBuilder
	// WithMethodsToCall adds MethodsToCall (property field)
	WithMethodsToCall(...ExtensionObjectDefinition) CallRequestBuilder
	// Build builds the CallRequest or returns an error if something is wrong
	Build() (CallRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CallRequest
}

// NewCallRequestBuilder() creates a CallRequestBuilder
func NewCallRequestBuilder() CallRequestBuilder {
	return &_CallRequestBuilder{_CallRequest: new(_CallRequest)}
}

type _CallRequestBuilder struct {
	*_CallRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CallRequestBuilder) = (*_CallRequestBuilder)(nil)

func (b *_CallRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CallRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfMethodsToCall int32, methodsToCall []ExtensionObjectDefinition) CallRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithNoOfMethodsToCall(noOfMethodsToCall).WithMethodsToCall(methodsToCall...)
}

func (b *_CallRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) CallRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CallRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CallRequestBuilder {
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

func (b *_CallRequestBuilder) WithNoOfMethodsToCall(noOfMethodsToCall int32) CallRequestBuilder {
	b.NoOfMethodsToCall = noOfMethodsToCall
	return b
}

func (b *_CallRequestBuilder) WithMethodsToCall(methodsToCall ...ExtensionObjectDefinition) CallRequestBuilder {
	b.MethodsToCall = methodsToCall
	return b
}

func (b *_CallRequestBuilder) Build() (CallRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CallRequest.deepCopy(), nil
}

func (b *_CallRequestBuilder) MustBuild() CallRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CallRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CallRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CallRequestBuilder) DeepCopy() any {
	_copy := b.CreateCallRequestBuilder().(*_CallRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCallRequestBuilder creates a CallRequestBuilder
func (b *_CallRequest) CreateCallRequestBuilder() CallRequestBuilder {
	if b == nil {
		return NewCallRequestBuilder()
	}
	return &_CallRequestBuilder{_CallRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CallRequest) GetIdentifier() string {
	return "712"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CallRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CallRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CallRequest) GetNoOfMethodsToCall() int32 {
	return m.NoOfMethodsToCall
}

func (m *_CallRequest) GetMethodsToCall() []ExtensionObjectDefinition {
	return m.MethodsToCall
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCallRequest(structType any) CallRequest {
	if casted, ok := structType.(CallRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CallRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CallRequest) GetTypeName() string {
	return "CallRequest"
}

func (m *_CallRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfMethodsToCall)
	lengthInBits += 32

	// Array field
	if len(m.MethodsToCall) > 0 {
		for _curItem, element := range m.MethodsToCall {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MethodsToCall), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CallRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CallRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__callRequest CallRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CallRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CallRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfMethodsToCall, err := ReadSimpleField(ctx, "noOfMethodsToCall", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMethodsToCall' field"))
	}
	m.NoOfMethodsToCall = noOfMethodsToCall

	methodsToCall, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "methodsToCall", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("706")), readBuffer), uint64(noOfMethodsToCall))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'methodsToCall' field"))
	}
	m.MethodsToCall = methodsToCall

	if closeErr := readBuffer.CloseContext("CallRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CallRequest")
	}

	return m, nil
}

func (m *_CallRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CallRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CallRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CallRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfMethodsToCall", m.GetNoOfMethodsToCall(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMethodsToCall' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "methodsToCall", m.GetMethodsToCall(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'methodsToCall' field")
		}

		if popErr := writeBuffer.PopContext("CallRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CallRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CallRequest) IsCallRequest() {}

func (m *_CallRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CallRequest) deepCopy() *_CallRequest {
	if m == nil {
		return nil
	}
	_CallRequestCopy := &_CallRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfMethodsToCall,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.MethodsToCall),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CallRequestCopy
}

func (m *_CallRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
