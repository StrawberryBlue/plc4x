//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUReadInputRegistersRequest struct {
	StartingAddress uint16
	Quantity        uint16
	Parent          *ModbusPDU
}

// The corresponding interface
type IModbusPDUReadInputRegistersRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadInputRegistersRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUReadInputRegistersRequest) FunctionFlag() uint8 {
	return 0x04
}

func (m *ModbusPDUReadInputRegistersRequest) Response() bool {
	return false
}

func (m *ModbusPDUReadInputRegistersRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadInputRegistersRequest(startingAddress uint16, quantity uint16) *ModbusPDU {
	child := &ModbusPDUReadInputRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Parent:          NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUReadInputRegistersRequest(structType interface{}) *ModbusPDUReadInputRegistersRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadInputRegistersRequest {
		if casted, ok := typ.(ModbusPDUReadInputRegistersRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadInputRegistersRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadInputRegistersRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadInputRegistersRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadInputRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadInputRegistersRequest"
}

func (m *ModbusPDUReadInputRegistersRequest) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadInputRegistersRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadInputRegistersRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

	// Simple Field (startingAddress)
	startingAddress, _startingAddressErr := io.ReadUint16(16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}

	// Simple Field (quantity)
	quantity, _quantityErr := io.ReadUint16(16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadInputRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Parent:          &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUReadInputRegistersRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := io.WriteUint16(16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := io.WriteUint16(16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReadInputRegistersRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "startingAddress":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.StartingAddress = data
			case "quantity":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Quantity = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (m *ModbusPDUReadInputRegistersRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.StartingAddress, xml.StartElement{Name: xml.Name{Local: "startingAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUReadInputRegistersRequest) String() string {
	return string(m.Box("ModbusPDUReadInputRegistersRequest", utils.DefaultWidth*2))
}

func (m ModbusPDUReadInputRegistersRequest) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ModbusPDUReadInputRegistersRequest"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("StartingAddress", m.StartingAddress, width-2))
	boxes = append(boxes, utils.BoxAnything("Quantity", m.Quantity, width-2))
	return utils.BoxString(name, string(utils.AlignBoxes(boxes, width-2)), width)
}
