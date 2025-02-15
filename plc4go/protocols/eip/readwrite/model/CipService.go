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

// CipService is the corresponding interface of CipService
type CipService interface {
	CipServiceContract
	CipServiceRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCipService is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipService()
	// CreateBuilder creates a CipServiceBuilder
	CreateCipServiceBuilder() CipServiceBuilder
}

// CipServiceContract provides a set of functions which can be overwritten by a sub struct
type CipServiceContract interface {
	// GetServiceLen() returns a parser argument
	GetServiceLen() uint16
	// IsCipService is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipService()
	// CreateBuilder creates a CipServiceBuilder
	CreateCipServiceBuilder() CipServiceBuilder
}

// CipServiceRequirements provides a set of functions which need to be implemented by a sub struct
type CipServiceRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetConnected returns Connected (discriminator field)
	GetConnected() bool
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetService returns Service (discriminator field)
	GetService() uint8
}

// _CipService is the data-structure of this message
type _CipService struct {
	_SubType CipService

	// Arguments.
	ServiceLen uint16
}

var _ CipServiceContract = (*_CipService)(nil)

// NewCipService factory function for _CipService
func NewCipService(serviceLen uint16) *_CipService {
	return &_CipService{ServiceLen: serviceLen}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipServiceBuilder is a builder for CipService
type CipServiceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() CipServiceBuilder
	// AsGetAttributeAllRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsGetAttributeAllRequest() interface {
		GetAttributeAllRequestBuilder
		Done() CipServiceBuilder
	}
	// AsGetAttributeAllResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsGetAttributeAllResponse() interface {
		GetAttributeAllResponseBuilder
		Done() CipServiceBuilder
	}
	// AsSetAttributeAllRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsSetAttributeAllRequest() interface {
		SetAttributeAllRequestBuilder
		Done() CipServiceBuilder
	}
	// AsSetAttributeAllResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsSetAttributeAllResponse() interface {
		SetAttributeAllResponseBuilder
		Done() CipServiceBuilder
	}
	// AsGetAttributeListRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsGetAttributeListRequest() interface {
		GetAttributeListRequestBuilder
		Done() CipServiceBuilder
	}
	// AsGetAttributeListResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsGetAttributeListResponse() interface {
		GetAttributeListResponseBuilder
		Done() CipServiceBuilder
	}
	// AsSetAttributeListRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsSetAttributeListRequest() interface {
		SetAttributeListRequestBuilder
		Done() CipServiceBuilder
	}
	// AsSetAttributeListResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsSetAttributeListResponse() interface {
		SetAttributeListResponseBuilder
		Done() CipServiceBuilder
	}
	// AsMultipleServiceRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsMultipleServiceRequest() interface {
		MultipleServiceRequestBuilder
		Done() CipServiceBuilder
	}
	// AsMultipleServiceResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsMultipleServiceResponse() interface {
		MultipleServiceResponseBuilder
		Done() CipServiceBuilder
	}
	// AsGetAttributeSingleRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsGetAttributeSingleRequest() interface {
		GetAttributeSingleRequestBuilder
		Done() CipServiceBuilder
	}
	// AsGetAttributeSingleResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsGetAttributeSingleResponse() interface {
		GetAttributeSingleResponseBuilder
		Done() CipServiceBuilder
	}
	// AsSetAttributeSingleRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsSetAttributeSingleRequest() interface {
		SetAttributeSingleRequestBuilder
		Done() CipServiceBuilder
	}
	// AsSetAttributeSingleResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsSetAttributeSingleResponse() interface {
		SetAttributeSingleResponseBuilder
		Done() CipServiceBuilder
	}
	// AsCipReadRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipReadRequest() interface {
		CipReadRequestBuilder
		Done() CipServiceBuilder
	}
	// AsCipReadResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipReadResponse() interface {
		CipReadResponseBuilder
		Done() CipServiceBuilder
	}
	// AsCipWriteRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipWriteRequest() interface {
		CipWriteRequestBuilder
		Done() CipServiceBuilder
	}
	// AsCipWriteResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipWriteResponse() interface {
		CipWriteResponseBuilder
		Done() CipServiceBuilder
	}
	// AsCipConnectionManagerCloseRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipConnectionManagerCloseRequest() interface {
		CipConnectionManagerCloseRequestBuilder
		Done() CipServiceBuilder
	}
	// AsCipConnectionManagerCloseResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipConnectionManagerCloseResponse() interface {
		CipConnectionManagerCloseResponseBuilder
		Done() CipServiceBuilder
	}
	// AsCipUnconnectedRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipUnconnectedRequest() interface {
		CipUnconnectedRequestBuilder
		Done() CipServiceBuilder
	}
	// AsCipConnectedRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipConnectedRequest() interface {
		CipConnectedRequestBuilder
		Done() CipServiceBuilder
	}
	// AsCipConnectedResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipConnectedResponse() interface {
		CipConnectedResponseBuilder
		Done() CipServiceBuilder
	}
	// AsCipConnectionManagerRequest converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipConnectionManagerRequest() interface {
		CipConnectionManagerRequestBuilder
		Done() CipServiceBuilder
	}
	// AsCipConnectionManagerResponse converts this build to a subType of CipService. It is always possible to return to current builder using Done()
	AsCipConnectionManagerResponse() interface {
		CipConnectionManagerResponseBuilder
		Done() CipServiceBuilder
	}
	// Build builds the CipService or returns an error if something is wrong
	PartialBuild() (CipServiceContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() CipServiceContract
	// Build builds the CipService or returns an error if something is wrong
	Build() (CipService, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipService
}

// NewCipServiceBuilder() creates a CipServiceBuilder
func NewCipServiceBuilder() CipServiceBuilder {
	return &_CipServiceBuilder{_CipService: new(_CipService)}
}

type _CipServiceChildBuilder interface {
	utils.Copyable
	setParent(CipServiceContract)
	buildForCipService() (CipService, error)
}

type _CipServiceBuilder struct {
	*_CipService

	childBuilder _CipServiceChildBuilder

	err *utils.MultiError
}

var _ (CipServiceBuilder) = (*_CipServiceBuilder)(nil)

func (b *_CipServiceBuilder) WithMandatoryFields() CipServiceBuilder {
	return b
}

func (b *_CipServiceBuilder) PartialBuild() (CipServiceContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipService.deepCopy(), nil
}

func (b *_CipServiceBuilder) PartialMustBuild() CipServiceContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipServiceBuilder) AsGetAttributeAllRequest() interface {
	GetAttributeAllRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		GetAttributeAllRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewGetAttributeAllRequestBuilder().(*_GetAttributeAllRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsGetAttributeAllResponse() interface {
	GetAttributeAllResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		GetAttributeAllResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewGetAttributeAllResponseBuilder().(*_GetAttributeAllResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsSetAttributeAllRequest() interface {
	SetAttributeAllRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SetAttributeAllRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewSetAttributeAllRequestBuilder().(*_SetAttributeAllRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsSetAttributeAllResponse() interface {
	SetAttributeAllResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SetAttributeAllResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewSetAttributeAllResponseBuilder().(*_SetAttributeAllResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsGetAttributeListRequest() interface {
	GetAttributeListRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		GetAttributeListRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewGetAttributeListRequestBuilder().(*_GetAttributeListRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsGetAttributeListResponse() interface {
	GetAttributeListResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		GetAttributeListResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewGetAttributeListResponseBuilder().(*_GetAttributeListResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsSetAttributeListRequest() interface {
	SetAttributeListRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SetAttributeListRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewSetAttributeListRequestBuilder().(*_SetAttributeListRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsSetAttributeListResponse() interface {
	SetAttributeListResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SetAttributeListResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewSetAttributeListResponseBuilder().(*_SetAttributeListResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsMultipleServiceRequest() interface {
	MultipleServiceRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MultipleServiceRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewMultipleServiceRequestBuilder().(*_MultipleServiceRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsMultipleServiceResponse() interface {
	MultipleServiceResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MultipleServiceResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewMultipleServiceResponseBuilder().(*_MultipleServiceResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsGetAttributeSingleRequest() interface {
	GetAttributeSingleRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		GetAttributeSingleRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewGetAttributeSingleRequestBuilder().(*_GetAttributeSingleRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsGetAttributeSingleResponse() interface {
	GetAttributeSingleResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		GetAttributeSingleResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewGetAttributeSingleResponseBuilder().(*_GetAttributeSingleResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsSetAttributeSingleRequest() interface {
	SetAttributeSingleRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SetAttributeSingleRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewSetAttributeSingleRequestBuilder().(*_SetAttributeSingleRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsSetAttributeSingleResponse() interface {
	SetAttributeSingleResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SetAttributeSingleResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewSetAttributeSingleResponseBuilder().(*_SetAttributeSingleResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipReadRequest() interface {
	CipReadRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipReadRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipReadRequestBuilder().(*_CipReadRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipReadResponse() interface {
	CipReadResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipReadResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipReadResponseBuilder().(*_CipReadResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipWriteRequest() interface {
	CipWriteRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipWriteRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipWriteRequestBuilder().(*_CipWriteRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipWriteResponse() interface {
	CipWriteResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipWriteResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipWriteResponseBuilder().(*_CipWriteResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipConnectionManagerCloseRequest() interface {
	CipConnectionManagerCloseRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipConnectionManagerCloseRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipConnectionManagerCloseRequestBuilder().(*_CipConnectionManagerCloseRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipConnectionManagerCloseResponse() interface {
	CipConnectionManagerCloseResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipConnectionManagerCloseResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipConnectionManagerCloseResponseBuilder().(*_CipConnectionManagerCloseResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipUnconnectedRequest() interface {
	CipUnconnectedRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipUnconnectedRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipUnconnectedRequestBuilder().(*_CipUnconnectedRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipConnectedRequest() interface {
	CipConnectedRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipConnectedRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipConnectedRequestBuilder().(*_CipConnectedRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipConnectedResponse() interface {
	CipConnectedResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipConnectedResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipConnectedResponseBuilder().(*_CipConnectedResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipConnectionManagerRequest() interface {
	CipConnectionManagerRequestBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipConnectionManagerRequestBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipConnectionManagerRequestBuilder().(*_CipConnectionManagerRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) AsCipConnectionManagerResponse() interface {
	CipConnectionManagerResponseBuilder
	Done() CipServiceBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipConnectionManagerResponseBuilder
		Done() CipServiceBuilder
	}); ok {
		return cb
	}
	cb := NewCipConnectionManagerResponseBuilder().(*_CipConnectionManagerResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CipServiceBuilder) Build() (CipService, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCipService()
}

func (b *_CipServiceBuilder) MustBuild() CipService {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipServiceBuilder) DeepCopy() any {
	_copy := b.CreateCipServiceBuilder().(*_CipServiceBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_CipServiceChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipServiceBuilder creates a CipServiceBuilder
func (b *_CipService) CreateCipServiceBuilder() CipServiceBuilder {
	if b == nil {
		return NewCipServiceBuilder()
	}
	return &_CipServiceBuilder{_CipService: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipService(structType any) CipService {
	if casted, ok := structType.(CipService); ok {
		return casted
	}
	if casted, ok := structType.(*CipService); ok {
		return *casted
	}
	return nil
}

func (m *_CipService) GetTypeName() string {
	return "CipService"
}

func (m *_CipService) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (response)
	lengthInBits += 1
	// Discriminator Field (service)
	lengthInBits += 7

	return lengthInBits
}

func (m *_CipService) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CipServiceParse[T CipService](ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (T, error) {
	return CipServiceParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipServiceParseWithBufferProducer[T CipService](connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CipServiceParseWithBuffer[T](ctx, readBuffer, connected, serviceLen)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CipServiceParseWithBuffer[T CipService](ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (T, error) {
	v, err := (&_CipService{ServiceLen: serviceLen}).parse(ctx, readBuffer, connected, serviceLen)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_CipService) parse(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (__cipService CipService, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipService"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipService")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	response, err := ReadDiscriminatorField[bool](ctx, "response", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'response' field"))
	}

	service, err := ReadDiscriminatorField[uint8](ctx, "service", ReadUnsignedByte(readBuffer, uint8(7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'service' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CipService
	switch {
	case service == 0x01 && response == bool(false): // GetAttributeAllRequest
		if _child, err = new(_GetAttributeAllRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeAllRequest for type-switch of CipService")
		}
	case service == 0x01 && response == bool(true): // GetAttributeAllResponse
		if _child, err = new(_GetAttributeAllResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeAllResponse for type-switch of CipService")
		}
	case service == 0x02 && response == bool(false): // SetAttributeAllRequest
		if _child, err = new(_SetAttributeAllRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeAllRequest for type-switch of CipService")
		}
	case service == 0x02 && response == bool(true): // SetAttributeAllResponse
		if _child, err = new(_SetAttributeAllResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeAllResponse for type-switch of CipService")
		}
	case service == 0x03 && response == bool(false): // GetAttributeListRequest
		if _child, err = new(_GetAttributeListRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeListRequest for type-switch of CipService")
		}
	case service == 0x03 && response == bool(true): // GetAttributeListResponse
		if _child, err = new(_GetAttributeListResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeListResponse for type-switch of CipService")
		}
	case service == 0x04 && response == bool(false): // SetAttributeListRequest
		if _child, err = new(_SetAttributeListRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeListRequest for type-switch of CipService")
		}
	case service == 0x04 && response == bool(true): // SetAttributeListResponse
		if _child, err = new(_SetAttributeListResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeListResponse for type-switch of CipService")
		}
	case service == 0x0A && response == bool(false): // MultipleServiceRequest
		if _child, err = new(_MultipleServiceRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MultipleServiceRequest for type-switch of CipService")
		}
	case service == 0x0A && response == bool(true): // MultipleServiceResponse
		if _child, err = new(_MultipleServiceResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MultipleServiceResponse for type-switch of CipService")
		}
	case service == 0x0E && response == bool(false): // GetAttributeSingleRequest
		if _child, err = new(_GetAttributeSingleRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeSingleRequest for type-switch of CipService")
		}
	case service == 0x0E && response == bool(true): // GetAttributeSingleResponse
		if _child, err = new(_GetAttributeSingleResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeSingleResponse for type-switch of CipService")
		}
	case service == 0x10 && response == bool(false): // SetAttributeSingleRequest
		if _child, err = new(_SetAttributeSingleRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeSingleRequest for type-switch of CipService")
		}
	case service == 0x10 && response == bool(true): // SetAttributeSingleResponse
		if _child, err = new(_SetAttributeSingleResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeSingleResponse for type-switch of CipService")
		}
	case service == 0x4C && response == bool(false): // CipReadRequest
		if _child, err = new(_CipReadRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipReadRequest for type-switch of CipService")
		}
	case service == 0x4C && response == bool(true): // CipReadResponse
		if _child, err = new(_CipReadResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipReadResponse for type-switch of CipService")
		}
	case service == 0x4D && response == bool(false): // CipWriteRequest
		if _child, err = new(_CipWriteRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipWriteRequest for type-switch of CipService")
		}
	case service == 0x4D && response == bool(true): // CipWriteResponse
		if _child, err = new(_CipWriteResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipWriteResponse for type-switch of CipService")
		}
	case service == 0x4E && response == bool(false): // CipConnectionManagerCloseRequest
		if _child, err = new(_CipConnectionManagerCloseRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerCloseRequest for type-switch of CipService")
		}
	case service == 0x4E && response == bool(true): // CipConnectionManagerCloseResponse
		if _child, err = new(_CipConnectionManagerCloseResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerCloseResponse for type-switch of CipService")
		}
	case service == 0x52 && response == bool(false) && connected == bool(false): // CipUnconnectedRequest
		if _child, err = new(_CipUnconnectedRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipUnconnectedRequest for type-switch of CipService")
		}
	case service == 0x52 && response == bool(false) && connected == bool(true): // CipConnectedRequest
		if _child, err = new(_CipConnectedRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectedRequest for type-switch of CipService")
		}
	case service == 0x52 && response == bool(true): // CipConnectedResponse
		if _child, err = new(_CipConnectedResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectedResponse for type-switch of CipService")
		}
	case service == 0x5B && response == bool(false): // CipConnectionManagerRequest
		if _child, err = new(_CipConnectionManagerRequest).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerRequest for type-switch of CipService")
		}
	case service == 0x5B && response == bool(true): // CipConnectionManagerResponse
		if _child, err = new(_CipConnectionManagerResponse).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerResponse for type-switch of CipService")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [service=%v, response=%v, connected=%v]", service, response, connected)
	}

	if closeErr := readBuffer.CloseContext("CipService"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipService")
	}

	return _child, nil
}

func (pm *_CipService) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CipService, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CipService"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CipService")
	}

	if err := WriteDiscriminatorField(ctx, "response", m.GetResponse(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'response' field")
	}

	if err := WriteDiscriminatorField(ctx, "service", m.GetService(), WriteUnsignedByte(writeBuffer, 7)); err != nil {
		return errors.Wrap(err, "Error serializing 'service' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CipService"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CipService")
	}
	return nil
}

////
// Arguments Getter

func (m *_CipService) GetServiceLen() uint16 {
	return m.ServiceLen
}

//
////

func (m *_CipService) IsCipService() {}

func (m *_CipService) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipService) deepCopy() *_CipService {
	if m == nil {
		return nil
	}
	_CipServiceCopy := &_CipService{
		nil, // will be set by child
		m.ServiceLen,
	}
	return _CipServiceCopy
}
