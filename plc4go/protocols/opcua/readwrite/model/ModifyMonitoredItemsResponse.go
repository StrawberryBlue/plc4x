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

// ModifyMonitoredItemsResponse is the corresponding interface of ModifyMonitoredItemsResponse
type ModifyMonitoredItemsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsModifyMonitoredItemsResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModifyMonitoredItemsResponse()
	// CreateBuilder creates a ModifyMonitoredItemsResponseBuilder
	CreateModifyMonitoredItemsResponseBuilder() ModifyMonitoredItemsResponseBuilder
}

// _ModifyMonitoredItemsResponse is the data-structure of this message
type _ModifyMonitoredItemsResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

var _ ModifyMonitoredItemsResponse = (*_ModifyMonitoredItemsResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ModifyMonitoredItemsResponse)(nil)

// NewModifyMonitoredItemsResponse factory function for _ModifyMonitoredItemsResponse
func NewModifyMonitoredItemsResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_ModifyMonitoredItemsResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for ModifyMonitoredItemsResponse must not be nil")
	}
	_result := &_ModifyMonitoredItemsResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfResults:                       noOfResults,
		Results:                           results,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModifyMonitoredItemsResponseBuilder is a builder for ModifyMonitoredItemsResponse
type ModifyMonitoredItemsResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) ModifyMonitoredItemsResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ExtensionObjectDefinition) ModifyMonitoredItemsResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ModifyMonitoredItemsResponseBuilder
	// WithNoOfResults adds NoOfResults (property field)
	WithNoOfResults(int32) ModifyMonitoredItemsResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...ExtensionObjectDefinition) ModifyMonitoredItemsResponseBuilder
	// WithNoOfDiagnosticInfos adds NoOfDiagnosticInfos (property field)
	WithNoOfDiagnosticInfos(int32) ModifyMonitoredItemsResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) ModifyMonitoredItemsResponseBuilder
	// Build builds the ModifyMonitoredItemsResponse or returns an error if something is wrong
	Build() (ModifyMonitoredItemsResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModifyMonitoredItemsResponse
}

// NewModifyMonitoredItemsResponseBuilder() creates a ModifyMonitoredItemsResponseBuilder
func NewModifyMonitoredItemsResponseBuilder() ModifyMonitoredItemsResponseBuilder {
	return &_ModifyMonitoredItemsResponseBuilder{_ModifyMonitoredItemsResponse: new(_ModifyMonitoredItemsResponse)}
}

type _ModifyMonitoredItemsResponseBuilder struct {
	*_ModifyMonitoredItemsResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ModifyMonitoredItemsResponseBuilder) = (*_ModifyMonitoredItemsResponseBuilder)(nil)

func (b *_ModifyMonitoredItemsResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) ModifyMonitoredItemsResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithNoOfResults(noOfResults).WithResults(results...).WithNoOfDiagnosticInfos(noOfDiagnosticInfos).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithResponseHeader(responseHeader ExtensionObjectDefinition) ModifyMonitoredItemsResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ModifyMonitoredItemsResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithNoOfResults(noOfResults int32) ModifyMonitoredItemsResponseBuilder {
	b.NoOfResults = noOfResults
	return b
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithResults(results ...ExtensionObjectDefinition) ModifyMonitoredItemsResponseBuilder {
	b.Results = results
	return b
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithNoOfDiagnosticInfos(noOfDiagnosticInfos int32) ModifyMonitoredItemsResponseBuilder {
	b.NoOfDiagnosticInfos = noOfDiagnosticInfos
	return b
}

func (b *_ModifyMonitoredItemsResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) ModifyMonitoredItemsResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_ModifyMonitoredItemsResponseBuilder) Build() (ModifyMonitoredItemsResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModifyMonitoredItemsResponse.deepCopy(), nil
}

func (b *_ModifyMonitoredItemsResponseBuilder) MustBuild() ModifyMonitoredItemsResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ModifyMonitoredItemsResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ModifyMonitoredItemsResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ModifyMonitoredItemsResponseBuilder) DeepCopy() any {
	_copy := b.CreateModifyMonitoredItemsResponseBuilder().(*_ModifyMonitoredItemsResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModifyMonitoredItemsResponseBuilder creates a ModifyMonitoredItemsResponseBuilder
func (b *_ModifyMonitoredItemsResponse) CreateModifyMonitoredItemsResponseBuilder() ModifyMonitoredItemsResponseBuilder {
	if b == nil {
		return NewModifyMonitoredItemsResponseBuilder()
	}
	return &_ModifyMonitoredItemsResponseBuilder{_ModifyMonitoredItemsResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModifyMonitoredItemsResponse) GetIdentifier() string {
	return "766"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModifyMonitoredItemsResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModifyMonitoredItemsResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_ModifyMonitoredItemsResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_ModifyMonitoredItemsResponse) GetResults() []ExtensionObjectDefinition {
	return m.Results
}

func (m *_ModifyMonitoredItemsResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_ModifyMonitoredItemsResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModifyMonitoredItemsResponse(structType any) ModifyMonitoredItemsResponse {
	if casted, ok := structType.(ModifyMonitoredItemsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModifyMonitoredItemsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModifyMonitoredItemsResponse) GetTypeName() string {
	return "ModifyMonitoredItemsResponse"
}

func (m *_ModifyMonitoredItemsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ModifyMonitoredItemsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModifyMonitoredItemsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__modifyMonitoredItemsResponse ModifyMonitoredItemsResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModifyMonitoredItemsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModifyMonitoredItemsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadSimpleField(ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	m.NoOfResults = noOfResults

	results, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "results", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("760")), readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("ModifyMonitoredItemsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModifyMonitoredItemsResponse")
	}

	return m, nil
}

func (m *_ModifyMonitoredItemsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModifyMonitoredItemsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModifyMonitoredItemsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModifyMonitoredItemsResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfResults", m.GetNoOfResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("ModifyMonitoredItemsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModifyMonitoredItemsResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModifyMonitoredItemsResponse) IsModifyMonitoredItemsResponse() {}

func (m *_ModifyMonitoredItemsResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModifyMonitoredItemsResponse) deepCopy() *_ModifyMonitoredItemsResponse {
	if m == nil {
		return nil
	}
	_ModifyMonitoredItemsResponseCopy := &_ModifyMonitoredItemsResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfResults,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.Results),
		m.NoOfDiagnosticInfos,
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ModifyMonitoredItemsResponseCopy
}

func (m *_ModifyMonitoredItemsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
