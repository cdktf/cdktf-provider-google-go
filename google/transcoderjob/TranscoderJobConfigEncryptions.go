// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigEncryptions struct {
	// Identifier for this set of encryption options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job#id TranscoderJob#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// aes128 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job#aes128 TranscoderJob#aes128}
	Aes128 *TranscoderJobConfigEncryptionsAes128 `field:"optional" json:"aes128" yaml:"aes128"`
	// drm_systems block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job#drm_systems TranscoderJob#drm_systems}
	DrmSystems *TranscoderJobConfigEncryptionsDrmSystems `field:"optional" json:"drmSystems" yaml:"drmSystems"`
	// mpeg_cenc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job#mpeg_cenc TranscoderJob#mpeg_cenc}
	MpegCenc *TranscoderJobConfigEncryptionsMpegCenc `field:"optional" json:"mpegCenc" yaml:"mpegCenc"`
	// sample_aes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job#sample_aes TranscoderJob#sample_aes}
	SampleAes *TranscoderJobConfigEncryptionsSampleAes `field:"optional" json:"sampleAes" yaml:"sampleAes"`
	// secret_manager_key_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job#secret_manager_key_source TranscoderJob#secret_manager_key_source}
	SecretManagerKeySource *TranscoderJobConfigEncryptionsSecretManagerKeySource `field:"optional" json:"secretManagerKeySource" yaml:"secretManagerKeySource"`
}

