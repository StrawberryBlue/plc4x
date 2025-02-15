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

// BACnetConfirmedServiceRequestAtomicReadFile is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFile
type BACnetConfirmedServiceRequestAtomicReadFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetFileIdentifier returns FileIdentifier (property field)
	GetFileIdentifier() BACnetApplicationTagObjectIdentifier
	// GetAccessMethod returns AccessMethod (property field)
	GetAccessMethod() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// IsBACnetConfirmedServiceRequestAtomicReadFile is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicReadFile()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileBuilder
	CreateBACnetConfirmedServiceRequestAtomicReadFileBuilder() BACnetConfirmedServiceRequestAtomicReadFileBuilder
}

// _BACnetConfirmedServiceRequestAtomicReadFile is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFile struct {
	BACnetConfirmedServiceRequestContract
	FileIdentifier BACnetApplicationTagObjectIdentifier
	AccessMethod   BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
}

var _ BACnetConfirmedServiceRequestAtomicReadFile = (*_BACnetConfirmedServiceRequestAtomicReadFile)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestAtomicReadFile)(nil)

// NewBACnetConfirmedServiceRequestAtomicReadFile factory function for _BACnetConfirmedServiceRequestAtomicReadFile
func NewBACnetConfirmedServiceRequestAtomicReadFile(fileIdentifier BACnetApplicationTagObjectIdentifier, accessMethod BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAtomicReadFile {
	if fileIdentifier == nil {
		panic("fileIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetConfirmedServiceRequestAtomicReadFile must not be nil")
	}
	if accessMethod == nil {
		panic("accessMethod of type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord for BACnetConfirmedServiceRequestAtomicReadFile must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAtomicReadFile{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		FileIdentifier:                        fileIdentifier,
		AccessMethod:                          accessMethod,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAtomicReadFileBuilder is a builder for BACnetConfirmedServiceRequestAtomicReadFile
type BACnetConfirmedServiceRequestAtomicReadFileBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fileIdentifier BACnetApplicationTagObjectIdentifier, accessMethod BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) BACnetConfirmedServiceRequestAtomicReadFileBuilder
	// WithFileIdentifier adds FileIdentifier (property field)
	WithFileIdentifier(BACnetApplicationTagObjectIdentifier) BACnetConfirmedServiceRequestAtomicReadFileBuilder
	// WithFileIdentifierBuilder adds FileIdentifier (property field) which is build by the builder
	WithFileIdentifierBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestAtomicReadFileBuilder
	// WithAccessMethod adds AccessMethod (property field)
	WithAccessMethod(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) BACnetConfirmedServiceRequestAtomicReadFileBuilder
	// WithAccessMethodBuilder adds AccessMethod (property field) which is build by the builder
	WithAccessMethodBuilder(func(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) BACnetConfirmedServiceRequestAtomicReadFileBuilder
	// Build builds the BACnetConfirmedServiceRequestAtomicReadFile or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAtomicReadFile, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAtomicReadFile
}

// NewBACnetConfirmedServiceRequestAtomicReadFileBuilder() creates a BACnetConfirmedServiceRequestAtomicReadFileBuilder
func NewBACnetConfirmedServiceRequestAtomicReadFileBuilder() BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	return &_BACnetConfirmedServiceRequestAtomicReadFileBuilder{_BACnetConfirmedServiceRequestAtomicReadFile: new(_BACnetConfirmedServiceRequestAtomicReadFile)}
}

type _BACnetConfirmedServiceRequestAtomicReadFileBuilder struct {
	*_BACnetConfirmedServiceRequestAtomicReadFile

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAtomicReadFileBuilder) = (*_BACnetConfirmedServiceRequestAtomicReadFileBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) WithMandatoryFields(fileIdentifier BACnetApplicationTagObjectIdentifier, accessMethod BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	return b.WithFileIdentifier(fileIdentifier).WithAccessMethod(accessMethod)
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) WithFileIdentifier(fileIdentifier BACnetApplicationTagObjectIdentifier) BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	b.FileIdentifier = fileIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) WithFileIdentifierBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	builder := builderSupplier(b.FileIdentifier.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.FileIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) WithAccessMethod(accessMethod BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	b.AccessMethod = accessMethod
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) WithAccessMethodBuilder(builderSupplier func(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder) BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	builder := builderSupplier(b.AccessMethod.CreateBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder())
	var err error
	b.AccessMethod, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) Build() (BACnetConfirmedServiceRequestAtomicReadFile, error) {
	if b.FileIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fileIdentifier' not set"))
	}
	if b.AccessMethod == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'accessMethod' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestAtomicReadFile.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) MustBuild() BACnetConfirmedServiceRequestAtomicReadFile {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestAtomicReadFileBuilder().(*_BACnetConfirmedServiceRequestAtomicReadFileBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestAtomicReadFileBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileBuilder
func (b *_BACnetConfirmedServiceRequestAtomicReadFile) CreateBACnetConfirmedServiceRequestAtomicReadFileBuilder() BACnetConfirmedServiceRequestAtomicReadFileBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestAtomicReadFileBuilder()
	}
	return &_BACnetConfirmedServiceRequestAtomicReadFileBuilder{_BACnetConfirmedServiceRequestAtomicReadFile: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetFileIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.FileIdentifier
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetAccessMethod() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFile(structType any) BACnetConfirmedServiceRequestAtomicReadFile {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFile"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (fileIdentifier)
	lengthInBits += m.FileIdentifier.GetLengthInBits(ctx)

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAtomicReadFile BACnetConfirmedServiceRequestAtomicReadFile, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "fileIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileIdentifier' field"))
	}
	m.FileIdentifier = fileIdentifier

	accessMethod, err := ReadSimpleField[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](ctx, "accessMethod", ReadComplex[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessMethod' field"))
	}
	m.AccessMethod = accessMethod

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFile")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFile")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "fileIdentifier", m.GetFileIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileIdentifier' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](ctx, "accessMethod", m.GetAccessMethod(), WriteComplex[BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFile")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) IsBACnetConfirmedServiceRequestAtomicReadFile() {
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) deepCopy() *_BACnetConfirmedServiceRequestAtomicReadFile {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAtomicReadFileCopy := &_BACnetConfirmedServiceRequestAtomicReadFile{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.FileIdentifier.DeepCopy().(BACnetApplicationTagObjectIdentifier),
		m.AccessMethod.DeepCopy().(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestAtomicReadFileCopy
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) String() string {
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
