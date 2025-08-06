// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigOverlaysAnimationsAnimationFade struct {
	// Required. Type of fade animation: 'FADE_IN' or 'FADE_OUT'. The possible values are:.
	//
	// * 'FADE_TYPE_UNSPECIFIED': The fade type is not specified.
	//
	// * 'FADE_IN': Fade the overlay object into view.
	//
	// * 'FADE_OUT': Fade the overlay object out of view. Possible values: ["FADE_TYPE_UNSPECIFIED", "FADE_IN", "FADE_OUT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/transcoder_job_template#fade_type TranscoderJobTemplate#fade_type}
	FadeType *string `field:"required" json:"fadeType" yaml:"fadeType"`
	// The time to end the fade animation, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/transcoder_job_template#end_time_offset TranscoderJobTemplate#end_time_offset}
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The time to start the fade animation, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/transcoder_job_template#start_time_offset TranscoderJobTemplate#start_time_offset}
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
	// xy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/transcoder_job_template#xy TranscoderJobTemplate#xy}
	Xy *TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXy `field:"optional" json:"xy" yaml:"xy"`
}

