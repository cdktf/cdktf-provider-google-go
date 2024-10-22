// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall struct {
	// Transfer the call to a phone number in E.164 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/dialogflow_cx_page#phone_number DialogflowCxPage#phone_number}
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}

