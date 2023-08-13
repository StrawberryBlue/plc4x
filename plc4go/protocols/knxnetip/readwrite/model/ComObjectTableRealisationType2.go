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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ComObjectTableRealisationType2 is the corresponding interface of ComObjectTableRealisationType2
type ComObjectTableRealisationType2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ComObjectTable
	// GetNumEntries returns NumEntries (property field)
	GetNumEntries() uint8
	// GetRamFlagsTablePointer returns RamFlagsTablePointer (property field)
	GetRamFlagsTablePointer() uint8
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() []GroupObjectDescriptorRealisationType2
}

// ComObjectTableRealisationType2Exactly can be used when we want exactly this type and not a type which fulfills ComObjectTableRealisationType2.
// This is useful for switch cases.
type ComObjectTableRealisationType2Exactly interface {
	ComObjectTableRealisationType2
	isComObjectTableRealisationType2() bool
}

// _ComObjectTableRealisationType2 is the data-structure of this message
type _ComObjectTableRealisationType2 struct {
	*_ComObjectTable
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []GroupObjectDescriptorRealisationType2
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType2) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType2) InitializeParent(parent ComObjectTable) {}

func (m *_ComObjectTableRealisationType2) GetParent() ComObjectTable {
	return m._ComObjectTable
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType2) GetNumEntries() uint8 {
	return m.NumEntries
}

func (m *_ComObjectTableRealisationType2) GetRamFlagsTablePointer() uint8 {
	return m.RamFlagsTablePointer
}

func (m *_ComObjectTableRealisationType2) GetComObjectDescriptors() []GroupObjectDescriptorRealisationType2 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType2 factory function for _ComObjectTableRealisationType2
func NewComObjectTableRealisationType2(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType2) *_ComObjectTableRealisationType2 {
	_result := &_ComObjectTableRealisationType2{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		_ComObjectTable:      NewComObjectTable(),
	}
	_result._ComObjectTable._ComObjectTableChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType2(structType any) ComObjectTableRealisationType2 {
	if casted, ok := structType.(ComObjectTableRealisationType2); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType2); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType2) GetTypeName() string {
	return "ComObjectTableRealisationType2"
}

func (m *_ComObjectTableRealisationType2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for _curItem, element := range m.ComObjectDescriptors {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ComObjectDescriptors), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ComObjectTableRealisationType2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ComObjectTableRealisationType2Parse(ctx context.Context, theBytes []byte, firmwareType FirmwareType) (ComObjectTableRealisationType2, error) {
	return ComObjectTableRealisationType2ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), firmwareType)
}

func ComObjectTableRealisationType2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, firmwareType FirmwareType) (ComObjectTableRealisationType2, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numEntries)
	_numEntries, _numEntriesErr := readBuffer.ReadUint8("numEntries", 8)
	if _numEntriesErr != nil {
		return nil, errors.Wrap(_numEntriesErr, "Error parsing 'numEntries' field of ComObjectTableRealisationType2")
	}
	numEntries := _numEntries

	// Simple Field (ramFlagsTablePointer)
	_ramFlagsTablePointer, _ramFlagsTablePointerErr := readBuffer.ReadUint8("ramFlagsTablePointer", 8)
	if _ramFlagsTablePointerErr != nil {
		return nil, errors.Wrap(_ramFlagsTablePointerErr, "Error parsing 'ramFlagsTablePointer' field of ComObjectTableRealisationType2")
	}
	ramFlagsTablePointer := _ramFlagsTablePointer

	// Array field (comObjectDescriptors)
	if pullErr := readBuffer.PullContext("comObjectDescriptors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for comObjectDescriptors")
	}
	// Count array
	comObjectDescriptors := make([]GroupObjectDescriptorRealisationType2, utils.Max(numEntries, 0))
	// This happens when the size is set conditional to 0
	if len(comObjectDescriptors) == 0 {
		comObjectDescriptors = nil
	}
	{
		_numItems := uint16(utils.Max(numEntries, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := GroupObjectDescriptorRealisationType2ParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'comObjectDescriptors' field of ComObjectTableRealisationType2")
			}
			comObjectDescriptors[_curItem] = _item.(GroupObjectDescriptorRealisationType2)
		}
	}
	if closeErr := readBuffer.CloseContext("comObjectDescriptors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for comObjectDescriptors")
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType2")
	}

	// Create a partially initialized instance
	_child := &_ComObjectTableRealisationType2{
		_ComObjectTable:      &_ComObjectTable{},
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
	}
	_child._ComObjectTable._ComObjectTableChildRequirements = _child
	return _child, nil
}

func (m *_ComObjectTableRealisationType2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComObjectTableRealisationType2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType2")
		}

		// Simple Field (numEntries)
		numEntries := uint8(m.GetNumEntries())
		_numEntriesErr := writeBuffer.WriteUint8("numEntries", 8, (numEntries))
		if _numEntriesErr != nil {
			return errors.Wrap(_numEntriesErr, "Error serializing 'numEntries' field")
		}

		// Simple Field (ramFlagsTablePointer)
		ramFlagsTablePointer := uint8(m.GetRamFlagsTablePointer())
		_ramFlagsTablePointerErr := writeBuffer.WriteUint8("ramFlagsTablePointer", 8, (ramFlagsTablePointer))
		if _ramFlagsTablePointerErr != nil {
			return errors.Wrap(_ramFlagsTablePointerErr, "Error serializing 'ramFlagsTablePointer' field")
		}

		// Array Field (comObjectDescriptors)
		if pushErr := writeBuffer.PushContext("comObjectDescriptors", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for comObjectDescriptors")
		}
		for _curItem, _element := range m.GetComObjectDescriptors() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetComObjectDescriptors()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'comObjectDescriptors' field")
			}
		}
		if popErr := writeBuffer.PopContext("comObjectDescriptors", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for comObjectDescriptors")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType2")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType2) isComObjectTableRealisationType2() bool {
	return true
}

func (m *_ComObjectTableRealisationType2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
