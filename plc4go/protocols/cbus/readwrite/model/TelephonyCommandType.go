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

// TelephonyCommandType is an enum
type TelephonyCommandType uint8

type ITelephonyCommandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	TelephonyCommandType_EVENT TelephonyCommandType = 0x00
)

var TelephonyCommandTypeValues []TelephonyCommandType

func init() {
	_ = errors.New
	TelephonyCommandTypeValues = []TelephonyCommandType{
		TelephonyCommandType_EVENT,
	}
}

func (e TelephonyCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0xFF
		}
	default:
		{
			return 0
		}
	}
}

func TelephonyCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (enum TelephonyCommandType, ok bool) {
	for _, sizeValue := range TelephonyCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, true
		}
	}
	return 0, false
}
func TelephonyCommandTypeByValue(value uint8) (enum TelephonyCommandType, ok bool) {
	switch value {
	case 0x00:
		return TelephonyCommandType_EVENT, true
	}
	return 0, false
}

func TelephonyCommandTypeByName(value string) (enum TelephonyCommandType, ok bool) {
	switch value {
	case "EVENT":
		return TelephonyCommandType_EVENT, true
	}
	return 0, false
}

func TelephonyCommandTypeKnows(value uint8) bool {
	for _, typeValue := range TelephonyCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTelephonyCommandType(structType any) TelephonyCommandType {
	castFunc := func(typ any) TelephonyCommandType {
		if sTelephonyCommandType, ok := typ.(TelephonyCommandType); ok {
			return sTelephonyCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m TelephonyCommandType) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m TelephonyCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyCommandTypeParse(ctx context.Context, theBytes []byte) (TelephonyCommandType, error) {
	return TelephonyCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TelephonyCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyCommandType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("TelephonyCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TelephonyCommandType")
	}
	if enum, ok := TelephonyCommandTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TelephonyCommandType")
		return TelephonyCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e TelephonyCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TelephonyCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("TelephonyCommandType", 4, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e TelephonyCommandType) GetValue() uint8 {
	return uint8(e)
}

func (e TelephonyCommandType) GetNumberOfArguments() uint8 {
	return e.NumberOfArguments()
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TelephonyCommandType) PLC4XEnumName() string {
	switch e {
	case TelephonyCommandType_EVENT:
		return "EVENT"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e TelephonyCommandType) String() string {
	return e.PLC4XEnumName()
}
