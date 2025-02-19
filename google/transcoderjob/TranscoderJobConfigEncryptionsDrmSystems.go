// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigEncryptionsDrmSystems struct {
	// clearkey block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/transcoder_job#clearkey TranscoderJob#clearkey}
	Clearkey *TranscoderJobConfigEncryptionsDrmSystemsClearkey `field:"optional" json:"clearkey" yaml:"clearkey"`
	// fairplay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/transcoder_job#fairplay TranscoderJob#fairplay}
	Fairplay *TranscoderJobConfigEncryptionsDrmSystemsFairplay `field:"optional" json:"fairplay" yaml:"fairplay"`
	// playready block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/transcoder_job#playready TranscoderJob#playready}
	Playready *TranscoderJobConfigEncryptionsDrmSystemsPlayready `field:"optional" json:"playready" yaml:"playready"`
	// widevine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/transcoder_job#widevine TranscoderJob#widevine}
	Widevine *TranscoderJobConfigEncryptionsDrmSystemsWidevine `field:"optional" json:"widevine" yaml:"widevine"`
}

