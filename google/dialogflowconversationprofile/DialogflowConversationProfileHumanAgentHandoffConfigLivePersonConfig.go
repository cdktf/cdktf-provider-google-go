// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfig struct {
	// Account number of the LivePerson account to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#account_number DialogflowConversationProfile#account_number}
	AccountNumber *string `field:"required" json:"accountNumber" yaml:"accountNumber"`
}

