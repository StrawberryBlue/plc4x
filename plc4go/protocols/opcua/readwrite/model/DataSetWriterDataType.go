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

// DataSetWriterDataType is the corresponding interface of DataSetWriterDataType
type DataSetWriterDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetDataSetWriterId returns DataSetWriterId (property field)
	GetDataSetWriterId() uint16
	// GetDataSetFieldContentMask returns DataSetFieldContentMask (property field)
	GetDataSetFieldContentMask() DataSetFieldContentMask
	// GetKeyFrameCount returns KeyFrameCount (property field)
	GetKeyFrameCount() uint32
	// GetDataSetName returns DataSetName (property field)
	GetDataSetName() PascalString
	// GetNoOfDataSetWriterProperties returns NoOfDataSetWriterProperties (property field)
	GetNoOfDataSetWriterProperties() int32
	// GetDataSetWriterProperties returns DataSetWriterProperties (property field)
	GetDataSetWriterProperties() []ExtensionObjectDefinition
	// GetTransportSettings returns TransportSettings (property field)
	GetTransportSettings() ExtensionObject
	// GetMessageSettings returns MessageSettings (property field)
	GetMessageSettings() ExtensionObject
}

// DataSetWriterDataTypeExactly can be used when we want exactly this type and not a type which fulfills DataSetWriterDataType.
// This is useful for switch cases.
type DataSetWriterDataTypeExactly interface {
	DataSetWriterDataType
	isDataSetWriterDataType() bool
}

// _DataSetWriterDataType is the data-structure of this message
type _DataSetWriterDataType struct {
	*_ExtensionObjectDefinition
	Name                        PascalString
	Enabled                     bool
	DataSetWriterId             uint16
	DataSetFieldContentMask     DataSetFieldContentMask
	KeyFrameCount               uint32
	DataSetName                 PascalString
	NoOfDataSetWriterProperties int32
	DataSetWriterProperties     []ExtensionObjectDefinition
	TransportSettings           ExtensionObject
	MessageSettings             ExtensionObject
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetWriterDataType) GetIdentifier() string {
	return "15599"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetWriterDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DataSetWriterDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataSetWriterDataType) GetName() PascalString {
	return m.Name
}

func (m *_DataSetWriterDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_DataSetWriterDataType) GetDataSetWriterId() uint16 {
	return m.DataSetWriterId
}

func (m *_DataSetWriterDataType) GetDataSetFieldContentMask() DataSetFieldContentMask {
	return m.DataSetFieldContentMask
}

func (m *_DataSetWriterDataType) GetKeyFrameCount() uint32 {
	return m.KeyFrameCount
}

func (m *_DataSetWriterDataType) GetDataSetName() PascalString {
	return m.DataSetName
}

func (m *_DataSetWriterDataType) GetNoOfDataSetWriterProperties() int32 {
	return m.NoOfDataSetWriterProperties
}

func (m *_DataSetWriterDataType) GetDataSetWriterProperties() []ExtensionObjectDefinition {
	return m.DataSetWriterProperties
}

func (m *_DataSetWriterDataType) GetTransportSettings() ExtensionObject {
	return m.TransportSettings
}

