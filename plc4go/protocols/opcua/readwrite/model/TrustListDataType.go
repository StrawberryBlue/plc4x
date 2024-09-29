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

// TrustListDataType is the corresponding interface of TrustListDataType
type TrustListDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSpecifiedLists returns SpecifiedLists (property field)
	GetSpecifiedLists() uint32
	// GetNoOfTrustedCertificates returns NoOfTrustedCertificates (property field)
	GetNoOfTrustedCertificates() int32
	// GetTrustedCertificates returns TrustedCertificates (property field)
	GetTrustedCertificates() []PascalByteString
	// GetNoOfTrustedCrls returns NoOfTrustedCrls (property field)
	GetNoOfTrustedCrls() int32
	// GetTrustedCrls returns TrustedCrls (property field)
	GetTrustedCrls() []PascalByteString
	// GetNoOfIssuerCertificates returns NoOfIssuerCertificates (property field)
	GetNoOfIssuerCertificates() int32
	// GetIssuerCertificates returns IssuerCertificates (property field)
	GetIssuerCertificates() []PascalByteString
	// GetNoOfIssuerCrls returns NoOfIssuerCrls (property field)
	GetNoOfIssuerCrls() int32
	// GetIssuerCrls returns IssuerCrls (property field)
	GetIssuerCrls() []PascalByteString
	// IsTrustListDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTrustListDataType()
	// CreateBuilder creates a TrustListDataTypeBuilder
	CreateTrustListDataTypeBuilder() TrustListDataTypeBuilder
}

// _TrustListDataType is the data-structure of this message
type _TrustListDataType struct {
	ExtensionObjectDefinitionContract
	SpecifiedLists          uint32
	NoOfTrustedCertificates int32
	TrustedCertificates     []PascalByteString
	NoOfTrustedCrls         int32
	TrustedCrls             []PascalByteString
	NoOfIssuerCertificates  int32
	IssuerCertificates      []PascalByteString
	NoOfIssuerCrls          int32
	IssuerCrls              []PascalByteString
}

var _ TrustListDataType = (*_TrustListDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TrustListDataType)(nil)

