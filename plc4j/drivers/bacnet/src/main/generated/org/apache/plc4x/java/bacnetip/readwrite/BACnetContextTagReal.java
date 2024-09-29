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

public class BACnetContextTagReal extends BACnetContextTag implements Message {

  // Accessors for discriminator values.
  public BACnetDataType getDataType() {
    return BACnetDataType.REAL;
  }

  // Properties.
  protected final BACnetTagPayloadReal payload;

  // Arguments.
  protected final Short tagNumberArgument;

  public BACnetContextTagReal(
      BACnetTagHeader header, BACnetTagPayloadReal payload, Short tagNumberArgument) {
    super(header, tagNumberArgument);
    this.payload = payload;
    this.tagNumberArgument = tagNumberArgument;
  }

  public BACnetTagPayloadReal getPayload() {
    return payload;
  }

  public float getActualValue() {
    return (float) (getPayload().getValue());
  }

  @Override
  protected void serializeBACnetContextTagChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetContextTagReal");

    // Simple Field (payload)
    writeSimpleField("payload", payload, writeComplex(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    float actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetContextTagReal");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetContextTagReal _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetContextTagBuilder staticParseBACnetContextTagBuilder(
      ReadBuffer readBuffer, Short tagNumberArgument, BACnetDataType dataType)
      throws ParseException {
    readBuffer.pullContext("BACnetContextTagReal");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagPayloadReal payload =
        readSimpleField(
            "payload", readComplex(() -> BACnetTagPayloadReal.staticParse(readBuffer), readBuffer));
    float actualValue = readVirtualField("actualValue", float.class, payload.getValue());

    readBuffer.closeContext("BACnetContextTagReal");
    // Create the instance
    return new BACnetContextTagRealBuilderImpl(payload, tagNumberArgument);
  }

  public static class BACnetContextTagRealBuilderImpl
      implements BACnetContextTag.BACnetContextTagBuilder {
    private final BACnetTagPayloadReal payload;
    private final Short tagNumberArgument;

    public BACnetContextTagRealBuilderImpl(BACnetTagPayloadReal payload, Short tagNumberArgument) {
      this.payload = payload;
      this.tagNumberArgument = tagNumberArgument;
    }

    public BACnetContextTagReal build(BACnetTagHeader header, Short tagNumberArgument) {
      BACnetContextTagReal bACnetContextTagReal =
          new BACnetContextTagReal(header, payload, tagNumberArgument);
      return bACnetContextTagReal;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetContextTagReal)) {
      return false;
    }
    BACnetContextTagReal that = (BACnetContextTagReal) o;
    return (getPayload() == that.getPayload()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPayload());
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
