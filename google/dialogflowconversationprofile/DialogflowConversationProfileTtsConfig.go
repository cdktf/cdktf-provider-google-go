// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileTtsConfig struct {
	// An identifier which selects 'audio effects' profiles that are applied on (post synthesized) text to speech.
	//
	// Effects are applied on top of each other in the order they are given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#effects_profile_id DialogflowConversationProfile#effects_profile_id}
	EffectsProfileId *[]*string `field:"optional" json:"effectsProfileId" yaml:"effectsProfileId"`
	// Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20 semitones from the original pitch. -20 means decrease 20 semitones from the original pitch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#pitch DialogflowConversationProfile#pitch}
	Pitch *float64 `field:"optional" json:"pitch" yaml:"pitch"`
	// Speaking rate/speed, in the range [0.25, 4.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#speaking_rate DialogflowConversationProfile#speaking_rate}
	SpeakingRate *float64 `field:"optional" json:"speakingRate" yaml:"speakingRate"`
	// voice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#voice DialogflowConversationProfile#voice}
	Voice *DialogflowConversationProfileTtsConfigVoice `field:"optional" json:"voice" yaml:"voice"`
	// Volume gain (in dB) of the normal native volume supported by the specific voice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#volume_gain_db DialogflowConversationProfile#volume_gain_db}
	VolumeGainDb *float64 `field:"optional" json:"volumeGainDb" yaml:"volumeGainDb"`
}

