// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigOverlaysAnimationsAnimationFade struct {
	// Required. Type of fade animation: 'FADE_IN' or 'FADE_OUT'. The possible values are:.
	//
	// * 'FADE_TYPE_UNSPECIFIED': The fade type is not specified.
	//
	// * 'FADE_IN': Fade the overlay object into view.
	//
	// * 'FADE_OUT': Fade the overlay object out of view. Possible values: ["FADE_TYPE_UNSPECIFIED", "FADE_IN", "FADE_OUT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/transcoder_job#fade_type TranscoderJob#fade_type}
	FadeType *string `field:"required" json:"fadeType" yaml:"fadeType"`
	// The time to end the fade animation, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/transcoder_job#end_time_offset TranscoderJob#end_time_offset}
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The time to start the fade animation, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/transcoder_job#start_time_offset TranscoderJob#start_time_offset}
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
	// xy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/transcoder_job#xy TranscoderJob#xy}
	Xy *TranscoderJobConfigOverlaysAnimationsAnimationFadeXy `field:"optional" json:"xy" yaml:"xy"`
}

