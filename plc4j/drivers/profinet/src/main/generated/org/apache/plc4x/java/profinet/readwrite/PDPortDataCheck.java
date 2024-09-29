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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PDPortDataCheck extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.PD_PORT_DATA_CHECK;
  }

  // Constant values.
  public static final Integer PADFIELD = 0x0000;

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final int slotNumber;
  protected final int subSlotNumber;
  protected final PnIoCm_Block checkPeers;

  public PDPortDataCheck(
      short blockVersionHigh,
      short blockVersionLow,
      int slotNumber,
      int subSlotNumber,
      PnIoCm_Block checkPeers) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.slotNumber = slotNumber;
    this.subSlotNumber = subSlotNumber;
    this.checkPeers = checkPeers;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public int getSlotNumber() {
    return slotNumber;
  }

  public int getSubSlotNumber() {
    return subSlotNumber;
  }

  public PnIoCm_Block getCheckPeers() {
    return checkPeers;
  }

  public int getPadField() {
    return PADFIELD;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PDPortDataCheck");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (padField)
    writeConstField(
        "padField",
        PADFIELD,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (slotNumber)
    writeSimpleField(
        "slotNumber",
        slotNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (subSlotNumber)
    writeSimpleField(
        "subSlotNumber",
        subSlotNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (checkPeers)
    writeSimpleField(
        "checkPeers",
        checkPeers,
        writeComplex(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PDPortDataCheck");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PDPortDataCheck _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Const Field (padField)
    lengthInBits += 16;

    // Simple field (slotNumber)
    lengthInBits += 16;

    // Simple field (subSlotNumber)
    lengthInBits += 16;

    // Simple field (checkPeers)
    lengthInBits += checkPeers.getLengthInBits();

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PDPortDataCheck");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int padField =
        readConstField(
            "padField",
            readUnsignedInt(readBuffer, 16),
            PDPortDataCheck.PADFIELD,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int slotNumber =
        readSimpleField(
            "slotNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int subSlotNumber =
        readSimpleField(
            "subSlotNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    PnIoCm_Block checkPeers =
        readSimpleField(
            "checkPeers",
            readComplex(() -> PnIoCm_Block.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PDPortDataCheck");
    // Create the instance
    return new PDPortDataCheckBuilderImpl(
        blockVersionHigh, blockVersionLow, slotNumber, subSlotNumber, checkPeers);
  }

  public static class PDPortDataCheckBuilderImpl implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final int slotNumber;
    private final int subSlotNumber;
    private final PnIoCm_Block checkPeers;

    public PDPortDataCheckBuilderImpl(
        short blockVersionHigh,
        short blockVersionLow,
        int slotNumber,
        int subSlotNumber,
        PnIoCm_Block checkPeers) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.slotNumber = slotNumber;
      this.subSlotNumber = subSlotNumber;
      this.checkPeers = checkPeers;
    }

    public PDPortDataCheck build() {
      PDPortDataCheck pDPortDataCheck =
          new PDPortDataCheck(
              blockVersionHigh, blockVersionLow, slotNumber, subSlotNumber, checkPeers);
      return pDPortDataCheck;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PDPortDataCheck)) {
      return false;
    }
    PDPortDataCheck that = (PDPortDataCheck) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getSlotNumber() == that.getSlotNumber())
        && (getSubSlotNumber() == that.getSubSlotNumber())
        && (getCheckPeers() == that.getCheckPeers())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getSlotNumber(),
        getSubSlotNumber(),
        getCheckPeers());
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
