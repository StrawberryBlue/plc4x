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

// MonitoredItemCreateResult is the corresponding interface of MonitoredItemCreateResult
type MonitoredItemCreateResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetMonitoredItemId returns MonitoredItemId (property field)
	GetMonitoredItemId() uint32
	// GetRevisedSamplingInterval returns RevisedSamplingInterval (property field)
	GetRevisedSamplingInterval() float64
	// GetRevisedQueueSize returns RevisedQueueSize (property field)
	GetRevisedQueueSize() uint32
	// GetFilterResult returns FilterResult (property field)
	GetFilterResult() ExtensionObject
	// IsMonitoredItemCreateResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMonitoredItemCreateResult()
	// CreateBuilder creates a MonitoredItemCreateResultBuilder
	CreateMonitoredItemCreateResultBuilder() MonitoredItemCreateResultBuilder
}

// _MonitoredItemCreateResult is the data-structure of this message
type _MonitoredItemCreateResult struct {
	ExtensionObjectDefinitionContract
	StatusCode              StatusCode
	MonitoredItemId         uint32
	RevisedSamplingInterval float64
	RevisedQueueSize        uint32
	FilterResult            ExtensionObject
}

var _ MonitoredItemCreateResult = (*_MonitoredItemCreateResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_MonitoredItemCreateResult)(nil)

// NewMonitoredItemCreateResult factory function for _MonitoredItemCreateResult
func NewMonitoredItemCreateResult(statusCode StatusCode, monitoredItemId uint32, revisedSamplingInterval float64, revisedQueueSize uint32, filterResult ExtensionObject) *_MonitoredItemCreateResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for MonitoredItemCreateResult must not be nil")
	}
	if filterResult == nil {
		panic("filterResult of type ExtensionObject for MonitoredItemCreateResult must not be nil")
	}
	_result := &_MonitoredItemCreateResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		MonitoredItemId:                   monitoredItemId,
		RevisedSamplingInterval:           revisedSamplingInterval,
		RevisedQueueSize:                  revisedQueueSize,
		FilterResult:                      filterResult,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MonitoredItemCreateResultBuilder is a builder for MonitoredItemCreateResult
type MonitoredItemCreateResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, monitoredItemId uint32, revisedSamplingInterval float64, revisedQueueSize uint32, filterResult ExtensionObject) MonitoredItemCreateResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) MonitoredItemCreateResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) MonitoredItemCreateResultBuilder
	// WithMonitoredItemId adds MonitoredItemId (property field)
	WithMonitoredItemId(uint32) MonitoredItemCreateResultBuilder
	// WithRevisedSamplingInterval adds RevisedSamplingInterval (property field)
	WithRevisedSamplingInterval(float64) MonitoredItemCreateResultBuilder
	// WithRevisedQueueSize adds RevisedQueueSize (property field)
	WithRevisedQueueSize(uint32) MonitoredItemCreateResultBuilder
	// WithFilterResult adds FilterResult (property field)
	WithFilterResult(ExtensionObject) MonitoredItemCreateResultBuilder
	// WithFilterResultBuilder adds FilterResult (property field) which is build by the builder
	WithFilterResultBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) MonitoredItemCreateResultBuilder
	// Build builds the MonitoredItemCreateResult or returns an error if something is wrong
	Build() (MonitoredItemCreateResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MonitoredItemCreateResult
}

// NewMonitoredItemCreateResultBuilder() creates a MonitoredItemCreateResultBuilder
func NewMonitoredItemCreateResultBuilder() MonitoredItemCreateResultBuilder {
	return &_MonitoredItemCreateResultBuilder{_MonitoredItemCreateResult: new(_MonitoredItemCreateResult)}
}

type _MonitoredItemCreateResultBuilder struct {
	*_MonitoredItemCreateResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (MonitoredItemCreateResultBuilder) = (*_MonitoredItemCreateResultBuilder)(nil)

func (b *_MonitoredItemCreateResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_MonitoredItemCreateResultBuilder) WithMandatoryFields(statusCode StatusCode, monitoredItemId uint32, revisedSamplingInterval float64, revisedQueueSize uint32, filterResult ExtensionObject) MonitoredItemCreateResultBuilder {
	return b.WithStatusCode(statusCode).WithMonitoredItemId(monitoredItemId).WithRevisedSamplingInterval(revisedSamplingInterval).WithRevisedQueueSize(revisedQueueSize).WithFilterResult(filterResult)
}

func (b *_MonitoredItemCreateResultBuilder) WithStatusCode(statusCode StatusCode) MonitoredItemCreateResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_MonitoredItemCreateResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) MonitoredItemCreateResultBuilder {
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

func (b *_MonitoredItemCreateResultBuilder) WithMonitoredItemId(monitoredItemId uint32) MonitoredItemCreateResultBuilder {
	b.MonitoredItemId = monitoredItemId
	return b
}

func (b *_MonitoredItemCreateResultBuilder) WithRevisedSamplingInterval(revisedSamplingInterval float64) MonitoredItemCreateResultBuilder {
	b.RevisedSamplingInterval = revisedSamplingInterval
	return b
}

func (b *_MonitoredItemCreateResultBuilder) WithRevisedQueueSize(revisedQueueSize uint32) MonitoredItemCreateResultBuilder {
	b.RevisedQueueSize = revisedQueueSize
	return b
}

func (b *_MonitoredItemCreateResultBuilder) WithFilterResult(filterResult ExtensionObject) MonitoredItemCreateResultBuilder {
	b.FilterResult = filterResult
	return b
}

func (b *_MonitoredItemCreateResultBuilder) WithFilterResultBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) MonitoredItemCreateResultBuilder {
	builder := builderSupplier(b.FilterResult.CreateExtensionObjectBuilder())
	var err error
	b.FilterResult, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_MonitoredItemCreateResultBuilder) Build() (MonitoredItemCreateResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.FilterResult == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'filterResult' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MonitoredItemCreateResult.deepCopy(), nil
}

