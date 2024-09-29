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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetRecipientDevice extends BACnetRecipient implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagObjectIdentifier deviceValue;

  public BACnetRecipientDevice(
      BACnetTagHeader peekedTagHeader, BACnetContextTagObjectIdentifier deviceValue) {
    super(peekedTagHeader);
    this.deviceValue = deviceValue;
  }

  public BACnetContextTagObjectIdentifier getDeviceValue() {
    return deviceValue;
  }

  @Override
  protected void serializeBACnetRecipientChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetRecipientDevice");

    // Simple Field (deviceValue)
    writeSimpleField("deviceValue", deviceValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetRecipientDevice");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetRecipientDevice _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (deviceValue)
    lengthInBits += deviceValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetRecipientBuilder staticParseBACnetRecipientBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetRecipientDevice");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier deviceValue =
        readSimpleField(
            "deviceValue",
            readComplex(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    readBuffer.closeContext("BACnetRecipientDevice");
    // Create the instance
    return new BACnetRecipientDeviceBuilderImpl(deviceValue);
  }

  public static class BACnetRecipientDeviceBuilderImpl
      implements BACnetRecipient.BACnetRecipientBuilder {
    private final BACnetContextTagObjectIdentifier deviceValue;

    public BACnetRecipientDeviceBuilderImpl(BACnetContextTagObjectIdentifier deviceValue) {
      this.deviceValue = deviceValue;
    }

    public BACnetRecipientDevice build(BACnetTagHeader peekedTagHeader) {
      BACnetRecipientDevice bACnetRecipientDevice =
          new BACnetRecipientDevice(peekedTagHeader, deviceValue);
      return bACnetRecipientDevice;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetRecipientDevice)) {
      return false;
    }
    BACnetRecipientDevice that = (BACnetRecipientDevice) o;
    return (getDeviceValue() == that.getDeviceValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDeviceValue());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
