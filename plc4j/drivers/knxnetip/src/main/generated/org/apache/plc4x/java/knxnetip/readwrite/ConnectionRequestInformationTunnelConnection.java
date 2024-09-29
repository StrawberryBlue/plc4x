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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ConnectionRequestInformationTunnelConnection extends ConnectionRequestInformation
    implements Message {

  // Accessors for discriminator values.
  public Short getConnectionType() {
    return (short) 0x04;
  }

  // Properties.
  protected final KnxLayer knxLayer;

  public ConnectionRequestInformationTunnelConnection(KnxLayer knxLayer) {
    super();
    this.knxLayer = knxLayer;
  }

  public KnxLayer getKnxLayer() {
    return knxLayer;
  }

  @Override
  protected void serializeConnectionRequestInformationChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ConnectionRequestInformationTunnelConnection");

    // Simple Field (knxLayer)
    writeSimpleEnumField(
        "knxLayer",
        "KnxLayer",
        knxLayer,
        writeEnum(KnxLayer::getValue, KnxLayer::name, writeUnsignedShort(writeBuffer, 8)));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("ConnectionRequestInformationTunnelConnection");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ConnectionRequestInformationTunnelConnection _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (knxLayer)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ConnectionRequestInformationBuilder staticParseConnectionRequestInformationBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ConnectionRequestInformationTunnelConnection");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    KnxLayer knxLayer =
        readEnumField(
            "knxLayer",
            "KnxLayer",
            readEnum(KnxLayer::enumForValue, readUnsignedShort(readBuffer, 8)));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    readBuffer.closeContext("ConnectionRequestInformationTunnelConnection");
    // Create the instance
    return new ConnectionRequestInformationTunnelConnectionBuilderImpl(knxLayer);
  }

  public static class ConnectionRequestInformationTunnelConnectionBuilderImpl
      implements ConnectionRequestInformation.ConnectionRequestInformationBuilder {
    private final KnxLayer knxLayer;

    public ConnectionRequestInformationTunnelConnectionBuilderImpl(KnxLayer knxLayer) {
      this.knxLayer = knxLayer;
    }

    public ConnectionRequestInformationTunnelConnection build() {
      ConnectionRequestInformationTunnelConnection connectionRequestInformationTunnelConnection =
          new ConnectionRequestInformationTunnelConnection(knxLayer);
      return connectionRequestInformationTunnelConnection;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ConnectionRequestInformationTunnelConnection)) {
      return false;
    }
    ConnectionRequestInformationTunnelConnection that =
        (ConnectionRequestInformationTunnelConnection) o;
    return (getKnxLayer() == that.getKnxLayer()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getKnxLayer());
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
