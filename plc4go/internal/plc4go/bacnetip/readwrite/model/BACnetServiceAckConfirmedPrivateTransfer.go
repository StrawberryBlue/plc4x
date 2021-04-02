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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetServiceAckConfirmedPrivateTransfer struct {
	Parent *BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckConfirmedPrivateTransfer interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckConfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x12
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) InitializeParent(parent *BACnetServiceAck) {
}

func NewBACnetServiceAckConfirmedPrivateTransfer() *BACnetServiceAck {
	child := &BACnetServiceAckConfirmedPrivateTransfer{
		Parent: NewBACnetServiceAck(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetServiceAckConfirmedPrivateTransfer(structType interface{}) *BACnetServiceAckConfirmedPrivateTransfer {
	castFunc := func(typ interface{}) *BACnetServiceAckConfirmedPrivateTransfer {
		if casted, ok := typ.(BACnetServiceAckConfirmedPrivateTransfer); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckConfirmedPrivateTransfer); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckConfirmedPrivateTransfer(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckConfirmedPrivateTransfer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetServiceAckConfirmedPrivateTransfer"
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckConfirmedPrivateTransferParse(io *utils.ReadBuffer) (*BACnetServiceAck, error) {

	// Create a partially initialized instance
	_child := &BACnetServiceAckConfirmedPrivateTransfer{
		Parent: &BACnetServiceAck{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
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

func (m *BACnetServiceAckConfirmedPrivateTransfer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetServiceAckConfirmedPrivateTransfer) String() string {
	return string(m.Box("BACnetServiceAckConfirmedPrivateTransfer", utils.DefaultWidth*2))
}

func (m BACnetServiceAckConfirmedPrivateTransfer) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetServiceAckConfirmedPrivateTransfer"
	}
	boxes := make([]utils.AsciiBox, 0)
	return utils.BoxString(name, string(utils.AlignBoxes(boxes, width-2)), width)
}
