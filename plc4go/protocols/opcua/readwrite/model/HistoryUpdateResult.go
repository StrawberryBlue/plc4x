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

// HistoryUpdateResult is the corresponding interface of HistoryUpdateResult
type HistoryUpdateResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfOperationResults returns NoOfOperationResults (property field)
	GetNoOfOperationResults() int32
	// GetOperationResults returns OperationResults (property field)
	GetOperationResults() []StatusCode
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsHistoryUpdateResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryUpdateResult()
	// CreateBuilder creates a HistoryUpdateResultBuilder
	CreateHistoryUpdateResultBuilder() HistoryUpdateResultBuilder
}

// _HistoryUpdateResult is the data-structure of this message
type _HistoryUpdateResult struct {
	ExtensionObjectDefinitionContract
	StatusCode           StatusCode
	NoOfOperationResults int32
	OperationResults     []StatusCode
	NoOfDiagnosticInfos  int32
	DiagnosticInfos      []DiagnosticInfo
}

var _ HistoryUpdateResult = (*_HistoryUpdateResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryUpdateResult)(nil)

// NewHistoryUpdateResult factory function for _HistoryUpdateResult
func NewHistoryUpdateResult(statusCode StatusCode, noOfOperationResults int32, operationResults []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_HistoryUpdateResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for HistoryUpdateResult must not be nil")
	}
	_result := &_HistoryUpdateResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		NoOfOperationResults:              noOfOperationResults,
		OperationResults:                  operationResults,
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

// HistoryUpdateResultBuilder is a builder for HistoryUpdateResult
type HistoryUpdateResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, noOfOperationResults int32, operationResults []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) HistoryUpdateResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) HistoryUpdateResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) HistoryUpdateResultBuilder
	// WithNoOfOperationResults adds NoOfOperationResults (property field)
	WithNoOfOperationResults(int32) HistoryUpdateResultBuilder
	// WithOperationResults adds OperationResults (property field)
	WithOperationResults(...StatusCode) HistoryUpdateResultBuilder
	// WithNoOfDiagnosticInfos adds NoOfDiagnosticInfos (property field)
	WithNoOfDiagnosticInfos(int32) HistoryUpdateResultBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) HistoryUpdateResultBuilder
	// Build builds the HistoryUpdateResult or returns an error if something is wrong
	Build() (HistoryUpdateResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryUpdateResult
}

// NewHistoryUpdateResultBuilder() creates a HistoryUpdateResultBuilder
func NewHistoryUpdateResultBuilder() HistoryUpdateResultBuilder {
	return &_HistoryUpdateResultBuilder{_HistoryUpdateResult: new(_HistoryUpdateResult)}
}

type _HistoryUpdateResultBuilder struct {
	*_HistoryUpdateResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryUpdateResultBuilder) = (*_HistoryUpdateResultBuilder)(nil)

func (b *_HistoryUpdateResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_HistoryUpdateResultBuilder) WithMandatoryFields(statusCode StatusCode, noOfOperationResults int32, operationResults []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) HistoryUpdateResultBuilder {
	return b.WithStatusCode(statusCode).WithNoOfOperationResults(noOfOperationResults).WithOperationResults(operationResults...).WithNoOfDiagnosticInfos(noOfDiagnosticInfos).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_HistoryUpdateResultBuilder) WithStatusCode(statusCode StatusCode) HistoryUpdateResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_HistoryUpdateResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) HistoryUpdateResultBuilder {
	builder := builderSupplier(b.StatusCode.CreateStatusCodeBuilder())
	var err error
	b.StatusCode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_HistoryUpdateResultBuilder) WithNoOfOperationResults(noOfOperationResults int32) HistoryUpdateResultBuilder {
	b.NoOfOperationResults = noOfOperationResults
	return b
}

func (b *_HistoryUpdateResultBuilder) WithOperationResults(operationResults ...StatusCode) HistoryUpdateResultBuilder {
	b.OperationResults = operationResults
	return b
}

func (b *_HistoryUpdateResultBuilder) WithNoOfDiagnosticInfos(noOfDiagnosticInfos int32) HistoryUpdateResultBuilder {
	b.NoOfDiagnosticInfos = noOfDiagnosticInfos
	return b
}

func (b *_HistoryUpdateResultBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) HistoryUpdateResultBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_HistoryUpdateResultBuilder) Build() (HistoryUpdateResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryUpdateResult.deepCopy(), nil
}

func (b *_HistoryUpdateResultBuilder) MustBuild() HistoryUpdateResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_HistoryUpdateResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_HistoryUpdateResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryUpdateResultBuilder) DeepCopy() any {
	_copy := b.CreateHistoryUpdateResultBuilder().(*_HistoryUpdateResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryUpdateResultBuilder creates a HistoryUpdateResultBuilder
func (b *_HistoryUpdateResult) CreateHistoryUpdateResultBuilder() HistoryUpdateResultBuilder {
	if b == nil {
		return NewHistoryUpdateResultBuilder()
	}
	return &_HistoryUpdateResultBuilder{_HistoryUpdateResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateResult) GetIdentifier() string {
	return "697"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryUpdateResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_HistoryUpdateResult) GetNoOfOperationResults() int32 {
	return m.NoOfOperationResults
}

func (m *_HistoryUpdateResult) GetOperationResults() []StatusCode {
	return m.OperationResults
}

func (m *_HistoryUpdateResult) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_HistoryUpdateResult) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryUpdateResult(structType any) HistoryUpdateResult {
	if casted, ok := structType.(HistoryUpdateResult); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateResult); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateResult) GetTypeName() string {
	return "HistoryUpdateResult"
}

func (m *_HistoryUpdateResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfOperationResults)
	lengthInBits += 32

	// Array field
	if len(m.OperationResults) > 0 {
		for _curItem, element := range m.OperationResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OperationResults), _curItem)
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

func (m *_HistoryUpdateResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryUpdateResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyUpdateResult HistoryUpdateResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryUpdateResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	noOfOperationResults, err := ReadSimpleField(ctx, "noOfOperationResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfOperationResults' field"))
	}
	m.NoOfOperationResults = noOfOperationResults

	operationResults, err := ReadCountArrayField[StatusCode](ctx, "operationResults", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfOperationResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationResults' field"))
	}
	m.OperationResults = operationResults

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

	if closeErr := readBuffer.CloseContext("HistoryUpdateResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateResult")
	}

	return m, nil
}

func (m *_HistoryUpdateResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfOperationResults", m.GetNoOfOperationResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfOperationResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "operationResults", m.GetOperationResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'operationResults' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("HistoryUpdateResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryUpdateResult) IsHistoryUpdateResult() {}

func (m *_HistoryUpdateResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryUpdateResult) deepCopy() *_HistoryUpdateResult {
	if m == nil {
		return nil
	}
	_HistoryUpdateResultCopy := &_HistoryUpdateResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.NoOfOperationResults,
		utils.DeepCopySlice[StatusCode, StatusCode](m.OperationResults),
		m.NoOfDiagnosticInfos,
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryUpdateResultCopy
}

func (m *_HistoryUpdateResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
