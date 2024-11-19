// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowAdvancedSettings struct {
	// audio_export_gcs_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/dialogflow_cx_flow#audio_export_gcs_destination DialogflowCxFlow#audio_export_gcs_destination}
	AudioExportGcsDestination *DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination `field:"optional" json:"audioExportGcsDestination" yaml:"audioExportGcsDestination"`
	// dtmf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/dialogflow_cx_flow#dtmf_settings DialogflowCxFlow#dtmf_settings}
	DtmfSettings *DialogflowCxFlowAdvancedSettingsDtmfSettings `field:"optional" json:"dtmfSettings" yaml:"dtmfSettings"`
	// logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/dialogflow_cx_flow#logging_settings DialogflowCxFlow#logging_settings}
	LoggingSettings *DialogflowCxFlowAdvancedSettingsLoggingSettings `field:"optional" json:"loggingSettings" yaml:"loggingSettings"`
	// speech_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/dialogflow_cx_flow#speech_settings DialogflowCxFlow#speech_settings}
	SpeechSettings *DialogflowCxFlowAdvancedSettingsSpeechSettings `field:"optional" json:"speechSettings" yaml:"speechSettings"`
}

