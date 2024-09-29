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

public class BACnetObjectTypesSupportedTagged implements Message {

  // Properties.
  protected final BACnetTagHeader header;
  protected final BACnetTagPayloadBitString payload;

  // Arguments.
  protected final Short tagNumber;
  protected final TagClass tagClass;

  public BACnetObjectTypesSupportedTagged(
      BACnetTagHeader header,
      BACnetTagPayloadBitString payload,
      Short tagNumber,
      TagClass tagClass) {
    super();
    this.header = header;
    this.payload = payload;
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public BACnetTagPayloadBitString getPayload() {
    return payload;
  }

  public boolean getTimeValue() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (0))) ? getPayload().getData().get(0) : false));
  }

  public boolean getNotificationForwarder() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (1))) ? getPayload().getData().get(1) : false));
  }

  public boolean getAlertEnrollment() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (2))) ? getPayload().getData().get(2) : false));
  }

  public boolean getChannel() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (3))) ? getPayload().getData().get(3) : false));
  }

  public boolean getLightingOutput() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (4))) ? getPayload().getData().get(4) : false));
  }

  public boolean getBinaryLightingOutput() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (5))) ? getPayload().getData().get(5) : false));
  }

  public boolean getNetworkPort() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (6))) ? getPayload().getData().get(6) : false));
  }

  public boolean getElevatorGroup() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (7))) ? getPayload().getData().get(7) : false));
  }

  public boolean getEscalator() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (8))) ? getPayload().getData().get(8) : false));
  }

  public boolean getLift() {
    return (boolean)
        (((((COUNT(getPayload().getData())) > (9))) ? getPayload().getData().get(9) : false));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetObjectTypesSupportedTagged");

    // Simple Field (header)
    writeSimpleField("header", header, writeComplex(writeBuffer));

    // Simple Field (payload)
    writeSimpleField("payload", payload, writeComplex(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean timeValue = getTimeValue();
    writeBuffer.writeVirtual("timeValue", timeValue);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean notificationForwarder = getNotificationForwarder();
    writeBuffer.writeVirtual("notificationForwarder", notificationForwarder);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean alertEnrollment = getAlertEnrollment();
    writeBuffer.writeVirtual("alertEnrollment", alertEnrollment);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean channel = getChannel();
    writeBuffer.writeVirtual("channel", channel);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean lightingOutput = getLightingOutput();
    writeBuffer.writeVirtual("lightingOutput", lightingOutput);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean binaryLightingOutput = getBinaryLightingOutput();
    writeBuffer.writeVirtual("binaryLightingOutput", binaryLightingOutput);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean networkPort = getNetworkPort();
    writeBuffer.writeVirtual("networkPort", networkPort);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean elevatorGroup = getElevatorGroup();
    writeBuffer.writeVirtual("elevatorGroup", elevatorGroup);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean escalator = getEscalator();
    writeBuffer.writeVirtual("escalator", escalator);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean lift = getLift();
    writeBuffer.writeVirtual("lift", lift);

    writeBuffer.popContext("BACnetObjectTypesSupportedTagged");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetObjectTypesSupportedTagged _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetObjectTypesSupportedTagged staticParse(
      ReadBuffer readBuffer, Short tagNumber, TagClass tagClass) throws ParseException {
    readBuffer.pullContext("BACnetObjectTypesSupportedTagged");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header", readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getTagClass()) == (tagClass))) {
      throw new ParseValidationException("tag class doesn't match");
    }
    // Validation
    if (!((((header.getTagClass()) == (TagClass.APPLICATION_TAGS)))
        || (((header.getActualTagNumber()) == (tagNumber))))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }

    BACnetTagPayloadBitString payload =
        readSimpleField(
            "payload",
            readComplex(
                () ->
                    BACnetTagPayloadBitString.staticParse(
                        readBuffer, (long) (header.getActualLength())),
                readBuffer));
    boolean timeValue =
        readVirtualField(
            "timeValue",
            boolean.class,
            ((((COUNT(payload.getData())) > (0))) ? payload.getData().get(0) : false));
    boolean notificationForwarder =
        readVirtualField(
            "notificationForwarder",
            boolean.class,
            ((((COUNT(payload.getData())) > (1))) ? payload.getData().get(1) : false));
    boolean alertEnrollment =
        readVirtualField(
            "alertEnrollment",
            boolean.class,
            ((((COUNT(payload.getData())) > (2))) ? payload.getData().get(2) : false));
    boolean channel =
        readVirtualField(
            "channel",
            boolean.class,
            ((((COUNT(payload.getData())) > (3))) ? payload.getData().get(3) : false));
    boolean lightingOutput =
        readVirtualField(
            "lightingOutput",
            boolean.class,
            ((((COUNT(payload.getData())) > (4))) ? payload.getData().get(4) : false));
    boolean binaryLightingOutput =
        readVirtualField(
            "binaryLightingOutput",
            boolean.class,
            ((((COUNT(payload.getData())) > (5))) ? payload.getData().get(5) : false));
    boolean networkPort =
        readVirtualField(
            "networkPort",
            boolean.class,
            ((((COUNT(payload.getData())) > (6))) ? payload.getData().get(6) : false));
    boolean elevatorGroup =
        readVirtualField(
            "elevatorGroup",
            boolean.class,
            ((((COUNT(payload.getData())) > (7))) ? payload.getData().get(7) : false));
    boolean escalator =
        readVirtualField(
            "escalator",
            boolean.class,
            ((((COUNT(payload.getData())) > (8))) ? payload.getData().get(8) : false));
    boolean lift =
        readVirtualField(
            "lift",
            boolean.class,
            ((((COUNT(payload.getData())) > (9))) ? payload.getData().get(9) : false));

    readBuffer.closeContext("BACnetObjectTypesSupportedTagged");
    // Create the instance
    BACnetObjectTypesSupportedTagged _bACnetObjectTypesSupportedTagged;
    _bACnetObjectTypesSupportedTagged =
        new BACnetObjectTypesSupportedTagged(header, payload, tagNumber, tagClass);
    return _bACnetObjectTypesSupportedTagged;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetObjectTypesSupportedTagged)) {
      return false;
    }
    BACnetObjectTypesSupportedTagged that = (BACnetObjectTypesSupportedTagged) o;
    return (getHeader() == that.getHeader()) && (getPayload() == that.getPayload()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader(), getPayload());
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
