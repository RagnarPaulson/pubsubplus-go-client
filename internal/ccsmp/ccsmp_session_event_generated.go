// pubsubplus-go-client
//
// Copyright 2021-2022 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ccsmp

// Code generated by ccsmp_session_event_generator.go via go generate. DO NOT EDIT.

/*
#include "solclient/solClient.h"
*/
import "C"

const (
	// SolClientSessionEventUpNotice: The Session is established.
	SolClientSessionEventUpNotice SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_UP_NOTICE
	// SolClientSessionEventDownError: The Session was established and then went down.
	SolClientSessionEventDownError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_DOWN_ERROR
	// SolClientSessionEventConnectFailedError: The Session attempted to connect but was unsuccessful.
	SolClientSessionEventConnectFailedError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_CONNECT_FAILED_ERROR
	// SolClientSessionEventRejectedMsgError: The appliance rejected a published message.
	SolClientSessionEventRejectedMsgError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_REJECTED_MSG_ERROR
	// SolClientSessionEventSubscriptionError: The appliance rejected a subscription (add or remove).
	SolClientSessionEventSubscriptionError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_SUBSCRIPTION_ERROR
	// SolClientSessionEventRxMsgTooBigError: The API discarded a received message that exceeded the Session buffer size.
	SolClientSessionEventRxMsgTooBigError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_RX_MSG_TOO_BIG_ERROR
	// SolClientSessionEventAcknowledgement: The oldest transmitted Persistent/Non-Persistent message that has been acknowledged.
	SolClientSessionEventAcknowledgement SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_ACKNOWLEDGEMENT
	// SolClientSessionEventAssuredPublishingUp: Deprecated -- see notes in solClient_session_startAssuredPublishing. The AD Handshake (that is, Guaranteed Delivery handshake) has completed for the publisher and Guaranteed messages can be sent.
	SolClientSessionEventAssuredPublishingUp SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_ASSURED_PUBLISHING_UP
	// SolClientSessionEventAssuredConnectFailed: Deprecated -- see notes in solClient_session_startAssuredPublishing. The appliance rejected the AD Handshake to start Guaranteed publishing. Use ::SOLCLIENT_SESSION_EVENT_ASSURED_DELIVERY_DOWN instead.
	SolClientSessionEventAssuredConnectFailed SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_ASSURED_CONNECT_FAILED
	// SolClientSessionEventAssuredDeliveryDown: Guaranteed Delivery publishing is not available. The guaranteed delivery capability on the session has been disabled by some action on the appliance.
	SolClientSessionEventAssuredDeliveryDown SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_ASSURED_DELIVERY_DOWN
	// SolClientSessionEventTeUnsubscribeError: The Topic Endpoint unsubscribe command failed.
	SolClientSessionEventTeUnsubscribeError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_TE_UNSUBSCRIBE_ERROR
	// SolClientSessionEventTeUnsubscribeOk: The Topic Endpoint unsubscribe completed.
	SolClientSessionEventTeUnsubscribeOk SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_TE_UNSUBSCRIBE_OK
	// SolClientSessionEventCanSend: The send is no longer blocked.
	SolClientSessionEventCanSend SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_CAN_SEND
	// SolClientSessionEventReconnectingNotice: The Session has gone down, and an automatic reconnect attempt is in progress.
	SolClientSessionEventReconnectingNotice SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_RECONNECTING_NOTICE
	// SolClientSessionEventReconnectedNotice: The automatic reconnect of the Session was successful, and the Session was established again.
	SolClientSessionEventReconnectedNotice SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_RECONNECTED_NOTICE
	// SolClientSessionEventProvisionError: The endpoint create/delete command failed.
	SolClientSessionEventProvisionError SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_PROVISION_ERROR
	// SolClientSessionEventProvisionOk: The endpoint create/delete command completed.
	SolClientSessionEventProvisionOk SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_PROVISION_OK
	// SolClientSessionEventSubscriptionOk: The subscribe or unsubscribe operation has succeeded.
	SolClientSessionEventSubscriptionOk SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_SUBSCRIPTION_OK
	// SolClientSessionEventVirtualRouterNameChanged: The appliance's Virtual Router Name changed during a reconnect operation. This could render existing queues or temporary topics invalid.
	SolClientSessionEventVirtualRouterNameChanged SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_VIRTUAL_ROUTER_NAME_CHANGED
	// SolClientSessionEventModifypropOk: The session property modification completed.
	SolClientSessionEventModifypropOk SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_MODIFYPROP_OK
	// SolClientSessionEventModifypropFail: The session property modification failed.
	SolClientSessionEventModifypropFail SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_MODIFYPROP_FAIL
	// SolClientSessionEventRepublishUnackedMessages: After successfully reconnecting a disconnected session, the SDK received an unknown publisher flow name response when reconnecting the GD publisher flow. If configured to auto-retry (See ::SOLCLIENT_SESSION_PROP_GD_RECONNECT_FAIL_ACTION.) this event is generated to indicate how many unacknowledged messages are retransmitted on success. As the publisher state has been lost on failover, receiving this event may indicate that some messages have been duplicated in the system.
	SolClientSessionEventRepublishUnackedMessages SolClientSessionEvent = C.SOLCLIENT_SESSION_EVENT_REPUBLISH_UNACKED_MESSAGES
)