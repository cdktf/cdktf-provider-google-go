// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigElementaryStreamsVideoStreamH264 struct {
	// The video bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#bitrate_bps TranscoderJobTemplate#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// The target video frame rate in frames per second (FPS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#frame_rate TranscoderJobTemplate#frame_rate}
	FrameRate *float64 `field:"required" json:"frameRate" yaml:"frameRate"`
	// Target CRF level. The default is '21'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#crf_level TranscoderJobTemplate#crf_level}
	CrfLevel *float64 `field:"optional" json:"crfLevel" yaml:"crfLevel"`
	// The entropy coder to use. The default is 'cabac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#entropy_coder TranscoderJobTemplate#entropy_coder}
	EntropyCoder *string `field:"optional" json:"entropyCoder" yaml:"entropyCoder"`
	// Select the GOP size based on the specified duration. The default is '3s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#gop_duration TranscoderJobTemplate#gop_duration}
	GopDuration *string `field:"optional" json:"gopDuration" yaml:"gopDuration"`
	// The height of the video in pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#height_pixels TranscoderJobTemplate#height_pixels}
	HeightPixels *float64 `field:"optional" json:"heightPixels" yaml:"heightPixels"`
	// hlg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#hlg TranscoderJobTemplate#hlg}
	Hlg *TranscoderJobTemplateConfigElementaryStreamsVideoStreamH264Hlg `field:"optional" json:"hlg" yaml:"hlg"`
	// Pixel format to use. The default is 'yuv420p'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#pixel_format TranscoderJobTemplate#pixel_format}
	PixelFormat *string `field:"optional" json:"pixelFormat" yaml:"pixelFormat"`
	// Enforces the specified codec preset. The default is 'veryfast'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#preset TranscoderJobTemplate#preset}
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Enforces the specified codec profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#profile TranscoderJobTemplate#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Specify the mode. The default is 'vbr'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#rate_control_mode TranscoderJobTemplate#rate_control_mode}
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// sdr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#sdr TranscoderJobTemplate#sdr}
	Sdr *TranscoderJobTemplateConfigElementaryStreamsVideoStreamH264Sdr `field:"optional" json:"sdr" yaml:"sdr"`
	// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#vbv_fullness_bits TranscoderJobTemplate#vbv_fullness_bits}
	VbvFullnessBits *float64 `field:"optional" json:"vbvFullnessBits" yaml:"vbvFullnessBits"`
	// Size of the Video Buffering Verifier (VBV) buffer in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#vbv_size_bits TranscoderJobTemplate#vbv_size_bits}
	VbvSizeBits *float64 `field:"optional" json:"vbvSizeBits" yaml:"vbvSizeBits"`
	// The width of the video in pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/transcoder_job_template#width_pixels TranscoderJobTemplate#width_pixels}
	WidthPixels *float64 `field:"optional" json:"widthPixels" yaml:"widthPixels"`
}

