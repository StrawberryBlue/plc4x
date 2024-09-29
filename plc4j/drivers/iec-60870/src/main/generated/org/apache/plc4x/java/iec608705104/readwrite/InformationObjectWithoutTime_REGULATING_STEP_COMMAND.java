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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObjectWithoutTime_REGULATING_STEP_COMMAND
    extends InformationObjectWithoutTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.REGULATING_STEP_COMMAND;
  }

  // Properties.
  protected final RegulatingStepCommand rco;

  public InformationObjectWithoutTime_REGULATING_STEP_COMMAND(
      int address, RegulatingStepCommand rco) {
    super(address);
    this.rco = rco;
  }

  public RegulatingStepCommand getRco() {
    return rco;
  }

  @Override
  protected void serializeInformationObjectWithoutTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithoutTime_REGULATING_STEP_COMMAND");

    // Simple Field (rco)
    writeSimpleField(
        "rco", rco, writeComplex(writeBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithoutTime_REGULATING_STEP_COMMAND");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithoutTime_REGULATING_STEP_COMMAND _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (rco)
    lengthInBits += rco.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithoutTimeBuilder staticParseInformationObjectWithoutTimeBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext("InformationObjectWithoutTime_REGULATING_STEP_COMMAND");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RegulatingStepCommand rco =
        readSimpleField(
            "rco",
            readComplex(() -> RegulatingStepCommand.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithoutTime_REGULATING_STEP_COMMAND");
    // Create the instance
    return new InformationObjectWithoutTime_REGULATING_STEP_COMMANDBuilderImpl(rco);
  }

  public static class InformationObjectWithoutTime_REGULATING_STEP_COMMANDBuilderImpl
      implements InformationObjectWithoutTime.InformationObjectWithoutTimeBuilder {
    private final RegulatingStepCommand rco;

    public InformationObjectWithoutTime_REGULATING_STEP_COMMANDBuilderImpl(
        RegulatingStepCommand rco) {
      this.rco = rco;
    }

    public InformationObjectWithoutTime_REGULATING_STEP_COMMAND build(int address) {
      InformationObjectWithoutTime_REGULATING_STEP_COMMAND
          informationObjectWithoutTime_REGULATING_STEP_COMMAND =
              new InformationObjectWithoutTime_REGULATING_STEP_COMMAND(address, rco);
      return informationObjectWithoutTime_REGULATING_STEP_COMMAND;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithoutTime_REGULATING_STEP_COMMAND)) {
      return false;
    }
    InformationObjectWithoutTime_REGULATING_STEP_COMMAND that =
        (InformationObjectWithoutTime_REGULATING_STEP_COMMAND) o;
    return (getRco() == that.getRco()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRco());
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
