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

// Code generated by "plc4xGenerator -type=Address -prefix=pdu_"; DO NOT EDIT.

package pdu

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *Address) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *Address) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("Address"); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.AddrType)

		if err := writeBuffer.WriteString("addrType", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if d.AddrNet != nil {
		if err := writeBuffer.WriteUint16("addrNet", 16, *d.AddrNet); err != nil {
			return err
		}
	}
	if err := writeBuffer.WriteByteArray("addrAddress", d.AddrAddress); err != nil {
		return err
	}
	if d.AddrLen != nil {
		{
			_value := fmt.Sprintf("%v", d.AddrLen)

			if err := writeBuffer.WriteString("addrLen", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.AddrRoute != nil {
		{
			_value := fmt.Sprintf("%v", d.AddrRoute)

			if err := writeBuffer.WriteString("addrRoute", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.AddrIP != nil {
		if err := writeBuffer.WriteUint32("addrIP", 32, *d.AddrIP); err != nil {
			return err
		}
	}
	if d.AddrMask != nil {
		if err := writeBuffer.WriteUint32("addrMask", 32, *d.AddrMask); err != nil {
			return err
		}
	}
	if d.AddrHost != nil {
		if err := writeBuffer.WriteUint32("addrHost", 32, *d.AddrHost); err != nil {
			return err
		}
	}
	if d.AddrSubnet != nil {
		if err := writeBuffer.WriteUint32("addrSubnet", 32, *d.AddrSubnet); err != nil {
			return err
		}
	}
	if d.AddrPort != nil {
		if err := writeBuffer.WriteUint16("addrPort", 16, *d.AddrPort); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteString("_leafName", uint32(len(d._leafName)*8), d._leafName); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("Address"); err != nil {
		return err
	}
	return nil
}

func (d *Address) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
