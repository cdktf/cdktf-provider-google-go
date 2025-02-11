// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigOverlays struct {
	// animations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/transcoder_job_template#animations TranscoderJobTemplate#animations}
	Animations interface{} `field:"optional" json:"animations" yaml:"animations"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/transcoder_job_template#image TranscoderJobTemplate#image}
	Image *TranscoderJobTemplateConfigOverlaysImage `field:"optional" json:"image" yaml:"image"`
}

