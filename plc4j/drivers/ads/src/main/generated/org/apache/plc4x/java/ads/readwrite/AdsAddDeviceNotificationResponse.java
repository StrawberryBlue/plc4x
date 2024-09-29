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
package org.apache.plc4x.java.ads.readwrite;

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

public class AdsAddDeviceNotificationResponse extends AmsPacket implements Message {

  // Accessors for discriminator values.
  public CommandId getCommandId() {
    return CommandId.ADS_ADD_DEVICE_NOTIFICATION;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final ReturnCode result;
  protected final long notificationHandle;

  public AdsAddDeviceNotificationResponse(
      AmsNetId targetAmsNetId,
      int targetAmsPort,
      AmsNetId sourceAmsNetId,
      int sourceAmsPort,
      long errorCode,
      long invokeId,
      ReturnCode result,
      long notificationHandle) {
    super(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId);
    this.result = result;
    this.notificationHandle = notificationHandle;
  }

  public ReturnCode getResult() {
    return result;
  }

  public long getNotificationHandle() {
    return notificationHandle;
  }

  @Override
  protected void serializeAmsPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdsAddDeviceNotificationResponse");

    // Simple Field (result)
    writeSimpleEnumField(
        "result",
        "ReturnCode",
        result,
        writeEnum(ReturnCode::getValue, ReturnCode::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (notificationHandle)
    writeSimpleField("notificationHandle", notificationHandle, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("AdsAddDeviceNotificationResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsAddDeviceNotificationResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (result)
    lengthInBits += 32;

    // Simple field (notificationHandle)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static AmsPacketBuilder staticParseAmsPacketBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsAddDeviceNotificationResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReturnCode result =
        readEnumField(
            "result",
            "ReturnCode",
            readEnum(ReturnCode::enumForValue, readUnsignedLong(readBuffer, 32)));

    long notificationHandle =
        readSimpleField("notificationHandle", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("AdsAddDeviceNotificationResponse");
    // Create the instance
    return new AdsAddDeviceNotificationResponseBuilderImpl(result, notificationHandle);
  }

  public static class AdsAddDeviceNotificationResponseBuilderImpl
      implements AmsPacket.AmsPacketBuilder {
    private final ReturnCode result;
    private final long notificationHandle;

    public AdsAddDeviceNotificationResponseBuilderImpl(ReturnCode result, long notificationHandle) {
      this.result = result;
      this.notificationHandle = notificationHandle;
    }

    public AdsAddDeviceNotificationResponse build(
        AmsNetId targetAmsNetId,
        int targetAmsPort,
        AmsNetId sourceAmsNetId,
        int sourceAmsPort,
        long errorCode,
        long invokeId) {
      AdsAddDeviceNotificationResponse adsAddDeviceNotificationResponse =
          new AdsAddDeviceNotificationResponse(
              targetAmsNetId,
              targetAmsPort,
              sourceAmsNetId,
              sourceAmsPort,
              errorCode,
              invokeId,
              result,
              notificationHandle);
      return adsAddDeviceNotificationResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsAddDeviceNotificationResponse)) {
      return false;
    }
    AdsAddDeviceNotificationResponse that = (AdsAddDeviceNotificationResponse) o;
    return (getResult() == that.getResult())
        && (getNotificationHandle() == that.getNotificationHandle())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getResult(), getNotificationHandle());
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