// NewTrustListDataType factory function for _TrustListDataType
func NewTrustListDataType(specifiedLists uint32, noOfTrustedCertificates int32, trustedCertificates []PascalByteString, noOfTrustedCrls int32, trustedCrls []PascalByteString, noOfIssuerCertificates int32, issuerCertificates []PascalByteString, noOfIssuerCrls int32, issuerCrls []PascalByteString) *_TrustListDataType {
	_result := &_TrustListDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SpecifiedLists:                    specifiedLists,
		NoOfTrustedCertificates:           noOfTrustedCertificates,
		TrustedCertificates:               trustedCertificates,
		NoOfTrustedCrls:                   noOfTrustedCrls,
		TrustedCrls:                       trustedCrls,
		NoOfIssuerCertificates:            noOfIssuerCertificates,
		IssuerCertificates:                issuerCertificates,
		NoOfIssuerCrls:                    noOfIssuerCrls,
		IssuerCrls:                        issuerCrls,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TrustListDataTypeBuilder is a builder for TrustListDataType
type TrustListDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(specifiedLists uint32, noOfTrustedCertificates int32, trustedCertificates []PascalByteString, noOfTrustedCrls int32, trustedCrls []PascalByteString, noOfIssuerCertificates int32, issuerCertificates []PascalByteString, noOfIssuerCrls int32, issuerCrls []PascalByteString) TrustListDataTypeBuilder
	// WithSpecifiedLists adds SpecifiedLists (property field)
	WithSpecifiedLists(uint32) TrustListDataTypeBuilder
	// WithNoOfTrustedCertificates adds NoOfTrustedCertificates (property field)
	WithNoOfTrustedCertificates(int32) TrustListDataTypeBuilder
	// WithTrustedCertificates adds TrustedCertificates (property field)
	WithTrustedCertificates(...PascalByteString) TrustListDataTypeBuilder
	// WithNoOfTrustedCrls adds NoOfTrustedCrls (property field)
	WithNoOfTrustedCrls(int32) TrustListDataTypeBuilder
	// WithTrustedCrls adds TrustedCrls (property field)
	WithTrustedCrls(...PascalByteString) TrustListDataTypeBuilder
	// WithNoOfIssuerCertificates adds NoOfIssuerCertificates (property field)
	WithNoOfIssuerCertificates(int32) TrustListDataTypeBuilder
	// WithIssuerCertificates adds IssuerCertificates (property field)
	WithIssuerCertificates(...PascalByteString) TrustListDataTypeBuilder
	// WithNoOfIssuerCrls adds NoOfIssuerCrls (property field)
	WithNoOfIssuerCrls(int32) TrustListDataTypeBuilder
	// WithIssuerCrls adds IssuerCrls (property field)
	WithIssuerCrls(...PascalByteString) TrustListDataTypeBuilder
	// Build builds the TrustListDataType or returns an error if something is wrong
	Build() (TrustListDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TrustListDataType
}

// NewTrustListDataTypeBuilder() creates a TrustListDataTypeBuilder
func NewTrustListDataTypeBuilder() TrustListDataTypeBuilder {
	return &_TrustListDataTypeBuilder{_TrustListDataType: new(_TrustListDataType)}
}

type _TrustListDataTypeBuilder struct {
	*_TrustListDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TrustListDataTypeBuilder) = (*_TrustListDataTypeBuilder)(nil)

func (b *_TrustListDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_TrustListDataTypeBuilder) WithMandatoryFields(specifiedLists uint32, noOfTrustedCertificates int32, trustedCertificates []PascalByteString, noOfTrustedCrls int32, trustedCrls []PascalByteString, noOfIssuerCertificates int32, issuerCertificates []PascalByteString, noOfIssuerCrls int32, issuerCrls []PascalByteString) TrustListDataTypeBuilder {
	return b.WithSpecifiedLists(specifiedLists).WithNoOfTrustedCertificates(noOfTrustedCertificates).WithTrustedCertificates(trustedCertificates...).WithNoOfTrustedCrls(noOfTrustedCrls).WithTrustedCrls(trustedCrls...).WithNoOfIssuerCertificates(noOfIssuerCertificates).WithIssuerCertificates(issuerCertificates...).WithNoOfIssuerCrls(noOfIssuerCrls).WithIssuerCrls(issuerCrls...)
}

func (b *_TrustListDataTypeBuilder) WithSpecifiedLists(specifiedLists uint32) TrustListDataTypeBuilder {
	b.SpecifiedLists = specifiedLists
	return b
}

func (b *_TrustListDataTypeBuilder) WithNoOfTrustedCertificates(noOfTrustedCertificates int32) TrustListDataTypeBuilder {
	b.NoOfTrustedCertificates = noOfTrustedCertificates
	return b
}

func (b *_TrustListDataTypeBuilder) WithTrustedCertificates(trustedCertificates ...PascalByteString) TrustListDataTypeBuilder {
	b.TrustedCertificates = trustedCertificates
	return b
}

func (b *_TrustListDataTypeBuilder) WithNoOfTrustedCrls(noOfTrustedCrls int32) TrustListDataTypeBuilder {
	b.NoOfTrustedCrls = noOfTrustedCrls
	return b
}

func (b *_TrustListDataTypeBuilder) WithTrustedCrls(trustedCrls ...PascalByteString) TrustListDataTypeBuilder {
	b.TrustedCrls = trustedCrls
	return b
}

func (b *_TrustListDataTypeBuilder) WithNoOfIssuerCertificates(noOfIssuerCertificates int32) TrustListDataTypeBuilder {
	b.NoOfIssuerCertificates = noOfIssuerCertificates
	return b
}

func (b *_TrustListDataTypeBuilder) WithIssuerCertificates(issuerCertificates ...PascalByteString) TrustListDataTypeBuilder {
	b.IssuerCertificates = issuerCertificates
	return b
}

func (b *_TrustListDataTypeBuilder) WithNoOfIssuerCrls(noOfIssuerCrls int32) TrustListDataTypeBuilder {
	b.NoOfIssuerCrls = noOfIssuerCrls
	return b
}

func (b *_TrustListDataTypeBuilder) WithIssuerCrls(issuerCrls ...PascalByteString) TrustListDataTypeBuilder {
	b.IssuerCrls = issuerCrls
	return b
}

func (b *_TrustListDataTypeBuilder) Build() (TrustListDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TrustListDataType.deepCopy(), nil
}

func (b *_TrustListDataTypeBuilder) MustBuild() TrustListDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TrustListDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_TrustListDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TrustListDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateTrustListDataTypeBuilder().(*_TrustListDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTrustListDataTypeBuilder creates a TrustListDataTypeBuilder
func (b *_TrustListDataType) CreateTrustListDataTypeBuilder() TrustListDataTypeBuilder {
	if b == nil {
		return NewTrustListDataTypeBuilder()
	}
	return &_TrustListDataTypeBuilder{_TrustListDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TrustListDataType) GetIdentifier() string {
	return "12556"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TrustListDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TrustListDataType) GetSpecifiedLists() uint32 {
	return m.SpecifiedLists
}

func (m *_TrustListDataType) GetNoOfTrustedCertificates() int32 {
	return m.NoOfTrustedCertificates
}

func (m *_TrustListDataType) GetTrustedCertificates() []PascalByteString {
	return m.TrustedCertificates
}

func (m *_TrustListDataType) GetNoOfTrustedCrls() int32 {
	return m.NoOfTrustedCrls
}

func (m *_TrustListDataType) GetTrustedCrls() []PascalByteString {
	return m.TrustedCrls
}

func (m *_TrustListDataType) GetNoOfIssuerCertificates() int32 {
	return m.NoOfIssuerCertificates
}

func (m *_TrustListDataType) GetIssuerCertificates() []PascalByteString {
	return m.IssuerCertificates
}

func (m *_TrustListDataType) GetNoOfIssuerCrls() int32 {
	return m.NoOfIssuerCrls
}

func (m *_TrustListDataType) GetIssuerCrls() []PascalByteString {
	return m.IssuerCrls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTrustListDataType(structType any) TrustListDataType {
	if casted, ok := structType.(TrustListDataType); ok {
		return casted
	}
	if casted, ok := structType.(*TrustListDataType); ok {
		return *casted
	}
	return nil
}

func (m *_TrustListDataType) GetTypeName() string {
	return "TrustListDataType"
}

func (m *_TrustListDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (specifiedLists)
	lengthInBits += 32

	// Simple field (noOfTrustedCertificates)
	lengthInBits += 32

	// Array field
	if len(m.TrustedCertificates) > 0 {
		for _curItem, element := range m.TrustedCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TrustedCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfTrustedCrls)
	lengthInBits += 32

	// Array field
	if len(m.TrustedCrls) > 0 {
		for _curItem, element := range m.TrustedCrls {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TrustedCrls), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfIssuerCertificates)
	lengthInBits += 32

	// Array field
	if len(m.IssuerCertificates) > 0 {
		for _curItem, element := range m.IssuerCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.IssuerCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfIssuerCrls)
	lengthInBits += 32

	// Array field
	if len(m.IssuerCrls) > 0 {
		for _curItem, element := range m.IssuerCrls {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.IssuerCrls), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_TrustListDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TrustListDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__trustListDataType TrustListDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TrustListDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TrustListDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedLists, err := ReadSimpleField(ctx, "specifiedLists", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedLists' field"))
	}
	m.SpecifiedLists = specifiedLists

	noOfTrustedCertificates, err := ReadSimpleField(ctx, "noOfTrustedCertificates", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfTrustedCertificates' field"))
	}
	m.NoOfTrustedCertificates = noOfTrustedCertificates

	trustedCertificates, err := ReadCountArrayField[PascalByteString](ctx, "trustedCertificates", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer), uint64(noOfTrustedCertificates))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'trustedCertificates' field"))
	}
	m.TrustedCertificates = trustedCertificates

	noOfTrustedCrls, err := ReadSimpleField(ctx, "noOfTrustedCrls", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfTrustedCrls' field"))
	}
	m.NoOfTrustedCrls = noOfTrustedCrls

	trustedCrls, err := ReadCountArrayField[PascalByteString](ctx, "trustedCrls", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer), uint64(noOfTrustedCrls))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'trustedCrls' field"))
	}
	m.TrustedCrls = trustedCrls

	noOfIssuerCertificates, err := ReadSimpleField(ctx, "noOfIssuerCertificates", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfIssuerCertificates' field"))
	}
	m.NoOfIssuerCertificates = noOfIssuerCertificates

	issuerCertificates, err := ReadCountArrayField[PascalByteString](ctx, "issuerCertificates", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer), uint64(noOfIssuerCertificates))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issuerCertificates' field"))
	}
	m.IssuerCertificates = issuerCertificates

	noOfIssuerCrls, err := ReadSimpleField(ctx, "noOfIssuerCrls", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfIssuerCrls' field"))
	}
	m.NoOfIssuerCrls = noOfIssuerCrls

	issuerCrls, err := ReadCountArrayField[PascalByteString](ctx, "issuerCrls", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer), uint64(noOfIssuerCrls))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issuerCrls' field"))
	}
	m.IssuerCrls = issuerCrls

	if closeErr := readBuffer.CloseContext("TrustListDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TrustListDataType")
	}

	return m, nil
}

