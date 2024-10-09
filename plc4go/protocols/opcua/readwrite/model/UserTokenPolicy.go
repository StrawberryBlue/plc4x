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

// UserTokenPolicy is the corresponding interface of UserTokenPolicy
type UserTokenPolicy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// GetTokenType returns TokenType (property field)
	GetTokenType() UserTokenType
	// GetIssuedTokenType returns IssuedTokenType (property field)
	GetIssuedTokenType() PascalString
	// GetIssuerEndpointUrl returns IssuerEndpointUrl (property field)
	GetIssuerEndpointUrl() PascalString
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// IsUserTokenPolicy is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUserTokenPolicy()
	// CreateBuilder creates a UserTokenPolicyBuilder
	CreateUserTokenPolicyBuilder() UserTokenPolicyBuilder
}

// _UserTokenPolicy is the data-structure of this message
type _UserTokenPolicy struct {
	ExtensionObjectDefinitionContract
	PolicyId          PascalString
	TokenType         UserTokenType
	IssuedTokenType   PascalString
	IssuerEndpointUrl PascalString
	SecurityPolicyUri PascalString
}

var _ UserTokenPolicy = (*_UserTokenPolicy)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UserTokenPolicy)(nil)

// NewUserTokenPolicy factory function for _UserTokenPolicy
func NewUserTokenPolicy(policyId PascalString, tokenType UserTokenType, issuedTokenType PascalString, issuerEndpointUrl PascalString, securityPolicyUri PascalString) *_UserTokenPolicy {
	if policyId == nil {
		panic("policyId of type PascalString for UserTokenPolicy must not be nil")
	}
	if issuedTokenType == nil {
		panic("issuedTokenType of type PascalString for UserTokenPolicy must not be nil")
	}
	if issuerEndpointUrl == nil {
		panic("issuerEndpointUrl of type PascalString for UserTokenPolicy must not be nil")
	}
	if securityPolicyUri == nil {
		panic("securityPolicyUri of type PascalString for UserTokenPolicy must not be nil")
	}
	_result := &_UserTokenPolicy{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PolicyId:                          policyId,
		TokenType:                         tokenType,
		IssuedTokenType:                   issuedTokenType,
		IssuerEndpointUrl:                 issuerEndpointUrl,
		SecurityPolicyUri:                 securityPolicyUri,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UserTokenPolicyBuilder is a builder for UserTokenPolicy
type UserTokenPolicyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(policyId PascalString, tokenType UserTokenType, issuedTokenType PascalString, issuerEndpointUrl PascalString, securityPolicyUri PascalString) UserTokenPolicyBuilder
	// WithPolicyId adds PolicyId (property field)
	WithPolicyId(PascalString) UserTokenPolicyBuilder
	// WithPolicyIdBuilder adds PolicyId (property field) which is build by the builder
	WithPolicyIdBuilder(func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder
	// WithTokenType adds TokenType (property field)
	WithTokenType(UserTokenType) UserTokenPolicyBuilder
	// WithIssuedTokenType adds IssuedTokenType (property field)
	WithIssuedTokenType(PascalString) UserTokenPolicyBuilder
	// WithIssuedTokenTypeBuilder adds IssuedTokenType (property field) which is build by the builder
	WithIssuedTokenTypeBuilder(func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder
	// WithIssuerEndpointUrl adds IssuerEndpointUrl (property field)
	WithIssuerEndpointUrl(PascalString) UserTokenPolicyBuilder
	// WithIssuerEndpointUrlBuilder adds IssuerEndpointUrl (property field) which is build by the builder
	WithIssuerEndpointUrlBuilder(func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder
	// WithSecurityPolicyUri adds SecurityPolicyUri (property field)
	WithSecurityPolicyUri(PascalString) UserTokenPolicyBuilder
	// WithSecurityPolicyUriBuilder adds SecurityPolicyUri (property field) which is build by the builder
	WithSecurityPolicyUriBuilder(func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder
	// Build builds the UserTokenPolicy or returns an error if something is wrong
	Build() (UserTokenPolicy, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UserTokenPolicy
}

// NewUserTokenPolicyBuilder() creates a UserTokenPolicyBuilder
func NewUserTokenPolicyBuilder() UserTokenPolicyBuilder {
	return &_UserTokenPolicyBuilder{_UserTokenPolicy: new(_UserTokenPolicy)}
}

type _UserTokenPolicyBuilder struct {
	*_UserTokenPolicy

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UserTokenPolicyBuilder) = (*_UserTokenPolicyBuilder)(nil)

func (b *_UserTokenPolicyBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_UserTokenPolicyBuilder) WithMandatoryFields(policyId PascalString, tokenType UserTokenType, issuedTokenType PascalString, issuerEndpointUrl PascalString, securityPolicyUri PascalString) UserTokenPolicyBuilder {
	return b.WithPolicyId(policyId).WithTokenType(tokenType).WithIssuedTokenType(issuedTokenType).WithIssuerEndpointUrl(issuerEndpointUrl).WithSecurityPolicyUri(securityPolicyUri)
}

func (b *_UserTokenPolicyBuilder) WithPolicyId(policyId PascalString) UserTokenPolicyBuilder {
	b.PolicyId = policyId
	return b
}

func (b *_UserTokenPolicyBuilder) WithPolicyIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder {
	builder := builderSupplier(b.PolicyId.CreatePascalStringBuilder())
	var err error
	b.PolicyId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserTokenPolicyBuilder) WithTokenType(tokenType UserTokenType) UserTokenPolicyBuilder {
	b.TokenType = tokenType
	return b
}

func (b *_UserTokenPolicyBuilder) WithIssuedTokenType(issuedTokenType PascalString) UserTokenPolicyBuilder {
	b.IssuedTokenType = issuedTokenType
	return b
}

func (b *_UserTokenPolicyBuilder) WithIssuedTokenTypeBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder {
	builder := builderSupplier(b.IssuedTokenType.CreatePascalStringBuilder())
	var err error
	b.IssuedTokenType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserTokenPolicyBuilder) WithIssuerEndpointUrl(issuerEndpointUrl PascalString) UserTokenPolicyBuilder {
	b.IssuerEndpointUrl = issuerEndpointUrl
	return b
}

func (b *_UserTokenPolicyBuilder) WithIssuerEndpointUrlBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder {
	builder := builderSupplier(b.IssuerEndpointUrl.CreatePascalStringBuilder())
	var err error
	b.IssuerEndpointUrl, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserTokenPolicyBuilder) WithSecurityPolicyUri(securityPolicyUri PascalString) UserTokenPolicyBuilder {
	b.SecurityPolicyUri = securityPolicyUri
	return b
}

func (b *_UserTokenPolicyBuilder) WithSecurityPolicyUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserTokenPolicyBuilder {
	builder := builderSupplier(b.SecurityPolicyUri.CreatePascalStringBuilder())
	var err error
	b.SecurityPolicyUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserTokenPolicyBuilder) Build() (UserTokenPolicy, error) {
	if b.PolicyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'policyId' not set"))
	}
	if b.IssuedTokenType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'issuedTokenType' not set"))
	}
	if b.IssuerEndpointUrl == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'issuerEndpointUrl' not set"))
	}
	if b.SecurityPolicyUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityPolicyUri' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UserTokenPolicy.deepCopy(), nil
}

func (b *_UserTokenPolicyBuilder) MustBuild() UserTokenPolicy {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_UserTokenPolicyBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_UserTokenPolicyBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UserTokenPolicyBuilder) DeepCopy() any {
	_copy := b.CreateUserTokenPolicyBuilder().(*_UserTokenPolicyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUserTokenPolicyBuilder creates a UserTokenPolicyBuilder
func (b *_UserTokenPolicy) CreateUserTokenPolicyBuilder() UserTokenPolicyBuilder {
	if b == nil {
		return NewUserTokenPolicyBuilder()
	}
	return &_UserTokenPolicyBuilder{_UserTokenPolicy: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserTokenPolicy) GetIdentifier() string {
	return "306"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserTokenPolicy) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserTokenPolicy) GetPolicyId() PascalString {
	return m.PolicyId
}

func (m *_UserTokenPolicy) GetTokenType() UserTokenType {
	return m.TokenType
}

func (m *_UserTokenPolicy) GetIssuedTokenType() PascalString {
	return m.IssuedTokenType
}

func (m *_UserTokenPolicy) GetIssuerEndpointUrl() PascalString {
	return m.IssuerEndpointUrl
}

func (m *_UserTokenPolicy) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUserTokenPolicy(structType any) UserTokenPolicy {
	if casted, ok := structType.(UserTokenPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*UserTokenPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_UserTokenPolicy) GetTypeName() string {
	return "UserTokenPolicy"
}

func (m *_UserTokenPolicy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	// Simple field (tokenType)
	lengthInBits += 32

	// Simple field (issuedTokenType)
	lengthInBits += m.IssuedTokenType.GetLengthInBits(ctx)

	// Simple field (issuerEndpointUrl)
	lengthInBits += m.IssuerEndpointUrl.GetLengthInBits(ctx)

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UserTokenPolicy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UserTokenPolicy) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__userTokenPolicy UserTokenPolicy, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserTokenPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserTokenPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}
	m.PolicyId = policyId

	tokenType, err := ReadEnumField[UserTokenType](ctx, "tokenType", "UserTokenType", ReadEnum(UserTokenTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tokenType' field"))
	}
	m.TokenType = tokenType

	issuedTokenType, err := ReadSimpleField[PascalString](ctx, "issuedTokenType", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issuedTokenType' field"))
	}
	m.IssuedTokenType = issuedTokenType

	issuerEndpointUrl, err := ReadSimpleField[PascalString](ctx, "issuerEndpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issuerEndpointUrl' field"))
	}
	m.IssuerEndpointUrl = issuerEndpointUrl

	securityPolicyUri, err := ReadSimpleField[PascalString](ctx, "securityPolicyUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityPolicyUri' field"))
	}
	m.SecurityPolicyUri = securityPolicyUri

	if closeErr := readBuffer.CloseContext("UserTokenPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserTokenPolicy")
	}

	return m, nil
}

