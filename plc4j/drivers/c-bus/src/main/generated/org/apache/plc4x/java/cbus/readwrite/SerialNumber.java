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
package org.apache.plc4x.java.cbus.readwrite;

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

public class SerialNumber implements Message {

  // Properties.
  protected final byte octet1;
  protected final byte octet2;
  protected final byte octet3;
  protected final byte octet4;

  public SerialNumber(byte octet1, byte octet2, byte octet3, byte octet4) {
    super();
    this.octet1 = octet1;
    this.octet2 = octet2;
    this.octet3 = octet3;
    this.octet4 = octet4;
  }

  public byte getOctet1() {
    return octet1;
  }

  public byte getOctet2() {
    return octet2;
  }

  public byte getOctet3() {
    return octet3;
  }

  public byte getOctet4() {
    return octet4;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SerialNumber");

    // Simple Field (octet1)
    writeSimpleField("octet1", octet1, writeByte(writeBuffer, 8));

    // Simple Field (octet2)
    writeSimpleField("octet2", octet2, writeByte(writeBuffer, 8));

    // Simple Field (octet3)
    writeSimpleField("octet3", octet3, writeByte(writeBuffer, 8));

    // Simple Field (octet4)
    writeSimpleField("octet4", octet4, writeByte(writeBuffer, 8));

    writeBuffer.popContext("SerialNumber");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SerialNumber _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (octet1)
    lengthInBits += 8;

    // Simple field (octet2)
    lengthInBits += 8;

    // Simple field (octet3)
    lengthInBits += 8;

    // Simple field (octet4)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static SerialNumber staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("SerialNumber");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte octet1 = readSimpleField("octet1", readByte(readBuffer, 8));

    byte octet2 = readSimpleField("octet2", readByte(readBuffer, 8));

    byte octet3 = readSimpleField("octet3", readByte(readBuffer, 8));

    byte octet4 = readSimpleField("octet4", readByte(readBuffer, 8));

    readBuffer.closeContext("SerialNumber");
    // Create the instance
    SerialNumber _serialNumber;
    _serialNumber = new SerialNumber(octet1, octet2, octet3, octet4);
    return _serialNumber;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SerialNumber)) {
      return false;
    }
    SerialNumber that = (SerialNumber) o;
    return (getOctet1() == that.getOctet1())
        && (getOctet2() == that.getOctet2())
        && (getOctet3() == that.getOctet3())
        && (getOctet4() == that.getOctet4())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOctet1(), getOctet2(), getOctet3(), getOctet4());
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
