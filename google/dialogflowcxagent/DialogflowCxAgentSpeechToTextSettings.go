// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxagent


type DialogflowCxAgentSpeechToTextSettings struct {
	// Whether to use speech adaptation for speech recognition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/dialogflow_cx_agent#enable_speech_adaptation DialogflowCxAgent#enable_speech_adaptation}
	EnableSpeechAdaptation interface{} `field:"optional" json:"enableSpeechAdaptation" yaml:"enableSpeechAdaptation"`
}

