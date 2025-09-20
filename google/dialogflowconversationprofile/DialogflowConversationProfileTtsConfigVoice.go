// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileTtsConfigVoice struct {
	// The name of the voice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#name DialogflowConversationProfile#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The preferred gender of the voice. Possible values: ["SSML_VOICE_GENDER_UNSPECIFIED", "SSML_VOICE_GENDER_MALE", "SSML_VOICE_GENDER_FEMALE", "SSML_VOICE_GENDER_NEUTRAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#ssml_gender DialogflowConversationProfile#ssml_gender}
	SsmlGender *string `field:"optional" json:"ssmlGender" yaml:"ssmlGender"`
}

