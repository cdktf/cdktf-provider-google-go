// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxsecuritysettings


type DialogflowCxSecuritySettingsAudioExportSettings struct {
	// Filename pattern for exported audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/dialogflow_cx_security_settings#audio_export_pattern DialogflowCxSecuritySettings#audio_export_pattern}
	AudioExportPattern *string `field:"optional" json:"audioExportPattern" yaml:"audioExportPattern"`
	// File format for exported audio file.
	//
	// Currently only in telephony recordings.
	// * MULAW: G.711 mu-law PCM with 8kHz sample rate.
	// * MP3: MP3 file format.
	// * OGG: OGG Vorbis. Possible values: ["MULAW", "MP3", "OGG"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/dialogflow_cx_security_settings#audio_format DialogflowCxSecuritySettings#audio_format}
	AudioFormat *string `field:"optional" json:"audioFormat" yaml:"audioFormat"`
	// Enable audio redaction if it is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/dialogflow_cx_security_settings#enable_audio_redaction DialogflowCxSecuritySettings#enable_audio_redaction}
	EnableAudioRedaction interface{} `field:"optional" json:"enableAudioRedaction" yaml:"enableAudioRedaction"`
	// Cloud Storage bucket to export audio record to.
	//
	// Setting this field would grant the Storage Object Creator role to the Dialogflow Service Agent. API caller that tries to modify this field should have the permission of storage.buckets.setIamPolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/dialogflow_cx_security_settings#gcs_bucket DialogflowCxSecuritySettings#gcs_bucket}
	GcsBucket *string `field:"optional" json:"gcsBucket" yaml:"gcsBucket"`
}

