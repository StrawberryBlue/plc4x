//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package readwrite

import (
	"encoding/xml"
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/ads/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi"
)

type AdsXmlParserHelper struct {
}

func (m AdsXmlParserHelper) Parse(typeName string, xmlString string) (spi.Message, error) {
	switch typeName {
	case "AdsMultiRequestItem":
		var obj *model.AdsMultiRequestItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AmsTCPPacket":
		var obj *model.AmsTCPPacket
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "State":
		var obj *model.State
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AmsPacket":
		var obj *model.AmsPacket
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AmsSerialFrame":
		var obj *model.AmsSerialFrame
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AmsSerialAcknowledgeFrame":
		var obj *model.AmsSerialAcknowledgeFrame
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AdsData":
		var obj *model.AdsData
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AmsNetId":
		var obj *model.AmsNetId
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AdsStampHeader":
		var obj *model.AdsStampHeader
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AmsSerialResetFrame":
		var obj *model.AmsSerialResetFrame
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	case "AdsNotificationSample":
		var obj *model.AdsNotificationSample
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.New("error unmarshalling xml: " + err.Error())
		}
		return obj, nil
	}
	return nil, errors.New("Unsupported type " + typeName)
}
