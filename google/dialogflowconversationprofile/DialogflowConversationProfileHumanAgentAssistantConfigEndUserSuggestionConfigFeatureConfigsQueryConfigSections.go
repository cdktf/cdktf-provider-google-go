// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSections struct {
	// The selected sections chosen to return when requesting a summary of a conversation If not provided the default selection will be "{SITUATION, ACTION, RESULT}".
	//
	// Possible values: ["SECTION_TYPE_UNSPECIFIED", "SITUATION", "ACTION", "RESOLUTION", "REASON_FOR_CANCELLATION", "CUSTOMER_SATISFACTION", "ENTITIES"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#section_types DialogflowConversationProfile#section_types}
	SectionTypes *[]*string `field:"optional" json:"sectionTypes" yaml:"sectionTypes"`
}

