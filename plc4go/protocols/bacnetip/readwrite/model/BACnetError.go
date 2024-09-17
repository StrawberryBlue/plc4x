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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetError is the corresponding interface of BACnetError
type BACnetError interface {
	BACnetErrorContract
	BACnetErrorRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetError()
}

// BACnetErrorContract provides a set of functions which can be overwritten by a sub struct
type BACnetErrorContract interface {
	// IsBACnetError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetError()
}

// BACnetErrorRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetErrorRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetErrorChoice returns ErrorChoice (discriminator field)
	GetErrorChoice() BACnetConfirmedServiceChoice
}

// _BACnetError is the data-structure of this message
type _BACnetError struct {
	_SubType BACnetError
}

var _ BACnetErrorContract = (*_BACnetError)(nil)

// NewBACnetError factory function for _BACnetError
func NewBACnetError() *_BACnetError {
	return &_BACnetError{}
}

// Deprecated: use the interface for direct cast
func CastBACnetError(structType any) BACnetError {
	if casted, ok := structType.(BACnetError); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetError); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetError) GetTypeName() string {
	return "BACnetError"
}

func (m *_BACnetError) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_BACnetError) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetErrorParse[T BACnetError](ctx context.Context, theBytes []byte, errorChoice BACnetConfirmedServiceChoice) (T, error) {
	return BACnetErrorParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), errorChoice)
}

func BACnetErrorParseWithBufferProducer[T BACnetError](errorChoice BACnetConfirmedServiceChoice) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetErrorParseWithBuffer[T](ctx, readBuffer, errorChoice)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetErrorParseWithBuffer[T BACnetError](ctx context.Context, readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (T, error) {
	v, err := (&_BACnetError{}).parse(ctx, readBuffer, errorChoice)
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

func (m *_BACnetError) parse(ctx context.Context, readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (__bACnetError BACnetError, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetError
	switch {
	case errorChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE: // SubscribeCOVPropertyMultipleError
		if _child, err = (&_SubscribeCOVPropertyMultipleError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SubscribeCOVPropertyMultipleError for type-switch of BACnetError")
		}
	case errorChoice == BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT: // ChangeListAddError
		if _child, err = (&_ChangeListAddError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ChangeListAddError for type-switch of BACnetError")
		}
	case errorChoice == BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT: // ChangeListRemoveError
		if _child, err = (&_ChangeListRemoveError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ChangeListRemoveError for type-switch of BACnetError")
		}
	case errorChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT: // CreateObjectError
		if _child, err = (&_CreateObjectError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CreateObjectError for type-switch of BACnetError")
		}
	case errorChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE: // WritePropertyMultipleError
		if _child, err = (&_WritePropertyMultipleError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type WritePropertyMultipleError for type-switch of BACnetError")
		}
	case errorChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER: // ConfirmedPrivateTransferError
		if _child, err = (&_ConfirmedPrivateTransferError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConfirmedPrivateTransferError for type-switch of BACnetError")
		}
	case errorChoice == BACnetConfirmedServiceChoice_VT_CLOSE: // VTCloseError
		if _child, err = (&_VTCloseError{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VTCloseError for type-switch of BACnetError")
		}
	case 0 == 0: // BACnetErrorGeneral
		if _child, err = (&_BACnetErrorGeneral{}).parse(ctx, readBuffer, m, errorChoice); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetErrorGeneral for type-switch of BACnetError")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [errorChoice=%v]", errorChoice)
	}

	if closeErr := readBuffer.CloseContext("BACnetError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetError")
	}

	return _child, nil
}

func (pm *_BACnetError) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetError, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetError"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetError")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetError"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetError")
	}
	return nil
}

func (m *_BACnetError) IsBACnetError() {}