func (b *_MonitoredItemCreateResultBuilder) MustBuild() MonitoredItemCreateResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MonitoredItemCreateResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_MonitoredItemCreateResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_MonitoredItemCreateResultBuilder) DeepCopy() any {
	_copy := b.CreateMonitoredItemCreateResultBuilder().(*_MonitoredItemCreateResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMonitoredItemCreateResultBuilder creates a MonitoredItemCreateResultBuilder
func (b *_MonitoredItemCreateResult) CreateMonitoredItemCreateResultBuilder() MonitoredItemCreateResultBuilder {
	if b == nil {
		return NewMonitoredItemCreateResultBuilder()
	}
	return &_MonitoredItemCreateResultBuilder{_MonitoredItemCreateResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemCreateResult) GetIdentifier() string {
	return "748"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemCreateResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemCreateResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_MonitoredItemCreateResult) GetMonitoredItemId() uint32 {
	return m.MonitoredItemId
}

func (m *_MonitoredItemCreateResult) GetRevisedSamplingInterval() float64 {
	return m.RevisedSamplingInterval
}

func (m *_MonitoredItemCreateResult) GetRevisedQueueSize() uint32 {
	return m.RevisedQueueSize
}

func (m *_MonitoredItemCreateResult) GetFilterResult() ExtensionObject {
	return m.FilterResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMonitoredItemCreateResult(structType any) MonitoredItemCreateResult {
	if casted, ok := structType.(MonitoredItemCreateResult); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemCreateResult); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemCreateResult) GetTypeName() string {
	return "MonitoredItemCreateResult"
}

func (m *_MonitoredItemCreateResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (monitoredItemId)
	lengthInBits += 32

	// Simple field (revisedSamplingInterval)
	lengthInBits += 64

	// Simple field (revisedQueueSize)
	lengthInBits += 32

	// Simple field (filterResult)
	lengthInBits += m.FilterResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemCreateResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredItemCreateResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__monitoredItemCreateResult MonitoredItemCreateResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredItemCreateResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemCreateResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	monitoredItemId, err := ReadSimpleField(ctx, "monitoredItemId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItemId' field"))
	}
	m.MonitoredItemId = monitoredItemId

	revisedSamplingInterval, err := ReadSimpleField(ctx, "revisedSamplingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedSamplingInterval' field"))
	}
	m.RevisedSamplingInterval = revisedSamplingInterval

	revisedQueueSize, err := ReadSimpleField(ctx, "revisedQueueSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedQueueSize' field"))
	}
	m.RevisedQueueSize = revisedQueueSize

	filterResult, err := ReadSimpleField[ExtensionObject](ctx, "filterResult", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'filterResult' field"))
	}
	m.FilterResult = filterResult

	if closeErr := readBuffer.CloseContext("MonitoredItemCreateResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemCreateResult")
	}

	return m, nil
}

func (m *_MonitoredItemCreateResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemCreateResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemCreateResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemCreateResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[uint32](ctx, "monitoredItemId", m.GetMonitoredItemId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItemId' field")
		}

		if err := WriteSimpleField[float64](ctx, "revisedSamplingInterval", m.GetRevisedSamplingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisedSamplingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "revisedQueueSize", m.GetRevisedQueueSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisedQueueSize' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "filterResult", m.GetFilterResult(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'filterResult' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemCreateResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemCreateResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemCreateResult) IsMonitoredItemCreateResult() {}

func (m *_MonitoredItemCreateResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MonitoredItemCreateResult) deepCopy() *_MonitoredItemCreateResult {
	if m == nil {
		return nil
	}
	_MonitoredItemCreateResultCopy := &_MonitoredItemCreateResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.MonitoredItemId,
		m.RevisedSamplingInterval,
		m.RevisedQueueSize,
		m.FilterResult.DeepCopy().(ExtensionObject),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _MonitoredItemCreateResultCopy
}

func (m *_MonitoredItemCreateResult) String() string {
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
