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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type ComObjectValueType uint8

type IComObjectValueType interface {
	SizeInBytes() uint8
	Serialize(io utils.WriteBuffer) error
}

const (
	ComObjectValueType_BIT1   ComObjectValueType = 0x00
	ComObjectValueType_BIT2   ComObjectValueType = 0x01
	ComObjectValueType_BIT3   ComObjectValueType = 0x02
	ComObjectValueType_BIT4   ComObjectValueType = 0x03
	ComObjectValueType_BIT5   ComObjectValueType = 0x04
	ComObjectValueType_BIT6   ComObjectValueType = 0x05
	ComObjectValueType_BIT7   ComObjectValueType = 0x06
	ComObjectValueType_BYTE1  ComObjectValueType = 0x07
	ComObjectValueType_BYTE2  ComObjectValueType = 0x08
	ComObjectValueType_BYTE3  ComObjectValueType = 0x09
	ComObjectValueType_BYTE4  ComObjectValueType = 0x0A
	ComObjectValueType_BYTE6  ComObjectValueType = 0x0B
	ComObjectValueType_BYTE8  ComObjectValueType = 0x0C
	ComObjectValueType_BYTE10 ComObjectValueType = 0x0D
	ComObjectValueType_BYTE14 ComObjectValueType = 0x0E
)

func (e ComObjectValueType) SizeInBytes() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 1
		}
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x02:
		{ /* '0x02' */
			return 1
		}
	case 0x03:
		{ /* '0x03' */
			return 1
		}
	case 0x04:
		{ /* '0x04' */
			return 1
		}
	case 0x05:
		{ /* '0x05' */
			return 1
		}
	case 0x06:
		{ /* '0x06' */
			return 1
		}
	case 0x07:
		{ /* '0x07' */
			return 1
		}
	case 0x08:
		{ /* '0x08' */
			return 2
		}
	case 0x09:
		{ /* '0x09' */
			return 3
		}
	case 0x0A:
		{ /* '0x0A' */
			return 4
		}
	case 0x0B:
		{ /* '0x0B' */
			return 6
		}
	case 0x0C:
		{ /* '0x0C' */
			return 8
		}
	case 0x0D:
		{ /* '0x0D' */
			return 10
		}
	case 0x0E:
		{ /* '0x0E' */
			return 14
		}
	default:
		{
			return 0
		}
	}
}
func ComObjectValueTypeByValue(value uint8) ComObjectValueType {
	switch value {
	case 0x00:
		return ComObjectValueType_BIT1
	case 0x01:
		return ComObjectValueType_BIT2
	case 0x02:
		return ComObjectValueType_BIT3
	case 0x03:
		return ComObjectValueType_BIT4
	case 0x04:
		return ComObjectValueType_BIT5
	case 0x05:
		return ComObjectValueType_BIT6
	case 0x06:
		return ComObjectValueType_BIT7
	case 0x07:
		return ComObjectValueType_BYTE1
	case 0x08:
		return ComObjectValueType_BYTE2
	case 0x09:
		return ComObjectValueType_BYTE3
	case 0x0A:
		return ComObjectValueType_BYTE4
	case 0x0B:
		return ComObjectValueType_BYTE6
	case 0x0C:
		return ComObjectValueType_BYTE8
	case 0x0D:
		return ComObjectValueType_BYTE10
	case 0x0E:
		return ComObjectValueType_BYTE14
	}
	return 0
}

func ComObjectValueTypeByName(value string) ComObjectValueType {
	switch value {
	case "BIT1":
		return ComObjectValueType_BIT1
	case "BIT2":
		return ComObjectValueType_BIT2
	case "BIT3":
		return ComObjectValueType_BIT3
	case "BIT4":
		return ComObjectValueType_BIT4
	case "BIT5":
		return ComObjectValueType_BIT5
	case "BIT6":
		return ComObjectValueType_BIT6
	case "BIT7":
		return ComObjectValueType_BIT7
	case "BYTE1":
		return ComObjectValueType_BYTE1
	case "BYTE2":
		return ComObjectValueType_BYTE2
	case "BYTE3":
		return ComObjectValueType_BYTE3
	case "BYTE4":
		return ComObjectValueType_BYTE4
	case "BYTE6":
		return ComObjectValueType_BYTE6
	case "BYTE8":
		return ComObjectValueType_BYTE8
	case "BYTE10":
		return ComObjectValueType_BYTE10
	case "BYTE14":
		return ComObjectValueType_BYTE14
	}
	return 0
}

func CastComObjectValueType(structType interface{}) ComObjectValueType {
	castFunc := func(typ interface{}) ComObjectValueType {
		if sComObjectValueType, ok := typ.(ComObjectValueType); ok {
			return sComObjectValueType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ComObjectValueType) LengthInBits() uint16 {
	return 8
}

func (m ComObjectValueType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ComObjectValueTypeParse(io *utils.ReadBuffer) (ComObjectValueType, error) {
	val, err := io.ReadUint8(8)
	if err != nil {
		return 0, nil
	}
	return ComObjectValueTypeByValue(val), nil
}

func (e ComObjectValueType) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint8(8, uint8(e))
	return err
}

func (e ComObjectValueType) String() string {
	switch e {
	case ComObjectValueType_BIT1:
		return "BIT1"
	case ComObjectValueType_BIT2:
		return "BIT2"
	case ComObjectValueType_BIT3:
		return "BIT3"
	case ComObjectValueType_BIT4:
		return "BIT4"
	case ComObjectValueType_BIT5:
		return "BIT5"
	case ComObjectValueType_BIT6:
		return "BIT6"
	case ComObjectValueType_BIT7:
		return "BIT7"
	case ComObjectValueType_BYTE1:
		return "BYTE1"
	case ComObjectValueType_BYTE2:
		return "BYTE2"
	case ComObjectValueType_BYTE3:
		return "BYTE3"
	case ComObjectValueType_BYTE4:
		return "BYTE4"
	case ComObjectValueType_BYTE6:
		return "BYTE6"
	case ComObjectValueType_BYTE8:
		return "BYTE8"
	case ComObjectValueType_BYTE10:
		return "BYTE10"
	case ComObjectValueType_BYTE14:
		return "BYTE14"
	}
	return ""
}
