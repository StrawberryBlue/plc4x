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

// CBusMessageToServer is the corresponding interface of CBusMessageToServer
type CBusMessageToServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusMessage
	// GetRequest returns Request (property field)
	GetRequest() Request
	// IsCBusMessageToServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusMessageToServer()
	// CreateBuilder creates a CBusMessageToServerBuilder
	CreateCBusMessageToServerBuilder() CBusMessageToServerBuilder
}

// _CBusMessageToServer is the data-structure of this message
type _CBusMessageToServer struct {
	CBusMessageContract
	Request Request
}

var _ CBusMessageToServer = (*_CBusMessageToServer)(nil)
var _ CBusMessageRequirements = (*_CBusMessageToServer)(nil)

// NewCBusMessageToServer factory function for _CBusMessageToServer
func NewCBusMessageToServer(request Request, requestContext RequestContext, cBusOptions CBusOptions) *_CBusMessageToServer {
	if request == nil {
		panic("request of type Request for CBusMessageToServer must not be nil")
	}
	_result := &_CBusMessageToServer{
		CBusMessageContract: NewCBusMessage(requestContext, cBusOptions),
		Request:             request,
	}
	_result.CBusMessageContract.(*_CBusMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusMessageToServerBuilder is a builder for CBusMessageToServer
type CBusMessageToServerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(request Request) CBusMessageToServerBuilder
	// WithRequest adds Request (property field)
	WithRequest(Request) CBusMessageToServerBuilder
	// WithRequestBuilder adds Request (property field) which is build by the builder
	WithRequestBuilder(func(RequestBuilder) RequestBuilder) CBusMessageToServerBuilder
	// Build builds the CBusMessageToServer or returns an error if something is wrong
	Build() (CBusMessageToServer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusMessageToServer
}

// NewCBusMessageToServerBuilder() creates a CBusMessageToServerBuilder
func NewCBusMessageToServerBuilder() CBusMessageToServerBuilder {
	return &_CBusMessageToServerBuilder{_CBusMessageToServer: new(_CBusMessageToServer)}
}

type _CBusMessageToServerBuilder struct {
	*_CBusMessageToServer

	parentBuilder *_CBusMessageBuilder

	err *utils.MultiError
}

var _ (CBusMessageToServerBuilder) = (*_CBusMessageToServerBuilder)(nil)

func (b *_CBusMessageToServerBuilder) setParent(contract CBusMessageContract) {
	b.CBusMessageContract = contract
}

func (b *_CBusMessageToServerBuilder) WithMandatoryFields(request Request) CBusMessageToServerBuilder {
	return b.WithRequest(request)
}

func (b *_CBusMessageToServerBuilder) WithRequest(request Request) CBusMessageToServerBuilder {
	b.Request = request
	return b
}

func (b *_CBusMessageToServerBuilder) WithRequestBuilder(builderSupplier func(RequestBuilder) RequestBuilder) CBusMessageToServerBuilder {
	builder := builderSupplier(b.Request.CreateRequestBuilder())
	var err error
	b.Request, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestBuilder failed"))
	}
	return b
}

func (b *_CBusMessageToServerBuilder) Build() (CBusMessageToServer, error) {
	if b.Request == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'request' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusMessageToServer.deepCopy(), nil
}

func (b *_CBusMessageToServerBuilder) MustBuild() CBusMessageToServer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CBusMessageToServerBuilder) Done() CBusMessageBuilder {
	return b.parentBuilder
}

func (b *_CBusMessageToServerBuilder) buildForCBusMessage() (CBusMessage, error) {
	return b.Build()
}

func (b *_CBusMessageToServerBuilder) DeepCopy() any {
	_copy := b.CreateCBusMessageToServerBuilder().(*_CBusMessageToServerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusMessageToServerBuilder creates a CBusMessageToServerBuilder
func (b *_CBusMessageToServer) CreateCBusMessageToServerBuilder() CBusMessageToServerBuilder {
	if b == nil {
		return NewCBusMessageToServerBuilder()
	}
	return &_CBusMessageToServerBuilder{_CBusMessageToServer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CBusMessageToServer) GetIsResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusMessageToServer) GetParent() CBusMessageContract {
	return m.CBusMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusMessageToServer) GetRequest() Request {
	return m.Request
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusMessageToServer(structType any) CBusMessageToServer {
	if casted, ok := structType.(CBusMessageToServer); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessageToServer); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessageToServer) GetTypeName() string {
	return "CBusMessageToServer"
}

func (m *_CBusMessageToServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusMessageContract.(*_CBusMessage).getLengthInBits(ctx))

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusMessageToServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusMessageToServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusMessage, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (__cBusMessageToServer CBusMessageToServer, err error) {
	m.CBusMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessageToServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessageToServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	request, err := ReadSimpleField[Request](ctx, "request", ReadComplex[Request](RequestParseWithBufferProducer[Request]((CBusOptions)(cBusOptions)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'request' field"))
	}
	m.Request = request

	if closeErr := readBuffer.CloseContext("CBusMessageToServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessageToServer")
	}

	return m, nil
}

func (m *_CBusMessageToServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusMessageToServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusMessageToServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusMessageToServer")
		}

		if err := WriteSimpleField[Request](ctx, "request", m.GetRequest(), WriteComplex[Request](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'request' field")
		}

		if popErr := writeBuffer.PopContext("CBusMessageToServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusMessageToServer")
		}
		return nil
	}
	return m.CBusMessageContract.(*_CBusMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusMessageToServer) IsCBusMessageToServer() {}

func (m *_CBusMessageToServer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusMessageToServer) deepCopy() *_CBusMessageToServer {
	if m == nil {
		return nil
	}
	_CBusMessageToServerCopy := &_CBusMessageToServer{
		m.CBusMessageContract.(*_CBusMessage).deepCopy(),
		m.Request.DeepCopy().(Request),
	}
	m.CBusMessageContract.(*_CBusMessage)._SubType = m
	return _CBusMessageToServerCopy
}

func (m *_CBusMessageToServer) String() string {
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
