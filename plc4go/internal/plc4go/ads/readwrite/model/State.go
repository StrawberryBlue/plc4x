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
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	log "github.com/sirupsen/logrus"
	"io"
)

// The data-structure of this message
type State struct {
	InitCommand         bool
	UpdCommand          bool
	TimestampAdded      bool
	HighPriorityCommand bool
	SystemCommand       bool
	AdsCommand          bool
	NoReturn            bool
	Response            bool
	Broadcast           bool
	IState
}

// The corresponding interface
type IState interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

func NewState(initCommand bool, updCommand bool, timestampAdded bool, highPriorityCommand bool, systemCommand bool, adsCommand bool, noReturn bool, response bool, broadcast bool) *State {
	return &State{InitCommand: initCommand, UpdCommand: updCommand, TimestampAdded: timestampAdded, HighPriorityCommand: highPriorityCommand, SystemCommand: systemCommand, AdsCommand: adsCommand, NoReturn: noReturn, Response: response, Broadcast: broadcast}
}

func CastState(structType interface{}) *State {
	castFunc := func(typ interface{}) *State {
		if casted, ok := typ.(State); ok {
			return &casted
		}
		if casted, ok := typ.(*State); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *State) GetTypeName() string {
	return "State"
}

func (m *State) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (initCommand)
	lengthInBits += 1

	// Simple field (updCommand)
	lengthInBits += 1

	// Simple field (timestampAdded)
	lengthInBits += 1

	// Simple field (highPriorityCommand)
	lengthInBits += 1

	// Simple field (systemCommand)
	lengthInBits += 1

	// Simple field (adsCommand)
	lengthInBits += 1

	// Simple field (noReturn)
	lengthInBits += 1

	// Simple field (response)
	lengthInBits += 1

	// Simple field (broadcast)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 7

	return lengthInBits
}

func (m *State) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func StateParse(io *utils.ReadBuffer) (*State, error) {

	// Simple Field (initCommand)
	initCommand, _initCommandErr := io.ReadBit()
	if _initCommandErr != nil {
		return nil, errors.New("Error parsing 'initCommand' field " + _initCommandErr.Error())
	}

	// Simple Field (updCommand)
	updCommand, _updCommandErr := io.ReadBit()
	if _updCommandErr != nil {
		return nil, errors.New("Error parsing 'updCommand' field " + _updCommandErr.Error())
	}

	// Simple Field (timestampAdded)
	timestampAdded, _timestampAddedErr := io.ReadBit()
	if _timestampAddedErr != nil {
		return nil, errors.New("Error parsing 'timestampAdded' field " + _timestampAddedErr.Error())
	}

	// Simple Field (highPriorityCommand)
	highPriorityCommand, _highPriorityCommandErr := io.ReadBit()
	if _highPriorityCommandErr != nil {
		return nil, errors.New("Error parsing 'highPriorityCommand' field " + _highPriorityCommandErr.Error())
	}

	// Simple Field (systemCommand)
	systemCommand, _systemCommandErr := io.ReadBit()
	if _systemCommandErr != nil {
		return nil, errors.New("Error parsing 'systemCommand' field " + _systemCommandErr.Error())
	}

	// Simple Field (adsCommand)
	adsCommand, _adsCommandErr := io.ReadBit()
	if _adsCommandErr != nil {
		return nil, errors.New("Error parsing 'adsCommand' field " + _adsCommandErr.Error())
	}

	// Simple Field (noReturn)
	noReturn, _noReturnErr := io.ReadBit()
	if _noReturnErr != nil {
		return nil, errors.New("Error parsing 'noReturn' field " + _noReturnErr.Error())
	}

	// Simple Field (response)
	response, _responseErr := io.ReadBit()
	if _responseErr != nil {
		return nil, errors.New("Error parsing 'response' field " + _responseErr.Error())
	}

	// Simple Field (broadcast)
	broadcast, _broadcastErr := io.ReadBit()
	if _broadcastErr != nil {
		return nil, errors.New("Error parsing 'broadcast' field " + _broadcastErr.Error())
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadInt8(7)
		if _err != nil {
			return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
		}
		if reserved != int8(0x0) {
			log.WithFields(log.Fields{
				"expected value": int8(0x0),
				"got value":      reserved,
			}).Info("Got unexpected response.")
		}
	}

	// Create the instance
	return NewState(initCommand, updCommand, timestampAdded, highPriorityCommand, systemCommand, adsCommand, noReturn, response, broadcast), nil
}

