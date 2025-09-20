// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigElementaryStreamsAudioStream struct {
	// Audio bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/transcoder_job_template#bitrate_bps TranscoderJobTemplate#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// Number of audio channels. The default is '2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/transcoder_job_template#channel_count TranscoderJobTemplate#channel_count}
	ChannelCount *float64 `field:"optional" json:"channelCount" yaml:"channelCount"`
	// A list of channel names specifying layout of the audio channels.  The default is ["fl", "fr"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/transcoder_job_template#channel_layout TranscoderJobTemplate#channel_layout}
	ChannelLayout *[]*string `field:"optional" json:"channelLayout" yaml:"channelLayout"`
	// The codec for this audio stream. The default is 'aac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/transcoder_job_template#codec TranscoderJobTemplate#codec}
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The audio sample rate in Hertz. The default is '48000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/transcoder_job_template#sample_rate_hertz TranscoderJobTemplate#sample_rate_hertz}
	SampleRateHertz *float64 `field:"optional" json:"sampleRateHertz" yaml:"sampleRateHertz"`
}

