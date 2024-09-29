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
package org.apache.plc4x.java.opcua.readwrite;

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

public class Argument extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "298";
  }

  // Properties.
  protected final PascalString name;
  protected final NodeId dataType;
  protected final int valueRank;
  protected final int noOfArrayDimensions;
  protected final List<Long> arrayDimensions;
  protected final LocalizedText description;

  public Argument(
      PascalString name,
      NodeId dataType,
      int valueRank,
      int noOfArrayDimensions,
      List<Long> arrayDimensions,
      LocalizedText description) {
    super();
    this.name = name;
    this.dataType = dataType;
    this.valueRank = valueRank;
    this.noOfArrayDimensions = noOfArrayDimensions;
    this.arrayDimensions = arrayDimensions;
    this.description = description;
  }

  public PascalString getName() {
    return name;
  }

  public NodeId getDataType() {
    return dataType;
  }

  public int getValueRank() {
    return valueRank;
  }

  public int getNoOfArrayDimensions() {
    return noOfArrayDimensions;
  }

  public List<Long> getArrayDimensions() {
    return arrayDimensions;
  }

  public LocalizedText getDescription() {
    return description;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Argument");

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    // Simple Field (dataType)
    writeSimpleField("dataType", dataType, writeComplex(writeBuffer));

    // Simple Field (valueRank)
    writeSimpleField("valueRank", valueRank, writeSignedInt(writeBuffer, 32));

    // Simple Field (noOfArrayDimensions)
    writeSimpleField("noOfArrayDimensions", noOfArrayDimensions, writeSignedInt(writeBuffer, 32));

    // Array Field (arrayDimensions)
    writeSimpleTypeArrayField(
        "arrayDimensions", arrayDimensions, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    writeBuffer.popContext("Argument");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Argument _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (dataType)
    lengthInBits += dataType.getLengthInBits();

    // Simple field (valueRank)
    lengthInBits += 32;

    // Simple field (noOfArrayDimensions)
    lengthInBits += 32;

    // Array field
    if (arrayDimensions != null) {
      lengthInBits += 32 * arrayDimensions.size();
    }

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("Argument");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    NodeId dataType =
        readSimpleField("dataType", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    int valueRank = readSimpleField("valueRank", readSignedInt(readBuffer, 32));

    int noOfArrayDimensions = readSimpleField("noOfArrayDimensions", readSignedInt(readBuffer, 32));

    List<Long> arrayDimensions =
        readCountArrayField(
            "arrayDimensions", readUnsignedLong(readBuffer, 32), noOfArrayDimensions);

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("Argument");
    // Create the instance
    return new ArgumentBuilderImpl(
        name, dataType, valueRank, noOfArrayDimensions, arrayDimensions, description);
  }

  public static class ArgumentBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString name;
    private final NodeId dataType;
    private final int valueRank;
    private final int noOfArrayDimensions;
    private final List<Long> arrayDimensions;
    private final LocalizedText description;

    public ArgumentBuilderImpl(
        PascalString name,
        NodeId dataType,
        int valueRank,
        int noOfArrayDimensions,
        List<Long> arrayDimensions,
        LocalizedText description) {
      this.name = name;
      this.dataType = dataType;
      this.valueRank = valueRank;
      this.noOfArrayDimensions = noOfArrayDimensions;
      this.arrayDimensions = arrayDimensions;
      this.description = description;
    }

    public Argument build() {
      Argument argument =
          new Argument(
              name, dataType, valueRank, noOfArrayDimensions, arrayDimensions, description);
      return argument;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Argument)) {
      return false;
    }
    Argument that = (Argument) o;
    return (getName() == that.getName())
        && (getDataType() == that.getDataType())
        && (getValueRank() == that.getValueRank())
        && (getNoOfArrayDimensions() == that.getNoOfArrayDimensions())
        && (getArrayDimensions() == that.getArrayDimensions())
        && (getDescription() == that.getDescription())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getName(),
        getDataType(),
        getValueRank(),
        getNoOfArrayDimensions(),
        getArrayDimensions(),
        getDescription());
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
