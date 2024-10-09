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

// Code generated by "plc4xGenerator -type=RouterInfo -prefix=netservice_"; DO NOT EDIT.

package netservice

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *RouterInfo) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *RouterInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("RouterInfo"); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("snet", uint32(len(d.snet.String())*8), d.snet.String()); err != nil {
		return err
	}
	if d.address != nil {
		if err := writeBuffer.WriteString("address", uint32(len(d.address.String())*8), d.address.String()); err != nil {
			return err
		}
	}
	if err := writeBuffer.PushContext("dnets", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _name, elem := range d.dnets {
		name := fmt.Sprintf("%v", &_name)
		_value := fmt.Sprintf("%v", elem)

		if err := writeBuffer.WriteString(name, uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("dnets", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("status", uint32(len(d.status.String())*8), d.status.String()); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("RouterInfo"); err != nil {
		return err
	}
	return nil
}

func (d *RouterInfo) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	wb := utils.NewWriteBufferBoxBased(utils.WithWriteBufferBoxBasedMergeSingleBoxes(), utils.WithWriteBufferBoxBasedOmitEmptyBoxes())
	if err := wb.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