func (m *_TrustListDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TrustListDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TrustListDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TrustListDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "specifiedLists", m.GetSpecifiedLists(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'specifiedLists' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfTrustedCertificates", m.GetNoOfTrustedCertificates(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfTrustedCertificates' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "trustedCertificates", m.GetTrustedCertificates(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'trustedCertificates' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfTrustedCrls", m.GetNoOfTrustedCrls(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfTrustedCrls' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "trustedCrls", m.GetTrustedCrls(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'trustedCrls' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfIssuerCertificates", m.GetNoOfIssuerCertificates(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfIssuerCertificates' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "issuerCertificates", m.GetIssuerCertificates(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'issuerCertificates' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfIssuerCrls", m.GetNoOfIssuerCrls(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfIssuerCrls' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "issuerCrls", m.GetIssuerCrls(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'issuerCrls' field")
		}

		if popErr := writeBuffer.PopContext("TrustListDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TrustListDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TrustListDataType) IsTrustListDataType() {}

func (m *_TrustListDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TrustListDataType) deepCopy() *_TrustListDataType {
	if m == nil {
		return nil
	}
	_TrustListDataTypeCopy := &_TrustListDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SpecifiedLists,
		m.NoOfTrustedCertificates,
		utils.DeepCopySlice[PascalByteString, PascalByteString](m.TrustedCertificates),
		m.NoOfTrustedCrls,
		utils.DeepCopySlice[PascalByteString, PascalByteString](m.TrustedCrls),
		m.NoOfIssuerCertificates,
		utils.DeepCopySlice[PascalByteString, PascalByteString](m.IssuerCertificates),
		m.NoOfIssuerCrls,
		utils.DeepCopySlice[PascalByteString, PascalByteString](m.IssuerCrls),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TrustListDataTypeCopy
}

func (m *_TrustListDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
