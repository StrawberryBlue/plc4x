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

// CEMIAdditionalInformation is the corresponding interface of CEMIAdditionalInformation
type CEMIAdditionalInformation interface {
	CEMIAdditionalInformationContract
	CEMIAdditionalInformationRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsCEMIAdditionalInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMIAdditionalInformation()
}

// CEMIAdditionalInformationContract provides a set of functions which can be overwritten by a sub struct
type CEMIAdditionalInformationContract interface {
	// IsCEMIAdditionalInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMIAdditionalInformation()
}

// CEMIAdditionalInformationRequirements provides a set of functions which need to be implemented by a sub struct
type CEMIAdditionalInformationRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetAdditionalInformationType returns AdditionalInformationType (discriminator field)
	GetAdditionalInformationType() uint8
}

// _CEMIAdditionalInformation is the data-structure of this message
type _CEMIAdditionalInformation struct {
	_SubType CEMIAdditionalInformation
}

var _ CEMIAdditionalInformationContract = (*_CEMIAdditionalInformation)(nil)

// NewCEMIAdditionalInformation factory function for _CEMIAdditionalInformation
func NewCEMIAdditionalInformation() *_CEMIAdditionalInformation {
	return &_CEMIAdditionalInformation{}
}

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformation(structType any) CEMIAdditionalInformation {
	if casted, ok := structType.(CEMIAdditionalInformation); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformation); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformation) GetTypeName() string {
	return "CEMIAdditionalInformation"
}

func (m *_CEMIAdditionalInformation) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (additionalInformationType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CEMIAdditionalInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CEMIAdditionalInformationParse[T CEMIAdditionalInformation](ctx context.Context, theBytes []byte) (T, error) {
	return CEMIAdditionalInformationParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func CEMIAdditionalInformationParseWithBufferProducer[T CEMIAdditionalInformation]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CEMIAdditionalInformationParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CEMIAdditionalInformationParseWithBuffer[T CEMIAdditionalInformation](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_CEMIAdditionalInformation{}).parse(ctx, readBuffer)
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

func (m *_CEMIAdditionalInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__cEMIAdditionalInformation CEMIAdditionalInformation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	additionalInformationType, err := ReadDiscriminatorField[uint8](ctx, "additionalInformationType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformationType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CEMIAdditionalInformation
	switch {
	case additionalInformationType == 0x03: // CEMIAdditionalInformationBusmonitorInfo
		if _child, err = (&_CEMIAdditionalInformationBusmonitorInfo{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CEMIAdditionalInformationBusmonitorInfo for type-switch of CEMIAdditionalInformation")
		}
	case additionalInformationType == 0x04: // CEMIAdditionalInformationRelativeTimestamp
		if _child, err = (&_CEMIAdditionalInformationRelativeTimestamp{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CEMIAdditionalInformationRelativeTimestamp for type-switch of CEMIAdditionalInformation")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [additionalInformationType=%v]", additionalInformationType)
	}

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformation")
	}

	return _child, nil
}

func (pm *_CEMIAdditionalInformation) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CEMIAdditionalInformation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CEMIAdditionalInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformation")
	}

	if err := WriteDiscriminatorField(ctx, "additionalInformationType", m.GetAdditionalInformationType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'additionalInformationType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMIAdditionalInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformation")
	}
	return nil
}

func (m *_CEMIAdditionalInformation) IsCEMIAdditionalInformation() {}
