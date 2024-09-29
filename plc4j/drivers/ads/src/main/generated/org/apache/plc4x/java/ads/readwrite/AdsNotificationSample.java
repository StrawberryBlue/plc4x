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

public class AdsNotificationSample implements Message {

  // Properties.
  protected final long notificationHandle;
  protected final long sampleSize;
  protected final byte[] data;

  public AdsNotificationSample(long notificationHandle, long sampleSize, byte[] data) {
    super();
    this.notificationHandle = notificationHandle;
    this.sampleSize = sampleSize;
    this.data = data;
  }

  public long getNotificationHandle() {
    return notificationHandle;
  }

  public long getSampleSize() {
    return sampleSize;
  }

  public byte[] getData() {
    return data;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdsNotificationSample");

    // Simple Field (notificationHandle)
    writeSimpleField("notificationHandle", notificationHandle, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (sampleSize)
    writeSimpleField("sampleSize", sampleSize, writeUnsignedLong(writeBuffer, 32));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("AdsNotificationSample");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AdsNotificationSample _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (notificationHandle)
    lengthInBits += 32;

    // Simple field (sampleSize)
    lengthInBits += 32;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static AdsNotificationSample staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AdsNotificationSample");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long notificationHandle =
        readSimpleField("notificationHandle", readUnsignedLong(readBuffer, 32));

    long sampleSize = readSimpleField("sampleSize", readUnsignedLong(readBuffer, 32));

    byte[] data = readBuffer.readByteArray("data", Math.toIntExact(sampleSize));

    readBuffer.closeContext("AdsNotificationSample");
    // Create the instance
    AdsNotificationSample _adsNotificationSample;
    _adsNotificationSample = new AdsNotificationSample(notificationHandle, sampleSize, data);
    return _adsNotificationSample;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsNotificationSample)) {
      return false;
    }
    AdsNotificationSample that = (AdsNotificationSample) o;
    return (getNotificationHandle() == that.getNotificationHandle())
        && (getSampleSize() == that.getSampleSize())
        && (getData() == that.getData())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNotificationHandle(), getSampleSize(), getData());
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
