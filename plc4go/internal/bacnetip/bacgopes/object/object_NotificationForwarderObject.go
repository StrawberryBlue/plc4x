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

package object

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type NotificationForwarderObject struct {
	Object
}

func NewNotificationForwarderObject(options ...Option) (*NotificationForwarderObject, error) {
	n := new(NotificationForwarderObject)
	objectType := "notificationForwarder"
	properties := []Property{
		NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
		NewReadableProperty("reliability", V2P(NewReliability)),
		NewReadableProperty("outOfService", V2P(NewBoolean)),
		NewReadableProperty("recipientList", ListOfP(NewDestination)),
		NewWritableProperty("subscribedRecipients", ListOfP(NewEventNotificationSubscription)),
		NewReadableProperty("processIdentifierFilter", V2P(NewProcessIdSelection)),
		NewOptionalProperty("portFilter", ArrayOfP(NewPortPermission, 0, 0)),
		NewReadableProperty("localForwardingOnly", V2P(NewBoolean)),
		NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
	}
	var err error
	n.Object, err = NewObject(Combine(options, WithObjectType(objectType), WithObjectExtraProperties(properties))...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating object")
	}
	if _, err := RegisterObjectType(NKW(KWCls, n)); err != nil {
		return nil, errors.Wrap(err, "error registering object type")
	}
	return n, nil
}
