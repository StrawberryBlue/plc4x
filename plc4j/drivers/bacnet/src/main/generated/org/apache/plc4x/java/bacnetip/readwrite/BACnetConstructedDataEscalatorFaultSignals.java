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

public class BACnetConstructedDataEscalatorFaultSignals extends BACnetConstructedData
    implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return BACnetObjectType.ESCALATOR;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.FAULT_SIGNALS;
  }

  // Properties.
  protected final List<BACnetEscalatorFaultTagged> faultSignals;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataEscalatorFaultSignals(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      List<BACnetEscalatorFaultTagged> faultSignals,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.faultSignals = faultSignals;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public List<BACnetEscalatorFaultTagged> getFaultSignals() {
    return faultSignals;
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataEscalatorFaultSignals");

    // Array Field (faultSignals)
    writeComplexTypeArrayField("faultSignals", faultSignals, writeBuffer);

    writeBuffer.popContext("BACnetConstructedDataEscalatorFaultSignals");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataEscalatorFaultSignals _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (faultSignals != null) {
      for (Message element : faultSignals) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataEscalatorFaultSignals");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    List<BACnetEscalatorFaultTagged> faultSignals =
        readTerminatedArrayField(
            "faultSignals",
            readComplex(
                () ->
                    BACnetEscalatorFaultTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .isBACnetConstructedDataClosingTag(readBuffer, false, tagNumber))));

    readBuffer.closeContext("BACnetConstructedDataEscalatorFaultSignals");
    // Create the instance
    return new BACnetConstructedDataEscalatorFaultSignalsBuilderImpl(
        faultSignals, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataEscalatorFaultSignalsBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final List<BACnetEscalatorFaultTagged> faultSignals;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataEscalatorFaultSignalsBuilderImpl(
        List<BACnetEscalatorFaultTagged> faultSignals,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.faultSignals = faultSignals;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataEscalatorFaultSignals build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataEscalatorFaultSignals bACnetConstructedDataEscalatorFaultSignals =
          new BACnetConstructedDataEscalatorFaultSignals(
              openingTag, peekedTagHeader, closingTag, faultSignals, tagNumber, arrayIndexArgument);
      return bACnetConstructedDataEscalatorFaultSignals;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataEscalatorFaultSignals)) {
      return false;
    }
    BACnetConstructedDataEscalatorFaultSignals that =
        (BACnetConstructedDataEscalatorFaultSignals) o;
    return (getFaultSignals() == that.getFaultSignals()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFaultSignals());
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
