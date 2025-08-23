// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigElementaryStreamsAudioStream struct {
	// Audio bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/transcoder_job#bitrate_bps TranscoderJob#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// Number of audio channels. The default is '2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/transcoder_job#channel_count TranscoderJob#channel_count}
	ChannelCount *float64 `field:"optional" json:"channelCount" yaml:"channelCount"`
	// A list of channel names specifying layout of the audio channels. The default is ["fl", "fr"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/transcoder_job#channel_layout TranscoderJob#channel_layout}
	ChannelLayout *[]*string `field:"optional" json:"channelLayout" yaml:"channelLayout"`
	// The codec for this audio stream. The default is 'aac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/transcoder_job#codec TranscoderJob#codec}
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The audio sample rate in Hertz. The default is '48000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/transcoder_job#sample_rate_hertz TranscoderJob#sample_rate_hertz}
	SampleRateHertz *float64 `field:"optional" json:"sampleRateHertz" yaml:"sampleRateHertz"`
}

