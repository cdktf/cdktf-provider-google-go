// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigMuxStreams struct {
	// The container format. The default is 'mp4'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/transcoder_job#container TranscoderJob#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// List of ElementaryStream.key values multiplexed in this stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/transcoder_job#elementary_streams TranscoderJob#elementary_streams}
	ElementaryStreams *[]*string `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// Identifier of the encryption configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/transcoder_job#encryption_id TranscoderJob#encryption_id}
	EncryptionId *string `field:"optional" json:"encryptionId" yaml:"encryptionId"`
	// The name of the generated file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/transcoder_job#file_name TranscoderJob#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// A unique key for this multiplexed stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/transcoder_job#key TranscoderJob#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// segment_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/transcoder_job#segment_settings TranscoderJob#segment_settings}
	SegmentSettings *TranscoderJobConfigMuxStreamsSegmentSettings `field:"optional" json:"segmentSettings" yaml:"segmentSettings"`
}

