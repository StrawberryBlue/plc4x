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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7VarRequestParameterItemAddress extends S7VarRequestParameterItem implements Message {

  // Accessors for discriminator values.
  public Short getItemType() {
    return (short) 0x12;
  }

  // Properties.
  protected final S7Address address;

  public S7VarRequestParameterItemAddress(S7Address address) {
    super();
    this.address = address;
  }

  public S7Address getAddress() {
    return address;
  }

  @Override
  protected void serializeS7VarRequestParameterItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7VarRequestParameterItemAddress");

    // Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short itemLength = (short) (getAddress().getLengthInBytes());
    writeImplicitField("itemLength", itemLength, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (address)
    writeSimpleField("address", address, writeComplex(writeBuffer));

    writeBuffer.popContext("S7VarRequestParameterItemAddress");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7VarRequestParameterItemAddress _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (itemLength)
    lengthInBits += 8;

    // Simple field (address)
    lengthInBits += address.getLengthInBits();

    return lengthInBits;
  }

  public static S7VarRequestParameterItemBuilder staticParseS7VarRequestParameterItemBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("S7VarRequestParameterItemAddress");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short itemLength = readImplicitField("itemLength", readUnsignedShort(readBuffer, 8));

    S7Address address =
        readSimpleField(
            "address", readComplex(() -> S7Address.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("S7VarRequestParameterItemAddress");
    // Create the instance
    return new S7VarRequestParameterItemAddressBuilderImpl(address);
  }

  public static class S7VarRequestParameterItemAddressBuilderImpl
      implements S7VarRequestParameterItem.S7VarRequestParameterItemBuilder {
    private final S7Address address;

    public S7VarRequestParameterItemAddressBuilderImpl(S7Address address) {
      this.address = address;
    }

    public S7VarRequestParameterItemAddress build() {
      S7VarRequestParameterItemAddress s7VarRequestParameterItemAddress =
          new S7VarRequestParameterItemAddress(address);
      return s7VarRequestParameterItemAddress;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7VarRequestParameterItemAddress)) {
      return false;
    }
    S7VarRequestParameterItemAddress that = (S7VarRequestParameterItemAddress) o;
    return (getAddress() == that.getAddress()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAddress());
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