func (m *_DataSetWriterDataType) GetMessageSettings() ExtensionObject {
	return m.MessageSettings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDataSetWriterDataType factory function for _DataSetWriterDataType
func NewDataSetWriterDataType(name PascalString, enabled bool, dataSetWriterId uint16, dataSetFieldContentMask DataSetFieldContentMask, keyFrameCount uint32, dataSetName PascalString, noOfDataSetWriterProperties int32, dataSetWriterProperties []ExtensionObjectDefinition, transportSettings ExtensionObject, messageSettings ExtensionObject) *_DataSetWriterDataType {
	_result := &_DataSetWriterDataType{
		Name:                        name,
		Enabled:                     enabled,
		DataSetWriterId:             dataSetWriterId,
		DataSetFieldContentMask:     dataSetFieldContentMask,
		KeyFrameCount:               keyFrameCount,
		DataSetName:                 dataSetName,
		NoOfDataSetWriterProperties: noOfDataSetWriterProperties,
		DataSetWriterProperties:     dataSetWriterProperties,
		TransportSettings:           transportSettings,
		MessageSettings:             messageSettings,
		_ExtensionObjectDefinition:  NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataSetWriterDataType(structType any) DataSetWriterDataType {
	if casted, ok := structType.(DataSetWriterDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetWriterDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetWriterDataType) GetTypeName() string {
	return "DataSetWriterDataType"
}

func (m *_DataSetWriterDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (dataSetWriterId)
	lengthInBits += 16

	// Simple field (dataSetFieldContentMask)
	lengthInBits += 32

	// Simple field (keyFrameCount)
	lengthInBits += 32

	// Simple field (dataSetName)
	lengthInBits += m.DataSetName.GetLengthInBits(ctx)

	// Simple field (noOfDataSetWriterProperties)
	lengthInBits += 32

	// Array field
	if len(m.DataSetWriterProperties) > 0 {
		for _curItem, element := range m.DataSetWriterProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataSetWriterProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportSettings)
	lengthInBits += m.TransportSettings.GetLengthInBits(ctx)

	// Simple field (messageSettings)
	lengthInBits += m.MessageSettings.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataSetWriterDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSetWriterDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (DataSetWriterDataType, error) {
	return DataSetWriterDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DataSetWriterDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DataSetWriterDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DataSetWriterDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetWriterDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of DataSetWriterDataType")
	}
	name := _name.(PascalString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DataSetWriterDataType")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (enabled)
	_enabled, _enabledErr := readBuffer.ReadBit("enabled")
	if _enabledErr != nil {
		return nil, errors.Wrap(_enabledErr, "Error parsing 'enabled' field of DataSetWriterDataType")
	}
	enabled := _enabled

	// Simple Field (dataSetWriterId)
	_dataSetWriterId, _dataSetWriterIdErr := readBuffer.ReadUint16("dataSetWriterId", 16)
	if _dataSetWriterIdErr != nil {
		return nil, errors.Wrap(_dataSetWriterIdErr, "Error parsing 'dataSetWriterId' field of DataSetWriterDataType")
	}
	dataSetWriterId := _dataSetWriterId

	// Simple Field (dataSetFieldContentMask)
	if pullErr := readBuffer.PullContext("dataSetFieldContentMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataSetFieldContentMask")
	}
	_dataSetFieldContentMask, _dataSetFieldContentMaskErr := DataSetFieldContentMaskParseWithBuffer(ctx, readBuffer)
	if _dataSetFieldContentMaskErr != nil {
		return nil, errors.Wrap(_dataSetFieldContentMaskErr, "Error parsing 'dataSetFieldContentMask' field of DataSetWriterDataType")
	}
	dataSetFieldContentMask := _dataSetFieldContentMask
	if closeErr := readBuffer.CloseContext("dataSetFieldContentMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataSetFieldContentMask")
	}

	// Simple Field (keyFrameCount)
	_keyFrameCount, _keyFrameCountErr := readBuffer.ReadUint32("keyFrameCount", 32)
	if _keyFrameCountErr != nil {
		return nil, errors.Wrap(_keyFrameCountErr, "Error parsing 'keyFrameCount' field of DataSetWriterDataType")
	}
	keyFrameCount := _keyFrameCount

	// Simple Field (dataSetName)
	if pullErr := readBuffer.PullContext("dataSetName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataSetName")
	}
	_dataSetName, _dataSetNameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _dataSetNameErr != nil {
		return nil, errors.Wrap(_dataSetNameErr, "Error parsing 'dataSetName' field of DataSetWriterDataType")
	}
	dataSetName := _dataSetName.(PascalString)
	if closeErr := readBuffer.CloseContext("dataSetName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataSetName")
	}

	// Simple Field (noOfDataSetWriterProperties)
	_noOfDataSetWriterProperties, _noOfDataSetWriterPropertiesErr := readBuffer.ReadInt32("noOfDataSetWriterProperties", 32)
	if _noOfDataSetWriterPropertiesErr != nil {
		return nil, errors.Wrap(_noOfDataSetWriterPropertiesErr, "Error parsing 'noOfDataSetWriterProperties' field of DataSetWriterDataType")
	}
	noOfDataSetWriterProperties := _noOfDataSetWriterProperties

	// Array field (dataSetWriterProperties)
	if pullErr := readBuffer.PullContext("dataSetWriterProperties", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataSetWriterProperties")
	}
	// Count array
	dataSetWriterProperties := make([]ExtensionObjectDefinition, utils.Max(noOfDataSetWriterProperties, 0))
	// This happens when the size is set conditional to 0
	if len(dataSetWriterProperties) == 0 {
		dataSetWriterProperties = nil
	}
	{
		_numItems := uint16(utils.Max(noOfDataSetWriterProperties, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "14535")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'dataSetWriterProperties' field of DataSetWriterDataType")
			}
			dataSetWriterProperties[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("dataSetWriterProperties", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataSetWriterProperties")
	}

	// Simple Field (transportSettings)
	if pullErr := readBuffer.PullContext("transportSettings"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSettings")
	}
	_transportSettings, _transportSettingsErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _transportSettingsErr != nil {
		return nil, errors.Wrap(_transportSettingsErr, "Error parsing 'transportSettings' field of DataSetWriterDataType")
	}
	transportSettings := _transportSettings.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("transportSettings"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSettings")
	}

	// Simple Field (messageSettings)
	if pullErr := readBuffer.PullContext("messageSettings"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageSettings")
	}
	_messageSettings, _messageSettingsErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _messageSettingsErr != nil {
		return nil, errors.Wrap(_messageSettingsErr, "Error parsing 'messageSettings' field of DataSetWriterDataType")
	}
	messageSettings := _messageSettings.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("messageSettings"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageSettings")
	}

	if closeErr := readBuffer.CloseContext("DataSetWriterDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetWriterDataType")
	}

	// Create a partially initialized instance
	_child := &_DataSetWriterDataType{
		_ExtensionObjectDefinition:  &_ExtensionObjectDefinition{},
		Name:                        name,
		Enabled:                     enabled,
		DataSetWriterId:             dataSetWriterId,
		DataSetFieldContentMask:     dataSetFieldContentMask,
		KeyFrameCount:               keyFrameCount,
		DataSetName:                 dataSetName,
		NoOfDataSetWriterProperties: noOfDataSetWriterProperties,
		DataSetWriterProperties:     dataSetWriterProperties,
		TransportSettings:           transportSettings,
		MessageSettings:             messageSettings,
		reservedField0:              reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DataSetWriterDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetWriterDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetWriterDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetWriterDataType")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (enabled)
		enabled := bool(m.GetEnabled())
		_enabledErr := writeBuffer.WriteBit("enabled", (enabled))
		if _enabledErr != nil {
			return errors.Wrap(_enabledErr, "Error serializing 'enabled' field")
		}

		// Simple Field (dataSetWriterId)
		dataSetWriterId := uint16(m.GetDataSetWriterId())
		_dataSetWriterIdErr := writeBuffer.WriteUint16("dataSetWriterId", 16, (dataSetWriterId))
		if _dataSetWriterIdErr != nil {
			return errors.Wrap(_dataSetWriterIdErr, "Error serializing 'dataSetWriterId' field")
		}

		// Simple Field (dataSetFieldContentMask)
		if pushErr := writeBuffer.PushContext("dataSetFieldContentMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataSetFieldContentMask")
		}
		_dataSetFieldContentMaskErr := writeBuffer.WriteSerializable(ctx, m.GetDataSetFieldContentMask())
		if popErr := writeBuffer.PopContext("dataSetFieldContentMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataSetFieldContentMask")
		}
		if _dataSetFieldContentMaskErr != nil {
			return errors.Wrap(_dataSetFieldContentMaskErr, "Error serializing 'dataSetFieldContentMask' field")
		}

		// Simple Field (keyFrameCount)
		keyFrameCount := uint32(m.GetKeyFrameCount())
		_keyFrameCountErr := writeBuffer.WriteUint32("keyFrameCount", 32, (keyFrameCount))
		if _keyFrameCountErr != nil {
			return errors.Wrap(_keyFrameCountErr, "Error serializing 'keyFrameCount' field")
		}

		// Simple Field (dataSetName)
		if pushErr := writeBuffer.PushContext("dataSetName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataSetName")
		}
		_dataSetNameErr := writeBuffer.WriteSerializable(ctx, m.GetDataSetName())
		if popErr := writeBuffer.PopContext("dataSetName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataSetName")
		}
		if _dataSetNameErr != nil {
			return errors.Wrap(_dataSetNameErr, "Error serializing 'dataSetName' field")
		}

		// Simple Field (noOfDataSetWriterProperties)
		noOfDataSetWriterProperties := int32(m.GetNoOfDataSetWriterProperties())
		_noOfDataSetWriterPropertiesErr := writeBuffer.WriteInt32("noOfDataSetWriterProperties", 32, (noOfDataSetWriterProperties))
		if _noOfDataSetWriterPropertiesErr != nil {
			return errors.Wrap(_noOfDataSetWriterPropertiesErr, "Error serializing 'noOfDataSetWriterProperties' field")
		}

		// Array Field (dataSetWriterProperties)
		if pushErr := writeBuffer.PushContext("dataSetWriterProperties", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataSetWriterProperties")
		}
		for _curItem, _element := range m.GetDataSetWriterProperties() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDataSetWriterProperties()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'dataSetWriterProperties' field")
			}
		}
		if popErr := writeBuffer.PopContext("dataSetWriterProperties", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataSetWriterProperties")
		}

		// Simple Field (transportSettings)
		if pushErr := writeBuffer.PushContext("transportSettings"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportSettings")
		}
		_transportSettingsErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSettings())
		if popErr := writeBuffer.PopContext("transportSettings"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportSettings")
		}
		if _transportSettingsErr != nil {
			return errors.Wrap(_transportSettingsErr, "Error serializing 'transportSettings' field")
		}

		// Simple Field (messageSettings)
		if pushErr := writeBuffer.PushContext("messageSettings"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messageSettings")
		}
		_messageSettingsErr := writeBuffer.WriteSerializable(ctx, m.GetMessageSettings())
		if popErr := writeBuffer.PopContext("messageSettings"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messageSettings")
		}
		if _messageSettingsErr != nil {
			return errors.Wrap(_messageSettingsErr, "Error serializing 'messageSettings' field")
		}

		if popErr := writeBuffer.PopContext("DataSetWriterDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetWriterDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetWriterDataType) isDataSetWriterDataType() bool {
	return true
}

func (m *_DataSetWriterDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