func (m *State) Serialize(io utils.WriteBuffer) error {

	// Simple Field (initCommand)
	initCommand := bool(m.InitCommand)
	_initCommandErr := io.WriteBit((initCommand))
	if _initCommandErr != nil {
		return errors.New("Error serializing 'initCommand' field " + _initCommandErr.Error())
	}

	// Simple Field (updCommand)
	updCommand := bool(m.UpdCommand)
	_updCommandErr := io.WriteBit((updCommand))
	if _updCommandErr != nil {
		return errors.New("Error serializing 'updCommand' field " + _updCommandErr.Error())
	}

	// Simple Field (timestampAdded)
	timestampAdded := bool(m.TimestampAdded)
	_timestampAddedErr := io.WriteBit((timestampAdded))
	if _timestampAddedErr != nil {
		return errors.New("Error serializing 'timestampAdded' field " + _timestampAddedErr.Error())
	}

	// Simple Field (highPriorityCommand)
	highPriorityCommand := bool(m.HighPriorityCommand)
	_highPriorityCommandErr := io.WriteBit((highPriorityCommand))
	if _highPriorityCommandErr != nil {
		return errors.New("Error serializing 'highPriorityCommand' field " + _highPriorityCommandErr.Error())
	}

	// Simple Field (systemCommand)
	systemCommand := bool(m.SystemCommand)
	_systemCommandErr := io.WriteBit((systemCommand))
	if _systemCommandErr != nil {
		return errors.New("Error serializing 'systemCommand' field " + _systemCommandErr.Error())
	}

	// Simple Field (adsCommand)
	adsCommand := bool(m.AdsCommand)
	_adsCommandErr := io.WriteBit((adsCommand))
	if _adsCommandErr != nil {
		return errors.New("Error serializing 'adsCommand' field " + _adsCommandErr.Error())
	}

	// Simple Field (noReturn)
	noReturn := bool(m.NoReturn)
	_noReturnErr := io.WriteBit((noReturn))
	if _noReturnErr != nil {
		return errors.New("Error serializing 'noReturn' field " + _noReturnErr.Error())
	}

	// Simple Field (response)
	response := bool(m.Response)
	_responseErr := io.WriteBit((response))
	if _responseErr != nil {
		return errors.New("Error serializing 'response' field " + _responseErr.Error())
	}

	// Simple Field (broadcast)
	broadcast := bool(m.Broadcast)
	_broadcastErr := io.WriteBit((broadcast))
	if _broadcastErr != nil {
		return errors.New("Error serializing 'broadcast' field " + _broadcastErr.Error())
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteInt8(7, int8(0x0))
		if _err != nil {
			return errors.New("Error serializing 'reserved' field " + _err.Error())
		}
	}

	return nil
}

func (m *State) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "initCommand":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.InitCommand = data
			case "updCommand":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.UpdCommand = data
			case "timestampAdded":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TimestampAdded = data
			case "highPriorityCommand":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.HighPriorityCommand = data
			case "systemCommand":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SystemCommand = data
			case "adsCommand":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdsCommand = data
			case "noReturn":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.NoReturn = data
			case "response":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Response = data
			case "broadcast":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Broadcast = data
			}
		}
	}
}

func (m *State) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.ads.readwrite.State"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.InitCommand, xml.StartElement{Name: xml.Name{Local: "initCommand"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.UpdCommand, xml.StartElement{Name: xml.Name{Local: "updCommand"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TimestampAdded, xml.StartElement{Name: xml.Name{Local: "timestampAdded"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.HighPriorityCommand, xml.StartElement{Name: xml.Name{Local: "highPriorityCommand"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SystemCommand, xml.StartElement{Name: xml.Name{Local: "systemCommand"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.AdsCommand, xml.StartElement{Name: xml.Name{Local: "adsCommand"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.NoReturn, xml.StartElement{Name: xml.Name{Local: "noReturn"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Response, xml.StartElement{Name: xml.Name{Local: "response"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Broadcast, xml.StartElement{Name: xml.Name{Local: "broadcast"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}
