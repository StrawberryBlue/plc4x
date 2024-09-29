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

public class InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE
    extends InformationObjectWithoutTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.MEASURED_VALUE_NORMALISED_VALUE;
  }

  // Properties.
  protected final NormalizedValue nva;
  protected final QualityDescriptor qds;

  public InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE(
      int address, NormalizedValue nva, QualityDescriptor qds) {
    super(address);
    this.nva = nva;
    this.qds = qds;
  }

  public NormalizedValue getNva() {
    return nva;
  }

  public QualityDescriptor getQds() {
    return qds;
  }

  @Override
  protected void serializeInformationObjectWithoutTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE");

    // Simple Field (nva)
    writeSimpleField(
        "nva", nva, writeComplex(writeBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (qds)
    writeSimpleField(
        "qds", qds, writeComplex(writeBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nva)
    lengthInBits += nva.getLengthInBits();

    // Simple field (qds)
    lengthInBits += qds.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithoutTimeBuilder staticParseInformationObjectWithoutTimeBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext("InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NormalizedValue nva =
        readSimpleField(
            "nva",
            readComplex(() -> NormalizedValue.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    QualityDescriptor qds =
        readSimpleField(
            "qds",
            readComplex(() -> QualityDescriptor.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE");
    // Create the instance
    return new InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUEBuilderImpl(nva, qds);
  }

  public static class InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUEBuilderImpl
      implements InformationObjectWithoutTime.InformationObjectWithoutTimeBuilder {
    private final NormalizedValue nva;
    private final QualityDescriptor qds;

    public InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUEBuilderImpl(
        NormalizedValue nva, QualityDescriptor qds) {
      this.nva = nva;
      this.qds = qds;
    }

    public InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE build(int address) {
      InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE
          informationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE =
              new InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE(address, nva, qds);
      return informationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE)) {
      return false;
    }
    InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE that =
        (InformationObjectWithoutTime_MEASURED_VALUE_NORMALISED_VALUE) o;
    return (getNva() == that.getNva()) && (getQds() == that.getQds()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNva(), getQds());
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
