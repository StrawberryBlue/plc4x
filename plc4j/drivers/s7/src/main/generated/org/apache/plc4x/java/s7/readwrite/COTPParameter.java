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

public abstract class COTPParameter implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getParameterType();

  public COTPParameter() {
    super();
  }

  protected abstract void serializeCOTPParameterChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("COTPParameter");

    // Discriminator Field (parameterType) (Used as input to a switch field)
    writeDiscriminatorField(
        "parameterType", getParameterType(), writeUnsignedShort(writeBuffer, 8));

    // Implicit Field (parameterLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short parameterLength = (short) ((getLengthInBytes()) - (2));
    writeImplicitField("parameterLength", parameterLength, writeUnsignedShort(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeCOTPParameterChild(writeBuffer);

    writeBuffer.popContext("COTPParameter");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    COTPParameter _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (parameterType)
    lengthInBits += 8;

    // Implicit Field (parameterLength)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static COTPParameter staticParse(ReadBuffer readBuffer, Short rest) throws ParseException {
    readBuffer.pullContext("COTPParameter");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short parameterType = readDiscriminatorField("parameterType", readUnsignedShort(readBuffer, 8));

    short parameterLength = readImplicitField("parameterLength", readUnsignedShort(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    COTPParameterBuilder builder = null;
    if (EvaluationHelper.equals(parameterType, (short) 0xC0)) {
      builder = COTPParameterTpduSize.staticParseCOTPParameterBuilder(readBuffer, rest);
    } else if (EvaluationHelper.equals(parameterType, (short) 0xC1)) {
      builder = COTPParameterCallingTsap.staticParseCOTPParameterBuilder(readBuffer, rest);
    } else if (EvaluationHelper.equals(parameterType, (short) 0xC2)) {
      builder = COTPParameterCalledTsap.staticParseCOTPParameterBuilder(readBuffer, rest);
    } else if (EvaluationHelper.equals(parameterType, (short) 0xC3)) {
      builder = COTPParameterChecksum.staticParseCOTPParameterBuilder(readBuffer, rest);
    } else if (EvaluationHelper.equals(parameterType, (short) 0xE0)) {
      builder =
          COTPParameterDisconnectAdditionalInformation.staticParseCOTPParameterBuilder(
              readBuffer, rest);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "parameterType="
              + parameterType
              + "]");
    }

    readBuffer.closeContext("COTPParameter");
    // Create the instance
    COTPParameter _cOTPParameter = builder.build();
    return _cOTPParameter;
  }

  public interface COTPParameterBuilder {
    COTPParameter build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof COTPParameter)) {
      return false;
    }
    COTPParameter that = (COTPParameter) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
