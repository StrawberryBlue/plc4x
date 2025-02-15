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

// ProjectInstallationIdentifier is the corresponding interface of ProjectInstallationIdentifier
type ProjectInstallationIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetProjectNumber returns ProjectNumber (property field)
	GetProjectNumber() uint8
	// GetInstallationNumber returns InstallationNumber (property field)
	GetInstallationNumber() uint8
	// IsProjectInstallationIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsProjectInstallationIdentifier()
	// CreateBuilder creates a ProjectInstallationIdentifierBuilder
	CreateProjectInstallationIdentifierBuilder() ProjectInstallationIdentifierBuilder
}

// _ProjectInstallationIdentifier is the data-structure of this message
type _ProjectInstallationIdentifier struct {
	ProjectNumber      uint8
	InstallationNumber uint8
}

var _ ProjectInstallationIdentifier = (*_ProjectInstallationIdentifier)(nil)

// NewProjectInstallationIdentifier factory function for _ProjectInstallationIdentifier
func NewProjectInstallationIdentifier(projectNumber uint8, installationNumber uint8) *_ProjectInstallationIdentifier {
	return &_ProjectInstallationIdentifier{ProjectNumber: projectNumber, InstallationNumber: installationNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ProjectInstallationIdentifierBuilder is a builder for ProjectInstallationIdentifier
type ProjectInstallationIdentifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(projectNumber uint8, installationNumber uint8) ProjectInstallationIdentifierBuilder
	// WithProjectNumber adds ProjectNumber (property field)
	WithProjectNumber(uint8) ProjectInstallationIdentifierBuilder
	// WithInstallationNumber adds InstallationNumber (property field)
	WithInstallationNumber(uint8) ProjectInstallationIdentifierBuilder
	// Build builds the ProjectInstallationIdentifier or returns an error if something is wrong
	Build() (ProjectInstallationIdentifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ProjectInstallationIdentifier
}

// NewProjectInstallationIdentifierBuilder() creates a ProjectInstallationIdentifierBuilder
func NewProjectInstallationIdentifierBuilder() ProjectInstallationIdentifierBuilder {
	return &_ProjectInstallationIdentifierBuilder{_ProjectInstallationIdentifier: new(_ProjectInstallationIdentifier)}
}

type _ProjectInstallationIdentifierBuilder struct {
	*_ProjectInstallationIdentifier

	err *utils.MultiError
}

var _ (ProjectInstallationIdentifierBuilder) = (*_ProjectInstallationIdentifierBuilder)(nil)

func (b *_ProjectInstallationIdentifierBuilder) WithMandatoryFields(projectNumber uint8, installationNumber uint8) ProjectInstallationIdentifierBuilder {
	return b.WithProjectNumber(projectNumber).WithInstallationNumber(installationNumber)
}

func (b *_ProjectInstallationIdentifierBuilder) WithProjectNumber(projectNumber uint8) ProjectInstallationIdentifierBuilder {
	b.ProjectNumber = projectNumber
	return b
}

func (b *_ProjectInstallationIdentifierBuilder) WithInstallationNumber(installationNumber uint8) ProjectInstallationIdentifierBuilder {
	b.InstallationNumber = installationNumber
	return b
}

func (b *_ProjectInstallationIdentifierBuilder) Build() (ProjectInstallationIdentifier, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ProjectInstallationIdentifier.deepCopy(), nil
}

func (b *_ProjectInstallationIdentifierBuilder) MustBuild() ProjectInstallationIdentifier {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ProjectInstallationIdentifierBuilder) DeepCopy() any {
	_copy := b.CreateProjectInstallationIdentifierBuilder().(*_ProjectInstallationIdentifierBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateProjectInstallationIdentifierBuilder creates a ProjectInstallationIdentifierBuilder
func (b *_ProjectInstallationIdentifier) CreateProjectInstallationIdentifierBuilder() ProjectInstallationIdentifierBuilder {
	if b == nil {
		return NewProjectInstallationIdentifierBuilder()
	}
	return &_ProjectInstallationIdentifierBuilder{_ProjectInstallationIdentifier: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ProjectInstallationIdentifier) GetProjectNumber() uint8 {
	return m.ProjectNumber
}

func (m *_ProjectInstallationIdentifier) GetInstallationNumber() uint8 {
	return m.InstallationNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastProjectInstallationIdentifier(structType any) ProjectInstallationIdentifier {
	if casted, ok := structType.(ProjectInstallationIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*ProjectInstallationIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_ProjectInstallationIdentifier) GetTypeName() string {
	return "ProjectInstallationIdentifier"
}

func (m *_ProjectInstallationIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (projectNumber)
	lengthInBits += 8

	// Simple field (installationNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ProjectInstallationIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ProjectInstallationIdentifierParse(ctx context.Context, theBytes []byte) (ProjectInstallationIdentifier, error) {
	return ProjectInstallationIdentifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ProjectInstallationIdentifierParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ProjectInstallationIdentifier, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ProjectInstallationIdentifier, error) {
		return ProjectInstallationIdentifierParseWithBuffer(ctx, readBuffer)
	}
}

func ProjectInstallationIdentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ProjectInstallationIdentifier, error) {
	v, err := (&_ProjectInstallationIdentifier{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ProjectInstallationIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__projectInstallationIdentifier ProjectInstallationIdentifier, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ProjectInstallationIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ProjectInstallationIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	projectNumber, err := ReadSimpleField(ctx, "projectNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'projectNumber' field"))
	}
	m.ProjectNumber = projectNumber

	installationNumber, err := ReadSimpleField(ctx, "installationNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'installationNumber' field"))
	}
	m.InstallationNumber = installationNumber

	if closeErr := readBuffer.CloseContext("ProjectInstallationIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ProjectInstallationIdentifier")
	}

	return m, nil
}

func (m *_ProjectInstallationIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ProjectInstallationIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ProjectInstallationIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ProjectInstallationIdentifier")
	}

	if err := WriteSimpleField[uint8](ctx, "projectNumber", m.GetProjectNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'projectNumber' field")
	}

	if err := WriteSimpleField[uint8](ctx, "installationNumber", m.GetInstallationNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'installationNumber' field")
	}

	if popErr := writeBuffer.PopContext("ProjectInstallationIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ProjectInstallationIdentifier")
	}
	return nil
}

func (m *_ProjectInstallationIdentifier) IsProjectInstallationIdentifier() {}

func (m *_ProjectInstallationIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ProjectInstallationIdentifier) deepCopy() *_ProjectInstallationIdentifier {
	if m == nil {
		return nil
	}
	_ProjectInstallationIdentifierCopy := &_ProjectInstallationIdentifier{
		m.ProjectNumber,
		m.InstallationNumber,
	}
	return _ProjectInstallationIdentifierCopy
}

func (m *_ProjectInstallationIdentifier) String() string {
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
