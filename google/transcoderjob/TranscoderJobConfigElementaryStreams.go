// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigElementaryStreams struct {
	// audio_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/transcoder_job#audio_stream TranscoderJob#audio_stream}
	AudioStream *TranscoderJobConfigElementaryStreamsAudioStream `field:"optional" json:"audioStream" yaml:"audioStream"`
	// A unique key for this atom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/transcoder_job#key TranscoderJob#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// video_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/transcoder_job#video_stream TranscoderJob#video_stream}
	VideoStream *TranscoderJobConfigElementaryStreamsVideoStream `field:"optional" json:"videoStream" yaml:"videoStream"`
}

