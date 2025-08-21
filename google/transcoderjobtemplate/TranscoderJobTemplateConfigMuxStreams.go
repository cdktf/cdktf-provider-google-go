// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigMuxStreams struct {
	// The container format. The default is 'mp4'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job_template#container TranscoderJobTemplate#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// List of ElementaryStream.key values multiplexed in this stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job_template#elementary_streams TranscoderJobTemplate#elementary_streams}
	ElementaryStreams *[]*string `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// Identifier of the encryption configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job_template#encryption_id TranscoderJobTemplate#encryption_id}
	EncryptionId *string `field:"optional" json:"encryptionId" yaml:"encryptionId"`
	// The name of the generated file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job_template#file_name TranscoderJobTemplate#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// A unique key for this multiplexed stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job_template#key TranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// segment_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/transcoder_job_template#segment_settings TranscoderJobTemplate#segment_settings}
	SegmentSettings *TranscoderJobTemplateConfigMuxStreamsSegmentSettings `field:"optional" json:"segmentSettings" yaml:"segmentSettings"`
}

