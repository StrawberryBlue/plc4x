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
package org.apache.plc4x.java.eip.readwrite;

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

public class ClassSegment implements Message {

  // Properties.
  protected final byte pathSegmentType;
  protected final byte logicalSegmentType;
  protected final byte logicalSegmentFormat;
  protected final short classSegment;

  public ClassSegment(
      byte pathSegmentType,
      byte logicalSegmentType,
      byte logicalSegmentFormat,
      short classSegment) {
    super();
    this.pathSegmentType = pathSegmentType;
    this.logicalSegmentType = logicalSegmentType;
    this.logicalSegmentFormat = logicalSegmentFormat;
    this.classSegment = classSegment;
  }

  public byte getPathSegmentType() {
    return pathSegmentType;
  }

  public byte getLogicalSegmentType() {
    return logicalSegmentType;
  }

  public byte getLogicalSegmentFormat() {
    return logicalSegmentFormat;
  }

  public short getClassSegment() {
    return classSegment;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ClassSegment");

    // Simple Field (pathSegmentType)
    writeSimpleField("pathSegmentType", pathSegmentType, writeUnsignedByte(writeBuffer, 3));

    // Simple Field (logicalSegmentType)
    writeSimpleField("logicalSegmentType", logicalSegmentType, writeUnsignedByte(writeBuffer, 3));

    // Simple Field (logicalSegmentFormat)
    writeSimpleField(
        "logicalSegmentFormat", logicalSegmentFormat, writeUnsignedByte(writeBuffer, 2));

    // Simple Field (classSegment)
    writeSimpleField("classSegment", classSegment, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("ClassSegment");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ClassSegment _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (pathSegmentType)
    lengthInBits += 3;

    // Simple field (logicalSegmentType)
    lengthInBits += 3;

    // Simple field (logicalSegmentFormat)
    lengthInBits += 2;

    // Simple field (classSegment)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ClassSegment staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ClassSegment");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte pathSegmentType = readSimpleField("pathSegmentType", readUnsignedByte(readBuffer, 3));

    byte logicalSegmentType =
        readSimpleField("logicalSegmentType", readUnsignedByte(readBuffer, 3));

    byte logicalSegmentFormat =
        readSimpleField("logicalSegmentFormat", readUnsignedByte(readBuffer, 2));

    short classSegment = readSimpleField("classSegment", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("ClassSegment");
    // Create the instance
    ClassSegment _classSegment;
    _classSegment =
        new ClassSegment(pathSegmentType, logicalSegmentType, logicalSegmentFormat, classSegment);
    return _classSegment;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ClassSegment)) {
      return false;
    }
    ClassSegment that = (ClassSegment) o;
    return (getPathSegmentType() == that.getPathSegmentType())
        && (getLogicalSegmentType() == that.getLogicalSegmentType())
        && (getLogicalSegmentFormat() == that.getLogicalSegmentFormat())
        && (getClassSegment() == that.getClassSegment())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getPathSegmentType(),
        getLogicalSegmentType(),
        getLogicalSegmentFormat(),
        getClassSegment());
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