func (m *_UserTokenPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UserTokenPolicy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UserTokenPolicy"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UserTokenPolicy")
		}

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
		}

		if err := WriteSimpleEnumField[UserTokenType](ctx, "tokenType", "UserTokenType", m.GetTokenType(), WriteEnum[UserTokenType, uint32](UserTokenType.GetValue, UserTokenType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'tokenType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "issuedTokenType", m.GetIssuedTokenType(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'issuedTokenType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "issuerEndpointUrl", m.GetIssuerEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'issuerEndpointUrl' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityPolicyUri", m.GetSecurityPolicyUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityPolicyUri' field")
		}

		if popErr := writeBuffer.PopContext("UserTokenPolicy"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UserTokenPolicy")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserTokenPolicy) IsUserTokenPolicy() {}

func (m *_UserTokenPolicy) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UserTokenPolicy) deepCopy() *_UserTokenPolicy {
	if m == nil {
		return nil
	}
	_UserTokenPolicyCopy := &_UserTokenPolicy{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.PolicyId.DeepCopy().(PascalString),
		m.TokenType,
		m.IssuedTokenType.DeepCopy().(PascalString),
		m.IssuerEndpointUrl.DeepCopy().(PascalString),
		m.SecurityPolicyUri.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UserTokenPolicyCopy
}

func (m *_UserTokenPolicy) String() string {
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
