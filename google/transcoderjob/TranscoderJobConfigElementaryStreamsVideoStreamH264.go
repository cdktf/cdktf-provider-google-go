// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigElementaryStreamsVideoStreamH264 struct {
	// The video bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#bitrate_bps TranscoderJob#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// The target video frame rate in frames per second (FPS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#frame_rate TranscoderJob#frame_rate}
	FrameRate *float64 `field:"required" json:"frameRate" yaml:"frameRate"`
	// Target CRF level. The default is '21'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#crf_level TranscoderJob#crf_level}
	CrfLevel *float64 `field:"optional" json:"crfLevel" yaml:"crfLevel"`
	// The entropy coder to use. The default is 'cabac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#entropy_coder TranscoderJob#entropy_coder}
	EntropyCoder *string `field:"optional" json:"entropyCoder" yaml:"entropyCoder"`
	// Select the GOP size based on the specified duration. The default is '3s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#gop_duration TranscoderJob#gop_duration}
	GopDuration *string `field:"optional" json:"gopDuration" yaml:"gopDuration"`
	// The height of the video in pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#height_pixels TranscoderJob#height_pixels}
	HeightPixels *float64 `field:"optional" json:"heightPixels" yaml:"heightPixels"`
	// hlg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#hlg TranscoderJob#hlg}
	Hlg *TranscoderJobConfigElementaryStreamsVideoStreamH264Hlg `field:"optional" json:"hlg" yaml:"hlg"`
	// Pixel format to use. The default is 'yuv420p'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#pixel_format TranscoderJob#pixel_format}
	PixelFormat *string `field:"optional" json:"pixelFormat" yaml:"pixelFormat"`
	// Enforces the specified codec preset. The default is 'veryfast'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#preset TranscoderJob#preset}
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Enforces the specified codec profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#profile TranscoderJob#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Specify the mode. The default is 'vbr'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#rate_control_mode TranscoderJob#rate_control_mode}
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// sdr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#sdr TranscoderJob#sdr}
	Sdr *TranscoderJobConfigElementaryStreamsVideoStreamH264Sdr `field:"optional" json:"sdr" yaml:"sdr"`
	// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#vbv_fullness_bits TranscoderJob#vbv_fullness_bits}
	VbvFullnessBits *float64 `field:"optional" json:"vbvFullnessBits" yaml:"vbvFullnessBits"`
	// Size of the Video Buffering Verifier (VBV) buffer in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#vbv_size_bits TranscoderJob#vbv_size_bits}
	VbvSizeBits *float64 `field:"optional" json:"vbvSizeBits" yaml:"vbvSizeBits"`
	// The width of the video in pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/transcoder_job#width_pixels TranscoderJob#width_pixels}
	WidthPixels *float64 `field:"optional" json:"widthPixels" yaml:"widthPixels"`
}

