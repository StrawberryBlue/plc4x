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
package org.apache.plc4x.java.s7.readwrite;

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

public class TPKTPacket implements Message {

  // Constant values.
  public static final Short PROTOCOLID = 0x03;

  // Properties.
  protected final COTPPacket payload;

  public TPKTPacket(COTPPacket payload) {
    super();
    this.payload = payload;
  }

  public COTPPacket getPayload() {
    return payload;
  }

  public short getProtocolId() {
    return PROTOCOLID;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TPKTPacket");

    // Const Field (protocolId)
    writeConstField(
        "protocolId",
        PROTOCOLID,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (len) (Used for parsing, but its value is not stored as it's implicitly given
    // by the objects content)
    int len = (int) ((getPayload().getLengthInBytes()) + (4));
    writeImplicitField(
        "len",
        len,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (payload)
    writeSimpleField(
        "payload",
        payload,
        writeComplex(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("TPKTPacket");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    TPKTPacket _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (protocolId)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Implicit Field (len)
    lengthInBits += 16;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static TPKTPacket staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("TPKTPacket");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short protocolId =
        readConstField(
            "protocolId",
            readUnsignedShort(readBuffer, 8),
            TPKTPacket.PROTOCOLID,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Short reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 8),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int len =
        readImplicitField(
            "len", readUnsignedInt(readBuffer, 16), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    COTPPacket payload =
        readSimpleField(
            "payload",
            readComplex(() -> COTPPacket.staticParse(readBuffer, (int) ((len) - (4))), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("TPKTPacket");
    // Create the instance
    TPKTPacket _tPKTPacket;
    _tPKTPacket = new TPKTPacket(payload);
    return _tPKTPacket;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TPKTPacket)) {
      return false;
    }
    TPKTPacket that = (TPKTPacket) o;
    return (getPayload() == that.getPayload()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getPayload());
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
