// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigElementaryStreams struct {
	// audio_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/transcoder_job_template#audio_stream TranscoderJobTemplate#audio_stream}
	AudioStream *TranscoderJobTemplateConfigElementaryStreamsAudioStream `field:"optional" json:"audioStream" yaml:"audioStream"`
	// A unique key for this atom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/transcoder_job_template#key TranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// video_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/transcoder_job_template#video_stream TranscoderJobTemplate#video_stream}
	VideoStream *TranscoderJobTemplateConfigElementaryStreamsVideoStream `field:"optional" json:"videoStream" yaml:"videoStream"`
}

