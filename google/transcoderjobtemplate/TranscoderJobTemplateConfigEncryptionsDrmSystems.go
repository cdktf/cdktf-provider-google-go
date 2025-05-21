// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigEncryptionsDrmSystems struct {
	// clearkey block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/transcoder_job_template#clearkey TranscoderJobTemplate#clearkey}
	Clearkey *TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey `field:"optional" json:"clearkey" yaml:"clearkey"`
	// fairplay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/transcoder_job_template#fairplay TranscoderJobTemplate#fairplay}
	Fairplay *TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay `field:"optional" json:"fairplay" yaml:"fairplay"`
	// playready block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/transcoder_job_template#playready TranscoderJobTemplate#playready}
	Playready *TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready `field:"optional" json:"playready" yaml:"playready"`
	// widevine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/transcoder_job_template#widevine TranscoderJobTemplate#widevine}
	Widevine *TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine `field:"optional" json:"widevine" yaml:"widevine"`
}

