// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob


type TranscoderJobConfigOverlaysAnimations struct {
	// animation_fade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/transcoder_job#animation_fade TranscoderJob#animation_fade}
	AnimationFade *TranscoderJobConfigOverlaysAnimationsAnimationFade `field:"optional" json:"animationFade" yaml:"animationFade"`
}

