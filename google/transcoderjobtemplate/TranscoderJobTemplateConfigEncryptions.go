// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigEncryptions struct {
	// Identifier for this set of encryption options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#id TranscoderJobTemplate#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// aes128 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#aes128 TranscoderJobTemplate#aes128}
	Aes128 *TranscoderJobTemplateConfigEncryptionsAes128 `field:"optional" json:"aes128" yaml:"aes128"`
	// drm_systems block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#drm_systems TranscoderJobTemplate#drm_systems}
	DrmSystems *TranscoderJobTemplateConfigEncryptionsDrmSystems `field:"optional" json:"drmSystems" yaml:"drmSystems"`
	// mpeg_cenc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#mpeg_cenc TranscoderJobTemplate#mpeg_cenc}
	MpegCenc *TranscoderJobTemplateConfigEncryptionsMpegCenc `field:"optional" json:"mpegCenc" yaml:"mpegCenc"`
	// sample_aes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#sample_aes TranscoderJobTemplate#sample_aes}
	SampleAes *TranscoderJobTemplateConfigEncryptionsSampleAes `field:"optional" json:"sampleAes" yaml:"sampleAes"`
	// secret_manager_key_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/transcoder_job_template#secret_manager_key_source TranscoderJobTemplate#secret_manager_key_source}
	SecretManagerKeySource *TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource `field:"optional" json:"secretManagerKeySource" yaml:"secretManagerKeySource"`
}

