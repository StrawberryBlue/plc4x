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

public class BACnetFaultParameterFaultStatusFlags extends BACnetFaultParameter implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetDeviceObjectPropertyReferenceEnclosed statusFlagsReference;
  protected final BACnetClosingTag closingTag;

  public BACnetFaultParameterFaultStatusFlags(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetDeviceObjectPropertyReferenceEnclosed statusFlagsReference,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.statusFlagsReference = statusFlagsReference;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetDeviceObjectPropertyReferenceEnclosed getStatusFlagsReference() {
    return statusFlagsReference;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetFaultParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetFaultParameterFaultStatusFlags");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, writeComplex(writeBuffer));

    // Simple Field (statusFlagsReference)
    writeSimpleField("statusFlagsReference", statusFlagsReference, writeComplex(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetFaultParameterFaultStatusFlags");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetFaultParameterFaultStatusFlags _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (statusFlagsReference)
    lengthInBits += statusFlagsReference.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetFaultParameterBuilder staticParseBACnetFaultParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetFaultParameterFaultStatusFlags");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            readComplex(() -> BACnetOpeningTag.staticParse(readBuffer, (short) (5)), readBuffer));

    BACnetDeviceObjectPropertyReferenceEnclosed statusFlagsReference =
        readSimpleField(
            "statusFlagsReference",
            readComplex(
                () ->
                    BACnetDeviceObjectPropertyReferenceEnclosed.staticParse(
                        readBuffer, (short) (1)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            readComplex(() -> BACnetClosingTag.staticParse(readBuffer, (short) (5)), readBuffer));

    readBuffer.closeContext("BACnetFaultParameterFaultStatusFlags");
    // Create the instance
    return new BACnetFaultParameterFaultStatusFlagsBuilderImpl(
        openingTag, statusFlagsReference, closingTag);
  }

  public static class BACnetFaultParameterFaultStatusFlagsBuilderImpl
      implements BACnetFaultParameter.BACnetFaultParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetDeviceObjectPropertyReferenceEnclosed statusFlagsReference;
    private final BACnetClosingTag closingTag;

    public BACnetFaultParameterFaultStatusFlagsBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetDeviceObjectPropertyReferenceEnclosed statusFlagsReference,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.statusFlagsReference = statusFlagsReference;
      this.closingTag = closingTag;
    }

    public BACnetFaultParameterFaultStatusFlags build(BACnetTagHeader peekedTagHeader) {
      BACnetFaultParameterFaultStatusFlags bACnetFaultParameterFaultStatusFlags =
          new BACnetFaultParameterFaultStatusFlags(
              peekedTagHeader, openingTag, statusFlagsReference, closingTag);
      return bACnetFaultParameterFaultStatusFlags;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetFaultParameterFaultStatusFlags)) {
      return false;
    }
    BACnetFaultParameterFaultStatusFlags that = (BACnetFaultParameterFaultStatusFlags) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getStatusFlagsReference() == that.getStatusFlagsReference())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getOpeningTag(), getStatusFlagsReference(), getClosingTag());
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
