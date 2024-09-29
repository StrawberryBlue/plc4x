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

public class RelativePathElement extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "539";
  }

  // Properties.
  protected final NodeId referenceTypeId;
  protected final boolean includeSubtypes;
  protected final boolean isInverse;
  protected final QualifiedName targetName;

  public RelativePathElement(
      NodeId referenceTypeId,
      boolean includeSubtypes,
      boolean isInverse,
      QualifiedName targetName) {
    super();
    this.referenceTypeId = referenceTypeId;
    this.includeSubtypes = includeSubtypes;
    this.isInverse = isInverse;
    this.targetName = targetName;
  }

  public NodeId getReferenceTypeId() {
    return referenceTypeId;
  }

  public boolean getIncludeSubtypes() {
    return includeSubtypes;
  }

  public boolean getIsInverse() {
    return isInverse;
  }

  public QualifiedName getTargetName() {
    return targetName;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RelativePathElement");

    // Simple Field (referenceTypeId)
    writeSimpleField("referenceTypeId", referenceTypeId, writeComplex(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 6));

    // Simple Field (includeSubtypes)
    writeSimpleField("includeSubtypes", includeSubtypes, writeBoolean(writeBuffer));

    // Simple Field (isInverse)
    writeSimpleField("isInverse", isInverse, writeBoolean(writeBuffer));

    // Simple Field (targetName)
    writeSimpleField("targetName", targetName, writeComplex(writeBuffer));

    writeBuffer.popContext("RelativePathElement");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RelativePathElement _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (referenceTypeId)
    lengthInBits += referenceTypeId.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 6;

    // Simple field (includeSubtypes)
    lengthInBits += 1;

    // Simple field (isInverse)
    lengthInBits += 1;

    // Simple field (targetName)
    lengthInBits += targetName.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RelativePathElement");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId referenceTypeId =
        readSimpleField(
            "referenceTypeId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 6), (byte) 0x00);

    boolean includeSubtypes = readSimpleField("includeSubtypes", readBoolean(readBuffer));

    boolean isInverse = readSimpleField("isInverse", readBoolean(readBuffer));

    QualifiedName targetName =
        readSimpleField(
            "targetName", readComplex(() -> QualifiedName.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("RelativePathElement");
    // Create the instance
    return new RelativePathElementBuilderImpl(
        referenceTypeId, includeSubtypes, isInverse, targetName);
  }

  public static class RelativePathElementBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId referenceTypeId;
    private final boolean includeSubtypes;
    private final boolean isInverse;
    private final QualifiedName targetName;

    public RelativePathElementBuilderImpl(
        NodeId referenceTypeId,
        boolean includeSubtypes,
        boolean isInverse,
        QualifiedName targetName) {
      this.referenceTypeId = referenceTypeId;
      this.includeSubtypes = includeSubtypes;
      this.isInverse = isInverse;
      this.targetName = targetName;
    }

    public RelativePathElement build() {
      RelativePathElement relativePathElement =
          new RelativePathElement(referenceTypeId, includeSubtypes, isInverse, targetName);
      return relativePathElement;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RelativePathElement)) {
      return false;
    }
    RelativePathElement that = (RelativePathElement) o;
    return (getReferenceTypeId() == that.getReferenceTypeId())
        && (getIncludeSubtypes() == that.getIncludeSubtypes())
        && (getIsInverse() == that.getIsInverse())
        && (getTargetName() == that.getTargetName())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getReferenceTypeId(),
        getIncludeSubtypes(),
        getIsInverse(),
        getTargetName());
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
